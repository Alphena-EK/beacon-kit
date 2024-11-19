// Code generated by mockery v2.48.0. DO NOT EDIT.

package mocks

import (
	common "github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"

	mock "github.com/stretchr/testify/mock"
)

// BeaconBlockHeader is an autogenerated mock type for the BeaconBlockHeader type
type BeaconBlockHeader[BeaconBlockHeaderT any] struct {
	mock.Mock
}

type BeaconBlockHeader_Expecter[BeaconBlockHeaderT any] struct {
	mock *mock.Mock
}

func (_m *BeaconBlockHeader[BeaconBlockHeaderT]) EXPECT() *BeaconBlockHeader_Expecter[BeaconBlockHeaderT] {
	return &BeaconBlockHeader_Expecter[BeaconBlockHeaderT]{mock: &_m.Mock}
}

// GetBodyRoot provides a mock function with given fields:
func (_m *BeaconBlockHeader[BeaconBlockHeaderT]) GetBodyRoot() common.Root {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBodyRoot")
	}

	var r0 common.Root
	if rf, ok := ret.Get(0).(func() common.Root); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Root)
		}
	}

	return r0
}

// BeaconBlockHeader_GetBodyRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBodyRoot'
type BeaconBlockHeader_GetBodyRoot_Call[BeaconBlockHeaderT any] struct {
	*mock.Call
}

// GetBodyRoot is a helper method to define mock.On call
func (_e *BeaconBlockHeader_Expecter[BeaconBlockHeaderT]) GetBodyRoot() *BeaconBlockHeader_GetBodyRoot_Call[BeaconBlockHeaderT] {
	return &BeaconBlockHeader_GetBodyRoot_Call[BeaconBlockHeaderT]{Call: _e.mock.On("GetBodyRoot")}
}

func (_c *BeaconBlockHeader_GetBodyRoot_Call[BeaconBlockHeaderT]) Run(run func()) *BeaconBlockHeader_GetBodyRoot_Call[BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlockHeader_GetBodyRoot_Call[BeaconBlockHeaderT]) Return(_a0 common.Root) *BeaconBlockHeader_GetBodyRoot_Call[BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_GetBodyRoot_Call[BeaconBlockHeaderT]) RunAndReturn(run func() common.Root) *BeaconBlockHeader_GetBodyRoot_Call[BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// GetParentBlockRoot provides a mock function with given fields:
func (_m *BeaconBlockHeader[BeaconBlockHeaderT]) GetParentBlockRoot() common.Root {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetParentBlockRoot")
	}

	var r0 common.Root
	if rf, ok := ret.Get(0).(func() common.Root); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Root)
		}
	}

	return r0
}

// BeaconBlockHeader_GetParentBlockRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetParentBlockRoot'
type BeaconBlockHeader_GetParentBlockRoot_Call[BeaconBlockHeaderT any] struct {
	*mock.Call
}

// GetParentBlockRoot is a helper method to define mock.On call
func (_e *BeaconBlockHeader_Expecter[BeaconBlockHeaderT]) GetParentBlockRoot() *BeaconBlockHeader_GetParentBlockRoot_Call[BeaconBlockHeaderT] {
	return &BeaconBlockHeader_GetParentBlockRoot_Call[BeaconBlockHeaderT]{Call: _e.mock.On("GetParentBlockRoot")}
}

func (_c *BeaconBlockHeader_GetParentBlockRoot_Call[BeaconBlockHeaderT]) Run(run func()) *BeaconBlockHeader_GetParentBlockRoot_Call[BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlockHeader_GetParentBlockRoot_Call[BeaconBlockHeaderT]) Return(_a0 common.Root) *BeaconBlockHeader_GetParentBlockRoot_Call[BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_GetParentBlockRoot_Call[BeaconBlockHeaderT]) RunAndReturn(run func() common.Root) *BeaconBlockHeader_GetParentBlockRoot_Call[BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// GetProposerIndex provides a mock function with given fields:
func (_m *BeaconBlockHeader[BeaconBlockHeaderT]) GetProposerIndex() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetProposerIndex")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// BeaconBlockHeader_GetProposerIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProposerIndex'
type BeaconBlockHeader_GetProposerIndex_Call[BeaconBlockHeaderT any] struct {
	*mock.Call
}

// GetProposerIndex is a helper method to define mock.On call
func (_e *BeaconBlockHeader_Expecter[BeaconBlockHeaderT]) GetProposerIndex() *BeaconBlockHeader_GetProposerIndex_Call[BeaconBlockHeaderT] {
	return &BeaconBlockHeader_GetProposerIndex_Call[BeaconBlockHeaderT]{Call: _e.mock.On("GetProposerIndex")}
}

func (_c *BeaconBlockHeader_GetProposerIndex_Call[BeaconBlockHeaderT]) Run(run func()) *BeaconBlockHeader_GetProposerIndex_Call[BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlockHeader_GetProposerIndex_Call[BeaconBlockHeaderT]) Return(_a0 math.U64) *BeaconBlockHeader_GetProposerIndex_Call[BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_GetProposerIndex_Call[BeaconBlockHeaderT]) RunAndReturn(run func() math.U64) *BeaconBlockHeader_GetProposerIndex_Call[BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// GetSlot provides a mock function with given fields:
func (_m *BeaconBlockHeader[BeaconBlockHeaderT]) GetSlot() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSlot")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// BeaconBlockHeader_GetSlot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSlot'
type BeaconBlockHeader_GetSlot_Call[BeaconBlockHeaderT any] struct {
	*mock.Call
}

// GetSlot is a helper method to define mock.On call
func (_e *BeaconBlockHeader_Expecter[BeaconBlockHeaderT]) GetSlot() *BeaconBlockHeader_GetSlot_Call[BeaconBlockHeaderT] {
	return &BeaconBlockHeader_GetSlot_Call[BeaconBlockHeaderT]{Call: _e.mock.On("GetSlot")}
}

func (_c *BeaconBlockHeader_GetSlot_Call[BeaconBlockHeaderT]) Run(run func()) *BeaconBlockHeader_GetSlot_Call[BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlockHeader_GetSlot_Call[BeaconBlockHeaderT]) Return(_a0 math.U64) *BeaconBlockHeader_GetSlot_Call[BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_GetSlot_Call[BeaconBlockHeaderT]) RunAndReturn(run func() math.U64) *BeaconBlockHeader_GetSlot_Call[BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// GetStateRoot provides a mock function with given fields:
func (_m *BeaconBlockHeader[BeaconBlockHeaderT]) GetStateRoot() common.Root {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStateRoot")
	}

	var r0 common.Root
	if rf, ok := ret.Get(0).(func() common.Root); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Root)
		}
	}

	return r0
}

// BeaconBlockHeader_GetStateRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStateRoot'
type BeaconBlockHeader_GetStateRoot_Call[BeaconBlockHeaderT any] struct {
	*mock.Call
}

// GetStateRoot is a helper method to define mock.On call
func (_e *BeaconBlockHeader_Expecter[BeaconBlockHeaderT]) GetStateRoot() *BeaconBlockHeader_GetStateRoot_Call[BeaconBlockHeaderT] {
	return &BeaconBlockHeader_GetStateRoot_Call[BeaconBlockHeaderT]{Call: _e.mock.On("GetStateRoot")}
}

func (_c *BeaconBlockHeader_GetStateRoot_Call[BeaconBlockHeaderT]) Run(run func()) *BeaconBlockHeader_GetStateRoot_Call[BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlockHeader_GetStateRoot_Call[BeaconBlockHeaderT]) Return(_a0 common.Root) *BeaconBlockHeader_GetStateRoot_Call[BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_GetStateRoot_Call[BeaconBlockHeaderT]) RunAndReturn(run func() common.Root) *BeaconBlockHeader_GetStateRoot_Call[BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// HashTreeRoot provides a mock function with given fields:
func (_m *BeaconBlockHeader[BeaconBlockHeaderT]) HashTreeRoot() common.Root {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HashTreeRoot")
	}

	var r0 common.Root
	if rf, ok := ret.Get(0).(func() common.Root); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Root)
		}
	}

	return r0
}

// BeaconBlockHeader_HashTreeRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HashTreeRoot'
type BeaconBlockHeader_HashTreeRoot_Call[BeaconBlockHeaderT any] struct {
	*mock.Call
}

// HashTreeRoot is a helper method to define mock.On call
func (_e *BeaconBlockHeader_Expecter[BeaconBlockHeaderT]) HashTreeRoot() *BeaconBlockHeader_HashTreeRoot_Call[BeaconBlockHeaderT] {
	return &BeaconBlockHeader_HashTreeRoot_Call[BeaconBlockHeaderT]{Call: _e.mock.On("HashTreeRoot")}
}

func (_c *BeaconBlockHeader_HashTreeRoot_Call[BeaconBlockHeaderT]) Run(run func()) *BeaconBlockHeader_HashTreeRoot_Call[BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlockHeader_HashTreeRoot_Call[BeaconBlockHeaderT]) Return(_a0 common.Root) *BeaconBlockHeader_HashTreeRoot_Call[BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_HashTreeRoot_Call[BeaconBlockHeaderT]) RunAndReturn(run func() common.Root) *BeaconBlockHeader_HashTreeRoot_Call[BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// MarshalSSZ provides a mock function with given fields:
func (_m *BeaconBlockHeader[BeaconBlockHeaderT]) MarshalSSZ() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MarshalSSZ")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeaconBlockHeader_MarshalSSZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarshalSSZ'
type BeaconBlockHeader_MarshalSSZ_Call[BeaconBlockHeaderT any] struct {
	*mock.Call
}

// MarshalSSZ is a helper method to define mock.On call
func (_e *BeaconBlockHeader_Expecter[BeaconBlockHeaderT]) MarshalSSZ() *BeaconBlockHeader_MarshalSSZ_Call[BeaconBlockHeaderT] {
	return &BeaconBlockHeader_MarshalSSZ_Call[BeaconBlockHeaderT]{Call: _e.mock.On("MarshalSSZ")}
}

func (_c *BeaconBlockHeader_MarshalSSZ_Call[BeaconBlockHeaderT]) Run(run func()) *BeaconBlockHeader_MarshalSSZ_Call[BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BeaconBlockHeader_MarshalSSZ_Call[BeaconBlockHeaderT]) Return(_a0 []byte, _a1 error) *BeaconBlockHeader_MarshalSSZ_Call[BeaconBlockHeaderT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BeaconBlockHeader_MarshalSSZ_Call[BeaconBlockHeaderT]) RunAndReturn(run func() ([]byte, error)) *BeaconBlockHeader_MarshalSSZ_Call[BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// New provides a mock function with given fields: slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot
func (_m *BeaconBlockHeader[BeaconBlockHeaderT]) New(slot math.U64, proposerIndex math.U64, parentBlockRoot common.Root, stateRoot common.Root, bodyRoot common.Root) BeaconBlockHeaderT {
	ret := _m.Called(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)

	if len(ret) == 0 {
		panic("no return value specified for New")
	}

	var r0 BeaconBlockHeaderT
	if rf, ok := ret.Get(0).(func(math.U64, math.U64, common.Root, common.Root, common.Root) BeaconBlockHeaderT); ok {
		r0 = rf(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(BeaconBlockHeaderT)
		}
	}

	return r0
}

// BeaconBlockHeader_New_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'New'
type BeaconBlockHeader_New_Call[BeaconBlockHeaderT any] struct {
	*mock.Call
}

// New is a helper method to define mock.On call
//   - slot math.U64
//   - proposerIndex math.U64
//   - parentBlockRoot common.Root
//   - stateRoot common.Root
//   - bodyRoot common.Root
func (_e *BeaconBlockHeader_Expecter[BeaconBlockHeaderT]) New(slot interface{}, proposerIndex interface{}, parentBlockRoot interface{}, stateRoot interface{}, bodyRoot interface{}) *BeaconBlockHeader_New_Call[BeaconBlockHeaderT] {
	return &BeaconBlockHeader_New_Call[BeaconBlockHeaderT]{Call: _e.mock.On("New", slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)}
}

func (_c *BeaconBlockHeader_New_Call[BeaconBlockHeaderT]) Run(run func(slot math.U64, proposerIndex math.U64, parentBlockRoot common.Root, stateRoot common.Root, bodyRoot common.Root)) *BeaconBlockHeader_New_Call[BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(math.U64), args[1].(math.U64), args[2].(common.Root), args[3].(common.Root), args[4].(common.Root))
	})
	return _c
}

func (_c *BeaconBlockHeader_New_Call[BeaconBlockHeaderT]) Return(_a0 BeaconBlockHeaderT) *BeaconBlockHeader_New_Call[BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_New_Call[BeaconBlockHeaderT]) RunAndReturn(run func(math.U64, math.U64, common.Root, common.Root, common.Root) BeaconBlockHeaderT) *BeaconBlockHeader_New_Call[BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// SetStateRoot provides a mock function with given fields: _a0
func (_m *BeaconBlockHeader[BeaconBlockHeaderT]) SetStateRoot(_a0 common.Root) {
	_m.Called(_a0)
}

// BeaconBlockHeader_SetStateRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetStateRoot'
type BeaconBlockHeader_SetStateRoot_Call[BeaconBlockHeaderT any] struct {
	*mock.Call
}

// SetStateRoot is a helper method to define mock.On call
//   - _a0 common.Root
func (_e *BeaconBlockHeader_Expecter[BeaconBlockHeaderT]) SetStateRoot(_a0 interface{}) *BeaconBlockHeader_SetStateRoot_Call[BeaconBlockHeaderT] {
	return &BeaconBlockHeader_SetStateRoot_Call[BeaconBlockHeaderT]{Call: _e.mock.On("SetStateRoot", _a0)}
}

func (_c *BeaconBlockHeader_SetStateRoot_Call[BeaconBlockHeaderT]) Run(run func(_a0 common.Root)) *BeaconBlockHeader_SetStateRoot_Call[BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Root))
	})
	return _c
}

func (_c *BeaconBlockHeader_SetStateRoot_Call[BeaconBlockHeaderT]) Return() *BeaconBlockHeader_SetStateRoot_Call[BeaconBlockHeaderT] {
	_c.Call.Return()
	return _c
}

func (_c *BeaconBlockHeader_SetStateRoot_Call[BeaconBlockHeaderT]) RunAndReturn(run func(common.Root)) *BeaconBlockHeader_SetStateRoot_Call[BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// UnmarshalSSZ provides a mock function with given fields: _a0
func (_m *BeaconBlockHeader[BeaconBlockHeaderT]) UnmarshalSSZ(_a0 []byte) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UnmarshalSSZ")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BeaconBlockHeader_UnmarshalSSZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnmarshalSSZ'
type BeaconBlockHeader_UnmarshalSSZ_Call[BeaconBlockHeaderT any] struct {
	*mock.Call
}

// UnmarshalSSZ is a helper method to define mock.On call
//   - _a0 []byte
func (_e *BeaconBlockHeader_Expecter[BeaconBlockHeaderT]) UnmarshalSSZ(_a0 interface{}) *BeaconBlockHeader_UnmarshalSSZ_Call[BeaconBlockHeaderT] {
	return &BeaconBlockHeader_UnmarshalSSZ_Call[BeaconBlockHeaderT]{Call: _e.mock.On("UnmarshalSSZ", _a0)}
}

func (_c *BeaconBlockHeader_UnmarshalSSZ_Call[BeaconBlockHeaderT]) Run(run func(_a0 []byte)) *BeaconBlockHeader_UnmarshalSSZ_Call[BeaconBlockHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *BeaconBlockHeader_UnmarshalSSZ_Call[BeaconBlockHeaderT]) Return(_a0 error) *BeaconBlockHeader_UnmarshalSSZ_Call[BeaconBlockHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BeaconBlockHeader_UnmarshalSSZ_Call[BeaconBlockHeaderT]) RunAndReturn(run func([]byte) error) *BeaconBlockHeader_UnmarshalSSZ_Call[BeaconBlockHeaderT] {
	_c.Call.Return(run)
	return _c
}

// NewBeaconBlockHeader creates a new instance of BeaconBlockHeader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBeaconBlockHeader[BeaconBlockHeaderT any](t interface {
	mock.TestingT
	Cleanup(func())
}) *BeaconBlockHeader[BeaconBlockHeaderT] {
	mock := &BeaconBlockHeader[BeaconBlockHeaderT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
