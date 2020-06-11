// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "k8s.io/api/apps/v1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockDeploymentReconciler is a mock of DeploymentReconciler interface.
type MockDeploymentReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentReconcilerMockRecorder
}

// MockDeploymentReconcilerMockRecorder is the mock recorder for MockDeploymentReconciler.
type MockDeploymentReconcilerMockRecorder struct {
	mock *MockDeploymentReconciler
}

// NewMockDeploymentReconciler creates a new mock instance.
func NewMockDeploymentReconciler(ctrl *gomock.Controller) *MockDeploymentReconciler {
	mock := &MockDeploymentReconciler{ctrl: ctrl}
	mock.recorder = &MockDeploymentReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentReconciler) EXPECT() *MockDeploymentReconcilerMockRecorder {
	return m.recorder
}

// ReconcileDeployment mocks base method.
func (m *MockDeploymentReconciler) ReconcileDeployment(obj *v1.Deployment) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDeployment", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileDeployment indicates an expected call of ReconcileDeployment.
func (mr *MockDeploymentReconcilerMockRecorder) ReconcileDeployment(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDeployment", reflect.TypeOf((*MockDeploymentReconciler)(nil).ReconcileDeployment), obj)
}

// MockDeploymentDeletionReconciler is a mock of DeploymentDeletionReconciler interface.
type MockDeploymentDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentDeletionReconcilerMockRecorder
}

// MockDeploymentDeletionReconcilerMockRecorder is the mock recorder for MockDeploymentDeletionReconciler.
type MockDeploymentDeletionReconcilerMockRecorder struct {
	mock *MockDeploymentDeletionReconciler
}

// NewMockDeploymentDeletionReconciler creates a new mock instance.
func NewMockDeploymentDeletionReconciler(ctrl *gomock.Controller) *MockDeploymentDeletionReconciler {
	mock := &MockDeploymentDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockDeploymentDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentDeletionReconciler) EXPECT() *MockDeploymentDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileDeploymentDeletion mocks base method.
func (m *MockDeploymentDeletionReconciler) ReconcileDeploymentDeletion(req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileDeploymentDeletion", req)
}

// ReconcileDeploymentDeletion indicates an expected call of ReconcileDeploymentDeletion.
func (mr *MockDeploymentDeletionReconcilerMockRecorder) ReconcileDeploymentDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDeploymentDeletion", reflect.TypeOf((*MockDeploymentDeletionReconciler)(nil).ReconcileDeploymentDeletion), req)
}

// MockDeploymentFinalizer is a mock of DeploymentFinalizer interface.
type MockDeploymentFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentFinalizerMockRecorder
}

// MockDeploymentFinalizerMockRecorder is the mock recorder for MockDeploymentFinalizer.
type MockDeploymentFinalizerMockRecorder struct {
	mock *MockDeploymentFinalizer
}

// NewMockDeploymentFinalizer creates a new mock instance.
func NewMockDeploymentFinalizer(ctrl *gomock.Controller) *MockDeploymentFinalizer {
	mock := &MockDeploymentFinalizer{ctrl: ctrl}
	mock.recorder = &MockDeploymentFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentFinalizer) EXPECT() *MockDeploymentFinalizerMockRecorder {
	return m.recorder
}

// ReconcileDeployment mocks base method.
func (m *MockDeploymentFinalizer) ReconcileDeployment(obj *v1.Deployment) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDeployment", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileDeployment indicates an expected call of ReconcileDeployment.
func (mr *MockDeploymentFinalizerMockRecorder) ReconcileDeployment(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDeployment", reflect.TypeOf((*MockDeploymentFinalizer)(nil).ReconcileDeployment), obj)
}

// DeploymentFinalizerName mocks base method.
func (m *MockDeploymentFinalizer) DeploymentFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeploymentFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// DeploymentFinalizerName indicates an expected call of DeploymentFinalizerName.
func (mr *MockDeploymentFinalizerMockRecorder) DeploymentFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentFinalizerName", reflect.TypeOf((*MockDeploymentFinalizer)(nil).DeploymentFinalizerName))
}

// FinalizeDeployment mocks base method.
func (m *MockDeploymentFinalizer) FinalizeDeployment(obj *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeDeployment", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeDeployment indicates an expected call of FinalizeDeployment.
func (mr *MockDeploymentFinalizerMockRecorder) FinalizeDeployment(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeDeployment", reflect.TypeOf((*MockDeploymentFinalizer)(nil).FinalizeDeployment), obj)
}

// MockDeploymentReconcileLoop is a mock of DeploymentReconcileLoop interface.
type MockDeploymentReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentReconcileLoopMockRecorder
}

// MockDeploymentReconcileLoopMockRecorder is the mock recorder for MockDeploymentReconcileLoop.
type MockDeploymentReconcileLoopMockRecorder struct {
	mock *MockDeploymentReconcileLoop
}

// NewMockDeploymentReconcileLoop creates a new mock instance.
func NewMockDeploymentReconcileLoop(ctrl *gomock.Controller) *MockDeploymentReconcileLoop {
	mock := &MockDeploymentReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockDeploymentReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentReconcileLoop) EXPECT() *MockDeploymentReconcileLoopMockRecorder {
	return m.recorder
}

// RunDeploymentReconciler mocks base method.
func (m *MockDeploymentReconcileLoop) RunDeploymentReconciler(ctx context.Context, rec controller.DeploymentReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunDeploymentReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunDeploymentReconciler indicates an expected call of RunDeploymentReconciler.
func (mr *MockDeploymentReconcileLoopMockRecorder) RunDeploymentReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunDeploymentReconciler", reflect.TypeOf((*MockDeploymentReconcileLoop)(nil).RunDeploymentReconciler), varargs...)
}

// MockReplicaSetReconciler is a mock of ReplicaSetReconciler interface.
type MockReplicaSetReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockReplicaSetReconcilerMockRecorder
}

// MockReplicaSetReconcilerMockRecorder is the mock recorder for MockReplicaSetReconciler.
type MockReplicaSetReconcilerMockRecorder struct {
	mock *MockReplicaSetReconciler
}

// NewMockReplicaSetReconciler creates a new mock instance.
func NewMockReplicaSetReconciler(ctrl *gomock.Controller) *MockReplicaSetReconciler {
	mock := &MockReplicaSetReconciler{ctrl: ctrl}
	mock.recorder = &MockReplicaSetReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicaSetReconciler) EXPECT() *MockReplicaSetReconcilerMockRecorder {
	return m.recorder
}

// ReconcileReplicaSet mocks base method.
func (m *MockReplicaSetReconciler) ReconcileReplicaSet(obj *v1.ReplicaSet) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileReplicaSet", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileReplicaSet indicates an expected call of ReconcileReplicaSet.
func (mr *MockReplicaSetReconcilerMockRecorder) ReconcileReplicaSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileReplicaSet", reflect.TypeOf((*MockReplicaSetReconciler)(nil).ReconcileReplicaSet), obj)
}

// MockReplicaSetDeletionReconciler is a mock of ReplicaSetDeletionReconciler interface.
type MockReplicaSetDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockReplicaSetDeletionReconcilerMockRecorder
}

// MockReplicaSetDeletionReconcilerMockRecorder is the mock recorder for MockReplicaSetDeletionReconciler.
type MockReplicaSetDeletionReconcilerMockRecorder struct {
	mock *MockReplicaSetDeletionReconciler
}

// NewMockReplicaSetDeletionReconciler creates a new mock instance.
func NewMockReplicaSetDeletionReconciler(ctrl *gomock.Controller) *MockReplicaSetDeletionReconciler {
	mock := &MockReplicaSetDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockReplicaSetDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicaSetDeletionReconciler) EXPECT() *MockReplicaSetDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileReplicaSetDeletion mocks base method.
func (m *MockReplicaSetDeletionReconciler) ReconcileReplicaSetDeletion(req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileReplicaSetDeletion", req)
}

// ReconcileReplicaSetDeletion indicates an expected call of ReconcileReplicaSetDeletion.
func (mr *MockReplicaSetDeletionReconcilerMockRecorder) ReconcileReplicaSetDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileReplicaSetDeletion", reflect.TypeOf((*MockReplicaSetDeletionReconciler)(nil).ReconcileReplicaSetDeletion), req)
}

// MockReplicaSetFinalizer is a mock of ReplicaSetFinalizer interface.
type MockReplicaSetFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockReplicaSetFinalizerMockRecorder
}

// MockReplicaSetFinalizerMockRecorder is the mock recorder for MockReplicaSetFinalizer.
type MockReplicaSetFinalizerMockRecorder struct {
	mock *MockReplicaSetFinalizer
}

// NewMockReplicaSetFinalizer creates a new mock instance.
func NewMockReplicaSetFinalizer(ctrl *gomock.Controller) *MockReplicaSetFinalizer {
	mock := &MockReplicaSetFinalizer{ctrl: ctrl}
	mock.recorder = &MockReplicaSetFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicaSetFinalizer) EXPECT() *MockReplicaSetFinalizerMockRecorder {
	return m.recorder
}

// ReconcileReplicaSet mocks base method.
func (m *MockReplicaSetFinalizer) ReconcileReplicaSet(obj *v1.ReplicaSet) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileReplicaSet", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileReplicaSet indicates an expected call of ReconcileReplicaSet.
func (mr *MockReplicaSetFinalizerMockRecorder) ReconcileReplicaSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileReplicaSet", reflect.TypeOf((*MockReplicaSetFinalizer)(nil).ReconcileReplicaSet), obj)
}

// ReplicaSetFinalizerName mocks base method.
func (m *MockReplicaSetFinalizer) ReplicaSetFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicaSetFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ReplicaSetFinalizerName indicates an expected call of ReplicaSetFinalizerName.
func (mr *MockReplicaSetFinalizerMockRecorder) ReplicaSetFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicaSetFinalizerName", reflect.TypeOf((*MockReplicaSetFinalizer)(nil).ReplicaSetFinalizerName))
}

// FinalizeReplicaSet mocks base method.
func (m *MockReplicaSetFinalizer) FinalizeReplicaSet(obj *v1.ReplicaSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeReplicaSet", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeReplicaSet indicates an expected call of FinalizeReplicaSet.
func (mr *MockReplicaSetFinalizerMockRecorder) FinalizeReplicaSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeReplicaSet", reflect.TypeOf((*MockReplicaSetFinalizer)(nil).FinalizeReplicaSet), obj)
}

// MockReplicaSetReconcileLoop is a mock of ReplicaSetReconcileLoop interface.
type MockReplicaSetReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockReplicaSetReconcileLoopMockRecorder
}

// MockReplicaSetReconcileLoopMockRecorder is the mock recorder for MockReplicaSetReconcileLoop.
type MockReplicaSetReconcileLoopMockRecorder struct {
	mock *MockReplicaSetReconcileLoop
}

// NewMockReplicaSetReconcileLoop creates a new mock instance.
func NewMockReplicaSetReconcileLoop(ctrl *gomock.Controller) *MockReplicaSetReconcileLoop {
	mock := &MockReplicaSetReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockReplicaSetReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicaSetReconcileLoop) EXPECT() *MockReplicaSetReconcileLoopMockRecorder {
	return m.recorder
}

// RunReplicaSetReconciler mocks base method.
func (m *MockReplicaSetReconcileLoop) RunReplicaSetReconciler(ctx context.Context, rec controller.ReplicaSetReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunReplicaSetReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunReplicaSetReconciler indicates an expected call of RunReplicaSetReconciler.
func (mr *MockReplicaSetReconcileLoopMockRecorder) RunReplicaSetReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunReplicaSetReconciler", reflect.TypeOf((*MockReplicaSetReconcileLoop)(nil).RunReplicaSetReconciler), varargs...)
}

// MockDaemonSetReconciler is a mock of DaemonSetReconciler interface.
type MockDaemonSetReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockDaemonSetReconcilerMockRecorder
}

// MockDaemonSetReconcilerMockRecorder is the mock recorder for MockDaemonSetReconciler.
type MockDaemonSetReconcilerMockRecorder struct {
	mock *MockDaemonSetReconciler
}

// NewMockDaemonSetReconciler creates a new mock instance.
func NewMockDaemonSetReconciler(ctrl *gomock.Controller) *MockDaemonSetReconciler {
	mock := &MockDaemonSetReconciler{ctrl: ctrl}
	mock.recorder = &MockDaemonSetReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDaemonSetReconciler) EXPECT() *MockDaemonSetReconcilerMockRecorder {
	return m.recorder
}

// ReconcileDaemonSet mocks base method.
func (m *MockDaemonSetReconciler) ReconcileDaemonSet(obj *v1.DaemonSet) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDaemonSet", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileDaemonSet indicates an expected call of ReconcileDaemonSet.
func (mr *MockDaemonSetReconcilerMockRecorder) ReconcileDaemonSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDaemonSet", reflect.TypeOf((*MockDaemonSetReconciler)(nil).ReconcileDaemonSet), obj)
}

// MockDaemonSetDeletionReconciler is a mock of DaemonSetDeletionReconciler interface.
type MockDaemonSetDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockDaemonSetDeletionReconcilerMockRecorder
}

// MockDaemonSetDeletionReconcilerMockRecorder is the mock recorder for MockDaemonSetDeletionReconciler.
type MockDaemonSetDeletionReconcilerMockRecorder struct {
	mock *MockDaemonSetDeletionReconciler
}

// NewMockDaemonSetDeletionReconciler creates a new mock instance.
func NewMockDaemonSetDeletionReconciler(ctrl *gomock.Controller) *MockDaemonSetDeletionReconciler {
	mock := &MockDaemonSetDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockDaemonSetDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDaemonSetDeletionReconciler) EXPECT() *MockDaemonSetDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileDaemonSetDeletion mocks base method.
func (m *MockDaemonSetDeletionReconciler) ReconcileDaemonSetDeletion(req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileDaemonSetDeletion", req)
}

// ReconcileDaemonSetDeletion indicates an expected call of ReconcileDaemonSetDeletion.
func (mr *MockDaemonSetDeletionReconcilerMockRecorder) ReconcileDaemonSetDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDaemonSetDeletion", reflect.TypeOf((*MockDaemonSetDeletionReconciler)(nil).ReconcileDaemonSetDeletion), req)
}

// MockDaemonSetFinalizer is a mock of DaemonSetFinalizer interface.
type MockDaemonSetFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockDaemonSetFinalizerMockRecorder
}

// MockDaemonSetFinalizerMockRecorder is the mock recorder for MockDaemonSetFinalizer.
type MockDaemonSetFinalizerMockRecorder struct {
	mock *MockDaemonSetFinalizer
}

// NewMockDaemonSetFinalizer creates a new mock instance.
func NewMockDaemonSetFinalizer(ctrl *gomock.Controller) *MockDaemonSetFinalizer {
	mock := &MockDaemonSetFinalizer{ctrl: ctrl}
	mock.recorder = &MockDaemonSetFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDaemonSetFinalizer) EXPECT() *MockDaemonSetFinalizerMockRecorder {
	return m.recorder
}

// ReconcileDaemonSet mocks base method.
func (m *MockDaemonSetFinalizer) ReconcileDaemonSet(obj *v1.DaemonSet) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDaemonSet", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileDaemonSet indicates an expected call of ReconcileDaemonSet.
func (mr *MockDaemonSetFinalizerMockRecorder) ReconcileDaemonSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDaemonSet", reflect.TypeOf((*MockDaemonSetFinalizer)(nil).ReconcileDaemonSet), obj)
}

// DaemonSetFinalizerName mocks base method.
func (m *MockDaemonSetFinalizer) DaemonSetFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DaemonSetFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// DaemonSetFinalizerName indicates an expected call of DaemonSetFinalizerName.
func (mr *MockDaemonSetFinalizerMockRecorder) DaemonSetFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DaemonSetFinalizerName", reflect.TypeOf((*MockDaemonSetFinalizer)(nil).DaemonSetFinalizerName))
}

// FinalizeDaemonSet mocks base method.
func (m *MockDaemonSetFinalizer) FinalizeDaemonSet(obj *v1.DaemonSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeDaemonSet", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeDaemonSet indicates an expected call of FinalizeDaemonSet.
func (mr *MockDaemonSetFinalizerMockRecorder) FinalizeDaemonSet(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeDaemonSet", reflect.TypeOf((*MockDaemonSetFinalizer)(nil).FinalizeDaemonSet), obj)
}

// MockDaemonSetReconcileLoop is a mock of DaemonSetReconcileLoop interface.
type MockDaemonSetReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockDaemonSetReconcileLoopMockRecorder
}

// MockDaemonSetReconcileLoopMockRecorder is the mock recorder for MockDaemonSetReconcileLoop.
type MockDaemonSetReconcileLoopMockRecorder struct {
	mock *MockDaemonSetReconcileLoop
}

// NewMockDaemonSetReconcileLoop creates a new mock instance.
func NewMockDaemonSetReconcileLoop(ctrl *gomock.Controller) *MockDaemonSetReconcileLoop {
	mock := &MockDaemonSetReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockDaemonSetReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDaemonSetReconcileLoop) EXPECT() *MockDaemonSetReconcileLoopMockRecorder {
	return m.recorder
}

// RunDaemonSetReconciler mocks base method.
func (m *MockDaemonSetReconcileLoop) RunDaemonSetReconciler(ctx context.Context, rec controller.DaemonSetReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunDaemonSetReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunDaemonSetReconciler indicates an expected call of RunDaemonSetReconciler.
func (mr *MockDaemonSetReconcileLoopMockRecorder) RunDaemonSetReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunDaemonSetReconciler", reflect.TypeOf((*MockDaemonSetReconcileLoop)(nil).RunDaemonSetReconciler), varargs...)
}