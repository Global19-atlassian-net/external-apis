// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockDestinationRuleReconciler is a mock of DestinationRuleReconciler interface.
type MockDestinationRuleReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationRuleReconcilerMockRecorder
}

// MockDestinationRuleReconcilerMockRecorder is the mock recorder for MockDestinationRuleReconciler.
type MockDestinationRuleReconcilerMockRecorder struct {
	mock *MockDestinationRuleReconciler
}

// NewMockDestinationRuleReconciler creates a new mock instance.
func NewMockDestinationRuleReconciler(ctrl *gomock.Controller) *MockDestinationRuleReconciler {
	mock := &MockDestinationRuleReconciler{ctrl: ctrl}
	mock.recorder = &MockDestinationRuleReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDestinationRuleReconciler) EXPECT() *MockDestinationRuleReconcilerMockRecorder {
	return m.recorder
}

// ReconcileDestinationRule mocks base method.
func (m *MockDestinationRuleReconciler) ReconcileDestinationRule(obj *v1alpha3.DestinationRule) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDestinationRule", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileDestinationRule indicates an expected call of ReconcileDestinationRule.
func (mr *MockDestinationRuleReconcilerMockRecorder) ReconcileDestinationRule(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDestinationRule", reflect.TypeOf((*MockDestinationRuleReconciler)(nil).ReconcileDestinationRule), obj)
}

// MockDestinationRuleDeletionReconciler is a mock of DestinationRuleDeletionReconciler interface.
type MockDestinationRuleDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationRuleDeletionReconcilerMockRecorder
}

// MockDestinationRuleDeletionReconcilerMockRecorder is the mock recorder for MockDestinationRuleDeletionReconciler.
type MockDestinationRuleDeletionReconcilerMockRecorder struct {
	mock *MockDestinationRuleDeletionReconciler
}

// NewMockDestinationRuleDeletionReconciler creates a new mock instance.
func NewMockDestinationRuleDeletionReconciler(ctrl *gomock.Controller) *MockDestinationRuleDeletionReconciler {
	mock := &MockDestinationRuleDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockDestinationRuleDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDestinationRuleDeletionReconciler) EXPECT() *MockDestinationRuleDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileDestinationRuleDeletion mocks base method.
func (m *MockDestinationRuleDeletionReconciler) ReconcileDestinationRuleDeletion(req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileDestinationRuleDeletion", req)
}

// ReconcileDestinationRuleDeletion indicates an expected call of ReconcileDestinationRuleDeletion.
func (mr *MockDestinationRuleDeletionReconcilerMockRecorder) ReconcileDestinationRuleDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDestinationRuleDeletion", reflect.TypeOf((*MockDestinationRuleDeletionReconciler)(nil).ReconcileDestinationRuleDeletion), req)
}

// MockDestinationRuleFinalizer is a mock of DestinationRuleFinalizer interface.
type MockDestinationRuleFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationRuleFinalizerMockRecorder
}

// MockDestinationRuleFinalizerMockRecorder is the mock recorder for MockDestinationRuleFinalizer.
type MockDestinationRuleFinalizerMockRecorder struct {
	mock *MockDestinationRuleFinalizer
}

// NewMockDestinationRuleFinalizer creates a new mock instance.
func NewMockDestinationRuleFinalizer(ctrl *gomock.Controller) *MockDestinationRuleFinalizer {
	mock := &MockDestinationRuleFinalizer{ctrl: ctrl}
	mock.recorder = &MockDestinationRuleFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDestinationRuleFinalizer) EXPECT() *MockDestinationRuleFinalizerMockRecorder {
	return m.recorder
}

// ReconcileDestinationRule mocks base method.
func (m *MockDestinationRuleFinalizer) ReconcileDestinationRule(obj *v1alpha3.DestinationRule) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileDestinationRule", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileDestinationRule indicates an expected call of ReconcileDestinationRule.
func (mr *MockDestinationRuleFinalizerMockRecorder) ReconcileDestinationRule(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileDestinationRule", reflect.TypeOf((*MockDestinationRuleFinalizer)(nil).ReconcileDestinationRule), obj)
}

// DestinationRuleFinalizerName mocks base method.
func (m *MockDestinationRuleFinalizer) DestinationRuleFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestinationRuleFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// DestinationRuleFinalizerName indicates an expected call of DestinationRuleFinalizerName.
func (mr *MockDestinationRuleFinalizerMockRecorder) DestinationRuleFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestinationRuleFinalizerName", reflect.TypeOf((*MockDestinationRuleFinalizer)(nil).DestinationRuleFinalizerName))
}

// FinalizeDestinationRule mocks base method.
func (m *MockDestinationRuleFinalizer) FinalizeDestinationRule(obj *v1alpha3.DestinationRule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeDestinationRule", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeDestinationRule indicates an expected call of FinalizeDestinationRule.
func (mr *MockDestinationRuleFinalizerMockRecorder) FinalizeDestinationRule(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeDestinationRule", reflect.TypeOf((*MockDestinationRuleFinalizer)(nil).FinalizeDestinationRule), obj)
}

// MockDestinationRuleReconcileLoop is a mock of DestinationRuleReconcileLoop interface.
type MockDestinationRuleReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockDestinationRuleReconcileLoopMockRecorder
}

// MockDestinationRuleReconcileLoopMockRecorder is the mock recorder for MockDestinationRuleReconcileLoop.
type MockDestinationRuleReconcileLoopMockRecorder struct {
	mock *MockDestinationRuleReconcileLoop
}

// NewMockDestinationRuleReconcileLoop creates a new mock instance.
func NewMockDestinationRuleReconcileLoop(ctrl *gomock.Controller) *MockDestinationRuleReconcileLoop {
	mock := &MockDestinationRuleReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockDestinationRuleReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDestinationRuleReconcileLoop) EXPECT() *MockDestinationRuleReconcileLoopMockRecorder {
	return m.recorder
}

// RunDestinationRuleReconciler mocks base method.
func (m *MockDestinationRuleReconcileLoop) RunDestinationRuleReconciler(ctx context.Context, rec controller.DestinationRuleReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunDestinationRuleReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunDestinationRuleReconciler indicates an expected call of RunDestinationRuleReconciler.
func (mr *MockDestinationRuleReconcileLoopMockRecorder) RunDestinationRuleReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunDestinationRuleReconciler", reflect.TypeOf((*MockDestinationRuleReconcileLoop)(nil).RunDestinationRuleReconciler), varargs...)
}

// MockEnvoyFilterReconciler is a mock of EnvoyFilterReconciler interface.
type MockEnvoyFilterReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockEnvoyFilterReconcilerMockRecorder
}

// MockEnvoyFilterReconcilerMockRecorder is the mock recorder for MockEnvoyFilterReconciler.
type MockEnvoyFilterReconcilerMockRecorder struct {
	mock *MockEnvoyFilterReconciler
}

// NewMockEnvoyFilterReconciler creates a new mock instance.
func NewMockEnvoyFilterReconciler(ctrl *gomock.Controller) *MockEnvoyFilterReconciler {
	mock := &MockEnvoyFilterReconciler{ctrl: ctrl}
	mock.recorder = &MockEnvoyFilterReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvoyFilterReconciler) EXPECT() *MockEnvoyFilterReconcilerMockRecorder {
	return m.recorder
}

// ReconcileEnvoyFilter mocks base method.
func (m *MockEnvoyFilterReconciler) ReconcileEnvoyFilter(obj *v1alpha3.EnvoyFilter) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileEnvoyFilter", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileEnvoyFilter indicates an expected call of ReconcileEnvoyFilter.
func (mr *MockEnvoyFilterReconcilerMockRecorder) ReconcileEnvoyFilter(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterReconciler)(nil).ReconcileEnvoyFilter), obj)
}

// MockEnvoyFilterDeletionReconciler is a mock of EnvoyFilterDeletionReconciler interface.
type MockEnvoyFilterDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockEnvoyFilterDeletionReconcilerMockRecorder
}

// MockEnvoyFilterDeletionReconcilerMockRecorder is the mock recorder for MockEnvoyFilterDeletionReconciler.
type MockEnvoyFilterDeletionReconcilerMockRecorder struct {
	mock *MockEnvoyFilterDeletionReconciler
}

// NewMockEnvoyFilterDeletionReconciler creates a new mock instance.
func NewMockEnvoyFilterDeletionReconciler(ctrl *gomock.Controller) *MockEnvoyFilterDeletionReconciler {
	mock := &MockEnvoyFilterDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockEnvoyFilterDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvoyFilterDeletionReconciler) EXPECT() *MockEnvoyFilterDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileEnvoyFilterDeletion mocks base method.
func (m *MockEnvoyFilterDeletionReconciler) ReconcileEnvoyFilterDeletion(req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileEnvoyFilterDeletion", req)
}

// ReconcileEnvoyFilterDeletion indicates an expected call of ReconcileEnvoyFilterDeletion.
func (mr *MockEnvoyFilterDeletionReconcilerMockRecorder) ReconcileEnvoyFilterDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileEnvoyFilterDeletion", reflect.TypeOf((*MockEnvoyFilterDeletionReconciler)(nil).ReconcileEnvoyFilterDeletion), req)
}

// MockEnvoyFilterFinalizer is a mock of EnvoyFilterFinalizer interface.
type MockEnvoyFilterFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockEnvoyFilterFinalizerMockRecorder
}

// MockEnvoyFilterFinalizerMockRecorder is the mock recorder for MockEnvoyFilterFinalizer.
type MockEnvoyFilterFinalizerMockRecorder struct {
	mock *MockEnvoyFilterFinalizer
}

// NewMockEnvoyFilterFinalizer creates a new mock instance.
func NewMockEnvoyFilterFinalizer(ctrl *gomock.Controller) *MockEnvoyFilterFinalizer {
	mock := &MockEnvoyFilterFinalizer{ctrl: ctrl}
	mock.recorder = &MockEnvoyFilterFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvoyFilterFinalizer) EXPECT() *MockEnvoyFilterFinalizerMockRecorder {
	return m.recorder
}

// ReconcileEnvoyFilter mocks base method.
func (m *MockEnvoyFilterFinalizer) ReconcileEnvoyFilter(obj *v1alpha3.EnvoyFilter) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileEnvoyFilter", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileEnvoyFilter indicates an expected call of ReconcileEnvoyFilter.
func (mr *MockEnvoyFilterFinalizerMockRecorder) ReconcileEnvoyFilter(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterFinalizer)(nil).ReconcileEnvoyFilter), obj)
}

// EnvoyFilterFinalizerName mocks base method.
func (m *MockEnvoyFilterFinalizer) EnvoyFilterFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvoyFilterFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// EnvoyFilterFinalizerName indicates an expected call of EnvoyFilterFinalizerName.
func (mr *MockEnvoyFilterFinalizerMockRecorder) EnvoyFilterFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvoyFilterFinalizerName", reflect.TypeOf((*MockEnvoyFilterFinalizer)(nil).EnvoyFilterFinalizerName))
}

// FinalizeEnvoyFilter mocks base method.
func (m *MockEnvoyFilterFinalizer) FinalizeEnvoyFilter(obj *v1alpha3.EnvoyFilter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeEnvoyFilter", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeEnvoyFilter indicates an expected call of FinalizeEnvoyFilter.
func (mr *MockEnvoyFilterFinalizerMockRecorder) FinalizeEnvoyFilter(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeEnvoyFilter", reflect.TypeOf((*MockEnvoyFilterFinalizer)(nil).FinalizeEnvoyFilter), obj)
}

// MockEnvoyFilterReconcileLoop is a mock of EnvoyFilterReconcileLoop interface.
type MockEnvoyFilterReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockEnvoyFilterReconcileLoopMockRecorder
}

// MockEnvoyFilterReconcileLoopMockRecorder is the mock recorder for MockEnvoyFilterReconcileLoop.
type MockEnvoyFilterReconcileLoopMockRecorder struct {
	mock *MockEnvoyFilterReconcileLoop
}

// NewMockEnvoyFilterReconcileLoop creates a new mock instance.
func NewMockEnvoyFilterReconcileLoop(ctrl *gomock.Controller) *MockEnvoyFilterReconcileLoop {
	mock := &MockEnvoyFilterReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockEnvoyFilterReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvoyFilterReconcileLoop) EXPECT() *MockEnvoyFilterReconcileLoopMockRecorder {
	return m.recorder
}

// RunEnvoyFilterReconciler mocks base method.
func (m *MockEnvoyFilterReconcileLoop) RunEnvoyFilterReconciler(ctx context.Context, rec controller.EnvoyFilterReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunEnvoyFilterReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunEnvoyFilterReconciler indicates an expected call of RunEnvoyFilterReconciler.
func (mr *MockEnvoyFilterReconcileLoopMockRecorder) RunEnvoyFilterReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunEnvoyFilterReconciler", reflect.TypeOf((*MockEnvoyFilterReconcileLoop)(nil).RunEnvoyFilterReconciler), varargs...)
}

// MockGatewayReconciler is a mock of GatewayReconciler interface.
type MockGatewayReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayReconcilerMockRecorder
}

// MockGatewayReconcilerMockRecorder is the mock recorder for MockGatewayReconciler.
type MockGatewayReconcilerMockRecorder struct {
	mock *MockGatewayReconciler
}

// NewMockGatewayReconciler creates a new mock instance.
func NewMockGatewayReconciler(ctrl *gomock.Controller) *MockGatewayReconciler {
	mock := &MockGatewayReconciler{ctrl: ctrl}
	mock.recorder = &MockGatewayReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayReconciler) EXPECT() *MockGatewayReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGateway mocks base method.
func (m *MockGatewayReconciler) ReconcileGateway(obj *v1alpha3.Gateway) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGateway", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileGateway indicates an expected call of ReconcileGateway.
func (mr *MockGatewayReconcilerMockRecorder) ReconcileGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGateway", reflect.TypeOf((*MockGatewayReconciler)(nil).ReconcileGateway), obj)
}

// MockGatewayDeletionReconciler is a mock of GatewayDeletionReconciler interface.
type MockGatewayDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayDeletionReconcilerMockRecorder
}

// MockGatewayDeletionReconcilerMockRecorder is the mock recorder for MockGatewayDeletionReconciler.
type MockGatewayDeletionReconcilerMockRecorder struct {
	mock *MockGatewayDeletionReconciler
}

// NewMockGatewayDeletionReconciler creates a new mock instance.
func NewMockGatewayDeletionReconciler(ctrl *gomock.Controller) *MockGatewayDeletionReconciler {
	mock := &MockGatewayDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockGatewayDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayDeletionReconciler) EXPECT() *MockGatewayDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileGatewayDeletion mocks base method.
func (m *MockGatewayDeletionReconciler) ReconcileGatewayDeletion(req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileGatewayDeletion", req)
}

// ReconcileGatewayDeletion indicates an expected call of ReconcileGatewayDeletion.
func (mr *MockGatewayDeletionReconcilerMockRecorder) ReconcileGatewayDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGatewayDeletion", reflect.TypeOf((*MockGatewayDeletionReconciler)(nil).ReconcileGatewayDeletion), req)
}

// MockGatewayFinalizer is a mock of GatewayFinalizer interface.
type MockGatewayFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayFinalizerMockRecorder
}

// MockGatewayFinalizerMockRecorder is the mock recorder for MockGatewayFinalizer.
type MockGatewayFinalizerMockRecorder struct {
	mock *MockGatewayFinalizer
}

// NewMockGatewayFinalizer creates a new mock instance.
func NewMockGatewayFinalizer(ctrl *gomock.Controller) *MockGatewayFinalizer {
	mock := &MockGatewayFinalizer{ctrl: ctrl}
	mock.recorder = &MockGatewayFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayFinalizer) EXPECT() *MockGatewayFinalizerMockRecorder {
	return m.recorder
}

// ReconcileGateway mocks base method.
func (m *MockGatewayFinalizer) ReconcileGateway(obj *v1alpha3.Gateway) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileGateway", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileGateway indicates an expected call of ReconcileGateway.
func (mr *MockGatewayFinalizerMockRecorder) ReconcileGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileGateway", reflect.TypeOf((*MockGatewayFinalizer)(nil).ReconcileGateway), obj)
}

// GatewayFinalizerName mocks base method.
func (m *MockGatewayFinalizer) GatewayFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GatewayFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GatewayFinalizerName indicates an expected call of GatewayFinalizerName.
func (mr *MockGatewayFinalizerMockRecorder) GatewayFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GatewayFinalizerName", reflect.TypeOf((*MockGatewayFinalizer)(nil).GatewayFinalizerName))
}

// FinalizeGateway mocks base method.
func (m *MockGatewayFinalizer) FinalizeGateway(obj *v1alpha3.Gateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeGateway", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeGateway indicates an expected call of FinalizeGateway.
func (mr *MockGatewayFinalizerMockRecorder) FinalizeGateway(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeGateway", reflect.TypeOf((*MockGatewayFinalizer)(nil).FinalizeGateway), obj)
}

// MockGatewayReconcileLoop is a mock of GatewayReconcileLoop interface.
type MockGatewayReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayReconcileLoopMockRecorder
}

// MockGatewayReconcileLoopMockRecorder is the mock recorder for MockGatewayReconcileLoop.
type MockGatewayReconcileLoopMockRecorder struct {
	mock *MockGatewayReconcileLoop
}

// NewMockGatewayReconcileLoop creates a new mock instance.
func NewMockGatewayReconcileLoop(ctrl *gomock.Controller) *MockGatewayReconcileLoop {
	mock := &MockGatewayReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockGatewayReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayReconcileLoop) EXPECT() *MockGatewayReconcileLoopMockRecorder {
	return m.recorder
}

// RunGatewayReconciler mocks base method.
func (m *MockGatewayReconcileLoop) RunGatewayReconciler(ctx context.Context, rec controller.GatewayReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunGatewayReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunGatewayReconciler indicates an expected call of RunGatewayReconciler.
func (mr *MockGatewayReconcileLoopMockRecorder) RunGatewayReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunGatewayReconciler", reflect.TypeOf((*MockGatewayReconcileLoop)(nil).RunGatewayReconciler), varargs...)
}

// MockServiceEntryReconciler is a mock of ServiceEntryReconciler interface.
type MockServiceEntryReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockServiceEntryReconcilerMockRecorder
}

// MockServiceEntryReconcilerMockRecorder is the mock recorder for MockServiceEntryReconciler.
type MockServiceEntryReconcilerMockRecorder struct {
	mock *MockServiceEntryReconciler
}

// NewMockServiceEntryReconciler creates a new mock instance.
func NewMockServiceEntryReconciler(ctrl *gomock.Controller) *MockServiceEntryReconciler {
	mock := &MockServiceEntryReconciler{ctrl: ctrl}
	mock.recorder = &MockServiceEntryReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceEntryReconciler) EXPECT() *MockServiceEntryReconcilerMockRecorder {
	return m.recorder
}

// ReconcileServiceEntry mocks base method.
func (m *MockServiceEntryReconciler) ReconcileServiceEntry(obj *v1alpha3.ServiceEntry) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileServiceEntry", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileServiceEntry indicates an expected call of ReconcileServiceEntry.
func (mr *MockServiceEntryReconcilerMockRecorder) ReconcileServiceEntry(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileServiceEntry", reflect.TypeOf((*MockServiceEntryReconciler)(nil).ReconcileServiceEntry), obj)
}

// MockServiceEntryDeletionReconciler is a mock of ServiceEntryDeletionReconciler interface.
type MockServiceEntryDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockServiceEntryDeletionReconcilerMockRecorder
}

// MockServiceEntryDeletionReconcilerMockRecorder is the mock recorder for MockServiceEntryDeletionReconciler.
type MockServiceEntryDeletionReconcilerMockRecorder struct {
	mock *MockServiceEntryDeletionReconciler
}

// NewMockServiceEntryDeletionReconciler creates a new mock instance.
func NewMockServiceEntryDeletionReconciler(ctrl *gomock.Controller) *MockServiceEntryDeletionReconciler {
	mock := &MockServiceEntryDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockServiceEntryDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceEntryDeletionReconciler) EXPECT() *MockServiceEntryDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileServiceEntryDeletion mocks base method.
func (m *MockServiceEntryDeletionReconciler) ReconcileServiceEntryDeletion(req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileServiceEntryDeletion", req)
}

// ReconcileServiceEntryDeletion indicates an expected call of ReconcileServiceEntryDeletion.
func (mr *MockServiceEntryDeletionReconcilerMockRecorder) ReconcileServiceEntryDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileServiceEntryDeletion", reflect.TypeOf((*MockServiceEntryDeletionReconciler)(nil).ReconcileServiceEntryDeletion), req)
}

// MockServiceEntryFinalizer is a mock of ServiceEntryFinalizer interface.
type MockServiceEntryFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockServiceEntryFinalizerMockRecorder
}

// MockServiceEntryFinalizerMockRecorder is the mock recorder for MockServiceEntryFinalizer.
type MockServiceEntryFinalizerMockRecorder struct {
	mock *MockServiceEntryFinalizer
}

// NewMockServiceEntryFinalizer creates a new mock instance.
func NewMockServiceEntryFinalizer(ctrl *gomock.Controller) *MockServiceEntryFinalizer {
	mock := &MockServiceEntryFinalizer{ctrl: ctrl}
	mock.recorder = &MockServiceEntryFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceEntryFinalizer) EXPECT() *MockServiceEntryFinalizerMockRecorder {
	return m.recorder
}

// ReconcileServiceEntry mocks base method.
func (m *MockServiceEntryFinalizer) ReconcileServiceEntry(obj *v1alpha3.ServiceEntry) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileServiceEntry", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileServiceEntry indicates an expected call of ReconcileServiceEntry.
func (mr *MockServiceEntryFinalizerMockRecorder) ReconcileServiceEntry(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileServiceEntry", reflect.TypeOf((*MockServiceEntryFinalizer)(nil).ReconcileServiceEntry), obj)
}

// ServiceEntryFinalizerName mocks base method.
func (m *MockServiceEntryFinalizer) ServiceEntryFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceEntryFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ServiceEntryFinalizerName indicates an expected call of ServiceEntryFinalizerName.
func (mr *MockServiceEntryFinalizerMockRecorder) ServiceEntryFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceEntryFinalizerName", reflect.TypeOf((*MockServiceEntryFinalizer)(nil).ServiceEntryFinalizerName))
}

// FinalizeServiceEntry mocks base method.
func (m *MockServiceEntryFinalizer) FinalizeServiceEntry(obj *v1alpha3.ServiceEntry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeServiceEntry", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeServiceEntry indicates an expected call of FinalizeServiceEntry.
func (mr *MockServiceEntryFinalizerMockRecorder) FinalizeServiceEntry(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeServiceEntry", reflect.TypeOf((*MockServiceEntryFinalizer)(nil).FinalizeServiceEntry), obj)
}

// MockServiceEntryReconcileLoop is a mock of ServiceEntryReconcileLoop interface.
type MockServiceEntryReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockServiceEntryReconcileLoopMockRecorder
}

// MockServiceEntryReconcileLoopMockRecorder is the mock recorder for MockServiceEntryReconcileLoop.
type MockServiceEntryReconcileLoopMockRecorder struct {
	mock *MockServiceEntryReconcileLoop
}

// NewMockServiceEntryReconcileLoop creates a new mock instance.
func NewMockServiceEntryReconcileLoop(ctrl *gomock.Controller) *MockServiceEntryReconcileLoop {
	mock := &MockServiceEntryReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockServiceEntryReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceEntryReconcileLoop) EXPECT() *MockServiceEntryReconcileLoopMockRecorder {
	return m.recorder
}

// RunServiceEntryReconciler mocks base method.
func (m *MockServiceEntryReconcileLoop) RunServiceEntryReconciler(ctx context.Context, rec controller.ServiceEntryReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunServiceEntryReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunServiceEntryReconciler indicates an expected call of RunServiceEntryReconciler.
func (mr *MockServiceEntryReconcileLoopMockRecorder) RunServiceEntryReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunServiceEntryReconciler", reflect.TypeOf((*MockServiceEntryReconcileLoop)(nil).RunServiceEntryReconciler), varargs...)
}

// MockVirtualServiceReconciler is a mock of VirtualServiceReconciler interface.
type MockVirtualServiceReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualServiceReconcilerMockRecorder
}

// MockVirtualServiceReconcilerMockRecorder is the mock recorder for MockVirtualServiceReconciler.
type MockVirtualServiceReconcilerMockRecorder struct {
	mock *MockVirtualServiceReconciler
}

// NewMockVirtualServiceReconciler creates a new mock instance.
func NewMockVirtualServiceReconciler(ctrl *gomock.Controller) *MockVirtualServiceReconciler {
	mock := &MockVirtualServiceReconciler{ctrl: ctrl}
	mock.recorder = &MockVirtualServiceReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualServiceReconciler) EXPECT() *MockVirtualServiceReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualService mocks base method.
func (m *MockVirtualServiceReconciler) ReconcileVirtualService(obj *v1alpha3.VirtualService) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualService", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileVirtualService indicates an expected call of ReconcileVirtualService.
func (mr *MockVirtualServiceReconcilerMockRecorder) ReconcileVirtualService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualService", reflect.TypeOf((*MockVirtualServiceReconciler)(nil).ReconcileVirtualService), obj)
}

// MockVirtualServiceDeletionReconciler is a mock of VirtualServiceDeletionReconciler interface.
type MockVirtualServiceDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualServiceDeletionReconcilerMockRecorder
}

// MockVirtualServiceDeletionReconcilerMockRecorder is the mock recorder for MockVirtualServiceDeletionReconciler.
type MockVirtualServiceDeletionReconcilerMockRecorder struct {
	mock *MockVirtualServiceDeletionReconciler
}

// NewMockVirtualServiceDeletionReconciler creates a new mock instance.
func NewMockVirtualServiceDeletionReconciler(ctrl *gomock.Controller) *MockVirtualServiceDeletionReconciler {
	mock := &MockVirtualServiceDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockVirtualServiceDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualServiceDeletionReconciler) EXPECT() *MockVirtualServiceDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualServiceDeletion mocks base method.
func (m *MockVirtualServiceDeletionReconciler) ReconcileVirtualServiceDeletion(req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileVirtualServiceDeletion", req)
}

// ReconcileVirtualServiceDeletion indicates an expected call of ReconcileVirtualServiceDeletion.
func (mr *MockVirtualServiceDeletionReconcilerMockRecorder) ReconcileVirtualServiceDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualServiceDeletion", reflect.TypeOf((*MockVirtualServiceDeletionReconciler)(nil).ReconcileVirtualServiceDeletion), req)
}

// MockVirtualServiceFinalizer is a mock of VirtualServiceFinalizer interface.
type MockVirtualServiceFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualServiceFinalizerMockRecorder
}

// MockVirtualServiceFinalizerMockRecorder is the mock recorder for MockVirtualServiceFinalizer.
type MockVirtualServiceFinalizerMockRecorder struct {
	mock *MockVirtualServiceFinalizer
}

// NewMockVirtualServiceFinalizer creates a new mock instance.
func NewMockVirtualServiceFinalizer(ctrl *gomock.Controller) *MockVirtualServiceFinalizer {
	mock := &MockVirtualServiceFinalizer{ctrl: ctrl}
	mock.recorder = &MockVirtualServiceFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualServiceFinalizer) EXPECT() *MockVirtualServiceFinalizerMockRecorder {
	return m.recorder
}

// ReconcileVirtualService mocks base method.
func (m *MockVirtualServiceFinalizer) ReconcileVirtualService(obj *v1alpha3.VirtualService) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualService", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileVirtualService indicates an expected call of ReconcileVirtualService.
func (mr *MockVirtualServiceFinalizerMockRecorder) ReconcileVirtualService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualService", reflect.TypeOf((*MockVirtualServiceFinalizer)(nil).ReconcileVirtualService), obj)
}

// VirtualServiceFinalizerName mocks base method.
func (m *MockVirtualServiceFinalizer) VirtualServiceFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualServiceFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// VirtualServiceFinalizerName indicates an expected call of VirtualServiceFinalizerName.
func (mr *MockVirtualServiceFinalizerMockRecorder) VirtualServiceFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualServiceFinalizerName", reflect.TypeOf((*MockVirtualServiceFinalizer)(nil).VirtualServiceFinalizerName))
}

// FinalizeVirtualService mocks base method.
func (m *MockVirtualServiceFinalizer) FinalizeVirtualService(obj *v1alpha3.VirtualService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeVirtualService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeVirtualService indicates an expected call of FinalizeVirtualService.
func (mr *MockVirtualServiceFinalizerMockRecorder) FinalizeVirtualService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeVirtualService", reflect.TypeOf((*MockVirtualServiceFinalizer)(nil).FinalizeVirtualService), obj)
}

// MockVirtualServiceReconcileLoop is a mock of VirtualServiceReconcileLoop interface.
type MockVirtualServiceReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualServiceReconcileLoopMockRecorder
}

// MockVirtualServiceReconcileLoopMockRecorder is the mock recorder for MockVirtualServiceReconcileLoop.
type MockVirtualServiceReconcileLoopMockRecorder struct {
	mock *MockVirtualServiceReconcileLoop
}

// NewMockVirtualServiceReconcileLoop creates a new mock instance.
func NewMockVirtualServiceReconcileLoop(ctrl *gomock.Controller) *MockVirtualServiceReconcileLoop {
	mock := &MockVirtualServiceReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockVirtualServiceReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualServiceReconcileLoop) EXPECT() *MockVirtualServiceReconcileLoopMockRecorder {
	return m.recorder
}

// RunVirtualServiceReconciler mocks base method.
func (m *MockVirtualServiceReconcileLoop) RunVirtualServiceReconciler(ctx context.Context, rec controller.VirtualServiceReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunVirtualServiceReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunVirtualServiceReconciler indicates an expected call of RunVirtualServiceReconciler.
func (mr *MockVirtualServiceReconcileLoopMockRecorder) RunVirtualServiceReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunVirtualServiceReconciler", reflect.TypeOf((*MockVirtualServiceReconcileLoop)(nil).RunVirtualServiceReconciler), varargs...)
}