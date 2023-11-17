// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-application-networking-k8s/pkg/k8s (interfaces: FinalizerManager)

// Package k8s is a generated GoMock package.
package k8s

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockFinalizerManager is a mock of FinalizerManager interface.
type MockFinalizerManager struct {
	ctrl     *gomock.Controller
	recorder *MockFinalizerManagerMockRecorder
}

// MockFinalizerManagerMockRecorder is the mock recorder for MockFinalizerManager.
type MockFinalizerManagerMockRecorder struct {
	mock *MockFinalizerManager
}

// NewMockFinalizerManager creates a new mock instance.
func NewMockFinalizerManager(ctrl *gomock.Controller) *MockFinalizerManager {
	mock := &MockFinalizerManager{ctrl: ctrl}
	mock.recorder = &MockFinalizerManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFinalizerManager) EXPECT() *MockFinalizerManagerMockRecorder {
	return m.recorder
}

// AddFinalizers mocks base method.
func (m *MockFinalizerManager) AddFinalizers(arg0 context.Context, arg1 client.Object, arg2 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddFinalizers", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFinalizers indicates an expected call of AddFinalizers.
func (mr *MockFinalizerManagerMockRecorder) AddFinalizers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFinalizers", reflect.TypeOf((*MockFinalizerManager)(nil).AddFinalizers), varargs...)
}

// RemoveFinalizers mocks base method.
func (m *MockFinalizerManager) RemoveFinalizers(arg0 context.Context, arg1 client.Object, arg2 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveFinalizers", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFinalizers indicates an expected call of RemoveFinalizers.
func (mr *MockFinalizerManagerMockRecorder) RemoveFinalizers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFinalizers", reflect.TypeOf((*MockFinalizerManager)(nil).RemoveFinalizers), varargs...)
}
