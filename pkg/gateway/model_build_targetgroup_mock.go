// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/gateway/model_build_targetgroup.go

// Package gateway is a generated GoMock package.
package gateway

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/aws-application-networking-k8s/pkg/apis/applicationnetworking/v1alpha1"
	core "github.com/aws/aws-application-networking-k8s/pkg/model/core"
	lattice "github.com/aws/aws-application-networking-k8s/pkg/model/lattice"
	gomock "github.com/golang/mock/gomock"
)

// MockSvcExportTargetGroupModelBuilder is a mock of SvcExportTargetGroupModelBuilder interface.
type MockSvcExportTargetGroupModelBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockSvcExportTargetGroupModelBuilderMockRecorder
}

// MockSvcExportTargetGroupModelBuilderMockRecorder is the mock recorder for MockSvcExportTargetGroupModelBuilder.
type MockSvcExportTargetGroupModelBuilderMockRecorder struct {
	mock *MockSvcExportTargetGroupModelBuilder
}

// NewMockSvcExportTargetGroupModelBuilder creates a new mock instance.
func NewMockSvcExportTargetGroupModelBuilder(ctrl *gomock.Controller) *MockSvcExportTargetGroupModelBuilder {
	mock := &MockSvcExportTargetGroupModelBuilder{ctrl: ctrl}
	mock.recorder = &MockSvcExportTargetGroupModelBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSvcExportTargetGroupModelBuilder) EXPECT() *MockSvcExportTargetGroupModelBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockSvcExportTargetGroupModelBuilder) Build(ctx context.Context, svcExport *v1alpha1.ServiceExport) (core.Stack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", ctx, svcExport)
	ret0, _ := ret[0].(core.Stack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build.
func (mr *MockSvcExportTargetGroupModelBuilderMockRecorder) Build(ctx, svcExport interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockSvcExportTargetGroupModelBuilder)(nil).Build), ctx, svcExport)
}

// BuildTargetGroup mocks base method.
func (m *MockSvcExportTargetGroupModelBuilder) BuildTargetGroup(ctx context.Context, svcExport *v1alpha1.ServiceExport) (*lattice.TargetGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildTargetGroup", ctx, svcExport)
	ret0, _ := ret[0].(*lattice.TargetGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildTargetGroup indicates an expected call of BuildTargetGroup.
func (mr *MockSvcExportTargetGroupModelBuilderMockRecorder) BuildTargetGroup(ctx, svcExport interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildTargetGroup", reflect.TypeOf((*MockSvcExportTargetGroupModelBuilder)(nil).BuildTargetGroup), ctx, svcExport)
}

// MockBackendRefTargetGroupModelBuilder is a mock of BackendRefTargetGroupModelBuilder interface.
type MockBackendRefTargetGroupModelBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBackendRefTargetGroupModelBuilderMockRecorder
}

// MockBackendRefTargetGroupModelBuilderMockRecorder is the mock recorder for MockBackendRefTargetGroupModelBuilder.
type MockBackendRefTargetGroupModelBuilderMockRecorder struct {
	mock *MockBackendRefTargetGroupModelBuilder
}

// NewMockBackendRefTargetGroupModelBuilder creates a new mock instance.
func NewMockBackendRefTargetGroupModelBuilder(ctrl *gomock.Controller) *MockBackendRefTargetGroupModelBuilder {
	mock := &MockBackendRefTargetGroupModelBuilder{ctrl: ctrl}
	mock.recorder = &MockBackendRefTargetGroupModelBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendRefTargetGroupModelBuilder) EXPECT() *MockBackendRefTargetGroupModelBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockBackendRefTargetGroupModelBuilder) Build(ctx context.Context, route core.Route, backendRef core.BackendRef, stack core.Stack) (core.Stack, *lattice.TargetGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", ctx, route, backendRef, stack)
	ret0, _ := ret[0].(core.Stack)
	ret1, _ := ret[1].(*lattice.TargetGroup)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Build indicates an expected call of Build.
func (mr *MockBackendRefTargetGroupModelBuilderMockRecorder) Build(ctx, route, backendRef, stack interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockBackendRefTargetGroupModelBuilder)(nil).Build), ctx, route, backendRef, stack)
}
