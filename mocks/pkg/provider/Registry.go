// Code generated by mockery v2.38.0. DO NOT EDIT.

package provider

import (
	io_prometheus_client "github.com/prometheus/client_model/go"
	mock "github.com/stretchr/testify/mock"

	prometheus "github.com/prometheus/client_golang/prometheus"
)

// Registry is an autogenerated mock type for the Registry type
type Registry struct {
	mock.Mock
}

type Registry_Expecter struct {
	mock *mock.Mock
}

func (_m *Registry) EXPECT() *Registry_Expecter {
	return &Registry_Expecter{mock: &_m.Mock}
}

// Collect provides a mock function with given fields: _a0
func (_m *Registry) Collect(_a0 chan<- prometheus.Metric) {
	_m.Called(_a0)
}

// Registry_Collect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Collect'
type Registry_Collect_Call struct {
	*mock.Call
}

// Collect is a helper method to define mock.On call
//   - _a0 chan<- prometheus.Metric
func (_e *Registry_Expecter) Collect(_a0 interface{}) *Registry_Collect_Call {
	return &Registry_Collect_Call{Call: _e.mock.On("Collect", _a0)}
}

func (_c *Registry_Collect_Call) Run(run func(_a0 chan<- prometheus.Metric)) *Registry_Collect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(chan<- prometheus.Metric))
	})
	return _c
}

func (_c *Registry_Collect_Call) Return() *Registry_Collect_Call {
	_c.Call.Return()
	return _c
}

func (_c *Registry_Collect_Call) RunAndReturn(run func(chan<- prometheus.Metric)) *Registry_Collect_Call {
	_c.Call.Return(run)
	return _c
}

// Describe provides a mock function with given fields: _a0
func (_m *Registry) Describe(_a0 chan<- *prometheus.Desc) {
	_m.Called(_a0)
}

// Registry_Describe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Describe'
type Registry_Describe_Call struct {
	*mock.Call
}

// Describe is a helper method to define mock.On call
//   - _a0 chan<- *prometheus.Desc
func (_e *Registry_Expecter) Describe(_a0 interface{}) *Registry_Describe_Call {
	return &Registry_Describe_Call{Call: _e.mock.On("Describe", _a0)}
}

func (_c *Registry_Describe_Call) Run(run func(_a0 chan<- *prometheus.Desc)) *Registry_Describe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(chan<- *prometheus.Desc))
	})
	return _c
}

func (_c *Registry_Describe_Call) Return() *Registry_Describe_Call {
	_c.Call.Return()
	return _c
}

func (_c *Registry_Describe_Call) RunAndReturn(run func(chan<- *prometheus.Desc)) *Registry_Describe_Call {
	_c.Call.Return(run)
	return _c
}

// Gather provides a mock function with given fields:
func (_m *Registry) Gather() ([]*io_prometheus_client.MetricFamily, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Gather")
	}

	var r0 []*io_prometheus_client.MetricFamily
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*io_prometheus_client.MetricFamily, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*io_prometheus_client.MetricFamily); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*io_prometheus_client.MetricFamily)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Registry_Gather_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Gather'
type Registry_Gather_Call struct {
	*mock.Call
}

// Gather is a helper method to define mock.On call
func (_e *Registry_Expecter) Gather() *Registry_Gather_Call {
	return &Registry_Gather_Call{Call: _e.mock.On("Gather")}
}

func (_c *Registry_Gather_Call) Run(run func()) *Registry_Gather_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Registry_Gather_Call) Return(_a0 []*io_prometheus_client.MetricFamily, _a1 error) *Registry_Gather_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Registry_Gather_Call) RunAndReturn(run func() ([]*io_prometheus_client.MetricFamily, error)) *Registry_Gather_Call {
	_c.Call.Return(run)
	return _c
}

// MustRegister provides a mock function with given fields: _a0
func (_m *Registry) MustRegister(_a0 ...prometheus.Collector) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Registry_MustRegister_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MustRegister'
type Registry_MustRegister_Call struct {
	*mock.Call
}

// MustRegister is a helper method to define mock.On call
//   - _a0 ...prometheus.Collector
func (_e *Registry_Expecter) MustRegister(_a0 ...interface{}) *Registry_MustRegister_Call {
	return &Registry_MustRegister_Call{Call: _e.mock.On("MustRegister",
		append([]interface{}{}, _a0...)...)}
}

func (_c *Registry_MustRegister_Call) Run(run func(_a0 ...prometheus.Collector)) *Registry_MustRegister_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]prometheus.Collector, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(prometheus.Collector)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Registry_MustRegister_Call) Return() *Registry_MustRegister_Call {
	_c.Call.Return()
	return _c
}

func (_c *Registry_MustRegister_Call) RunAndReturn(run func(...prometheus.Collector)) *Registry_MustRegister_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: _a0
func (_m *Registry) Register(_a0 prometheus.Collector) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(prometheus.Collector) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Registry_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type Registry_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - _a0 prometheus.Collector
func (_e *Registry_Expecter) Register(_a0 interface{}) *Registry_Register_Call {
	return &Registry_Register_Call{Call: _e.mock.On("Register", _a0)}
}

func (_c *Registry_Register_Call) Run(run func(_a0 prometheus.Collector)) *Registry_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(prometheus.Collector))
	})
	return _c
}

func (_c *Registry_Register_Call) Return(_a0 error) *Registry_Register_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Registry_Register_Call) RunAndReturn(run func(prometheus.Collector) error) *Registry_Register_Call {
	_c.Call.Return(run)
	return _c
}

// Unregister provides a mock function with given fields: _a0
func (_m *Registry) Unregister(_a0 prometheus.Collector) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Unregister")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(prometheus.Collector) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Registry_Unregister_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unregister'
type Registry_Unregister_Call struct {
	*mock.Call
}

// Unregister is a helper method to define mock.On call
//   - _a0 prometheus.Collector
func (_e *Registry_Expecter) Unregister(_a0 interface{}) *Registry_Unregister_Call {
	return &Registry_Unregister_Call{Call: _e.mock.On("Unregister", _a0)}
}

func (_c *Registry_Unregister_Call) Run(run func(_a0 prometheus.Collector)) *Registry_Unregister_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(prometheus.Collector))
	})
	return _c
}

func (_c *Registry_Unregister_Call) Return(_a0 bool) *Registry_Unregister_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Registry_Unregister_Call) RunAndReturn(run func(prometheus.Collector) bool) *Registry_Unregister_Call {
	_c.Call.Return(run)
	return _c
}

// NewRegistry creates a new instance of Registry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRegistry(t interface {
	mock.TestingT
	Cleanup(func())
}) *Registry {
	mock := &Registry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
