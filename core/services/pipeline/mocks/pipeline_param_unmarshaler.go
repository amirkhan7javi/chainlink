// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PipelineParamUnmarshaler is an autogenerated mock type for the PipelineParamUnmarshaler type
type PipelineParamUnmarshaler struct {
	mock.Mock
}

// UnmarshalPipelineParam provides a mock function with given fields: val
func (_m *PipelineParamUnmarshaler) UnmarshalPipelineParam(val interface{}) error {
	ret := _m.Called(val)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
