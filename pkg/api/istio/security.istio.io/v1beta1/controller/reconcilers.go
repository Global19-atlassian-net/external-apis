// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	security_istio_io_v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the AuthorizationPolicy Resource.
// implemented by the user
type AuthorizationPolicyReconciler interface {
	ReconcileAuthorizationPolicy(obj *security_istio_io_v1beta1.AuthorizationPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the AuthorizationPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type AuthorizationPolicyDeletionReconciler interface {
	ReconcileAuthorizationPolicyDeletion(req reconcile.Request)
}

type AuthorizationPolicyReconcilerFuncs struct {
	OnReconcileAuthorizationPolicy         func(obj *security_istio_io_v1beta1.AuthorizationPolicy) (reconcile.Result, error)
	OnReconcileAuthorizationPolicyDeletion func(req reconcile.Request)
}

func (f *AuthorizationPolicyReconcilerFuncs) ReconcileAuthorizationPolicy(obj *security_istio_io_v1beta1.AuthorizationPolicy) (reconcile.Result, error) {
	if f.OnReconcileAuthorizationPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileAuthorizationPolicy(obj)
}

func (f *AuthorizationPolicyReconcilerFuncs) ReconcileAuthorizationPolicyDeletion(req reconcile.Request) {
	if f.OnReconcileAuthorizationPolicyDeletion == nil {
		return
	}
	f.OnReconcileAuthorizationPolicyDeletion(req)
}

// Reconcile and finalize the AuthorizationPolicy Resource
// implemented by the user
type AuthorizationPolicyFinalizer interface {
	AuthorizationPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	AuthorizationPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeAuthorizationPolicy(obj *security_istio_io_v1beta1.AuthorizationPolicy) error
}

type AuthorizationPolicyReconcileLoop interface {
	RunAuthorizationPolicyReconciler(ctx context.Context, rec AuthorizationPolicyReconciler, predicates ...predicate.Predicate) error
}

type authorizationPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewAuthorizationPolicyReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) AuthorizationPolicyReconcileLoop {
	return &authorizationPolicyReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &security_istio_io_v1beta1.AuthorizationPolicy{}, options),
	}
}

func (c *authorizationPolicyReconcileLoop) RunAuthorizationPolicyReconciler(ctx context.Context, reconciler AuthorizationPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericAuthorizationPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(AuthorizationPolicyFinalizer); ok {
		reconcilerWrapper = genericAuthorizationPolicyFinalizer{
			genericAuthorizationPolicyReconciler: genericReconciler,
			finalizingReconciler:                 finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericAuthorizationPolicyHandler implements a generic reconcile.Reconciler
type genericAuthorizationPolicyReconciler struct {
	reconciler AuthorizationPolicyReconciler
}

func (r genericAuthorizationPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*security_istio_io_v1beta1.AuthorizationPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: AuthorizationPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileAuthorizationPolicy(obj)
}

func (r genericAuthorizationPolicyReconciler) ReconcileDeletion(request reconcile.Request) {
	if deletionReconciler, ok := r.reconciler.(AuthorizationPolicyDeletionReconciler); ok {
		deletionReconciler.ReconcileAuthorizationPolicyDeletion(request)
	}
}

// genericAuthorizationPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericAuthorizationPolicyFinalizer struct {
	genericAuthorizationPolicyReconciler
	finalizingReconciler AuthorizationPolicyFinalizer
}

func (r genericAuthorizationPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.AuthorizationPolicyFinalizerName()
}

func (r genericAuthorizationPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*security_istio_io_v1beta1.AuthorizationPolicy)
	if !ok {
		return errors.Errorf("internal error: AuthorizationPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeAuthorizationPolicy(obj)
}