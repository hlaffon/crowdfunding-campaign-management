package server

import (
	"log"
	"net/http"
	"strconv"
	"crowdfunding-campaign-management/model"

	"crowdfunding-campaign-management/internal/store"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db store.Storage
}

func NewHandler(db store.Storage) *Handler {
	return &Handler{db: db}
}

// @Summary Return the amount per day
// @Description Get the amount raised per day
// @Tags project
// @Accept */*
// @Param        id  path      int  true  "id"
// @Produce json
// @Success 200 {object} model.AmountPerDay
// @Failure		400	{object}	model.HTTPError
// @Failure		404	{object}	model.HTTPError
// @Failure		500	{object}	model.HTTPError
// @Router			/projects/{id}/amountPerDay [get]
func (h *Handler) AmountPerDay(ctx *gin.Context) {
	id := ctx.Param("id")
	projectId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("invalid project id: %v\n", err)
		ctx.JSON(http.StatusBadRequest, wrapErrorMessage(http.StatusBadRequest, "invalid project id"))
		return
	}
	res, err := h.db.GetTotalAmountPerDay(ctx, int(projectId))
	if err != nil {
		log.Printf("error requesting data : %v\n", err)
		ctx.JSON(http.StatusInternalServerError, wrapErrorMessage(http.StatusInternalServerError, "error requesting data"))
		return
	}
	if res == nil {
		log.Printf("no data found for project %d", projectId)
		ctx.JSON(http.StatusNotFound, wrapErrorMessage(http.StatusNotFound, "no data found"))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// @Summary Return the contributions number per day
// @Description Get the contributions number per day
// @Tags project
// @Accept */*
// @Param        id  path      int  true  "id"
// @Produce json
// @Success 200 {object} model.NumberPerDay
// @Failure		400	{object}	model.HTTPError
// @Failure		404	{object}	model.HTTPError
// @Failure		500	{object}	model.HTTPError
// @Router			/projects/{id}/contributionsPerDay [get]
func (h *Handler) ContributionsPerDay(ctx *gin.Context) {
	id := ctx.Param("id")
	projectId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("invalid project id: %v\n", err)
		ctx.JSON(http.StatusBadRequest, wrapErrorMessage(http.StatusBadRequest, "invalid project id"))
		return
	}
	res, err := h.db.GetNewContributionsPerDay(ctx, int(projectId))
	if err != nil {
		log.Printf("error requesting data : %v\n", err)
		ctx.JSON(http.StatusInternalServerError, wrapErrorMessage(http.StatusInternalServerError, "error requesting data"))
		return
	}
	if res == nil {
		log.Printf("no data found for project %d", projectId)
		ctx.JSON(http.StatusNotFound, wrapErrorMessage(http.StatusNotFound, "no data found"))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// @Summary Return the contributors number per day
// @Description Get the distinct contributors per day
// @Tags project
// @Accept */*
// @Param        id  path      int  true  "id"
// @Produce json
// @Success 200 {object} model.NumberPerDay
// @Failure		400	{object}	model.HTTPError
// @Failure		404	{object}	model.HTTPError
// @Failure		500	{object}	model.HTTPError
// @Router			/projects/{id}/contributorsPerDay [get]
func (h *Handler) ContributorsPerDay(ctx *gin.Context) {
	id := ctx.Param("id")
	projectId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("invalid project id: %v\n", err)
		ctx.JSON(http.StatusBadRequest, wrapErrorMessage(http.StatusBadRequest, "invalid project id"))
		return
	}
	res, err := h.db.GetNewContributorsPerDay(ctx, int(projectId))
	if err != nil {
		log.Printf("error requesting data : %v\n", err)
		ctx.JSON(http.StatusInternalServerError, wrapErrorMessage(http.StatusInternalServerError, "error requesting data"))
		return
	}
	if res == nil {
		log.Printf("no data found for project %d", projectId)
		ctx.JSON(http.StatusNotFound, wrapErrorMessage(http.StatusNotFound, "no data found"))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// @Summary Return the cumulated percentage of goal raised per day
// @Description Get the cumulated percentage of goal raised per day
// @Tags project
// @Accept */*
// @Param        id  path      int  true  "id"
// @Produce json
// @Success 200 {object} model.AmountPerDay
// @Failure		400	{object}	model.HTTPError
// @Failure		404	{object}	model.HTTPError
// @Failure		500	{object}	model.HTTPError
// @Router			/projects/{id}/percentagePerDay [get]
func (h *Handler) PercentagePerDay(ctx *gin.Context) {
	id := ctx.Param("id")
	projectId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("invalid project id: %v\n", err)
		ctx.JSON(http.StatusBadRequest, wrapErrorMessage(http.StatusBadRequest, "invalid project id"))
		return
	}
	res, err := h.db.GetPercentageOfGoalPerDay(ctx, int(projectId))
	if err != nil {
		log.Printf("error requesting data : %v\n", err)
		ctx.JSON(http.StatusInternalServerError, wrapErrorMessage(http.StatusInternalServerError, "error requesting data"))
		return
	}
	if res == nil {
		log.Printf("no data found for project %d", projectId)
		ctx.JSON(http.StatusNotFound, wrapErrorMessage(http.StatusNotFound, "no data found"))
		return
	}

	var cumulatedPercentage []*model.AmountPerDay
	var cp float64
	for _, p := range res {
		cp += p.Total
		cumulatedPercentage = append(cumulatedPercentage, &model.AmountPerDay{Date: p.Date, Total: cp})
	}

	ctx.JSON(http.StatusOK, cumulatedPercentage)
}

// @Summary Return the number of views per day
// @Description Get the total number of visits (views) per day
// @Tags project
// @Accept */*
// @Param        id  path      int  true  "id"
// @Produce json
// @Success 200 {object} model.NumberPerDay
// @Failure		400	{object}	model.HTTPError
// @Failure		404	{object}	model.HTTPError
// @Failure		500	{object}	model.HTTPError
// @Router			/projects/{id}/visitsPerDay [get]
func (h *Handler) VisitsPerDay(ctx *gin.Context) {
	id := ctx.Param("id")
	projectId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("invalid project id: %v\n", err)
		ctx.JSON(http.StatusBadRequest, wrapErrorMessage(http.StatusBadRequest, "invalid project id"))
		return
	}
	res, err := h.db.GetVisitsPerDay(ctx, int(projectId))
	if err != nil {
		log.Printf("error requesting data : %v\n", err)
		ctx.JSON(http.StatusInternalServerError, wrapErrorMessage(http.StatusInternalServerError, "error requesting data"))
		return
	}
	if res == nil {
		log.Printf("no data found for project %d", projectId)
		ctx.JSON(http.StatusNotFound, wrapErrorMessage(http.StatusNotFound, "no data found"))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// @Summary Return the average contribution
// @Description Get the average contribution in euros for a given project
// @Tags project
// @Accept */*
// @Param        id  path      int  true  "id"
// @Produce json
// @Success 200 {float64} float64
// @Failure		400	{object}	model.HTTPError
// @Failure		404	{object}	model.HTTPError
// @Failure		500	{object}	model.HTTPError
// @Router			/projects/{id}/averageContribution [get]
func (h *Handler) AverageContribution(ctx *gin.Context) {
	id := ctx.Param("id")
	projectId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("invalid project id: %v\n", err)
		ctx.JSON(http.StatusBadRequest, wrapErrorMessage(http.StatusBadRequest, "invalid project id"))
		return
	}
	res, err := h.db.GetAverageContribution(ctx, int(projectId))
	if err != nil {
		log.Printf("error requesting data : %v\n", err)
		ctx.JSON(http.StatusInternalServerError, wrapErrorMessage(http.StatusInternalServerError, "error requesting data"))
		return
	}
	if res == nil {
		log.Printf("no data found for project %d", projectId)
		ctx.JSON(http.StatusNotFound, wrapErrorMessage(http.StatusNotFound, "no data found"))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// @Summary Return the conversion rate per day
// @Description Get the conversion rate per day (number of contributions divided by number of views)
// @Tags project
// @Accept */*
// @Param        id  path      int  true  "id"
// @Produce json
// @Success 200 {object} model.AmountPerDay
// @Failure		400	{object}	model.HTTPError
// @Failure		404	{object}	model.HTTPError
// @Failure		500	{object}	model.HTTPError
// @Router			/projects/{id}/conversionRatePerDay [get]
func (h *Handler) ConversionRatePerDay(ctx *gin.Context) {
	id := ctx.Param("id")
	projectId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("invalid project id: %v\n", err)
		ctx.JSON(http.StatusBadRequest, wrapErrorMessage(http.StatusBadRequest, "invalid project id"))
		return
	}
	res, err := h.db.GetConversionRatePerDay(ctx, int(projectId))
	if err != nil {
		log.Printf("error requesting data : %v\n", err)
		ctx.JSON(http.StatusInternalServerError, wrapErrorMessage(http.StatusInternalServerError, "error requesting data"))
		return
	}
	if res == nil {
		log.Printf("no data found for project %d", projectId)
		ctx.JSON(http.StatusNotFound, wrapErrorMessage(http.StatusNotFound, "no data found"))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// @Summary Return the contribution which reached a milestone
// @Description Get the contribution which reached a milestone (aka percentage of goal)
// @Tags project
// @Accept */*
// @Param        id  path      int  true  "id"
// @Param        milestone  path      int  true  "milestone"
// @Produce json
// @Success 200 {object} model.MilestoneContribution
// @Failure		400	{object}	model.HTTPError
// @Failure		404	{object}	model.HTTPError
// @Failure		500	{object}	model.HTTPError
// @Router			/projects/{id}/contributionMilestones/{milestone} [get]
func (h *Handler) ContributionMilestones(ctx *gin.Context) {
	id := ctx.Param("id")
	mile := ctx.Param("milestone")
	projectId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("invalid project id: %v\n", err)
		ctx.JSON(http.StatusBadRequest, wrapErrorMessage(http.StatusBadRequest, "invalid project id"))
		return
	}
	milestone, err := strconv.ParseInt(mile, 10, 64)
	if err != nil {
		log.Printf("invalid milestone: %v\n", err)
		ctx.JSON(http.StatusBadRequest, wrapErrorMessage(http.StatusBadRequest, "invalid milestone"))
		return
	}
	res, err := h.db.GetMilestoneContribution(ctx, int(projectId), int(milestone))
	if err != nil {
		log.Printf("error requesting data : %v\n", err)
		ctx.JSON(http.StatusInternalServerError, wrapErrorMessage(http.StatusInternalServerError, "error requesting data"))
		return
	}
	if res == nil {
		log.Printf("no data found for project %d", projectId)
		ctx.JSON(http.StatusNotFound, wrapErrorMessage(http.StatusNotFound, "no data found"))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func wrapErrorMessage(code int, msg string) *model.HTTPError {
	return &model.HTTPError{
		Code:    code,
		Message: msg,
	}
}
