// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	certificates_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/certificates.smh.solo.io/v1alpha2"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the IssuedCertificate Resource.
// implemented by the user
type IssuedCertificateReconciler interface {
	ReconcileIssuedCertificate(obj *certificates_smh_solo_io_v1alpha2.IssuedCertificate) (reconcile.Result, error)
}

// Reconcile deletion events for the IssuedCertificate Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type IssuedCertificateDeletionReconciler interface {
	ReconcileIssuedCertificateDeletion(req reconcile.Request) error
}

type IssuedCertificateReconcilerFuncs struct {
	OnReconcileIssuedCertificate         func(obj *certificates_smh_solo_io_v1alpha2.IssuedCertificate) (reconcile.Result, error)
	OnReconcileIssuedCertificateDeletion func(req reconcile.Request) error
}

func (f *IssuedCertificateReconcilerFuncs) ReconcileIssuedCertificate(obj *certificates_smh_solo_io_v1alpha2.IssuedCertificate) (reconcile.Result, error) {
	if f.OnReconcileIssuedCertificate == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileIssuedCertificate(obj)
}

func (f *IssuedCertificateReconcilerFuncs) ReconcileIssuedCertificateDeletion(req reconcile.Request) error {
	if f.OnReconcileIssuedCertificateDeletion == nil {
		return nil
	}
	return f.OnReconcileIssuedCertificateDeletion(req)
}

// Reconcile and finalize the IssuedCertificate Resource
// implemented by the user
type IssuedCertificateFinalizer interface {
	IssuedCertificateReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	IssuedCertificateFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeIssuedCertificate(obj *certificates_smh_solo_io_v1alpha2.IssuedCertificate) error
}

type IssuedCertificateReconcileLoop interface {
	RunIssuedCertificateReconciler(ctx context.Context, rec IssuedCertificateReconciler, predicates ...predicate.Predicate) error
}

type issuedCertificateReconcileLoop struct {
	loop reconcile.Loop
}

func NewIssuedCertificateReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) IssuedCertificateReconcileLoop {
	return &issuedCertificateReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &certificates_smh_solo_io_v1alpha2.IssuedCertificate{}, options),
	}
}

func (c *issuedCertificateReconcileLoop) RunIssuedCertificateReconciler(ctx context.Context, reconciler IssuedCertificateReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericIssuedCertificateReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(IssuedCertificateFinalizer); ok {
		reconcilerWrapper = genericIssuedCertificateFinalizer{
			genericIssuedCertificateReconciler: genericReconciler,
			finalizingReconciler:               finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericIssuedCertificateHandler implements a generic reconcile.Reconciler
type genericIssuedCertificateReconciler struct {
	reconciler IssuedCertificateReconciler
}

func (r genericIssuedCertificateReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*certificates_smh_solo_io_v1alpha2.IssuedCertificate)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: IssuedCertificate handler received event for %T", object)
	}
	return r.reconciler.ReconcileIssuedCertificate(obj)
}

func (r genericIssuedCertificateReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(IssuedCertificateDeletionReconciler); ok {
		return deletionReconciler.ReconcileIssuedCertificateDeletion(request)
	}
	return nil
}

// genericIssuedCertificateFinalizer implements a generic reconcile.FinalizingReconciler
type genericIssuedCertificateFinalizer struct {
	genericIssuedCertificateReconciler
	finalizingReconciler IssuedCertificateFinalizer
}

func (r genericIssuedCertificateFinalizer) FinalizerName() string {
	return r.finalizingReconciler.IssuedCertificateFinalizerName()
}

func (r genericIssuedCertificateFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*certificates_smh_solo_io_v1alpha2.IssuedCertificate)
	if !ok {
		return errors.Errorf("internal error: IssuedCertificate handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeIssuedCertificate(obj)
}

// Reconcile Upsert events for the CertificateRequest Resource.
// implemented by the user
type CertificateRequestReconciler interface {
	ReconcileCertificateRequest(obj *certificates_smh_solo_io_v1alpha2.CertificateRequest) (reconcile.Result, error)
}

// Reconcile deletion events for the CertificateRequest Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type CertificateRequestDeletionReconciler interface {
	ReconcileCertificateRequestDeletion(req reconcile.Request) error
}

type CertificateRequestReconcilerFuncs struct {
	OnReconcileCertificateRequest         func(obj *certificates_smh_solo_io_v1alpha2.CertificateRequest) (reconcile.Result, error)
	OnReconcileCertificateRequestDeletion func(req reconcile.Request) error
}

func (f *CertificateRequestReconcilerFuncs) ReconcileCertificateRequest(obj *certificates_smh_solo_io_v1alpha2.CertificateRequest) (reconcile.Result, error) {
	if f.OnReconcileCertificateRequest == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileCertificateRequest(obj)
}

func (f *CertificateRequestReconcilerFuncs) ReconcileCertificateRequestDeletion(req reconcile.Request) error {
	if f.OnReconcileCertificateRequestDeletion == nil {
		return nil
	}
	return f.OnReconcileCertificateRequestDeletion(req)
}

// Reconcile and finalize the CertificateRequest Resource
// implemented by the user
type CertificateRequestFinalizer interface {
	CertificateRequestReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	CertificateRequestFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeCertificateRequest(obj *certificates_smh_solo_io_v1alpha2.CertificateRequest) error
}

type CertificateRequestReconcileLoop interface {
	RunCertificateRequestReconciler(ctx context.Context, rec CertificateRequestReconciler, predicates ...predicate.Predicate) error
}

type certificateRequestReconcileLoop struct {
	loop reconcile.Loop
}

func NewCertificateRequestReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) CertificateRequestReconcileLoop {
	return &certificateRequestReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &certificates_smh_solo_io_v1alpha2.CertificateRequest{}, options),
	}
}

func (c *certificateRequestReconcileLoop) RunCertificateRequestReconciler(ctx context.Context, reconciler CertificateRequestReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericCertificateRequestReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(CertificateRequestFinalizer); ok {
		reconcilerWrapper = genericCertificateRequestFinalizer{
			genericCertificateRequestReconciler: genericReconciler,
			finalizingReconciler:                finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericCertificateRequestHandler implements a generic reconcile.Reconciler
type genericCertificateRequestReconciler struct {
	reconciler CertificateRequestReconciler
}

func (r genericCertificateRequestReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*certificates_smh_solo_io_v1alpha2.CertificateRequest)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: CertificateRequest handler received event for %T", object)
	}
	return r.reconciler.ReconcileCertificateRequest(obj)
}

func (r genericCertificateRequestReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(CertificateRequestDeletionReconciler); ok {
		return deletionReconciler.ReconcileCertificateRequestDeletion(request)
	}
	return nil
}

// genericCertificateRequestFinalizer implements a generic reconcile.FinalizingReconciler
type genericCertificateRequestFinalizer struct {
	genericCertificateRequestReconciler
	finalizingReconciler CertificateRequestFinalizer
}

func (r genericCertificateRequestFinalizer) FinalizerName() string {
	return r.finalizingReconciler.CertificateRequestFinalizerName()
}

func (r genericCertificateRequestFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*certificates_smh_solo_io_v1alpha2.CertificateRequest)
	if !ok {
		return errors.Errorf("internal error: CertificateRequest handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeCertificateRequest(obj)
}