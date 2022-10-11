// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	solana "github.com/smartcontractkit/chainlink-solana/pkg/solana"
)

// ChainSet is an autogenerated mock type for the ChainSet type
type ChainSet struct {
	mock.Mock
}

// Chain provides a mock function with given fields: ctx, id
func (_m *ChainSet) Chain(ctx context.Context, id string) (solana.Chain, error) {
	ret := _m.Called(ctx, id)

	var r0 solana.Chain
	if rf, ok := ret.Get(0).(func(context.Context, string) solana.Chain); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(solana.Chain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *ChainSet) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Healthy provides a mock function with given fields:
func (_m *ChainSet) Healthy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *ChainSet) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *ChainSet) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewChainSet interface {
	mock.TestingT
	Cleanup(func())
}

// NewChainSet creates a new instance of ChainSet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChainSet(t mockConstructorTestingTNewChainSet) *ChainSet {
	mock := &ChainSet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
