// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"

// ExecutableSubWorkflow is an autogenerated mock type for the ExecutableSubWorkflow type
type ExecutableSubWorkflow struct {
	mock.Mock
}

type ExecutableSubWorkflow_FromNode struct {
	*mock.Call
}

func (_m ExecutableSubWorkflow_FromNode) Return(_a0 []string, _a1 error) *ExecutableSubWorkflow_FromNode {
	return &ExecutableSubWorkflow_FromNode{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ExecutableSubWorkflow) OnFromNode(name string) *ExecutableSubWorkflow_FromNode {
	c := _m.On("FromNode", name)
	return &ExecutableSubWorkflow_FromNode{Call: c}
}

func (_m *ExecutableSubWorkflow) OnFromNodeMatch(matchers ...interface{}) *ExecutableSubWorkflow_FromNode {
	c := _m.On("FromNode", matchers...)
	return &ExecutableSubWorkflow_FromNode{Call: c}
}

// FromNode provides a mock function with given fields: name
func (_m *ExecutableSubWorkflow) FromNode(name string) ([]string, error) {
	ret := _m.Called(name)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ExecutableSubWorkflow_GetConnections struct {
	*mock.Call
}

func (_m ExecutableSubWorkflow_GetConnections) Return(_a0 *v1alpha1.Connections) *ExecutableSubWorkflow_GetConnections {
	return &ExecutableSubWorkflow_GetConnections{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableSubWorkflow) OnGetConnections() *ExecutableSubWorkflow_GetConnections {
	c := _m.On("GetConnections")
	return &ExecutableSubWorkflow_GetConnections{Call: c}
}

func (_m *ExecutableSubWorkflow) OnGetConnectionsMatch(matchers ...interface{}) *ExecutableSubWorkflow_GetConnections {
	c := _m.On("GetConnections", matchers...)
	return &ExecutableSubWorkflow_GetConnections{Call: c}
}

// GetConnections provides a mock function with given fields:
func (_m *ExecutableSubWorkflow) GetConnections() *v1alpha1.Connections {
	ret := _m.Called()

	var r0 *v1alpha1.Connections
	if rf, ok := ret.Get(0).(func() *v1alpha1.Connections); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Connections)
		}
	}

	return r0
}

type ExecutableSubWorkflow_GetID struct {
	*mock.Call
}

func (_m ExecutableSubWorkflow_GetID) Return(_a0 string) *ExecutableSubWorkflow_GetID {
	return &ExecutableSubWorkflow_GetID{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableSubWorkflow) OnGetID() *ExecutableSubWorkflow_GetID {
	c := _m.On("GetID")
	return &ExecutableSubWorkflow_GetID{Call: c}
}

func (_m *ExecutableSubWorkflow) OnGetIDMatch(matchers ...interface{}) *ExecutableSubWorkflow_GetID {
	c := _m.On("GetID", matchers...)
	return &ExecutableSubWorkflow_GetID{Call: c}
}

// GetID provides a mock function with given fields:
func (_m *ExecutableSubWorkflow) GetID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ExecutableSubWorkflow_GetNode struct {
	*mock.Call
}

func (_m ExecutableSubWorkflow_GetNode) Return(_a0 v1alpha1.ExecutableNode, _a1 bool) *ExecutableSubWorkflow_GetNode {
	return &ExecutableSubWorkflow_GetNode{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ExecutableSubWorkflow) OnGetNode(nodeID string) *ExecutableSubWorkflow_GetNode {
	c := _m.On("GetNode", nodeID)
	return &ExecutableSubWorkflow_GetNode{Call: c}
}

func (_m *ExecutableSubWorkflow) OnGetNodeMatch(matchers ...interface{}) *ExecutableSubWorkflow_GetNode {
	c := _m.On("GetNode", matchers...)
	return &ExecutableSubWorkflow_GetNode{Call: c}
}

// GetNode provides a mock function with given fields: nodeID
func (_m *ExecutableSubWorkflow) GetNode(nodeID string) (v1alpha1.ExecutableNode, bool) {
	ret := _m.Called(nodeID)

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ExecutableNode); ok {
		r0 = rf(nodeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(nodeID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

type ExecutableSubWorkflow_GetNodes struct {
	*mock.Call
}

func (_m ExecutableSubWorkflow_GetNodes) Return(_a0 []string) *ExecutableSubWorkflow_GetNodes {
	return &ExecutableSubWorkflow_GetNodes{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableSubWorkflow) OnGetNodes() *ExecutableSubWorkflow_GetNodes {
	c := _m.On("GetNodes")
	return &ExecutableSubWorkflow_GetNodes{Call: c}
}

func (_m *ExecutableSubWorkflow) OnGetNodesMatch(matchers ...interface{}) *ExecutableSubWorkflow_GetNodes {
	c := _m.On("GetNodes", matchers...)
	return &ExecutableSubWorkflow_GetNodes{Call: c}
}

// GetNodes provides a mock function with given fields:
func (_m *ExecutableSubWorkflow) GetNodes() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

type ExecutableSubWorkflow_GetOnFailureNode struct {
	*mock.Call
}

func (_m ExecutableSubWorkflow_GetOnFailureNode) Return(_a0 v1alpha1.ExecutableNode) *ExecutableSubWorkflow_GetOnFailureNode {
	return &ExecutableSubWorkflow_GetOnFailureNode{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableSubWorkflow) OnGetOnFailureNode() *ExecutableSubWorkflow_GetOnFailureNode {
	c := _m.On("GetOnFailureNode")
	return &ExecutableSubWorkflow_GetOnFailureNode{Call: c}
}

func (_m *ExecutableSubWorkflow) OnGetOnFailureNodeMatch(matchers ...interface{}) *ExecutableSubWorkflow_GetOnFailureNode {
	c := _m.On("GetOnFailureNode", matchers...)
	return &ExecutableSubWorkflow_GetOnFailureNode{Call: c}
}

// GetOnFailureNode provides a mock function with given fields:
func (_m *ExecutableSubWorkflow) GetOnFailureNode() v1alpha1.ExecutableNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	return r0
}

type ExecutableSubWorkflow_GetOutputBindings struct {
	*mock.Call
}

func (_m ExecutableSubWorkflow_GetOutputBindings) Return(_a0 []*v1alpha1.Binding) *ExecutableSubWorkflow_GetOutputBindings {
	return &ExecutableSubWorkflow_GetOutputBindings{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableSubWorkflow) OnGetOutputBindings() *ExecutableSubWorkflow_GetOutputBindings {
	c := _m.On("GetOutputBindings")
	return &ExecutableSubWorkflow_GetOutputBindings{Call: c}
}

func (_m *ExecutableSubWorkflow) OnGetOutputBindingsMatch(matchers ...interface{}) *ExecutableSubWorkflow_GetOutputBindings {
	c := _m.On("GetOutputBindings", matchers...)
	return &ExecutableSubWorkflow_GetOutputBindings{Call: c}
}

// GetOutputBindings provides a mock function with given fields:
func (_m *ExecutableSubWorkflow) GetOutputBindings() []*v1alpha1.Binding {
	ret := _m.Called()

	var r0 []*v1alpha1.Binding
	if rf, ok := ret.Get(0).(func() []*v1alpha1.Binding); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1alpha1.Binding)
		}
	}

	return r0
}

type ExecutableSubWorkflow_GetOutputs struct {
	*mock.Call
}

func (_m ExecutableSubWorkflow_GetOutputs) Return(_a0 *v1alpha1.OutputVarMap) *ExecutableSubWorkflow_GetOutputs {
	return &ExecutableSubWorkflow_GetOutputs{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableSubWorkflow) OnGetOutputs() *ExecutableSubWorkflow_GetOutputs {
	c := _m.On("GetOutputs")
	return &ExecutableSubWorkflow_GetOutputs{Call: c}
}

func (_m *ExecutableSubWorkflow) OnGetOutputsMatch(matchers ...interface{}) *ExecutableSubWorkflow_GetOutputs {
	c := _m.On("GetOutputs", matchers...)
	return &ExecutableSubWorkflow_GetOutputs{Call: c}
}

// GetOutputs provides a mock function with given fields:
func (_m *ExecutableSubWorkflow) GetOutputs() *v1alpha1.OutputVarMap {
	ret := _m.Called()

	var r0 *v1alpha1.OutputVarMap
	if rf, ok := ret.Get(0).(func() *v1alpha1.OutputVarMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.OutputVarMap)
		}
	}

	return r0
}

type ExecutableSubWorkflow_StartNode struct {
	*mock.Call
}

func (_m ExecutableSubWorkflow_StartNode) Return(_a0 v1alpha1.ExecutableNode) *ExecutableSubWorkflow_StartNode {
	return &ExecutableSubWorkflow_StartNode{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableSubWorkflow) OnStartNode() *ExecutableSubWorkflow_StartNode {
	c := _m.On("StartNode")
	return &ExecutableSubWorkflow_StartNode{Call: c}
}

func (_m *ExecutableSubWorkflow) OnStartNodeMatch(matchers ...interface{}) *ExecutableSubWorkflow_StartNode {
	c := _m.On("StartNode", matchers...)
	return &ExecutableSubWorkflow_StartNode{Call: c}
}

// StartNode provides a mock function with given fields:
func (_m *ExecutableSubWorkflow) StartNode() v1alpha1.ExecutableNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	return r0
}
