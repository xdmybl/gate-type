// Code generated by engine gate build no edit

//go:generate mockgen -source ./reconcilers.go -destination mocks/reconcilers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	gate_v1 "github.com/xdmybl/gate-type/pkg/api/gate/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the CaCertificate Resource.
// implemented by the user
type CaCertificateReconciler interface {
	ReconcileCaCertificate(obj *gate_v1.CaCertificate) (reconcile.Result, error)
}

// Reconcile deletion events for the CaCertificate Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type CaCertificateDeletionReconciler interface {
	ReconcileCaCertificateDeletion(req reconcile.Request) error
}

type CaCertificateReconcilerFuncs struct {
	OnReconcileCaCertificate         func(obj *gate_v1.CaCertificate) (reconcile.Result, error)
	OnReconcileCaCertificateDeletion func(req reconcile.Request) error
}

func (f *CaCertificateReconcilerFuncs) ReconcileCaCertificate(obj *gate_v1.CaCertificate) (reconcile.Result, error) {
	if f.OnReconcileCaCertificate == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileCaCertificate(obj)
}

func (f *CaCertificateReconcilerFuncs) ReconcileCaCertificateDeletion(req reconcile.Request) error {
	if f.OnReconcileCaCertificateDeletion == nil {
		return nil
	}
	return f.OnReconcileCaCertificateDeletion(req)
}

// Reconcile and finalize the CaCertificate Resource
// implemented by the user
type CaCertificateFinalizer interface {
	CaCertificateReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	CaCertificateFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeCaCertificate(obj *gate_v1.CaCertificate) error
}

type CaCertificateReconcileLoop interface {
	RunCaCertificateReconciler(ctx context.Context, rec CaCertificateReconciler, predicates ...predicate.Predicate) error
}

type caCertificateReconcileLoop struct {
	loop reconcile.Loop
}

func NewCaCertificateReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) CaCertificateReconcileLoop {
	return &caCertificateReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &gate_v1.CaCertificate{}, options),
	}
}

func (c *caCertificateReconcileLoop) RunCaCertificateReconciler(ctx context.Context, reconciler CaCertificateReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericCaCertificateReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(CaCertificateFinalizer); ok {
		reconcilerWrapper = genericCaCertificateFinalizer{
			genericCaCertificateReconciler: genericReconciler,
			finalizingReconciler:           finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericCaCertificateHandler implements a generic reconcile.Reconciler
type genericCaCertificateReconciler struct {
	reconciler CaCertificateReconciler
}

func (r genericCaCertificateReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gate_v1.CaCertificate)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: CaCertificate handler received event for %T", object)
	}
	return r.reconciler.ReconcileCaCertificate(obj)
}

func (r genericCaCertificateReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(CaCertificateDeletionReconciler); ok {
		return deletionReconciler.ReconcileCaCertificateDeletion(request)
	}
	return nil
}

// genericCaCertificateFinalizer implements a generic reconcile.FinalizingReconciler
type genericCaCertificateFinalizer struct {
	genericCaCertificateReconciler
	finalizingReconciler CaCertificateFinalizer
}

func (r genericCaCertificateFinalizer) FinalizerName() string {
	return r.finalizingReconciler.CaCertificateFinalizerName()
}

func (r genericCaCertificateFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*gate_v1.CaCertificate)
	if !ok {
		return errors.Errorf("internal error: CaCertificate handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeCaCertificate(obj)
}

// Reconcile Upsert events for the SslCertificate Resource.
// implemented by the user
type SslCertificateReconciler interface {
	ReconcileSslCertificate(obj *gate_v1.SslCertificate) (reconcile.Result, error)
}

// Reconcile deletion events for the SslCertificate Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type SslCertificateDeletionReconciler interface {
	ReconcileSslCertificateDeletion(req reconcile.Request) error
}

type SslCertificateReconcilerFuncs struct {
	OnReconcileSslCertificate         func(obj *gate_v1.SslCertificate) (reconcile.Result, error)
	OnReconcileSslCertificateDeletion func(req reconcile.Request) error
}

func (f *SslCertificateReconcilerFuncs) ReconcileSslCertificate(obj *gate_v1.SslCertificate) (reconcile.Result, error) {
	if f.OnReconcileSslCertificate == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileSslCertificate(obj)
}

func (f *SslCertificateReconcilerFuncs) ReconcileSslCertificateDeletion(req reconcile.Request) error {
	if f.OnReconcileSslCertificateDeletion == nil {
		return nil
	}
	return f.OnReconcileSslCertificateDeletion(req)
}

// Reconcile and finalize the SslCertificate Resource
// implemented by the user
type SslCertificateFinalizer interface {
	SslCertificateReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	SslCertificateFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeSslCertificate(obj *gate_v1.SslCertificate) error
}

type SslCertificateReconcileLoop interface {
	RunSslCertificateReconciler(ctx context.Context, rec SslCertificateReconciler, predicates ...predicate.Predicate) error
}

type sslCertificateReconcileLoop struct {
	loop reconcile.Loop
}

func NewSslCertificateReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) SslCertificateReconcileLoop {
	return &sslCertificateReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &gate_v1.SslCertificate{}, options),
	}
}

func (c *sslCertificateReconcileLoop) RunSslCertificateReconciler(ctx context.Context, reconciler SslCertificateReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericSslCertificateReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(SslCertificateFinalizer); ok {
		reconcilerWrapper = genericSslCertificateFinalizer{
			genericSslCertificateReconciler: genericReconciler,
			finalizingReconciler:            finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericSslCertificateHandler implements a generic reconcile.Reconciler
type genericSslCertificateReconciler struct {
	reconciler SslCertificateReconciler
}

func (r genericSslCertificateReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gate_v1.SslCertificate)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: SslCertificate handler received event for %T", object)
	}
	return r.reconciler.ReconcileSslCertificate(obj)
}

func (r genericSslCertificateReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(SslCertificateDeletionReconciler); ok {
		return deletionReconciler.ReconcileSslCertificateDeletion(request)
	}
	return nil
}

// genericSslCertificateFinalizer implements a generic reconcile.FinalizingReconciler
type genericSslCertificateFinalizer struct {
	genericSslCertificateReconciler
	finalizingReconciler SslCertificateFinalizer
}

func (r genericSslCertificateFinalizer) FinalizerName() string {
	return r.finalizingReconciler.SslCertificateFinalizerName()
}

func (r genericSslCertificateFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*gate_v1.SslCertificate)
	if !ok {
		return errors.Errorf("internal error: SslCertificate handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeSslCertificate(obj)
}

// Reconcile Upsert events for the Upstream Resource.
// implemented by the user
type UpstreamReconciler interface {
	ReconcileUpstream(obj *gate_v1.Upstream) (reconcile.Result, error)
}

// Reconcile deletion events for the Upstream Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type UpstreamDeletionReconciler interface {
	ReconcileUpstreamDeletion(req reconcile.Request) error
}

type UpstreamReconcilerFuncs struct {
	OnReconcileUpstream         func(obj *gate_v1.Upstream) (reconcile.Result, error)
	OnReconcileUpstreamDeletion func(req reconcile.Request) error
}

func (f *UpstreamReconcilerFuncs) ReconcileUpstream(obj *gate_v1.Upstream) (reconcile.Result, error) {
	if f.OnReconcileUpstream == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileUpstream(obj)
}

func (f *UpstreamReconcilerFuncs) ReconcileUpstreamDeletion(req reconcile.Request) error {
	if f.OnReconcileUpstreamDeletion == nil {
		return nil
	}
	return f.OnReconcileUpstreamDeletion(req)
}

// Reconcile and finalize the Upstream Resource
// implemented by the user
type UpstreamFinalizer interface {
	UpstreamReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	UpstreamFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeUpstream(obj *gate_v1.Upstream) error
}

type UpstreamReconcileLoop interface {
	RunUpstreamReconciler(ctx context.Context, rec UpstreamReconciler, predicates ...predicate.Predicate) error
}

type upstreamReconcileLoop struct {
	loop reconcile.Loop
}

func NewUpstreamReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) UpstreamReconcileLoop {
	return &upstreamReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &gate_v1.Upstream{}, options),
	}
}

func (c *upstreamReconcileLoop) RunUpstreamReconciler(ctx context.Context, reconciler UpstreamReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericUpstreamReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(UpstreamFinalizer); ok {
		reconcilerWrapper = genericUpstreamFinalizer{
			genericUpstreamReconciler: genericReconciler,
			finalizingReconciler:      finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericUpstreamHandler implements a generic reconcile.Reconciler
type genericUpstreamReconciler struct {
	reconciler UpstreamReconciler
}

func (r genericUpstreamReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gate_v1.Upstream)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Upstream handler received event for %T", object)
	}
	return r.reconciler.ReconcileUpstream(obj)
}

func (r genericUpstreamReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(UpstreamDeletionReconciler); ok {
		return deletionReconciler.ReconcileUpstreamDeletion(request)
	}
	return nil
}

// genericUpstreamFinalizer implements a generic reconcile.FinalizingReconciler
type genericUpstreamFinalizer struct {
	genericUpstreamReconciler
	finalizingReconciler UpstreamFinalizer
}

func (r genericUpstreamFinalizer) FinalizerName() string {
	return r.finalizingReconciler.UpstreamFinalizerName()
}

func (r genericUpstreamFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*gate_v1.Upstream)
	if !ok {
		return errors.Errorf("internal error: Upstream handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeUpstream(obj)
}

// Reconcile Upsert events for the Gateway Resource.
// implemented by the user
type GatewayReconciler interface {
	ReconcileGateway(obj *gate_v1.Gateway) (reconcile.Result, error)
}

// Reconcile deletion events for the Gateway Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type GatewayDeletionReconciler interface {
	ReconcileGatewayDeletion(req reconcile.Request) error
}

type GatewayReconcilerFuncs struct {
	OnReconcileGateway         func(obj *gate_v1.Gateway) (reconcile.Result, error)
	OnReconcileGatewayDeletion func(req reconcile.Request) error
}

func (f *GatewayReconcilerFuncs) ReconcileGateway(obj *gate_v1.Gateway) (reconcile.Result, error) {
	if f.OnReconcileGateway == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileGateway(obj)
}

func (f *GatewayReconcilerFuncs) ReconcileGatewayDeletion(req reconcile.Request) error {
	if f.OnReconcileGatewayDeletion == nil {
		return nil
	}
	return f.OnReconcileGatewayDeletion(req)
}

// Reconcile and finalize the Gateway Resource
// implemented by the user
type GatewayFinalizer interface {
	GatewayReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	GatewayFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeGateway(obj *gate_v1.Gateway) error
}

type GatewayReconcileLoop interface {
	RunGatewayReconciler(ctx context.Context, rec GatewayReconciler, predicates ...predicate.Predicate) error
}

type gatewayReconcileLoop struct {
	loop reconcile.Loop
}

func NewGatewayReconcileLoop(name string, mgr manager.Manager, options reconcile.Options) GatewayReconcileLoop {
	return &gatewayReconcileLoop{
		// empty cluster indicates this reconciler is built for the local cluster
		loop: reconcile.NewLoop(name, "", mgr, &gate_v1.Gateway{}, options),
	}
}

func (c *gatewayReconcileLoop) RunGatewayReconciler(ctx context.Context, reconciler GatewayReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericGatewayReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(GatewayFinalizer); ok {
		reconcilerWrapper = genericGatewayFinalizer{
			genericGatewayReconciler: genericReconciler,
			finalizingReconciler:     finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericGatewayHandler implements a generic reconcile.Reconciler
type genericGatewayReconciler struct {
	reconciler GatewayReconciler
}

func (r genericGatewayReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*gate_v1.Gateway)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Gateway handler received event for %T", object)
	}
	return r.reconciler.ReconcileGateway(obj)
}

func (r genericGatewayReconciler) ReconcileDeletion(request reconcile.Request) error {
	if deletionReconciler, ok := r.reconciler.(GatewayDeletionReconciler); ok {
		return deletionReconciler.ReconcileGatewayDeletion(request)
	}
	return nil
}

// genericGatewayFinalizer implements a generic reconcile.FinalizingReconciler
type genericGatewayFinalizer struct {
	genericGatewayReconciler
	finalizingReconciler GatewayFinalizer
}

func (r genericGatewayFinalizer) FinalizerName() string {
	return r.finalizingReconciler.GatewayFinalizerName()
}

func (r genericGatewayFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*gate_v1.Gateway)
	if !ok {
		return errors.Errorf("internal error: Gateway handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeGateway(obj)
}
