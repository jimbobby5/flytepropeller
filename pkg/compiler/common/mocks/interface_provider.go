// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
import mock "github.com/stretchr/testify/mock"

// InterfaceProvider is an autogenerated mock type for the InterfaceProvider type
type InterfaceProvider struct {
	mock.Mock
}

type InterfaceProvider_GetExpectedInputs struct {
	*mock.Call
}

func (_m InterfaceProvider_GetExpectedInputs) Return(_a0 *core.ParameterMap) *InterfaceProvider_GetExpectedInputs {
	return &InterfaceProvider_GetExpectedInputs{Call: _m.Call.Return(_a0)}
}

func (_m *InterfaceProvider) OnGetExpectedInputs() *InterfaceProvider_GetExpectedInputs {
	c := _m.On("GetExpectedInputs")
	return &InterfaceProvider_GetExpectedInputs{Call: c}
}

func (_m *InterfaceProvider) OnGetExpectedInputsMatch(matchers ...interface{}) *InterfaceProvider_GetExpectedInputs {
	c := _m.On("GetExpectedInputs", matchers...)
	return &InterfaceProvider_GetExpectedInputs{Call: c}
}

// GetExpectedInputs provides a mock function with given fields:
func (_m *InterfaceProvider) GetExpectedInputs() *core.ParameterMap {
	ret := _m.Called()

	var r0 *core.ParameterMap
	if rf, ok := ret.Get(0).(func() *core.ParameterMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ParameterMap)
		}
	}

	return r0
}

type InterfaceProvider_GetExpectedOutputs struct {
	*mock.Call
}

func (_m InterfaceProvider_GetExpectedOutputs) Return(_a0 *core.VariableMap) *InterfaceProvider_GetExpectedOutputs {
	return &InterfaceProvider_GetExpectedOutputs{Call: _m.Call.Return(_a0)}
}

func (_m *InterfaceProvider) OnGetExpectedOutputs() *InterfaceProvider_GetExpectedOutputs {
	c := _m.On("GetExpectedOutputs")
	return &InterfaceProvider_GetExpectedOutputs{Call: c}
}

func (_m *InterfaceProvider) OnGetExpectedOutputsMatch(matchers ...interface{}) *InterfaceProvider_GetExpectedOutputs {
	c := _m.On("GetExpectedOutputs", matchers...)
	return &InterfaceProvider_GetExpectedOutputs{Call: c}
}

// GetExpectedOutputs provides a mock function with given fields:
func (_m *InterfaceProvider) GetExpectedOutputs() *core.VariableMap {
	ret := _m.Called()

	var r0 *core.VariableMap
	if rf, ok := ret.Get(0).(func() *core.VariableMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.VariableMap)
		}
	}

	return r0
}

type InterfaceProvider_GetID struct {
	*mock.Call
}

func (_m InterfaceProvider_GetID) Return(_a0 *core.Identifier) *InterfaceProvider_GetID {
	return &InterfaceProvider_GetID{Call: _m.Call.Return(_a0)}
}

func (_m *InterfaceProvider) OnGetID() *InterfaceProvider_GetID {
	c := _m.On("GetID")
	return &InterfaceProvider_GetID{Call: c}
}

func (_m *InterfaceProvider) OnGetIDMatch(matchers ...interface{}) *InterfaceProvider_GetID {
	c := _m.On("GetID", matchers...)
	return &InterfaceProvider_GetID{Call: c}
}

// GetID provides a mock function with given fields:
func (_m *InterfaceProvider) GetID() *core.Identifier {
	ret := _m.Called()

	var r0 *core.Identifier
	if rf, ok := ret.Get(0).(func() *core.Identifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identifier)
		}
	}

	return r0
}
