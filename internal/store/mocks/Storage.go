// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	context "context"
	model "crowdfunding-campaign-management/model"

	mock "github.com/stretchr/testify/mock"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

type Storage_Expecter struct {
	mock *mock.Mock
}

func (_m *Storage) EXPECT() *Storage_Expecter {
	return &Storage_Expecter{mock: &_m.Mock}
}

// GetAverageContribution provides a mock function with given fields: ctx, projectId
func (_m *Storage) GetAverageContribution(ctx context.Context, projectId int) (*float64, error) {
	ret := _m.Called(ctx, projectId)

	var r0 *float64
	if rf, ok := ret.Get(0).(func(context.Context, int) *float64); ok {
		r0 = rf(ctx, projectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*float64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, projectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetAverageContribution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAverageContribution'
type Storage_GetAverageContribution_Call struct {
	*mock.Call
}

// GetAverageContribution is a helper method to define mock.On call
//   - ctx context.Context
//   - projectId int
func (_e *Storage_Expecter) GetAverageContribution(ctx interface{}, projectId interface{}) *Storage_GetAverageContribution_Call {
	return &Storage_GetAverageContribution_Call{Call: _e.mock.On("GetAverageContribution", ctx, projectId)}
}

func (_c *Storage_GetAverageContribution_Call) Run(run func(ctx context.Context, projectId int)) *Storage_GetAverageContribution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *Storage_GetAverageContribution_Call) Return(_a0 *float64, _a1 error) *Storage_GetAverageContribution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetConversionRatePerDay provides a mock function with given fields: ctx, projectId
func (_m *Storage) GetConversionRatePerDay(ctx context.Context, projectId int) ([]*model.AmountPerDay, error) {
	ret := _m.Called(ctx, projectId)

	var r0 []*model.AmountPerDay
	if rf, ok := ret.Get(0).(func(context.Context, int) []*model.AmountPerDay); ok {
		r0 = rf(ctx, projectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.AmountPerDay)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, projectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetConversionRatePerDay_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConversionRatePerDay'
type Storage_GetConversionRatePerDay_Call struct {
	*mock.Call
}

// GetConversionRatePerDay is a helper method to define mock.On call
//   - ctx context.Context
//   - projectId int
func (_e *Storage_Expecter) GetConversionRatePerDay(ctx interface{}, projectId interface{}) *Storage_GetConversionRatePerDay_Call {
	return &Storage_GetConversionRatePerDay_Call{Call: _e.mock.On("GetConversionRatePerDay", ctx, projectId)}
}

func (_c *Storage_GetConversionRatePerDay_Call) Run(run func(ctx context.Context, projectId int)) *Storage_GetConversionRatePerDay_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *Storage_GetConversionRatePerDay_Call) Return(_a0 []*model.AmountPerDay, _a1 error) *Storage_GetConversionRatePerDay_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetMilestoneContribution provides a mock function with given fields: ctx, projectId, milestone
func (_m *Storage) GetMilestoneContribution(ctx context.Context, projectId int, milestone int) (*model.MilestoneContribution, error) {
	ret := _m.Called(ctx, projectId, milestone)

	var r0 *model.MilestoneContribution
	if rf, ok := ret.Get(0).(func(context.Context, int, int) *model.MilestoneContribution); ok {
		r0 = rf(ctx, projectId, milestone)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MilestoneContribution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, projectId, milestone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetMilestoneContribution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMilestoneContribution'
type Storage_GetMilestoneContribution_Call struct {
	*mock.Call
}

// GetMilestoneContribution is a helper method to define mock.On call
//   - ctx context.Context
//   - projectId int
//   - milestone int
func (_e *Storage_Expecter) GetMilestoneContribution(ctx interface{}, projectId interface{}, milestone interface{}) *Storage_GetMilestoneContribution_Call {
	return &Storage_GetMilestoneContribution_Call{Call: _e.mock.On("GetMilestoneContribution", ctx, projectId, milestone)}
}

func (_c *Storage_GetMilestoneContribution_Call) Run(run func(ctx context.Context, projectId int, milestone int)) *Storage_GetMilestoneContribution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *Storage_GetMilestoneContribution_Call) Return(_a0 *model.MilestoneContribution, _a1 error) *Storage_GetMilestoneContribution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetNewContributionsPerDay provides a mock function with given fields: ctx, projectId
func (_m *Storage) GetNewContributionsPerDay(ctx context.Context, projectId int) ([]*model.NumberPerDay, error) {
	ret := _m.Called(ctx, projectId)

	var r0 []*model.NumberPerDay
	if rf, ok := ret.Get(0).(func(context.Context, int) []*model.NumberPerDay); ok {
		r0 = rf(ctx, projectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.NumberPerDay)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, projectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetNewContributionsPerDay_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNewContributionsPerDay'
type Storage_GetNewContributionsPerDay_Call struct {
	*mock.Call
}

// GetNewContributionsPerDay is a helper method to define mock.On call
//   - ctx context.Context
//   - projectId int
func (_e *Storage_Expecter) GetNewContributionsPerDay(ctx interface{}, projectId interface{}) *Storage_GetNewContributionsPerDay_Call {
	return &Storage_GetNewContributionsPerDay_Call{Call: _e.mock.On("GetNewContributionsPerDay", ctx, projectId)}
}

func (_c *Storage_GetNewContributionsPerDay_Call) Run(run func(ctx context.Context, projectId int)) *Storage_GetNewContributionsPerDay_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *Storage_GetNewContributionsPerDay_Call) Return(_a0 []*model.NumberPerDay, _a1 error) *Storage_GetNewContributionsPerDay_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetNewContributorsPerDay provides a mock function with given fields: ctx, projectId
func (_m *Storage) GetNewContributorsPerDay(ctx context.Context, projectId int) ([]*model.NumberPerDay, error) {
	ret := _m.Called(ctx, projectId)

	var r0 []*model.NumberPerDay
	if rf, ok := ret.Get(0).(func(context.Context, int) []*model.NumberPerDay); ok {
		r0 = rf(ctx, projectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.NumberPerDay)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, projectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetNewContributorsPerDay_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNewContributorsPerDay'
type Storage_GetNewContributorsPerDay_Call struct {
	*mock.Call
}

// GetNewContributorsPerDay is a helper method to define mock.On call
//   - ctx context.Context
//   - projectId int
func (_e *Storage_Expecter) GetNewContributorsPerDay(ctx interface{}, projectId interface{}) *Storage_GetNewContributorsPerDay_Call {
	return &Storage_GetNewContributorsPerDay_Call{Call: _e.mock.On("GetNewContributorsPerDay", ctx, projectId)}
}

func (_c *Storage_GetNewContributorsPerDay_Call) Run(run func(ctx context.Context, projectId int)) *Storage_GetNewContributorsPerDay_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *Storage_GetNewContributorsPerDay_Call) Return(_a0 []*model.NumberPerDay, _a1 error) *Storage_GetNewContributorsPerDay_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPercentageOfGoalPerDay provides a mock function with given fields: ctx, projectId
func (_m *Storage) GetPercentageOfGoalPerDay(ctx context.Context, projectId int) ([]*model.AmountPerDay, error) {
	ret := _m.Called(ctx, projectId)

	var r0 []*model.AmountPerDay
	if rf, ok := ret.Get(0).(func(context.Context, int) []*model.AmountPerDay); ok {
		r0 = rf(ctx, projectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.AmountPerDay)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, projectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetPercentageOfGoalPerDay_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPercentageOfGoalPerDay'
type Storage_GetPercentageOfGoalPerDay_Call struct {
	*mock.Call
}

// GetPercentageOfGoalPerDay is a helper method to define mock.On call
//   - ctx context.Context
//   - projectId int
func (_e *Storage_Expecter) GetPercentageOfGoalPerDay(ctx interface{}, projectId interface{}) *Storage_GetPercentageOfGoalPerDay_Call {
	return &Storage_GetPercentageOfGoalPerDay_Call{Call: _e.mock.On("GetPercentageOfGoalPerDay", ctx, projectId)}
}

func (_c *Storage_GetPercentageOfGoalPerDay_Call) Run(run func(ctx context.Context, projectId int)) *Storage_GetPercentageOfGoalPerDay_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *Storage_GetPercentageOfGoalPerDay_Call) Return(_a0 []*model.AmountPerDay, _a1 error) *Storage_GetPercentageOfGoalPerDay_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTotalAmountPerDay provides a mock function with given fields: ctx, projectId
func (_m *Storage) GetTotalAmountPerDay(ctx context.Context, projectId int) ([]*model.AmountPerDay, error) {
	ret := _m.Called(ctx, projectId)

	var r0 []*model.AmountPerDay
	if rf, ok := ret.Get(0).(func(context.Context, int) []*model.AmountPerDay); ok {
		r0 = rf(ctx, projectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.AmountPerDay)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, projectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetTotalAmountPerDay_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTotalAmountPerDay'
type Storage_GetTotalAmountPerDay_Call struct {
	*mock.Call
}

// GetTotalAmountPerDay is a helper method to define mock.On call
//   - ctx context.Context
//   - projectId int
func (_e *Storage_Expecter) GetTotalAmountPerDay(ctx interface{}, projectId interface{}) *Storage_GetTotalAmountPerDay_Call {
	return &Storage_GetTotalAmountPerDay_Call{Call: _e.mock.On("GetTotalAmountPerDay", ctx, projectId)}
}

func (_c *Storage_GetTotalAmountPerDay_Call) Run(run func(ctx context.Context, projectId int)) *Storage_GetTotalAmountPerDay_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *Storage_GetTotalAmountPerDay_Call) Return(_a0 []*model.AmountPerDay, _a1 error) *Storage_GetTotalAmountPerDay_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetViewsPerDay provides a mock function with given fields: ctx, projectId
func (_m *Storage) GetVisitsPerDay(ctx context.Context, projectId int) ([]*model.NumberPerDay, error) {
	ret := _m.Called(ctx, projectId)

	var r0 []*model.NumberPerDay
	if rf, ok := ret.Get(0).(func(context.Context, int) []*model.NumberPerDay); ok {
		r0 = rf(ctx, projectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.NumberPerDay)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, projectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetViewsPerDay_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVisitsPerDay'
type Storage_GetViewsPerDay_Call struct {
	*mock.Call
}

// GetViewsPerDay is a helper method to define mock.On call
//   - ctx context.Context
//   - projectId int
func (_e *Storage_Expecter) GetViewsPerDay(ctx interface{}, projectId interface{}) *Storage_GetViewsPerDay_Call {
	return &Storage_GetViewsPerDay_Call{Call: _e.mock.On("GetVisitsPerDay", ctx, projectId)}
}

func (_c *Storage_GetViewsPerDay_Call) Run(run func(ctx context.Context, projectId int)) *Storage_GetViewsPerDay_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *Storage_GetViewsPerDay_Call) Return(_a0 []*model.NumberPerDay, _a1 error) *Storage_GetViewsPerDay_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetVisitorsPerDay provides a mock function with given fields: ctx, projectId
func (_m *Storage) GetVisitorsPerDay(ctx context.Context, projectId int) ([]*model.NumberPerDay, error) {
	ret := _m.Called(ctx, projectId)

	var r0 []*model.NumberPerDay
	if rf, ok := ret.Get(0).(func(context.Context, int) []*model.NumberPerDay); ok {
		r0 = rf(ctx, projectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.NumberPerDay)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, projectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_GetVisitorsPerDay_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVisitorsPerDay'
type Storage_GetVisitorsPerDay_Call struct {
	*mock.Call
}

// GetVisitorsPerDay is a helper method to define mock.On call
//   - ctx context.Context
//   - projectId int
func (_e *Storage_Expecter) GetVisitorsPerDay(ctx interface{}, projectId interface{}) *Storage_GetVisitorsPerDay_Call {
	return &Storage_GetVisitorsPerDay_Call{Call: _e.mock.On("GetVisitorsPerDay", ctx, projectId)}
}

func (_c *Storage_GetVisitorsPerDay_Call) Run(run func(ctx context.Context, projectId int)) *Storage_GetVisitorsPerDay_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *Storage_GetVisitorsPerDay_Call) Return(_a0 []*model.NumberPerDay, _a1 error) *Storage_GetVisitorsPerDay_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorage(t mockConstructorTestingTNewStorage) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
