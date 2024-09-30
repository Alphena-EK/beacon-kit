// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	backend "github.com/berachain/beacon-kit/mod/node-api/backend"
	bytes "github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"

	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"

	mock "github.com/stretchr/testify/mock"
)

// Validator is an autogenerated mock type for the Validator type
type Validator[WithdrawalCredentialsT backend.WithdrawalCredentials] struct {
	mock.Mock
}

type Validator_Expecter[WithdrawalCredentialsT backend.WithdrawalCredentials] struct {
	mock *mock.Mock
}

func (_m *Validator[WithdrawalCredentialsT]) EXPECT() *Validator_Expecter[WithdrawalCredentialsT] {
	return &Validator_Expecter[WithdrawalCredentialsT]{mock: &_m.Mock}
}

// GetActivationEligibilityEpoch provides a mock function with given fields:
func (_m *Validator[WithdrawalCredentialsT]) GetActivationEligibilityEpoch() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetActivationEligibilityEpoch")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// Validator_GetActivationEligibilityEpoch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetActivationEligibilityEpoch'
type Validator_GetActivationEligibilityEpoch_Call[WithdrawalCredentialsT backend.WithdrawalCredentials] struct {
	*mock.Call
}

// GetActivationEligibilityEpoch is a helper method to define mock.On call
func (_e *Validator_Expecter[WithdrawalCredentialsT]) GetActivationEligibilityEpoch() *Validator_GetActivationEligibilityEpoch_Call[WithdrawalCredentialsT] {
	return &Validator_GetActivationEligibilityEpoch_Call[WithdrawalCredentialsT]{Call: _e.mock.On("GetActivationEligibilityEpoch")}
}

func (_c *Validator_GetActivationEligibilityEpoch_Call[WithdrawalCredentialsT]) Run(run func()) *Validator_GetActivationEligibilityEpoch_Call[WithdrawalCredentialsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Validator_GetActivationEligibilityEpoch_Call[WithdrawalCredentialsT]) Return(_a0 math.U64) *Validator_GetActivationEligibilityEpoch_Call[WithdrawalCredentialsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Validator_GetActivationEligibilityEpoch_Call[WithdrawalCredentialsT]) RunAndReturn(run func() math.U64) *Validator_GetActivationEligibilityEpoch_Call[WithdrawalCredentialsT] {
	_c.Call.Return(run)
	return _c
}

// GetActivationEpoch provides a mock function with given fields:
func (_m *Validator[WithdrawalCredentialsT]) GetActivationEpoch() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetActivationEpoch")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// Validator_GetActivationEpoch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetActivationEpoch'
type Validator_GetActivationEpoch_Call[WithdrawalCredentialsT backend.WithdrawalCredentials] struct {
	*mock.Call
}

// GetActivationEpoch is a helper method to define mock.On call
func (_e *Validator_Expecter[WithdrawalCredentialsT]) GetActivationEpoch() *Validator_GetActivationEpoch_Call[WithdrawalCredentialsT] {
	return &Validator_GetActivationEpoch_Call[WithdrawalCredentialsT]{Call: _e.mock.On("GetActivationEpoch")}
}

func (_c *Validator_GetActivationEpoch_Call[WithdrawalCredentialsT]) Run(run func()) *Validator_GetActivationEpoch_Call[WithdrawalCredentialsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Validator_GetActivationEpoch_Call[WithdrawalCredentialsT]) Return(_a0 math.U64) *Validator_GetActivationEpoch_Call[WithdrawalCredentialsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Validator_GetActivationEpoch_Call[WithdrawalCredentialsT]) RunAndReturn(run func() math.U64) *Validator_GetActivationEpoch_Call[WithdrawalCredentialsT] {
	_c.Call.Return(run)
	return _c
}

// GetEffectiveBalance provides a mock function with given fields:
func (_m *Validator[WithdrawalCredentialsT]) GetEffectiveBalance() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetEffectiveBalance")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// Validator_GetEffectiveBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEffectiveBalance'
type Validator_GetEffectiveBalance_Call[WithdrawalCredentialsT backend.WithdrawalCredentials] struct {
	*mock.Call
}

// GetEffectiveBalance is a helper method to define mock.On call
func (_e *Validator_Expecter[WithdrawalCredentialsT]) GetEffectiveBalance() *Validator_GetEffectiveBalance_Call[WithdrawalCredentialsT] {
	return &Validator_GetEffectiveBalance_Call[WithdrawalCredentialsT]{Call: _e.mock.On("GetEffectiveBalance")}
}

func (_c *Validator_GetEffectiveBalance_Call[WithdrawalCredentialsT]) Run(run func()) *Validator_GetEffectiveBalance_Call[WithdrawalCredentialsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Validator_GetEffectiveBalance_Call[WithdrawalCredentialsT]) Return(_a0 math.U64) *Validator_GetEffectiveBalance_Call[WithdrawalCredentialsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Validator_GetEffectiveBalance_Call[WithdrawalCredentialsT]) RunAndReturn(run func() math.U64) *Validator_GetEffectiveBalance_Call[WithdrawalCredentialsT] {
	_c.Call.Return(run)
	return _c
}

// GetExitEpoch provides a mock function with given fields:
func (_m *Validator[WithdrawalCredentialsT]) GetExitEpoch() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetExitEpoch")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// Validator_GetExitEpoch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExitEpoch'
type Validator_GetExitEpoch_Call[WithdrawalCredentialsT backend.WithdrawalCredentials] struct {
	*mock.Call
}

// GetExitEpoch is a helper method to define mock.On call
func (_e *Validator_Expecter[WithdrawalCredentialsT]) GetExitEpoch() *Validator_GetExitEpoch_Call[WithdrawalCredentialsT] {
	return &Validator_GetExitEpoch_Call[WithdrawalCredentialsT]{Call: _e.mock.On("GetExitEpoch")}
}

func (_c *Validator_GetExitEpoch_Call[WithdrawalCredentialsT]) Run(run func()) *Validator_GetExitEpoch_Call[WithdrawalCredentialsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Validator_GetExitEpoch_Call[WithdrawalCredentialsT]) Return(_a0 math.U64) *Validator_GetExitEpoch_Call[WithdrawalCredentialsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Validator_GetExitEpoch_Call[WithdrawalCredentialsT]) RunAndReturn(run func() math.U64) *Validator_GetExitEpoch_Call[WithdrawalCredentialsT] {
	_c.Call.Return(run)
	return _c
}

// GetPubkey provides a mock function with given fields:
func (_m *Validator[WithdrawalCredentialsT]) GetPubkey() bytes.B48 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPubkey")
	}

	var r0 bytes.B48
	if rf, ok := ret.Get(0).(func() bytes.B48); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bytes.B48)
		}
	}

	return r0
}

// Validator_GetPubkey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPubkey'
type Validator_GetPubkey_Call[WithdrawalCredentialsT backend.WithdrawalCredentials] struct {
	*mock.Call
}

// GetPubkey is a helper method to define mock.On call
func (_e *Validator_Expecter[WithdrawalCredentialsT]) GetPubkey() *Validator_GetPubkey_Call[WithdrawalCredentialsT] {
	return &Validator_GetPubkey_Call[WithdrawalCredentialsT]{Call: _e.mock.On("GetPubkey")}
}

func (_c *Validator_GetPubkey_Call[WithdrawalCredentialsT]) Run(run func()) *Validator_GetPubkey_Call[WithdrawalCredentialsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Validator_GetPubkey_Call[WithdrawalCredentialsT]) Return(_a0 bytes.B48) *Validator_GetPubkey_Call[WithdrawalCredentialsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Validator_GetPubkey_Call[WithdrawalCredentialsT]) RunAndReturn(run func() bytes.B48) *Validator_GetPubkey_Call[WithdrawalCredentialsT] {
	_c.Call.Return(run)
	return _c
}

// GetWithdrawableEpoch provides a mock function with given fields:
func (_m *Validator[WithdrawalCredentialsT]) GetWithdrawableEpoch() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetWithdrawableEpoch")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// Validator_GetWithdrawableEpoch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWithdrawableEpoch'
type Validator_GetWithdrawableEpoch_Call[WithdrawalCredentialsT backend.WithdrawalCredentials] struct {
	*mock.Call
}

// GetWithdrawableEpoch is a helper method to define mock.On call
func (_e *Validator_Expecter[WithdrawalCredentialsT]) GetWithdrawableEpoch() *Validator_GetWithdrawableEpoch_Call[WithdrawalCredentialsT] {
	return &Validator_GetWithdrawableEpoch_Call[WithdrawalCredentialsT]{Call: _e.mock.On("GetWithdrawableEpoch")}
}

func (_c *Validator_GetWithdrawableEpoch_Call[WithdrawalCredentialsT]) Run(run func()) *Validator_GetWithdrawableEpoch_Call[WithdrawalCredentialsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Validator_GetWithdrawableEpoch_Call[WithdrawalCredentialsT]) Return(_a0 math.U64) *Validator_GetWithdrawableEpoch_Call[WithdrawalCredentialsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Validator_GetWithdrawableEpoch_Call[WithdrawalCredentialsT]) RunAndReturn(run func() math.U64) *Validator_GetWithdrawableEpoch_Call[WithdrawalCredentialsT] {
	_c.Call.Return(run)
	return _c
}

// GetWithdrawalCredentials provides a mock function with given fields:
func (_m *Validator[WithdrawalCredentialsT]) GetWithdrawalCredentials() WithdrawalCredentialsT {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetWithdrawalCredentials")
	}

	var r0 WithdrawalCredentialsT
	if rf, ok := ret.Get(0).(func() WithdrawalCredentialsT); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(WithdrawalCredentialsT)
	}

	return r0
}

// Validator_GetWithdrawalCredentials_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWithdrawalCredentials'
type Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT backend.WithdrawalCredentials] struct {
	*mock.Call
}

// GetWithdrawalCredentials is a helper method to define mock.On call
func (_e *Validator_Expecter[WithdrawalCredentialsT]) GetWithdrawalCredentials() *Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT] {
	return &Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT]{Call: _e.mock.On("GetWithdrawalCredentials")}
}

func (_c *Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT]) Run(run func()) *Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT]) Return(_a0 WithdrawalCredentialsT) *Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT]) RunAndReturn(run func() WithdrawalCredentialsT) *Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT] {
	_c.Call.Return(run)
	return _c
}

// IsSlashed provides a mock function with given fields:
func (_m *Validator[WithdrawalCredentialsT]) IsSlashed() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsSlashed")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Validator_IsSlashed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSlashed'
type Validator_IsSlashed_Call[WithdrawalCredentialsT backend.WithdrawalCredentials] struct {
	*mock.Call
}

// IsSlashed is a helper method to define mock.On call
func (_e *Validator_Expecter[WithdrawalCredentialsT]) IsSlashed() *Validator_IsSlashed_Call[WithdrawalCredentialsT] {
	return &Validator_IsSlashed_Call[WithdrawalCredentialsT]{Call: _e.mock.On("IsSlashed")}
}

func (_c *Validator_IsSlashed_Call[WithdrawalCredentialsT]) Run(run func()) *Validator_IsSlashed_Call[WithdrawalCredentialsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Validator_IsSlashed_Call[WithdrawalCredentialsT]) Return(_a0 bool) *Validator_IsSlashed_Call[WithdrawalCredentialsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Validator_IsSlashed_Call[WithdrawalCredentialsT]) RunAndReturn(run func() bool) *Validator_IsSlashed_Call[WithdrawalCredentialsT] {
	_c.Call.Return(run)
	return _c
}

// NewValidator creates a new instance of Validator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewValidator[WithdrawalCredentialsT backend.WithdrawalCredentials](t interface {
	mock.TestingT
	Cleanup(func())
}) *Validator[WithdrawalCredentialsT] {
	mock := &Validator[WithdrawalCredentialsT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
