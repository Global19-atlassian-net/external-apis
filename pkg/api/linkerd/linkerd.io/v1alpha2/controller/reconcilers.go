// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	linkerd_io_v1alpha2 "github.com/linkerd/linkerd2/controller/gen/apis/serviceprofile/v1alpha2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the ServiceProfile Resource.
// implemented by the user
type ServiceProfileReconciler interface {
	ReconcileServiceProfile(obj *linkerd_io_v1alpha2.ServiceProfile) (reconcile.Result, error)
}

// Reconcile deletion events for the ServiceProfile Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type ServiceProfileDeletionReconciler interface {
	ReconcileServiceProfileDeletion(req reconcile.Request) error
}

type ServiceProfileReconcilerFuncs struct {
	OnReconcileServiceProfile         func(obj *linkerd_io_v1alpha2.ServiceProfile) (reconcile.Result, error)
	OnReconcileServiceProfileDeletion func(req reconcile.Request) error
}

func (f *ServiceProfileReconcilerFuncs) ReconcileServiceProfile(obj *linkerd_io_v1alpha2.ServiceProfile) (reconcile.Result, error) {
	if f.OnReconcileServiceProfile == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileServiceProfile(obj)
}

func (f *ServiceProfileReconcilerFuncs) ReconcileServiceProfileDeletion(req reconcile.Request) error {
	if f.OnReconcileServiceProfileDeletion == nil {
		return nil
	}
	return f.OnReconcileServiceProfileDeletion(req)
}

// Reconcile and finalize the ServiceProfile Resource
// implemented by the user
type ServiceProfileFinalizer interface {
	ServiceProfileReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	ServiceProfileFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeServiceProfile(obj *linkerd_io_v1alpha2.ServiceProfile) error
}

type ServiceProfileReconcileLoop interface {
	RunServiceProfileReconciler(ctx context.Context, rec ServiceProfileReconciler, predicates ...predicate.Predicate) error
}

type serviceProfileReconcileLoop struct {
	loop reconcile.Loop
}

func NewServiceProfileReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) ServiceProfileReconcileLoop {
	return &serviceProfileReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &linkerd_io_v1alpha2.ServiceProfile{}, options),
	}
}

func (c *serviceProfileReconcileLoop) RunServiceProfileReconciler(ctx context.Context, reconciler ServiceProfileReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericServiceProfileReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(ServiceProfileFinalizer); ok {
		reconcilerWrapper = genericServiceProfileFinalizer{
			genericServiceProfileReconciler: genericReconciler,
			finalizingReconciler:            finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericServiceProfileHandler implements a generic reconcile.Reconciler
type genericServiceProfileReconciler struct {
	reconciler ServiceProfileReconciler
}

func (r genericServiceProfileReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*linkerd_io_v1alpha2.ServiceProfile)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: ServiceProfile handler received event for %T", object)
	}
	return r.reconciler.ReconcileServiceProfile(obj)
}

func (r genericServiceProfileReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(ServiceProfileDeletionReconciler); ok {
		return deletionReconciler.ReconcileServiceProfileDeletion(request)
	}
	return nil
}

// genericServiceProfileFinalizer implements a generic reconcile.FinalizingReconciler
type genericServiceProfileFinalizer struct {
	genericServiceProfileReconciler
	finalizingReconciler ServiceProfileFinalizer
}

func (r genericServiceProfileFinalizer) FinalizerName() string {
	return r.finalizingReconciler.ServiceProfileFinalizerName()
}

func (r genericServiceProfileFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*linkerd_io_v1alpha2.ServiceProfile)
	if !ok {
		return errors.Errorf("internal error: ServiceProfile handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeServiceProfile(obj)
}
