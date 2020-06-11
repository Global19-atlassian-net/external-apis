// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "k8s.io/api/core/v1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterSecretReconciler is a mock of MulticlusterSecretReconciler interface.
type MockMulticlusterSecretReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterSecretReconcilerMockRecorder
}

// MockMulticlusterSecretReconcilerMockRecorder is the mock recorder for MockMulticlusterSecretReconciler.
type MockMulticlusterSecretReconcilerMockRecorder struct {
	mock *MockMulticlusterSecretReconciler
}

// NewMockMulticlusterSecretReconciler creates a new mock instance.
func NewMockMulticlusterSecretReconciler(ctrl *gomock.Controller) *MockMulticlusterSecretReconciler {
	mock := &MockMulticlusterSecretReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterSecretReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterSecretReconciler) EXPECT() *MockMulticlusterSecretReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSecret mocks base method.
func (m *MockMulticlusterSecretReconciler) ReconcileSecret(clusterName string, obj *v1.Secret) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSecret", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileSecret indicates an expected call of ReconcileSecret.
func (mr *MockMulticlusterSecretReconcilerMockRecorder) ReconcileSecret(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSecret", reflect.TypeOf((*MockMulticlusterSecretReconciler)(nil).ReconcileSecret), clusterName, obj)
}

// MockMulticlusterSecretDeletionReconciler is a mock of MulticlusterSecretDeletionReconciler interface.
type MockMulticlusterSecretDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterSecretDeletionReconcilerMockRecorder
}

// MockMulticlusterSecretDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterSecretDeletionReconciler.
type MockMulticlusterSecretDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterSecretDeletionReconciler
}

// NewMockMulticlusterSecretDeletionReconciler creates a new mock instance.
func NewMockMulticlusterSecretDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterSecretDeletionReconciler {
	mock := &MockMulticlusterSecretDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterSecretDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterSecretDeletionReconciler) EXPECT() *MockMulticlusterSecretDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSecretDeletion mocks base method.
func (m *MockMulticlusterSecretDeletionReconciler) ReconcileSecretDeletion(clusterName string, req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileSecretDeletion", clusterName, req)
}

// ReconcileSecretDeletion indicates an expected call of ReconcileSecretDeletion.
func (mr *MockMulticlusterSecretDeletionReconcilerMockRecorder) ReconcileSecretDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSecretDeletion", reflect.TypeOf((*MockMulticlusterSecretDeletionReconciler)(nil).ReconcileSecretDeletion), clusterName, req)
}

// MockMulticlusterSecretReconcileLoop is a mock of MulticlusterSecretReconcileLoop interface.
type MockMulticlusterSecretReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterSecretReconcileLoopMockRecorder
}

// MockMulticlusterSecretReconcileLoopMockRecorder is the mock recorder for MockMulticlusterSecretReconcileLoop.
type MockMulticlusterSecretReconcileLoopMockRecorder struct {
	mock *MockMulticlusterSecretReconcileLoop
}

// NewMockMulticlusterSecretReconcileLoop creates a new mock instance.
func NewMockMulticlusterSecretReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterSecretReconcileLoop {
	mock := &MockMulticlusterSecretReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterSecretReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterSecretReconcileLoop) EXPECT() *MockMulticlusterSecretReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterSecretReconciler mocks base method.
func (m *MockMulticlusterSecretReconcileLoop) AddMulticlusterSecretReconciler(ctx context.Context, rec controller.MulticlusterSecretReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterSecretReconciler", varargs...)
}

// AddMulticlusterSecretReconciler indicates an expected call of AddMulticlusterSecretReconciler.
func (mr *MockMulticlusterSecretReconcileLoopMockRecorder) AddMulticlusterSecretReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterSecretReconciler", reflect.TypeOf((*MockMulticlusterSecretReconcileLoop)(nil).AddMulticlusterSecretReconciler), varargs...)
}

// MockMulticlusterServiceAccountReconciler is a mock of MulticlusterServiceAccountReconciler interface.
type MockMulticlusterServiceAccountReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterServiceAccountReconcilerMockRecorder
}

// MockMulticlusterServiceAccountReconcilerMockRecorder is the mock recorder for MockMulticlusterServiceAccountReconciler.
type MockMulticlusterServiceAccountReconcilerMockRecorder struct {
	mock *MockMulticlusterServiceAccountReconciler
}

// NewMockMulticlusterServiceAccountReconciler creates a new mock instance.
func NewMockMulticlusterServiceAccountReconciler(ctrl *gomock.Controller) *MockMulticlusterServiceAccountReconciler {
	mock := &MockMulticlusterServiceAccountReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterServiceAccountReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterServiceAccountReconciler) EXPECT() *MockMulticlusterServiceAccountReconcilerMockRecorder {
	return m.recorder
}

// ReconcileServiceAccount mocks base method.
func (m *MockMulticlusterServiceAccountReconciler) ReconcileServiceAccount(clusterName string, obj *v1.ServiceAccount) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileServiceAccount", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileServiceAccount indicates an expected call of ReconcileServiceAccount.
func (mr *MockMulticlusterServiceAccountReconcilerMockRecorder) ReconcileServiceAccount(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileServiceAccount", reflect.TypeOf((*MockMulticlusterServiceAccountReconciler)(nil).ReconcileServiceAccount), clusterName, obj)
}

// MockMulticlusterServiceAccountDeletionReconciler is a mock of MulticlusterServiceAccountDeletionReconciler interface.
type MockMulticlusterServiceAccountDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterServiceAccountDeletionReconcilerMockRecorder
}

// MockMulticlusterServiceAccountDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterServiceAccountDeletionReconciler.
type MockMulticlusterServiceAccountDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterServiceAccountDeletionReconciler
}

// NewMockMulticlusterServiceAccountDeletionReconciler creates a new mock instance.
func NewMockMulticlusterServiceAccountDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterServiceAccountDeletionReconciler {
	mock := &MockMulticlusterServiceAccountDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterServiceAccountDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterServiceAccountDeletionReconciler) EXPECT() *MockMulticlusterServiceAccountDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileServiceAccountDeletion mocks base method.
func (m *MockMulticlusterServiceAccountDeletionReconciler) ReconcileServiceAccountDeletion(clusterName string, req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileServiceAccountDeletion", clusterName, req)
}

// ReconcileServiceAccountDeletion indicates an expected call of ReconcileServiceAccountDeletion.
func (mr *MockMulticlusterServiceAccountDeletionReconcilerMockRecorder) ReconcileServiceAccountDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileServiceAccountDeletion", reflect.TypeOf((*MockMulticlusterServiceAccountDeletionReconciler)(nil).ReconcileServiceAccountDeletion), clusterName, req)
}

// MockMulticlusterServiceAccountReconcileLoop is a mock of MulticlusterServiceAccountReconcileLoop interface.
type MockMulticlusterServiceAccountReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterServiceAccountReconcileLoopMockRecorder
}

// MockMulticlusterServiceAccountReconcileLoopMockRecorder is the mock recorder for MockMulticlusterServiceAccountReconcileLoop.
type MockMulticlusterServiceAccountReconcileLoopMockRecorder struct {
	mock *MockMulticlusterServiceAccountReconcileLoop
}

// NewMockMulticlusterServiceAccountReconcileLoop creates a new mock instance.
func NewMockMulticlusterServiceAccountReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterServiceAccountReconcileLoop {
	mock := &MockMulticlusterServiceAccountReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterServiceAccountReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterServiceAccountReconcileLoop) EXPECT() *MockMulticlusterServiceAccountReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterServiceAccountReconciler mocks base method.
func (m *MockMulticlusterServiceAccountReconcileLoop) AddMulticlusterServiceAccountReconciler(ctx context.Context, rec controller.MulticlusterServiceAccountReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterServiceAccountReconciler", varargs...)
}

// AddMulticlusterServiceAccountReconciler indicates an expected call of AddMulticlusterServiceAccountReconciler.
func (mr *MockMulticlusterServiceAccountReconcileLoopMockRecorder) AddMulticlusterServiceAccountReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterServiceAccountReconciler", reflect.TypeOf((*MockMulticlusterServiceAccountReconcileLoop)(nil).AddMulticlusterServiceAccountReconciler), varargs...)
}

// MockMulticlusterConfigMapReconciler is a mock of MulticlusterConfigMapReconciler interface.
type MockMulticlusterConfigMapReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterConfigMapReconcilerMockRecorder
}

// MockMulticlusterConfigMapReconcilerMockRecorder is the mock recorder for MockMulticlusterConfigMapReconciler.
type MockMulticlusterConfigMapReconcilerMockRecorder struct {
	mock *MockMulticlusterConfigMapReconciler
}

// NewMockMulticlusterConfigMapReconciler creates a new mock instance.
func NewMockMulticlusterConfigMapReconciler(ctrl *gomock.Controller) *MockMulticlusterConfigMapReconciler {
	mock := &MockMulticlusterConfigMapReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterConfigMapReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterConfigMapReconciler) EXPECT() *MockMulticlusterConfigMapReconcilerMockRecorder {
	return m.recorder
}

// ReconcileConfigMap mocks base method.
func (m *MockMulticlusterConfigMapReconciler) ReconcileConfigMap(clusterName string, obj *v1.ConfigMap) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileConfigMap", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileConfigMap indicates an expected call of ReconcileConfigMap.
func (mr *MockMulticlusterConfigMapReconcilerMockRecorder) ReconcileConfigMap(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileConfigMap", reflect.TypeOf((*MockMulticlusterConfigMapReconciler)(nil).ReconcileConfigMap), clusterName, obj)
}

// MockMulticlusterConfigMapDeletionReconciler is a mock of MulticlusterConfigMapDeletionReconciler interface.
type MockMulticlusterConfigMapDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterConfigMapDeletionReconcilerMockRecorder
}

// MockMulticlusterConfigMapDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterConfigMapDeletionReconciler.
type MockMulticlusterConfigMapDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterConfigMapDeletionReconciler
}

// NewMockMulticlusterConfigMapDeletionReconciler creates a new mock instance.
func NewMockMulticlusterConfigMapDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterConfigMapDeletionReconciler {
	mock := &MockMulticlusterConfigMapDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterConfigMapDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterConfigMapDeletionReconciler) EXPECT() *MockMulticlusterConfigMapDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileConfigMapDeletion mocks base method.
func (m *MockMulticlusterConfigMapDeletionReconciler) ReconcileConfigMapDeletion(clusterName string, req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileConfigMapDeletion", clusterName, req)
}

// ReconcileConfigMapDeletion indicates an expected call of ReconcileConfigMapDeletion.
func (mr *MockMulticlusterConfigMapDeletionReconcilerMockRecorder) ReconcileConfigMapDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileConfigMapDeletion", reflect.TypeOf((*MockMulticlusterConfigMapDeletionReconciler)(nil).ReconcileConfigMapDeletion), clusterName, req)
}

// MockMulticlusterConfigMapReconcileLoop is a mock of MulticlusterConfigMapReconcileLoop interface.
type MockMulticlusterConfigMapReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterConfigMapReconcileLoopMockRecorder
}

// MockMulticlusterConfigMapReconcileLoopMockRecorder is the mock recorder for MockMulticlusterConfigMapReconcileLoop.
type MockMulticlusterConfigMapReconcileLoopMockRecorder struct {
	mock *MockMulticlusterConfigMapReconcileLoop
}

// NewMockMulticlusterConfigMapReconcileLoop creates a new mock instance.
func NewMockMulticlusterConfigMapReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterConfigMapReconcileLoop {
	mock := &MockMulticlusterConfigMapReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterConfigMapReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterConfigMapReconcileLoop) EXPECT() *MockMulticlusterConfigMapReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterConfigMapReconciler mocks base method.
func (m *MockMulticlusterConfigMapReconcileLoop) AddMulticlusterConfigMapReconciler(ctx context.Context, rec controller.MulticlusterConfigMapReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterConfigMapReconciler", varargs...)
}

// AddMulticlusterConfigMapReconciler indicates an expected call of AddMulticlusterConfigMapReconciler.
func (mr *MockMulticlusterConfigMapReconcileLoopMockRecorder) AddMulticlusterConfigMapReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterConfigMapReconciler", reflect.TypeOf((*MockMulticlusterConfigMapReconcileLoop)(nil).AddMulticlusterConfigMapReconciler), varargs...)
}

// MockMulticlusterServiceReconciler is a mock of MulticlusterServiceReconciler interface.
type MockMulticlusterServiceReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterServiceReconcilerMockRecorder
}

// MockMulticlusterServiceReconcilerMockRecorder is the mock recorder for MockMulticlusterServiceReconciler.
type MockMulticlusterServiceReconcilerMockRecorder struct {
	mock *MockMulticlusterServiceReconciler
}

// NewMockMulticlusterServiceReconciler creates a new mock instance.
func NewMockMulticlusterServiceReconciler(ctrl *gomock.Controller) *MockMulticlusterServiceReconciler {
	mock := &MockMulticlusterServiceReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterServiceReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterServiceReconciler) EXPECT() *MockMulticlusterServiceReconcilerMockRecorder {
	return m.recorder
}

// ReconcileService mocks base method.
func (m *MockMulticlusterServiceReconciler) ReconcileService(clusterName string, obj *v1.Service) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileService", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileService indicates an expected call of ReconcileService.
func (mr *MockMulticlusterServiceReconcilerMockRecorder) ReconcileService(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileService", reflect.TypeOf((*MockMulticlusterServiceReconciler)(nil).ReconcileService), clusterName, obj)
}

// MockMulticlusterServiceDeletionReconciler is a mock of MulticlusterServiceDeletionReconciler interface.
type MockMulticlusterServiceDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterServiceDeletionReconcilerMockRecorder
}

// MockMulticlusterServiceDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterServiceDeletionReconciler.
type MockMulticlusterServiceDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterServiceDeletionReconciler
}

// NewMockMulticlusterServiceDeletionReconciler creates a new mock instance.
func NewMockMulticlusterServiceDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterServiceDeletionReconciler {
	mock := &MockMulticlusterServiceDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterServiceDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterServiceDeletionReconciler) EXPECT() *MockMulticlusterServiceDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileServiceDeletion mocks base method.
func (m *MockMulticlusterServiceDeletionReconciler) ReconcileServiceDeletion(clusterName string, req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileServiceDeletion", clusterName, req)
}

// ReconcileServiceDeletion indicates an expected call of ReconcileServiceDeletion.
func (mr *MockMulticlusterServiceDeletionReconcilerMockRecorder) ReconcileServiceDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileServiceDeletion", reflect.TypeOf((*MockMulticlusterServiceDeletionReconciler)(nil).ReconcileServiceDeletion), clusterName, req)
}

// MockMulticlusterServiceReconcileLoop is a mock of MulticlusterServiceReconcileLoop interface.
type MockMulticlusterServiceReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterServiceReconcileLoopMockRecorder
}

// MockMulticlusterServiceReconcileLoopMockRecorder is the mock recorder for MockMulticlusterServiceReconcileLoop.
type MockMulticlusterServiceReconcileLoopMockRecorder struct {
	mock *MockMulticlusterServiceReconcileLoop
}

// NewMockMulticlusterServiceReconcileLoop creates a new mock instance.
func NewMockMulticlusterServiceReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterServiceReconcileLoop {
	mock := &MockMulticlusterServiceReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterServiceReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterServiceReconcileLoop) EXPECT() *MockMulticlusterServiceReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterServiceReconciler mocks base method.
func (m *MockMulticlusterServiceReconcileLoop) AddMulticlusterServiceReconciler(ctx context.Context, rec controller.MulticlusterServiceReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterServiceReconciler", varargs...)
}

// AddMulticlusterServiceReconciler indicates an expected call of AddMulticlusterServiceReconciler.
func (mr *MockMulticlusterServiceReconcileLoopMockRecorder) AddMulticlusterServiceReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterServiceReconciler", reflect.TypeOf((*MockMulticlusterServiceReconcileLoop)(nil).AddMulticlusterServiceReconciler), varargs...)
}

// MockMulticlusterPodReconciler is a mock of MulticlusterPodReconciler interface.
type MockMulticlusterPodReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterPodReconcilerMockRecorder
}

// MockMulticlusterPodReconcilerMockRecorder is the mock recorder for MockMulticlusterPodReconciler.
type MockMulticlusterPodReconcilerMockRecorder struct {
	mock *MockMulticlusterPodReconciler
}

// NewMockMulticlusterPodReconciler creates a new mock instance.
func NewMockMulticlusterPodReconciler(ctrl *gomock.Controller) *MockMulticlusterPodReconciler {
	mock := &MockMulticlusterPodReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterPodReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterPodReconciler) EXPECT() *MockMulticlusterPodReconcilerMockRecorder {
	return m.recorder
}

// ReconcilePod mocks base method.
func (m *MockMulticlusterPodReconciler) ReconcilePod(clusterName string, obj *v1.Pod) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePod", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcilePod indicates an expected call of ReconcilePod.
func (mr *MockMulticlusterPodReconcilerMockRecorder) ReconcilePod(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePod", reflect.TypeOf((*MockMulticlusterPodReconciler)(nil).ReconcilePod), clusterName, obj)
}

// MockMulticlusterPodDeletionReconciler is a mock of MulticlusterPodDeletionReconciler interface.
type MockMulticlusterPodDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterPodDeletionReconcilerMockRecorder
}

// MockMulticlusterPodDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterPodDeletionReconciler.
type MockMulticlusterPodDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterPodDeletionReconciler
}

// NewMockMulticlusterPodDeletionReconciler creates a new mock instance.
func NewMockMulticlusterPodDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterPodDeletionReconciler {
	mock := &MockMulticlusterPodDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterPodDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterPodDeletionReconciler) EXPECT() *MockMulticlusterPodDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcilePodDeletion mocks base method.
func (m *MockMulticlusterPodDeletionReconciler) ReconcilePodDeletion(clusterName string, req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcilePodDeletion", clusterName, req)
}

// ReconcilePodDeletion indicates an expected call of ReconcilePodDeletion.
func (mr *MockMulticlusterPodDeletionReconcilerMockRecorder) ReconcilePodDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePodDeletion", reflect.TypeOf((*MockMulticlusterPodDeletionReconciler)(nil).ReconcilePodDeletion), clusterName, req)
}

// MockMulticlusterPodReconcileLoop is a mock of MulticlusterPodReconcileLoop interface.
type MockMulticlusterPodReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterPodReconcileLoopMockRecorder
}

// MockMulticlusterPodReconcileLoopMockRecorder is the mock recorder for MockMulticlusterPodReconcileLoop.
type MockMulticlusterPodReconcileLoopMockRecorder struct {
	mock *MockMulticlusterPodReconcileLoop
}

// NewMockMulticlusterPodReconcileLoop creates a new mock instance.
func NewMockMulticlusterPodReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterPodReconcileLoop {
	mock := &MockMulticlusterPodReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterPodReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterPodReconcileLoop) EXPECT() *MockMulticlusterPodReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterPodReconciler mocks base method.
func (m *MockMulticlusterPodReconcileLoop) AddMulticlusterPodReconciler(ctx context.Context, rec controller.MulticlusterPodReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterPodReconciler", varargs...)
}

// AddMulticlusterPodReconciler indicates an expected call of AddMulticlusterPodReconciler.
func (mr *MockMulticlusterPodReconcileLoopMockRecorder) AddMulticlusterPodReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterPodReconciler", reflect.TypeOf((*MockMulticlusterPodReconcileLoop)(nil).AddMulticlusterPodReconciler), varargs...)
}

// MockMulticlusterNamespaceReconciler is a mock of MulticlusterNamespaceReconciler interface.
type MockMulticlusterNamespaceReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterNamespaceReconcilerMockRecorder
}

// MockMulticlusterNamespaceReconcilerMockRecorder is the mock recorder for MockMulticlusterNamespaceReconciler.
type MockMulticlusterNamespaceReconcilerMockRecorder struct {
	mock *MockMulticlusterNamespaceReconciler
}

// NewMockMulticlusterNamespaceReconciler creates a new mock instance.
func NewMockMulticlusterNamespaceReconciler(ctrl *gomock.Controller) *MockMulticlusterNamespaceReconciler {
	mock := &MockMulticlusterNamespaceReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterNamespaceReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterNamespaceReconciler) EXPECT() *MockMulticlusterNamespaceReconcilerMockRecorder {
	return m.recorder
}

// ReconcileNamespace mocks base method.
func (m *MockMulticlusterNamespaceReconciler) ReconcileNamespace(clusterName string, obj *v1.Namespace) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileNamespace", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileNamespace indicates an expected call of ReconcileNamespace.
func (mr *MockMulticlusterNamespaceReconcilerMockRecorder) ReconcileNamespace(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileNamespace", reflect.TypeOf((*MockMulticlusterNamespaceReconciler)(nil).ReconcileNamespace), clusterName, obj)
}

// MockMulticlusterNamespaceDeletionReconciler is a mock of MulticlusterNamespaceDeletionReconciler interface.
type MockMulticlusterNamespaceDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterNamespaceDeletionReconcilerMockRecorder
}

// MockMulticlusterNamespaceDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterNamespaceDeletionReconciler.
type MockMulticlusterNamespaceDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterNamespaceDeletionReconciler
}

// NewMockMulticlusterNamespaceDeletionReconciler creates a new mock instance.
func NewMockMulticlusterNamespaceDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterNamespaceDeletionReconciler {
	mock := &MockMulticlusterNamespaceDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterNamespaceDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterNamespaceDeletionReconciler) EXPECT() *MockMulticlusterNamespaceDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileNamespaceDeletion mocks base method.
func (m *MockMulticlusterNamespaceDeletionReconciler) ReconcileNamespaceDeletion(clusterName string, req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileNamespaceDeletion", clusterName, req)
}

// ReconcileNamespaceDeletion indicates an expected call of ReconcileNamespaceDeletion.
func (mr *MockMulticlusterNamespaceDeletionReconcilerMockRecorder) ReconcileNamespaceDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileNamespaceDeletion", reflect.TypeOf((*MockMulticlusterNamespaceDeletionReconciler)(nil).ReconcileNamespaceDeletion), clusterName, req)
}

// MockMulticlusterNamespaceReconcileLoop is a mock of MulticlusterNamespaceReconcileLoop interface.
type MockMulticlusterNamespaceReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterNamespaceReconcileLoopMockRecorder
}

// MockMulticlusterNamespaceReconcileLoopMockRecorder is the mock recorder for MockMulticlusterNamespaceReconcileLoop.
type MockMulticlusterNamespaceReconcileLoopMockRecorder struct {
	mock *MockMulticlusterNamespaceReconcileLoop
}

// NewMockMulticlusterNamespaceReconcileLoop creates a new mock instance.
func NewMockMulticlusterNamespaceReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterNamespaceReconcileLoop {
	mock := &MockMulticlusterNamespaceReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterNamespaceReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterNamespaceReconcileLoop) EXPECT() *MockMulticlusterNamespaceReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterNamespaceReconciler mocks base method.
func (m *MockMulticlusterNamespaceReconcileLoop) AddMulticlusterNamespaceReconciler(ctx context.Context, rec controller.MulticlusterNamespaceReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterNamespaceReconciler", varargs...)
}

// AddMulticlusterNamespaceReconciler indicates an expected call of AddMulticlusterNamespaceReconciler.
func (mr *MockMulticlusterNamespaceReconcileLoopMockRecorder) AddMulticlusterNamespaceReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterNamespaceReconciler", reflect.TypeOf((*MockMulticlusterNamespaceReconcileLoop)(nil).AddMulticlusterNamespaceReconciler), varargs...)
}

// MockMulticlusterNodeReconciler is a mock of MulticlusterNodeReconciler interface.
type MockMulticlusterNodeReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterNodeReconcilerMockRecorder
}

// MockMulticlusterNodeReconcilerMockRecorder is the mock recorder for MockMulticlusterNodeReconciler.
type MockMulticlusterNodeReconcilerMockRecorder struct {
	mock *MockMulticlusterNodeReconciler
}

// NewMockMulticlusterNodeReconciler creates a new mock instance.
func NewMockMulticlusterNodeReconciler(ctrl *gomock.Controller) *MockMulticlusterNodeReconciler {
	mock := &MockMulticlusterNodeReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterNodeReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterNodeReconciler) EXPECT() *MockMulticlusterNodeReconcilerMockRecorder {
	return m.recorder
}

// ReconcileNode mocks base method.
func (m *MockMulticlusterNodeReconciler) ReconcileNode(clusterName string, obj *v1.Node) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileNode", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileNode indicates an expected call of ReconcileNode.
func (mr *MockMulticlusterNodeReconcilerMockRecorder) ReconcileNode(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileNode", reflect.TypeOf((*MockMulticlusterNodeReconciler)(nil).ReconcileNode), clusterName, obj)
}

// MockMulticlusterNodeDeletionReconciler is a mock of MulticlusterNodeDeletionReconciler interface.
type MockMulticlusterNodeDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterNodeDeletionReconcilerMockRecorder
}

// MockMulticlusterNodeDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterNodeDeletionReconciler.
type MockMulticlusterNodeDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterNodeDeletionReconciler
}

// NewMockMulticlusterNodeDeletionReconciler creates a new mock instance.
func NewMockMulticlusterNodeDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterNodeDeletionReconciler {
	mock := &MockMulticlusterNodeDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterNodeDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterNodeDeletionReconciler) EXPECT() *MockMulticlusterNodeDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileNodeDeletion mocks base method.
func (m *MockMulticlusterNodeDeletionReconciler) ReconcileNodeDeletion(clusterName string, req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileNodeDeletion", clusterName, req)
}

// ReconcileNodeDeletion indicates an expected call of ReconcileNodeDeletion.
func (mr *MockMulticlusterNodeDeletionReconcilerMockRecorder) ReconcileNodeDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileNodeDeletion", reflect.TypeOf((*MockMulticlusterNodeDeletionReconciler)(nil).ReconcileNodeDeletion), clusterName, req)
}

// MockMulticlusterNodeReconcileLoop is a mock of MulticlusterNodeReconcileLoop interface.
type MockMulticlusterNodeReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterNodeReconcileLoopMockRecorder
}

// MockMulticlusterNodeReconcileLoopMockRecorder is the mock recorder for MockMulticlusterNodeReconcileLoop.
type MockMulticlusterNodeReconcileLoopMockRecorder struct {
	mock *MockMulticlusterNodeReconcileLoop
}

// NewMockMulticlusterNodeReconcileLoop creates a new mock instance.
func NewMockMulticlusterNodeReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterNodeReconcileLoop {
	mock := &MockMulticlusterNodeReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterNodeReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterNodeReconcileLoop) EXPECT() *MockMulticlusterNodeReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterNodeReconciler mocks base method.
func (m *MockMulticlusterNodeReconcileLoop) AddMulticlusterNodeReconciler(ctx context.Context, rec controller.MulticlusterNodeReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterNodeReconciler", varargs...)
}

// AddMulticlusterNodeReconciler indicates an expected call of AddMulticlusterNodeReconciler.
func (mr *MockMulticlusterNodeReconcileLoopMockRecorder) AddMulticlusterNodeReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterNodeReconciler", reflect.TypeOf((*MockMulticlusterNodeReconcileLoop)(nil).AddMulticlusterNodeReconciler), varargs...)
}