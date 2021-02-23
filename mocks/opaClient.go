// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rode/rode/opa (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	opa "github.com/rode/rode/opa"
	reflect "reflect"
)

// MockOpaClient is a mock of Client interface
type MockOpaClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpaClientMockRecorder
}

// MockOpaClientMockRecorder is the mock recorder for MockOpaClient
type MockOpaClientMockRecorder struct {
	mock *MockOpaClient
}

// NewMockOpaClient creates a new mock instance
func NewMockOpaClient(ctrl *gomock.Controller) *MockOpaClient {
	mock := &MockOpaClient{ctrl: ctrl}
	mock.recorder = &MockOpaClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpaClient) EXPECT() *MockOpaClientMockRecorder {
	return m.recorder
}

// EvaluatePolicy mocks base method
func (m *MockOpaClient) EvaluatePolicy(arg0 string, arg1 []byte) (*opa.EvaluatePolicyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EvaluatePolicy", arg0, arg1)
	ret0, _ := ret[0].(*opa.EvaluatePolicyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvaluatePolicy indicates an expected call of EvaluatePolicy
func (mr *MockOpaClientMockRecorder) EvaluatePolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvaluatePolicy", reflect.TypeOf((*MockOpaClient)(nil).EvaluatePolicy), arg0, arg1)
}

// InitializePolicy mocks base method
func (m *MockOpaClient) InitializePolicy(arg0 string) opa.ClientError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializePolicy", arg0)
	ret0, _ := ret[0].(opa.ClientError)
	return ret0
}

// InitializePolicy indicates an expected call of InitializePolicy
func (mr *MockOpaClientMockRecorder) InitializePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializePolicy", reflect.TypeOf((*MockOpaClient)(nil).InitializePolicy), arg0)
}