// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	s4 "github.com/smartcontractkit/chainlink/v2/core/services/s4"
	mock "github.com/stretchr/testify/mock"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// Constraints provides a mock function with given fields:
func (_m *Storage) Constraints() s4.Constraints {
	ret := _m.Called()

	var r0 s4.Constraints
	if rf, ok := ret.Get(0).(func() s4.Constraints); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(s4.Constraints)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, key
func (_m *Storage) Get(ctx context.Context, key *s4.Key) (*s4.Record, *s4.Metadata, error) {
	ret := _m.Called(ctx, key)

	var r0 *s4.Record
	var r1 *s4.Metadata
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *s4.Key) (*s4.Record, *s4.Metadata, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s4.Key) *s4.Record); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s4.Record)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s4.Key) *s4.Metadata); ok {
		r1 = rf(ctx, key)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*s4.Metadata)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *s4.Key) error); ok {
		r2 = rf(ctx, key)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Put provides a mock function with given fields: ctx, key, record, signature
func (_m *Storage) Put(ctx context.Context, key *s4.Key, record *s4.Record, signature []byte) error {
	ret := _m.Called(ctx, key, record, signature)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *s4.Key, *s4.Record, []byte) error); ok {
		r0 = rf(ctx, key, record, signature)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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
