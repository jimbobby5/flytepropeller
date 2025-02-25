// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import common "github.com/lyft/flytepropeller/pkg/compiler/common"
import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
import mock "github.com/stretchr/testify/mock"

// Node is an autogenerated mock type for the Node type
type Node struct {
	mock.Mock
}

type Node_GetBranchNode struct {
	*mock.Call
}

func (_m Node_GetBranchNode) Return(_a0 *core.BranchNode) *Node_GetBranchNode {
	return &Node_GetBranchNode{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnGetBranchNode() *Node_GetBranchNode {
	c := _m.On("GetBranchNode")
	return &Node_GetBranchNode{Call: c}
}

func (_m *Node) OnGetBranchNodeMatch(matchers ...interface{}) *Node_GetBranchNode {
	c := _m.On("GetBranchNode", matchers...)
	return &Node_GetBranchNode{Call: c}
}

// GetBranchNode provides a mock function with given fields:
func (_m *Node) GetBranchNode() *core.BranchNode {
	ret := _m.Called()

	var r0 *core.BranchNode
	if rf, ok := ret.Get(0).(func() *core.BranchNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.BranchNode)
		}
	}

	return r0
}

type Node_GetCoreNode struct {
	*mock.Call
}

func (_m Node_GetCoreNode) Return(_a0 *core.Node) *Node_GetCoreNode {
	return &Node_GetCoreNode{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnGetCoreNode() *Node_GetCoreNode {
	c := _m.On("GetCoreNode")
	return &Node_GetCoreNode{Call: c}
}

func (_m *Node) OnGetCoreNodeMatch(matchers ...interface{}) *Node_GetCoreNode {
	c := _m.On("GetCoreNode", matchers...)
	return &Node_GetCoreNode{Call: c}
}

// GetCoreNode provides a mock function with given fields:
func (_m *Node) GetCoreNode() *core.Node {
	ret := _m.Called()

	var r0 *core.Node
	if rf, ok := ret.Get(0).(func() *core.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Node)
		}
	}

	return r0
}

type Node_GetId struct {
	*mock.Call
}

func (_m Node_GetId) Return(_a0 string) *Node_GetId {
	return &Node_GetId{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnGetId() *Node_GetId {
	c := _m.On("GetId")
	return &Node_GetId{Call: c}
}

func (_m *Node) OnGetIdMatch(matchers ...interface{}) *Node_GetId {
	c := _m.On("GetId", matchers...)
	return &Node_GetId{Call: c}
}

// GetId provides a mock function with given fields:
func (_m *Node) GetId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type Node_GetInputs struct {
	*mock.Call
}

func (_m Node_GetInputs) Return(_a0 []*core.Binding) *Node_GetInputs {
	return &Node_GetInputs{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnGetInputs() *Node_GetInputs {
	c := _m.On("GetInputs")
	return &Node_GetInputs{Call: c}
}

func (_m *Node) OnGetInputsMatch(matchers ...interface{}) *Node_GetInputs {
	c := _m.On("GetInputs", matchers...)
	return &Node_GetInputs{Call: c}
}

// GetInputs provides a mock function with given fields:
func (_m *Node) GetInputs() []*core.Binding {
	ret := _m.Called()

	var r0 []*core.Binding
	if rf, ok := ret.Get(0).(func() []*core.Binding); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.Binding)
		}
	}

	return r0
}

type Node_GetInterface struct {
	*mock.Call
}

func (_m Node_GetInterface) Return(_a0 *core.TypedInterface) *Node_GetInterface {
	return &Node_GetInterface{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnGetInterface() *Node_GetInterface {
	c := _m.On("GetInterface")
	return &Node_GetInterface{Call: c}
}

func (_m *Node) OnGetInterfaceMatch(matchers ...interface{}) *Node_GetInterface {
	c := _m.On("GetInterface", matchers...)
	return &Node_GetInterface{Call: c}
}

// GetInterface provides a mock function with given fields:
func (_m *Node) GetInterface() *core.TypedInterface {
	ret := _m.Called()

	var r0 *core.TypedInterface
	if rf, ok := ret.Get(0).(func() *core.TypedInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.TypedInterface)
		}
	}

	return r0
}

type Node_GetMetadata struct {
	*mock.Call
}

func (_m Node_GetMetadata) Return(_a0 *core.NodeMetadata) *Node_GetMetadata {
	return &Node_GetMetadata{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnGetMetadata() *Node_GetMetadata {
	c := _m.On("GetMetadata")
	return &Node_GetMetadata{Call: c}
}

func (_m *Node) OnGetMetadataMatch(matchers ...interface{}) *Node_GetMetadata {
	c := _m.On("GetMetadata", matchers...)
	return &Node_GetMetadata{Call: c}
}

// GetMetadata provides a mock function with given fields:
func (_m *Node) GetMetadata() *core.NodeMetadata {
	ret := _m.Called()

	var r0 *core.NodeMetadata
	if rf, ok := ret.Get(0).(func() *core.NodeMetadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.NodeMetadata)
		}
	}

	return r0
}

type Node_GetOutputAliases struct {
	*mock.Call
}

func (_m Node_GetOutputAliases) Return(_a0 []*core.Alias) *Node_GetOutputAliases {
	return &Node_GetOutputAliases{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnGetOutputAliases() *Node_GetOutputAliases {
	c := _m.On("GetOutputAliases")
	return &Node_GetOutputAliases{Call: c}
}

func (_m *Node) OnGetOutputAliasesMatch(matchers ...interface{}) *Node_GetOutputAliases {
	c := _m.On("GetOutputAliases", matchers...)
	return &Node_GetOutputAliases{Call: c}
}

// GetOutputAliases provides a mock function with given fields:
func (_m *Node) GetOutputAliases() []*core.Alias {
	ret := _m.Called()

	var r0 []*core.Alias
	if rf, ok := ret.Get(0).(func() []*core.Alias); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.Alias)
		}
	}

	return r0
}

type Node_GetSubWorkflow struct {
	*mock.Call
}

func (_m Node_GetSubWorkflow) Return(_a0 common.Workflow) *Node_GetSubWorkflow {
	return &Node_GetSubWorkflow{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnGetSubWorkflow() *Node_GetSubWorkflow {
	c := _m.On("GetSubWorkflow")
	return &Node_GetSubWorkflow{Call: c}
}

func (_m *Node) OnGetSubWorkflowMatch(matchers ...interface{}) *Node_GetSubWorkflow {
	c := _m.On("GetSubWorkflow", matchers...)
	return &Node_GetSubWorkflow{Call: c}
}

// GetSubWorkflow provides a mock function with given fields:
func (_m *Node) GetSubWorkflow() common.Workflow {
	ret := _m.Called()

	var r0 common.Workflow
	if rf, ok := ret.Get(0).(func() common.Workflow); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Workflow)
		}
	}

	return r0
}

type Node_GetTask struct {
	*mock.Call
}

func (_m Node_GetTask) Return(_a0 common.Task) *Node_GetTask {
	return &Node_GetTask{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnGetTask() *Node_GetTask {
	c := _m.On("GetTask")
	return &Node_GetTask{Call: c}
}

func (_m *Node) OnGetTaskMatch(matchers ...interface{}) *Node_GetTask {
	c := _m.On("GetTask", matchers...)
	return &Node_GetTask{Call: c}
}

// GetTask provides a mock function with given fields:
func (_m *Node) GetTask() common.Task {
	ret := _m.Called()

	var r0 common.Task
	if rf, ok := ret.Get(0).(func() common.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Task)
		}
	}

	return r0
}

type Node_GetTaskNode struct {
	*mock.Call
}

func (_m Node_GetTaskNode) Return(_a0 *core.TaskNode) *Node_GetTaskNode {
	return &Node_GetTaskNode{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnGetTaskNode() *Node_GetTaskNode {
	c := _m.On("GetTaskNode")
	return &Node_GetTaskNode{Call: c}
}

func (_m *Node) OnGetTaskNodeMatch(matchers ...interface{}) *Node_GetTaskNode {
	c := _m.On("GetTaskNode", matchers...)
	return &Node_GetTaskNode{Call: c}
}

// GetTaskNode provides a mock function with given fields:
func (_m *Node) GetTaskNode() *core.TaskNode {
	ret := _m.Called()

	var r0 *core.TaskNode
	if rf, ok := ret.Get(0).(func() *core.TaskNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.TaskNode)
		}
	}

	return r0
}

type Node_GetUpstreamNodeIds struct {
	*mock.Call
}

func (_m Node_GetUpstreamNodeIds) Return(_a0 []string) *Node_GetUpstreamNodeIds {
	return &Node_GetUpstreamNodeIds{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnGetUpstreamNodeIds() *Node_GetUpstreamNodeIds {
	c := _m.On("GetUpstreamNodeIds")
	return &Node_GetUpstreamNodeIds{Call: c}
}

func (_m *Node) OnGetUpstreamNodeIdsMatch(matchers ...interface{}) *Node_GetUpstreamNodeIds {
	c := _m.On("GetUpstreamNodeIds", matchers...)
	return &Node_GetUpstreamNodeIds{Call: c}
}

// GetUpstreamNodeIds provides a mock function with given fields:
func (_m *Node) GetUpstreamNodeIds() []string {
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

type Node_GetWorkflowNode struct {
	*mock.Call
}

func (_m Node_GetWorkflowNode) Return(_a0 *core.WorkflowNode) *Node_GetWorkflowNode {
	return &Node_GetWorkflowNode{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnGetWorkflowNode() *Node_GetWorkflowNode {
	c := _m.On("GetWorkflowNode")
	return &Node_GetWorkflowNode{Call: c}
}

func (_m *Node) OnGetWorkflowNodeMatch(matchers ...interface{}) *Node_GetWorkflowNode {
	c := _m.On("GetWorkflowNode", matchers...)
	return &Node_GetWorkflowNode{Call: c}
}

// GetWorkflowNode provides a mock function with given fields:
func (_m *Node) GetWorkflowNode() *core.WorkflowNode {
	ret := _m.Called()

	var r0 *core.WorkflowNode
	if rf, ok := ret.Get(0).(func() *core.WorkflowNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.WorkflowNode)
		}
	}

	return r0
}
