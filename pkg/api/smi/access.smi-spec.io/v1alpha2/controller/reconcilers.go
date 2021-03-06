// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	access_smi_spec_io_v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the TrafficTarget Resource.
// implemented by the user
type TrafficTargetReconciler interface {
	ReconcileTrafficTarget(obj *access_smi_spec_io_v1alpha2.TrafficTarget) (reconcile.Result, error)
}

// Reconcile deletion events for the TrafficTarget Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type TrafficTargetDeletionReconciler interface {
	ReconcileTrafficTargetDeletion(req reconcile.Request) error
}

type TrafficTargetReconcilerFuncs struct {
	OnReconcileTrafficTarget         func(obj *access_smi_spec_io_v1alpha2.TrafficTarget) (reconcile.Result, error)
	OnReconcileTrafficTargetDeletion func(req reconcile.Request) error
}

func (f *TrafficTargetReconcilerFuncs) ReconcileTrafficTarget(obj *access_smi_spec_io_v1alpha2.TrafficTarget) (reconcile.Result, error) {
	if f.OnReconcileTrafficTarget == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileTrafficTarget(obj)
}

func (f *TrafficTargetReconcilerFuncs) ReconcileTrafficTargetDeletion(req reconcile.Request) error {
	if f.OnReconcileTrafficTargetDeletion == nil {
		return nil
	}
	return f.OnReconcileTrafficTargetDeletion(req)
}

// Reconcile and finalize the TrafficTarget Resource
// implemented by the user
type TrafficTargetFinalizer interface {
	TrafficTargetReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	TrafficTargetFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeTrafficTarget(obj *access_smi_spec_io_v1alpha2.TrafficTarget) error
}

type TrafficTargetReconcileLoop interface {
	RunTrafficTargetReconciler(ctx context.Context, rec TrafficTargetReconciler, predicates ...predicate.Predicate) error
}

type trafficTargetReconcileLoop struct {
	loop reconcile.Loop
}

func NewTrafficTargetReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) TrafficTargetReconcileLoop {
	return &trafficTargetReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &access_smi_spec_io_v1alpha2.TrafficTarget{}, options),
	}
}

func (c *trafficTargetReconcileLoop) RunTrafficTargetReconciler(ctx context.Context, reconciler TrafficTargetReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericTrafficTargetReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(TrafficTargetFinalizer); ok {
		reconcilerWrapper = genericTrafficTargetFinalizer{
			genericTrafficTargetReconciler: genericReconciler,
			finalizingReconciler:           finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericTrafficTargetHandler implements a generic reconcile.Reconciler
type genericTrafficTargetReconciler struct {
	reconciler TrafficTargetReconciler
}

func (r genericTrafficTargetReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*access_smi_spec_io_v1alpha2.TrafficTarget)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: TrafficTarget handler received event for %T", object)
	}
	return r.reconciler.ReconcileTrafficTarget(obj)
}

func (r genericTrafficTargetReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(TrafficTargetDeletionReconciler); ok {
		return deletionReconciler.ReconcileTrafficTargetDeletion(request)
	}
	return nil
}

// genericTrafficTargetFinalizer implements a generic reconcile.FinalizingReconciler
type genericTrafficTargetFinalizer struct {
	genericTrafficTargetReconciler
	finalizingReconciler TrafficTargetFinalizer
}

func (r genericTrafficTargetFinalizer) FinalizerName() string {
	return r.finalizingReconciler.TrafficTargetFinalizerName()
}

func (r genericTrafficTargetFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*access_smi_spec_io_v1alpha2.TrafficTarget)
	if !ok {
		return errors.Errorf("internal error: TrafficTarget handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeTrafficTarget(obj)
}
