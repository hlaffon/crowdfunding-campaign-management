package server_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	
	"net/http"
	"net/http/httptest"
	
	"crowdfunding-campaign-management/internal/server"
	"crowdfunding-campaign-management/internal/store/mocks"
	"crowdfunding-campaign-management/model"
)

func Test_NewHandler(t *testing.T) {
	database := mocks.NewStorage(t)
	h := server.NewHandler(database)
	assert.NotNil(t, h)
}

func TestHandler_AmountPerDay(t *testing.T) {
	t.Run("Invalid project Id", func(t *testing.T) {
		database := mocks.NewStorage(t)
		h := server.NewHandler(database)
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "toto",
			},
		}

		h.AmountPerDay(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var got *model.HTTPError
		err := json.Unmarshal(w.Body.Bytes(), &got)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, got.Code)
		assert.Equal(t, "invalid project id", got.Message)
	})

	t.Run("Error requesting DB", func(t *testing.T) {
		database := mocks.NewStorage(t)
		database.EXPECT().GetTotalAmountPerDay(mock.Anything, 12).Return(nil, fmt.Errorf("oops"))

		h := server.NewHandler(database)
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "12",
			},
		}
		h.AmountPerDay(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var got *model.HTTPError
		err := json.Unmarshal(w.Body.Bytes(), &got)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, got.Code)
		assert.Equal(t, "error requesting data", got.Message)
	})

	t.Run("No data in DB", func(t *testing.T) {
		database := mocks.NewStorage(t)
		database.EXPECT().GetTotalAmountPerDay(mock.Anything, 12).Return(nil, nil)

		h := server.NewHandler(database)
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "12",
			},
		}
		h.AmountPerDay(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		var got *model.HTTPError
		err := json.Unmarshal(w.Body.Bytes(), &got)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, got.Code)
		assert.Equal(t, "no data found", got.Message)
	})

	t.Run("OK", func(t *testing.T) {
		database := mocks.NewStorage(t)
		date := time.Now()
		expRes := []*model.AmountPerDay{
			{
				Date:  date,
				Total: 10,
			},
		}
		database.EXPECT().GetTotalAmountPerDay(mock.Anything, 12).Return(expRes, nil)

		h := server.NewHandler(database)
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "12",
			},
		}

		h.AmountPerDay(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		got := w.Body.String()
		exp, err := json.Marshal(expRes)
		assert.Nil(t, err)
		assert.Equal(t, string(exp), got)
	})
}

func TestHandler_PercentagePerDay(t *testing.T) {
	t.Run("Invalid project Id", func(t *testing.T) {
		database := mocks.NewStorage(t)
		h := server.NewHandler(database)
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "toto",
			},
		}
		h.PercentagePerDay(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var got *model.HTTPError
		err := json.Unmarshal(w.Body.Bytes(), &got)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, got.Code)
		assert.Equal(t, "invalid project id", got.Message)
	})

	t.Run("Error requesting DB", func(t *testing.T) {
		database := mocks.NewStorage(t)
		database.EXPECT().GetPercentageOfGoalPerDay(mock.Anything, 12).Return(nil, fmt.Errorf("oops"))

		h := server.NewHandler(database)
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "12",
			},
		}
		h.PercentagePerDay(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var got *model.HTTPError
		err := json.Unmarshal(w.Body.Bytes(), &got)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, got.Code)
		assert.Equal(t, "error requesting data", got.Message)
	})

	t.Run("No data in DB", func(t *testing.T) {
		database := mocks.NewStorage(t)
		database.EXPECT().GetPercentageOfGoalPerDay(mock.Anything, 12).Return(nil, nil)

		h := server.NewHandler(database)
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "12",
			},
		}
		h.PercentagePerDay(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		var got *model.HTTPError
		err := json.Unmarshal(w.Body.Bytes(), &got)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, got.Code)
		assert.Equal(t, "no data found", got.Message)
	})

	t.Run("OK", func(t *testing.T) {
		database := mocks.NewStorage(t)
		date := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		expPercentage := []*model.AmountPerDay{
			{
				Date:  date,
				Total: 10.1,
			},
			{
				Date:  date.AddDate(0, 0, 1),
				Total: 25.3,
			},
			{
				Date:  date.AddDate(0, 0, 2),
				Total: 42.0,
			},
		}
		expResult := []*model.AmountPerDay{
			{
				Date:  date,
				Total: 10.1,
			},
			{
				Date:  date.AddDate(0, 0, 1),
				Total: 35.4,
			},
			{
				Date:  date.AddDate(0, 0, 2),
				Total: 77.4,
			},
		}
		database.EXPECT().GetPercentageOfGoalPerDay(mock.Anything, 12).Return(expPercentage, nil)

		h := server.NewHandler(database)
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "12",
			},
		}

		h.PercentagePerDay(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		got := w.Body.String()
		exp, err := json.Marshal(expResult)
		assert.Nil(t, err)
		assert.Equal(t, string(exp), got)
	})
}

func TestHandler_ContributionMilestones(t *testing.T) {
	t.Run("Invalid project Id", func(t *testing.T) {
		database := mocks.NewStorage(t)
		h := server.NewHandler(database)
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "toto",
			},
		}
		h.ContributionMilestones(ctx)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		var got *model.HTTPError
		err := json.Unmarshal(w.Body.Bytes(), &got)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, got.Code)
		assert.Equal(t, "invalid project id", got.Message)
	})

	t.Run("Invalid milestone", func(t *testing.T) {
		database := mocks.NewStorage(t)
		h := server.NewHandler(database)
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "12",
			},
			{
				Key:   "milestone",
				Value: "43.5",
			},
		}
		h.ContributionMilestones(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var got *model.HTTPError
		err := json.Unmarshal(w.Body.Bytes(), &got)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, got.Code)
		assert.Equal(t, "invalid milestone", got.Message)
	})

	t.Run("Error requesting data", func(t *testing.T) {
		database := mocks.NewStorage(t)
		database.EXPECT().GetMilestoneContribution(mock.Anything, 12, 70).Return(nil, fmt.Errorf("oops"))
		h := server.NewHandler(database)
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "12",
			},
			{
				Key:   "milestone",
				Value: "70",
			},
		}
		h.ContributionMilestones(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var got *model.HTTPError
		err := json.Unmarshal(w.Body.Bytes(), &got)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, got.Code)
		assert.Equal(t, "error requesting data", got.Message)
	})

	t.Run("No data in DB", func(t *testing.T) {
		database := mocks.NewStorage(t)
		database.EXPECT().GetMilestoneContribution(mock.Anything, 12, 70).Return(nil, nil)
		h := server.NewHandler(database)
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "12",
			},
			{
				Key:   "milestone",
				Value: "70",
			},
		}

		h.ContributionMilestones(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		var got *model.HTTPError
		err := json.Unmarshal(w.Body.Bytes(), &got)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, got.Code)
		assert.Equal(t, "no data found", got.Message)
	})

	t.Run("OK", func(t *testing.T) {
		database := mocks.NewStorage(t)
		now := time.Now()
		expRes := &model.MilestoneContribution{
			CreatedDate:    now,
			Milestone:      70,
			ContributionId: 1234,
			UserId:         5678,
			Amount:         80,
			Goal:           1000,
			TotalRaised:    1050,
			ContributionNb: 22,
		}
		database.EXPECT().GetMilestoneContribution(mock.Anything, 12, 70).Return(expRes, nil)
		h := server.NewHandler(database)
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "12",
			},
			{
				Key:   "milestone",
				Value: "70",
			},
		}
		
		h.ContributionMilestones(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		got := w.Body.String()
		exp, err := json.Marshal(expRes)
		assert.Nil(t, err)
		assert.Equal(t, string(exp), got)
	})
}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		Method: "GET",
	}

	return ctx
}
