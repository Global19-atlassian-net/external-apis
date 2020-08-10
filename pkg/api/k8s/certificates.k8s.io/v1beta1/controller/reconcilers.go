// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	certificates_k8s_io_v1beta1 "k8s.io/api/certificates/v1beta1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the CertificateSigningRequest Resource.
// implemented by the user
type CertificateSigningRequestReconciler interface {
	ReconcileCertificateSigningRequest(obj *certificates_k8s_io_v1beta1.CertificateSigningRequest) (reconcile.Result, error)
}

// Reconcile deletion events for the CertificateSigningRequest Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type CertificateSigningRequestDeletionReconciler interface {
	ReconcileCertificateSigningRequestDeletion(req reconcile.Request) error
}

type CertificateSigningRequestReconcilerFuncs struct {
	OnReconcileCertificateSigningRequest         func(obj *certificates_k8s_io_v1beta1.CertificateSigningRequest) (reconcile.Result, error)
	OnReconcileCertificateSigningRequestDeletion func(req reconcile.Request) error
}

func (f *CertificateSigningRequestReconcilerFuncs) ReconcileCertificateSigningRequest(obj *certificates_k8s_io_v1beta1.CertificateSigningRequest) (reconcile.Result, error) {
	if f.OnReconcileCertificateSigningRequest == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileCertificateSigningRequest(obj)
}

func (f *CertificateSigningRequestReconcilerFuncs) ReconcileCertificateSigningRequestDeletion(req reconcile.Request) error {
	if f.OnReconcileCertificateSigningRequestDeletion == nil {
		return nil
	}
	return f.OnReconcileCertificateSigningRequestDeletion(req)
}

// Reconcile and finalize the CertificateSigningRequest Resource
// implemented by the user
type CertificateSigningRequestFinalizer interface {
	CertificateSigningRequestReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	CertificateSigningRequestFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeCertificateSigningRequest(obj *certificates_k8s_io_v1beta1.CertificateSigningRequest) error
}

type CertificateSigningRequestReconcileLoop interface {
	RunCertificateSigningRequestReconciler(ctx context.Context, rec CertificateSigningRequestReconciler, predicates ...predicate.Predicate) error
}

type certificateSigningRequestReconcileLoop struct {
	loop reconcile.Loop
}

func NewCertificateSigningRequestReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) CertificateSigningRequestReconcileLoop {
	return &certificateSigningRequestReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &certificates_k8s_io_v1beta1.CertificateSigningRequest{}, options),
	}
}

func (c *certificateSigningRequestReconcileLoop) RunCertificateSigningRequestReconciler(ctx context.Context, reconciler CertificateSigningRequestReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericCertificateSigningRequestReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(CertificateSigningRequestFinalizer); ok {
		reconcilerWrapper = genericCertificateSigningRequestFinalizer{
			genericCertificateSigningRequestReconciler: genericReconciler,
			finalizingReconciler:                       finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericCertificateSigningRequestHandler implements a generic reconcile.Reconciler
type genericCertificateSigningRequestReconciler struct {
	reconciler CertificateSigningRequestReconciler
}

func (r genericCertificateSigningRequestReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*certificates_k8s_io_v1beta1.CertificateSigningRequest)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: CertificateSigningRequest handler received event for %T", object)
	}
	return r.reconciler.ReconcileCertificateSigningRequest(obj)
}

func (r genericCertificateSigningRequestReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(CertificateSigningRequestDeletionReconciler); ok {
		return deletionReconciler.ReconcileCertificateSigningRequestDeletion(request)
	}
	return nil
}

// genericCertificateSigningRequestFinalizer implements a generic reconcile.FinalizingReconciler
type genericCertificateSigningRequestFinalizer struct {
	genericCertificateSigningRequestReconciler
	finalizingReconciler CertificateSigningRequestFinalizer
}

func (r genericCertificateSigningRequestFinalizer) FinalizerName() string {
	return r.finalizingReconciler.CertificateSigningRequestFinalizerName()
}

func (r genericCertificateSigningRequestFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*certificates_k8s_io_v1beta1.CertificateSigningRequest)
	if !ok {
		return errors.Errorf("internal error: CertificateSigningRequest handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeCertificateSigningRequest(obj)
}
