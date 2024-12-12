// Code generated by mockery v2.49.0. DO NOT EDIT.

package mocks

import (
	context "context"

	ctypes "github.com/berachain/beacon-kit/consensus-types/types"
	engineprimitives "github.com/berachain/beacon-kit/engine-primitives/engine-primitives"
	core "github.com/berachain/beacon-kit/state-transition/core"

	mock "github.com/stretchr/testify/mock"
)

// ExecutionEngine is an autogenerated mock type for the ExecutionEngine type
type ExecutionEngine[ExecutionPayloadHeaderT any, WithdrawalsT core.Withdrawals] struct {
	mock.Mock
}

type ExecutionEngine_Expecter[ExecutionPayloadHeaderT any, WithdrawalsT core.Withdrawals] struct {
	mock *mock.Mock
}

func (_m *ExecutionEngine[ExecutionPayloadHeaderT, WithdrawalsT]) EXPECT() *ExecutionEngine_Expecter[ExecutionPayloadHeaderT, WithdrawalsT] {
	return &ExecutionEngine_Expecter[ExecutionPayloadHeaderT, WithdrawalsT]{mock: &_m.Mock}
}

// VerifyAndNotifyNewPayload provides a mock function with given fields: ctx, req
func (_m *ExecutionEngine[ExecutionPayloadHeaderT, WithdrawalsT]) VerifyAndNotifyNewPayload(ctx context.Context, req *engineprimitives.NewPayloadRequest[*ctypes.ExecutionPayload, WithdrawalsT]) error {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for VerifyAndNotifyNewPayload")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *engineprimitives.NewPayloadRequest[*ctypes.ExecutionPayload,WithdrawalsT]) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExecutionEngine_VerifyAndNotifyNewPayload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyAndNotifyNewPayload'
type ExecutionEngine_VerifyAndNotifyNewPayload_Call[ExecutionPayloadHeaderT any, WithdrawalsT core.Withdrawals] struct {
	*mock.Call
}

// VerifyAndNotifyNewPayload is a helper method to define mock.On call
//   - ctx context.Context
//   - req *engineprimitives.NewPayloadRequest[ExecutionPayloadT,WithdrawalsT]
func (_e *ExecutionEngine_Expecter[ExecutionPayloadHeaderT, WithdrawalsT]) VerifyAndNotifyNewPayload(ctx interface{}, req interface{}) *ExecutionEngine_VerifyAndNotifyNewPayload_Call[ExecutionPayloadHeaderT, WithdrawalsT] {
	return &ExecutionEngine_VerifyAndNotifyNewPayload_Call[ExecutionPayloadHeaderT, WithdrawalsT]{Call: _e.mock.On("VerifyAndNotifyNewPayload", ctx, req)}
}

func (_c *ExecutionEngine_VerifyAndNotifyNewPayload_Call[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT]) Run(run func(ctx context.Context, req *engineprimitives.NewPayloadRequest[*ctypes.ExecutionPayload, WithdrawalsT])) *ExecutionEngine_VerifyAndNotifyNewPayload_Call[ExecutionPayloadHeaderT, WithdrawalsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*engineprimitives.NewPayloadRequest[ExecutionPayloadT, WithdrawalsT]))
	})
	return _c
}

func (_c *ExecutionEngine_VerifyAndNotifyNewPayload_Call[ExecutionPayloadHeaderT, WithdrawalsT]) Return(_a0 error) *ExecutionEngine_VerifyAndNotifyNewPayload_Call[ExecutionPayloadHeaderT, WithdrawalsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ExecutionEngine_VerifyAndNotifyNewPayload_Call[ExecutionPayloadHeaderT, WithdrawalsT]) RunAndReturn(run func(context.Context, *engineprimitives.NewPayloadRequest[WithdrawalsT]) error) *ExecutionEngine_VerifyAndNotifyNewPayload_Call[ExecutionPayloadHeaderT, WithdrawalsT] {
	_c.Call.Return(run)
	return _c
}

// NewExecutionEngine creates a new instance of ExecutionEngine. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExecutionEngine[ExecutionPayloadHeaderT any, WithdrawalsT core.Withdrawals](t interface {
	mock.TestingT
	Cleanup(func())
}) *ExecutionEngine[ExecutionPayloadHeaderT, WithdrawalsT] {
	mock := &ExecutionEngine[ExecutionPayloadHeaderT, WithdrawalsT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
