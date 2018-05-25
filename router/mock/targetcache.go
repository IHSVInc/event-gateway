// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/serverless/event-gateway/router (interfaces: Targeter)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	event "github.com/serverless/event-gateway/event"
	function "github.com/serverless/event-gateway/function"
	pathtree "github.com/serverless/event-gateway/internal/pathtree"
	router "github.com/serverless/event-gateway/router"
	subscription "github.com/serverless/event-gateway/subscription"
	reflect "reflect"
)

// MockTargeter is a mock of Targeter interface
type MockTargeter struct {
	ctrl     *gomock.Controller
	recorder *MockTargeterMockRecorder
}

// MockTargeterMockRecorder is the mock recorder for MockTargeter
type MockTargeterMockRecorder struct {
	mock *MockTargeter
}

// NewMockTargeter creates a new mock instance
func NewMockTargeter(ctrl *gomock.Controller) *MockTargeter {
	mock := &MockTargeter{ctrl: ctrl}
	mock.recorder = &MockTargeterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTargeter) EXPECT() *MockTargeterMockRecorder {
	return m.recorder
}

// Function mocks base method
func (m *MockTargeter) Function(arg0 string, arg1 function.ID) *function.Function {
	ret := m.ctrl.Call(m, "Function", arg0, arg1)
	ret0, _ := ret[0].(*function.Function)
	return ret0
}

// Function indicates an expected call of Function
func (mr *MockTargeterMockRecorder) Function(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Function", reflect.TypeOf((*MockTargeter)(nil).Function), arg0, arg1)
}

// HTTPBackingFunction mocks base method
func (m *MockTargeter) HTTPBackingFunction(arg0, arg1 string) (string, *function.ID, pathtree.Params, *subscription.CORS) {
	ret := m.ctrl.Call(m, "HTTPBackingFunction", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*function.ID)
	ret2, _ := ret[2].(pathtree.Params)
	ret3, _ := ret[3].(*subscription.CORS)
	return ret0, ret1, ret2, ret3
}

// HTTPBackingFunction indicates an expected call of HTTPBackingFunction
func (mr *MockTargeterMockRecorder) HTTPBackingFunction(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPBackingFunction", reflect.TypeOf((*MockTargeter)(nil).HTTPBackingFunction), arg0, arg1)
}

// SubscribersOfEvent mocks base method
func (m *MockTargeter) SubscribersOfEvent(arg0 string, arg1 event.TypeName) []router.FunctionInfo {
	ret := m.ctrl.Call(m, "SubscribersOfEvent", arg0, arg1)
	ret0, _ := ret[0].([]router.FunctionInfo)
	return ret0
}

// SubscribersOfEvent indicates an expected call of SubscribersOfEvent
func (mr *MockTargeterMockRecorder) SubscribersOfEvent(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribersOfEvent", reflect.TypeOf((*MockTargeter)(nil).SubscribersOfEvent), arg0, arg1)
}
