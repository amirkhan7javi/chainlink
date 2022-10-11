// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"

	pg "github.com/smartcontractkit/chainlink/core/services/pg"

	pipeline "github.com/smartcontractkit/chainlink/core/services/pipeline"

	time "time"

	uuid "github.com/satori/go.uuid"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// CreateRun provides a mock function with given fields: run, qopts
func (_m *ORM) CreateRun(run *pipeline.Run, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, run)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*pipeline.Run, ...pg.QOpt) error); ok {
		r0 = rf(run, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateSpec provides a mock function with given fields: _a0, maxTaskTimeout, qopts
func (_m *ORM) CreateSpec(_a0 pipeline.Pipeline, maxTaskTimeout models.Interval, qopts ...pg.QOpt) (int32, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, maxTaskTimeout)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int32
	if rf, ok := ret.Get(0).(func(pipeline.Pipeline, models.Interval, ...pg.QOpt) int32); ok {
		r0 = rf(_a0, maxTaskTimeout, qopts...)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pipeline.Pipeline, models.Interval, ...pg.QOpt) error); ok {
		r1 = rf(_a0, maxTaskTimeout, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRun provides a mock function with given fields: id
func (_m *ORM) DeleteRun(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRunsOlderThan provides a mock function with given fields: _a0, _a1
func (_m *ORM) DeleteRunsOlderThan(_a0 context.Context, _a1 time.Duration) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindRun provides a mock function with given fields: id
func (_m *ORM) FindRun(id int64) (pipeline.Run, error) {
	ret := _m.Called(id)

	var r0 pipeline.Run
	if rf, ok := ret.Get(0).(func(int64) pipeline.Run); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(pipeline.Run)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllRuns provides a mock function with given fields:
func (_m *ORM) GetAllRuns() ([]pipeline.Run, error) {
	ret := _m.Called()

	var r0 []pipeline.Run
	if rf, ok := ret.Get(0).(func() []pipeline.Run); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pipeline.Run)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQ provides a mock function with given fields:
func (_m *ORM) GetQ() pg.Q {
	ret := _m.Called()

	var r0 pg.Q
	if rf, ok := ret.Get(0).(func() pg.Q); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pg.Q)
	}

	return r0
}

// GetUnfinishedRuns provides a mock function with given fields: _a0, _a1, _a2
func (_m *ORM) GetUnfinishedRuns(_a0 context.Context, _a1 time.Time, _a2 func(pipeline.Run) error) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, func(pipeline.Run) error) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertFinishedRun provides a mock function with given fields: run, saveSuccessfulTaskRuns, qopts
func (_m *ORM) InsertFinishedRun(run *pipeline.Run, saveSuccessfulTaskRuns bool, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, run, saveSuccessfulTaskRuns)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*pipeline.Run, bool, ...pg.QOpt) error); ok {
		r0 = rf(run, saveSuccessfulTaskRuns, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertFinishedRuns provides a mock function with given fields: run, saveSuccessfulTaskRuns, qopts
func (_m *ORM) InsertFinishedRuns(run []*pipeline.Run, saveSuccessfulTaskRuns bool, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, run, saveSuccessfulTaskRuns)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*pipeline.Run, bool, ...pg.QOpt) error); ok {
		r0 = rf(run, saveSuccessfulTaskRuns, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertRun provides a mock function with given fields: run, qopts
func (_m *ORM) InsertRun(run *pipeline.Run, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, run)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*pipeline.Run, ...pg.QOpt) error); ok {
		r0 = rf(run, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreRun provides a mock function with given fields: run, qopts
func (_m *ORM) StoreRun(run *pipeline.Run, qopts ...pg.QOpt) (bool, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, run)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*pipeline.Run, ...pg.QOpt) bool); ok {
		r0 = rf(run, qopts...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pipeline.Run, ...pg.QOpt) error); ok {
		r1 = rf(run, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTaskRunResult provides a mock function with given fields: taskID, result
func (_m *ORM) UpdateTaskRunResult(taskID uuid.UUID, result pipeline.Result) (pipeline.Run, bool, error) {
	ret := _m.Called(taskID, result)

	var r0 pipeline.Run
	if rf, ok := ret.Get(0).(func(uuid.UUID, pipeline.Result) pipeline.Run); ok {
		r0 = rf(taskID, result)
	} else {
		r0 = ret.Get(0).(pipeline.Run)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(uuid.UUID, pipeline.Result) bool); ok {
		r1 = rf(taskID, result)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uuid.UUID, pipeline.Result) error); ok {
		r2 = rf(taskID, result)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewORM interface {
	mock.TestingT
	Cleanup(func())
}

// NewORM creates a new instance of ORM. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewORM(t mockConstructorTestingTNewORM) *ORM {
	mock := &ORM{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
