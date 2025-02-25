// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import handler "github.com/lyft/flytepropeller/pkg/controller/nodes/handler"
import mock "github.com/stretchr/testify/mock"

// NodeStateWriter is an autogenerated mock type for the NodeStateWriter type
type NodeStateWriter struct {
	mock.Mock
}

type NodeStateWriter_PutBranchNode struct {
	*mock.Call
}

func (_m NodeStateWriter_PutBranchNode) Return(_a0 error) *NodeStateWriter_PutBranchNode {
	return &NodeStateWriter_PutBranchNode{Call: _m.Call.Return(_a0)}
}

func (_m *NodeStateWriter) OnPutBranchNode(s handler.BranchNodeState) *NodeStateWriter_PutBranchNode {
	c := _m.On("PutBranchNode", s)
	return &NodeStateWriter_PutBranchNode{Call: c}
}

func (_m *NodeStateWriter) OnPutBranchNodeMatch(matchers ...interface{}) *NodeStateWriter_PutBranchNode {
	c := _m.On("PutBranchNode", matchers...)
	return &NodeStateWriter_PutBranchNode{Call: c}
}

// PutBranchNode provides a mock function with given fields: s
func (_m *NodeStateWriter) PutBranchNode(s handler.BranchNodeState) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(handler.BranchNodeState) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NodeStateWriter_PutDynamicNodeState struct {
	*mock.Call
}

func (_m NodeStateWriter_PutDynamicNodeState) Return(_a0 error) *NodeStateWriter_PutDynamicNodeState {
	return &NodeStateWriter_PutDynamicNodeState{Call: _m.Call.Return(_a0)}
}

func (_m *NodeStateWriter) OnPutDynamicNodeState(s handler.DynamicNodeState) *NodeStateWriter_PutDynamicNodeState {
	c := _m.On("PutDynamicNodeState", s)
	return &NodeStateWriter_PutDynamicNodeState{Call: c}
}

func (_m *NodeStateWriter) OnPutDynamicNodeStateMatch(matchers ...interface{}) *NodeStateWriter_PutDynamicNodeState {
	c := _m.On("PutDynamicNodeState", matchers...)
	return &NodeStateWriter_PutDynamicNodeState{Call: c}
}

// PutDynamicNodeState provides a mock function with given fields: s
func (_m *NodeStateWriter) PutDynamicNodeState(s handler.DynamicNodeState) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(handler.DynamicNodeState) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NodeStateWriter_PutTaskNodeState struct {
	*mock.Call
}

func (_m NodeStateWriter_PutTaskNodeState) Return(_a0 error) *NodeStateWriter_PutTaskNodeState {
	return &NodeStateWriter_PutTaskNodeState{Call: _m.Call.Return(_a0)}
}

func (_m *NodeStateWriter) OnPutTaskNodeState(s handler.TaskNodeState) *NodeStateWriter_PutTaskNodeState {
	c := _m.On("PutTaskNodeState", s)
	return &NodeStateWriter_PutTaskNodeState{Call: c}
}

func (_m *NodeStateWriter) OnPutTaskNodeStateMatch(matchers ...interface{}) *NodeStateWriter_PutTaskNodeState {
	c := _m.On("PutTaskNodeState", matchers...)
	return &NodeStateWriter_PutTaskNodeState{Call: c}
}

// PutTaskNodeState provides a mock function with given fields: s
func (_m *NodeStateWriter) PutTaskNodeState(s handler.TaskNodeState) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(handler.TaskNodeState) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NodeStateWriter_PutWorkflowNodeState struct {
	*mock.Call
}

func (_m NodeStateWriter_PutWorkflowNodeState) Return(_a0 error) *NodeStateWriter_PutWorkflowNodeState {
	return &NodeStateWriter_PutWorkflowNodeState{Call: _m.Call.Return(_a0)}
}

func (_m *NodeStateWriter) OnPutWorkflowNodeState(s handler.WorkflowNodeState) *NodeStateWriter_PutWorkflowNodeState {
	c := _m.On("PutWorkflowNodeState", s)
	return &NodeStateWriter_PutWorkflowNodeState{Call: c}
}

func (_m *NodeStateWriter) OnPutWorkflowNodeStateMatch(matchers ...interface{}) *NodeStateWriter_PutWorkflowNodeState {
	c := _m.On("PutWorkflowNodeState", matchers...)
	return &NodeStateWriter_PutWorkflowNodeState{Call: c}
}

// PutWorkflowNodeState provides a mock function with given fields: s
func (_m *NodeStateWriter) PutWorkflowNodeState(s handler.WorkflowNodeState) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(handler.WorkflowNodeState) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
