// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	split_smi_spec_io_v1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the TrafficSplit Resource across clusters.
// implemented by the user
type MulticlusterTrafficSplitReconciler interface {
	ReconcileTrafficSplit(clusterName string, obj *split_smi_spec_io_v1alpha1.TrafficSplit) (reconcile.Result, error)
}

// Reconcile deletion events for the TrafficSplit Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterTrafficSplitDeletionReconciler interface {
	ReconcileTrafficSplitDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterTrafficSplitReconcilerFuncs struct {
	OnReconcileTrafficSplit         func(clusterName string, obj *split_smi_spec_io_v1alpha1.TrafficSplit) (reconcile.Result, error)
	OnReconcileTrafficSplitDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterTrafficSplitReconcilerFuncs) ReconcileTrafficSplit(clusterName string, obj *split_smi_spec_io_v1alpha1.TrafficSplit) (reconcile.Result, error) {
	if f.OnReconcileTrafficSplit == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileTrafficSplit(clusterName, obj)
}

func (f *MulticlusterTrafficSplitReconcilerFuncs) ReconcileTrafficSplitDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileTrafficSplitDeletion == nil {
		return nil
	}
	return f.OnReconcileTrafficSplitDeletion(clusterName, req)
}

type MulticlusterTrafficSplitReconcileLoop interface {
	// AddMulticlusterTrafficSplitReconciler adds a MulticlusterTrafficSplitReconciler to the MulticlusterTrafficSplitReconcileLoop.
	AddMulticlusterTrafficSplitReconciler(ctx context.Context, rec MulticlusterTrafficSplitReconciler, predicates ...predicate.Predicate)
}

type multiclusterTrafficSplitReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterTrafficSplitReconcileLoop) AddMulticlusterTrafficSplitReconciler(ctx context.Context, rec MulticlusterTrafficSplitReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericTrafficSplitMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterTrafficSplitReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterTrafficSplitReconcileLoop {
	return &multiclusterTrafficSplitReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &split_smi_spec_io_v1alpha1.TrafficSplit{})}
}

type genericTrafficSplitMulticlusterReconciler struct {
	reconciler MulticlusterTrafficSplitReconciler
}

func (g genericTrafficSplitMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterTrafficSplitDeletionReconciler); ok {
		return deletionReconciler.ReconcileTrafficSplitDeletion(cluster, req)
	}
	return nil
}

func (g genericTrafficSplitMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*split_smi_spec_io_v1alpha1.TrafficSplit)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: TrafficSplit handler received event for %T", object)
	}
	return g.reconciler.ReconcileTrafficSplit(cluster, obj)
}
