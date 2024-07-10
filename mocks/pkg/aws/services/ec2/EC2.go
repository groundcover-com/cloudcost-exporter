// Code generated by mockery v2.43.2. DO NOT EDIT.

package ec2

import (
	context "context"

	serviceec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	mock "github.com/stretchr/testify/mock"
)

// EC2 is an autogenerated mock type for the EC2 type
type EC2 struct {
	mock.Mock
}

type EC2_Expecter struct {
	mock *mock.Mock
}

func (_m *EC2) EXPECT() *EC2_Expecter {
	return &EC2_Expecter{mock: &_m.Mock}
}

// DescribeInstances provides a mock function with given fields: ctx, e, optFns
func (_m *EC2) DescribeInstances(ctx context.Context, e *serviceec2.DescribeInstancesInput, optFns ...func(*serviceec2.Options)) (*serviceec2.DescribeInstancesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, e)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeInstances")
	}

	var r0 *serviceec2.DescribeInstancesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *serviceec2.DescribeInstancesInput, ...func(*serviceec2.Options)) (*serviceec2.DescribeInstancesOutput, error)); ok {
		return rf(ctx, e, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *serviceec2.DescribeInstancesInput, ...func(*serviceec2.Options)) *serviceec2.DescribeInstancesOutput); ok {
		r0 = rf(ctx, e, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*serviceec2.DescribeInstancesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *serviceec2.DescribeInstancesInput, ...func(*serviceec2.Options)) error); ok {
		r1 = rf(ctx, e, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EC2_DescribeInstances_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeInstances'
type EC2_DescribeInstances_Call struct {
	*mock.Call
}

// DescribeInstances is a helper method to define mock.On call
//   - ctx context.Context
//   - e *serviceec2.DescribeInstancesInput
//   - optFns ...func(*serviceec2.Options)
func (_e *EC2_Expecter) DescribeInstances(ctx interface{}, e interface{}, optFns ...interface{}) *EC2_DescribeInstances_Call {
	return &EC2_DescribeInstances_Call{Call: _e.mock.On("DescribeInstances",
		append([]interface{}{ctx, e}, optFns...)...)}
}

func (_c *EC2_DescribeInstances_Call) Run(run func(ctx context.Context, e *serviceec2.DescribeInstancesInput, optFns ...func(*serviceec2.Options))) *EC2_DescribeInstances_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*serviceec2.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*serviceec2.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*serviceec2.DescribeInstancesInput), variadicArgs...)
	})
	return _c
}

func (_c *EC2_DescribeInstances_Call) Return(_a0 *serviceec2.DescribeInstancesOutput, _a1 error) *EC2_DescribeInstances_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EC2_DescribeInstances_Call) RunAndReturn(run func(context.Context, *serviceec2.DescribeInstancesInput, ...func(*serviceec2.Options)) (*serviceec2.DescribeInstancesOutput, error)) *EC2_DescribeInstances_Call {
	_c.Call.Return(run)
	return _c
}

// DescribeRegions provides a mock function with given fields: ctx, e, optFns
func (_m *EC2) DescribeRegions(ctx context.Context, e *serviceec2.DescribeRegionsInput, optFns ...func(*serviceec2.Options)) (*serviceec2.DescribeRegionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, e)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeRegions")
	}

	var r0 *serviceec2.DescribeRegionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *serviceec2.DescribeRegionsInput, ...func(*serviceec2.Options)) (*serviceec2.DescribeRegionsOutput, error)); ok {
		return rf(ctx, e, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *serviceec2.DescribeRegionsInput, ...func(*serviceec2.Options)) *serviceec2.DescribeRegionsOutput); ok {
		r0 = rf(ctx, e, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*serviceec2.DescribeRegionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *serviceec2.DescribeRegionsInput, ...func(*serviceec2.Options)) error); ok {
		r1 = rf(ctx, e, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EC2_DescribeRegions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeRegions'
type EC2_DescribeRegions_Call struct {
	*mock.Call
}

// DescribeRegions is a helper method to define mock.On call
//   - ctx context.Context
//   - e *serviceec2.DescribeRegionsInput
//   - optFns ...func(*serviceec2.Options)
func (_e *EC2_Expecter) DescribeRegions(ctx interface{}, e interface{}, optFns ...interface{}) *EC2_DescribeRegions_Call {
	return &EC2_DescribeRegions_Call{Call: _e.mock.On("DescribeRegions",
		append([]interface{}{ctx, e}, optFns...)...)}
}

func (_c *EC2_DescribeRegions_Call) Run(run func(ctx context.Context, e *serviceec2.DescribeRegionsInput, optFns ...func(*serviceec2.Options))) *EC2_DescribeRegions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*serviceec2.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*serviceec2.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*serviceec2.DescribeRegionsInput), variadicArgs...)
	})
	return _c
}

func (_c *EC2_DescribeRegions_Call) Return(_a0 *serviceec2.DescribeRegionsOutput, _a1 error) *EC2_DescribeRegions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EC2_DescribeRegions_Call) RunAndReturn(run func(context.Context, *serviceec2.DescribeRegionsInput, ...func(*serviceec2.Options)) (*serviceec2.DescribeRegionsOutput, error)) *EC2_DescribeRegions_Call {
	_c.Call.Return(run)
	return _c
}

// DescribeSpotPriceHistory provides a mock function with given fields: ctx, input, optFns
func (_m *EC2) DescribeSpotPriceHistory(ctx context.Context, input *serviceec2.DescribeSpotPriceHistoryInput, optFns ...func(*serviceec2.Options)) (*serviceec2.DescribeSpotPriceHistoryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeSpotPriceHistory")
	}

	var r0 *serviceec2.DescribeSpotPriceHistoryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *serviceec2.DescribeSpotPriceHistoryInput, ...func(*serviceec2.Options)) (*serviceec2.DescribeSpotPriceHistoryOutput, error)); ok {
		return rf(ctx, input, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *serviceec2.DescribeSpotPriceHistoryInput, ...func(*serviceec2.Options)) *serviceec2.DescribeSpotPriceHistoryOutput); ok {
		r0 = rf(ctx, input, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*serviceec2.DescribeSpotPriceHistoryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *serviceec2.DescribeSpotPriceHistoryInput, ...func(*serviceec2.Options)) error); ok {
		r1 = rf(ctx, input, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EC2_DescribeSpotPriceHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeSpotPriceHistory'
type EC2_DescribeSpotPriceHistory_Call struct {
	*mock.Call
}

// DescribeSpotPriceHistory is a helper method to define mock.On call
//   - ctx context.Context
//   - input *serviceec2.DescribeSpotPriceHistoryInput
//   - optFns ...func(*serviceec2.Options)
func (_e *EC2_Expecter) DescribeSpotPriceHistory(ctx interface{}, input interface{}, optFns ...interface{}) *EC2_DescribeSpotPriceHistory_Call {
	return &EC2_DescribeSpotPriceHistory_Call{Call: _e.mock.On("DescribeSpotPriceHistory",
		append([]interface{}{ctx, input}, optFns...)...)}
}

func (_c *EC2_DescribeSpotPriceHistory_Call) Run(run func(ctx context.Context, input *serviceec2.DescribeSpotPriceHistoryInput, optFns ...func(*serviceec2.Options))) *EC2_DescribeSpotPriceHistory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*serviceec2.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*serviceec2.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*serviceec2.DescribeSpotPriceHistoryInput), variadicArgs...)
	})
	return _c
}

func (_c *EC2_DescribeSpotPriceHistory_Call) Return(_a0 *serviceec2.DescribeSpotPriceHistoryOutput, _a1 error) *EC2_DescribeSpotPriceHistory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EC2_DescribeSpotPriceHistory_Call) RunAndReturn(run func(context.Context, *serviceec2.DescribeSpotPriceHistoryInput, ...func(*serviceec2.Options)) (*serviceec2.DescribeSpotPriceHistoryOutput, error)) *EC2_DescribeSpotPriceHistory_Call {
	_c.Call.Return(run)
	return _c
}

// NewEC2 creates a new instance of EC2. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEC2(t interface {
	mock.TestingT
	Cleanup(func())
}) *EC2 {
	mock := &EC2{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
