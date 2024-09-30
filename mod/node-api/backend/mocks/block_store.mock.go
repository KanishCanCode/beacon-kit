// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	common "github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"

	mock "github.com/stretchr/testify/mock"
)

// BlockStore is an autogenerated mock type for the BlockStore type
type BlockStore[BeaconBlockT any] struct {
	mock.Mock
}

type BlockStore_Expecter[BeaconBlockT any] struct {
	mock *mock.Mock
}

func (_m *BlockStore[BeaconBlockT]) EXPECT() *BlockStore_Expecter[BeaconBlockT] {
	return &BlockStore_Expecter[BeaconBlockT]{mock: &_m.Mock}
}

// GetSlotByBlockRoot provides a mock function with given fields: root
func (_m *BlockStore[BeaconBlockT]) GetSlotByBlockRoot(root common.Root) (math.U64, error) {
	ret := _m.Called(root)

	if len(ret) == 0 {
		panic("no return value specified for GetSlotByBlockRoot")
	}

	var r0 math.U64
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Root) (math.U64, error)); ok {
		return rf(root)
	}
	if rf, ok := ret.Get(0).(func(common.Root) math.U64); ok {
		r0 = rf(root)
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	if rf, ok := ret.Get(1).(func(common.Root) error); ok {
		r1 = rf(root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockStore_GetSlotByBlockRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSlotByBlockRoot'
type BlockStore_GetSlotByBlockRoot_Call[BeaconBlockT any] struct {
	*mock.Call
}

// GetSlotByBlockRoot is a helper method to define mock.On call
//   - root common.Root
func (_e *BlockStore_Expecter[BeaconBlockT]) GetSlotByBlockRoot(root interface{}) *BlockStore_GetSlotByBlockRoot_Call[BeaconBlockT] {
	return &BlockStore_GetSlotByBlockRoot_Call[BeaconBlockT]{Call: _e.mock.On("GetSlotByBlockRoot", root)}
}

func (_c *BlockStore_GetSlotByBlockRoot_Call[BeaconBlockT]) Run(run func(root common.Root)) *BlockStore_GetSlotByBlockRoot_Call[BeaconBlockT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Root))
	})
	return _c
}

func (_c *BlockStore_GetSlotByBlockRoot_Call[BeaconBlockT]) Return(_a0 math.U64, _a1 error) *BlockStore_GetSlotByBlockRoot_Call[BeaconBlockT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlockStore_GetSlotByBlockRoot_Call[BeaconBlockT]) RunAndReturn(run func(common.Root) (math.U64, error)) *BlockStore_GetSlotByBlockRoot_Call[BeaconBlockT] {
	_c.Call.Return(run)
	return _c
}

// GetSlotByExecutionNumber provides a mock function with given fields: executionNumber
func (_m *BlockStore[BeaconBlockT]) GetSlotByExecutionNumber(executionNumber math.U64) (math.U64, error) {
	ret := _m.Called(executionNumber)

	if len(ret) == 0 {
		panic("no return value specified for GetSlotByExecutionNumber")
	}

	var r0 math.U64
	var r1 error
	if rf, ok := ret.Get(0).(func(math.U64) (math.U64, error)); ok {
		return rf(executionNumber)
	}
	if rf, ok := ret.Get(0).(func(math.U64) math.U64); ok {
		r0 = rf(executionNumber)
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	if rf, ok := ret.Get(1).(func(math.U64) error); ok {
		r1 = rf(executionNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockStore_GetSlotByExecutionNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSlotByExecutionNumber'
type BlockStore_GetSlotByExecutionNumber_Call[BeaconBlockT any] struct {
	*mock.Call
}

// GetSlotByExecutionNumber is a helper method to define mock.On call
//   - executionNumber math.U64
func (_e *BlockStore_Expecter[BeaconBlockT]) GetSlotByExecutionNumber(executionNumber interface{}) *BlockStore_GetSlotByExecutionNumber_Call[BeaconBlockT] {
	return &BlockStore_GetSlotByExecutionNumber_Call[BeaconBlockT]{Call: _e.mock.On("GetSlotByExecutionNumber", executionNumber)}
}

func (_c *BlockStore_GetSlotByExecutionNumber_Call[BeaconBlockT]) Run(run func(executionNumber math.U64)) *BlockStore_GetSlotByExecutionNumber_Call[BeaconBlockT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(math.U64))
	})
	return _c
}

func (_c *BlockStore_GetSlotByExecutionNumber_Call[BeaconBlockT]) Return(_a0 math.U64, _a1 error) *BlockStore_GetSlotByExecutionNumber_Call[BeaconBlockT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlockStore_GetSlotByExecutionNumber_Call[BeaconBlockT]) RunAndReturn(run func(math.U64) (math.U64, error)) *BlockStore_GetSlotByExecutionNumber_Call[BeaconBlockT] {
	_c.Call.Return(run)
	return _c
}

// GetSlotByStateRoot provides a mock function with given fields: root
func (_m *BlockStore[BeaconBlockT]) GetSlotByStateRoot(root common.Root) (math.U64, error) {
	ret := _m.Called(root)

	if len(ret) == 0 {
		panic("no return value specified for GetSlotByStateRoot")
	}

	var r0 math.U64
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Root) (math.U64, error)); ok {
		return rf(root)
	}
	if rf, ok := ret.Get(0).(func(common.Root) math.U64); ok {
		r0 = rf(root)
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	if rf, ok := ret.Get(1).(func(common.Root) error); ok {
		r1 = rf(root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockStore_GetSlotByStateRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSlotByStateRoot'
type BlockStore_GetSlotByStateRoot_Call[BeaconBlockT any] struct {
	*mock.Call
}

// GetSlotByStateRoot is a helper method to define mock.On call
//   - root common.Root
func (_e *BlockStore_Expecter[BeaconBlockT]) GetSlotByStateRoot(root interface{}) *BlockStore_GetSlotByStateRoot_Call[BeaconBlockT] {
	return &BlockStore_GetSlotByStateRoot_Call[BeaconBlockT]{Call: _e.mock.On("GetSlotByStateRoot", root)}
}

func (_c *BlockStore_GetSlotByStateRoot_Call[BeaconBlockT]) Run(run func(root common.Root)) *BlockStore_GetSlotByStateRoot_Call[BeaconBlockT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Root))
	})
	return _c
}

func (_c *BlockStore_GetSlotByStateRoot_Call[BeaconBlockT]) Return(_a0 math.U64, _a1 error) *BlockStore_GetSlotByStateRoot_Call[BeaconBlockT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlockStore_GetSlotByStateRoot_Call[BeaconBlockT]) RunAndReturn(run func(common.Root) (math.U64, error)) *BlockStore_GetSlotByStateRoot_Call[BeaconBlockT] {
	_c.Call.Return(run)
	return _c
}

// NewBlockStore creates a new instance of BlockStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlockStore[BeaconBlockT any](t interface {
	mock.TestingT
	Cleanup(func())
}) *BlockStore[BeaconBlockT] {
	mock := &BlockStore[BeaconBlockT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
