// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/k8s/batch/v1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "k8s.io/api/batch/v1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterJobReconciler is a mock of MulticlusterJobReconciler interface.
type MockMulticlusterJobReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterJobReconcilerMockRecorder
}

// MockMulticlusterJobReconcilerMockRecorder is the mock recorder for MockMulticlusterJobReconciler.
type MockMulticlusterJobReconcilerMockRecorder struct {
	mock *MockMulticlusterJobReconciler
}

// NewMockMulticlusterJobReconciler creates a new mock instance.
func NewMockMulticlusterJobReconciler(ctrl *gomock.Controller) *MockMulticlusterJobReconciler {
	mock := &MockMulticlusterJobReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterJobReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterJobReconciler) EXPECT() *MockMulticlusterJobReconcilerMockRecorder {
	return m.recorder
}

// ReconcileJob mocks base method.
func (m *MockMulticlusterJobReconciler) ReconcileJob(clusterName string, obj *v1.Job) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileJob", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileJob indicates an expected call of ReconcileJob.
func (mr *MockMulticlusterJobReconcilerMockRecorder) ReconcileJob(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileJob", reflect.TypeOf((*MockMulticlusterJobReconciler)(nil).ReconcileJob), clusterName, obj)
}

// MockMulticlusterJobDeletionReconciler is a mock of MulticlusterJobDeletionReconciler interface.
type MockMulticlusterJobDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterJobDeletionReconcilerMockRecorder
}

// MockMulticlusterJobDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterJobDeletionReconciler.
type MockMulticlusterJobDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterJobDeletionReconciler
}

// NewMockMulticlusterJobDeletionReconciler creates a new mock instance.
func NewMockMulticlusterJobDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterJobDeletionReconciler {
	mock := &MockMulticlusterJobDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterJobDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterJobDeletionReconciler) EXPECT() *MockMulticlusterJobDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileJobDeletion mocks base method.
func (m *MockMulticlusterJobDeletionReconciler) ReconcileJobDeletion(clusterName string, req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileJobDeletion", clusterName, req)
}

// ReconcileJobDeletion indicates an expected call of ReconcileJobDeletion.
func (mr *MockMulticlusterJobDeletionReconcilerMockRecorder) ReconcileJobDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileJobDeletion", reflect.TypeOf((*MockMulticlusterJobDeletionReconciler)(nil).ReconcileJobDeletion), clusterName, req)
}

// MockMulticlusterJobReconcileLoop is a mock of MulticlusterJobReconcileLoop interface.
type MockMulticlusterJobReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterJobReconcileLoopMockRecorder
}

// MockMulticlusterJobReconcileLoopMockRecorder is the mock recorder for MockMulticlusterJobReconcileLoop.
type MockMulticlusterJobReconcileLoopMockRecorder struct {
	mock *MockMulticlusterJobReconcileLoop
}

// NewMockMulticlusterJobReconcileLoop creates a new mock instance.
func NewMockMulticlusterJobReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterJobReconcileLoop {
	mock := &MockMulticlusterJobReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterJobReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterJobReconcileLoop) EXPECT() *MockMulticlusterJobReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterJobReconciler mocks base method.
func (m *MockMulticlusterJobReconcileLoop) AddMulticlusterJobReconciler(ctx context.Context, rec controller.MulticlusterJobReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterJobReconciler", varargs...)
}

// AddMulticlusterJobReconciler indicates an expected call of AddMulticlusterJobReconciler.
func (mr *MockMulticlusterJobReconcileLoopMockRecorder) AddMulticlusterJobReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterJobReconciler", reflect.TypeOf((*MockMulticlusterJobReconcileLoop)(nil).AddMulticlusterJobReconciler), varargs...)
}