// Code generated by engine gate build no edit

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the gate/v1 APIs
type MulticlusterClientset interface {
	// Cluster returns a Clientset for the given cluster
	Cluster(cluster string) (Clientset, error)
}

type multiclusterClientset struct {
	client multicluster.Client
}

func NewMulticlusterClientset(client multicluster.Client) MulticlusterClientset {
	return &multiclusterClientset{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

// clienset for the gate/v1 APIs
type Clientset interface {
	// clienset for the gate/v1/v1 APIs
	CaCertificates() CaCertificateClient
	// clienset for the gate/v1/v1 APIs
	SslCertificates() SslCertificateClient
	// clienset for the gate/v1/v1 APIs
	Upstreams() UpstreamClient
	// clienset for the gate/v1/v1 APIs
	Gateways() GatewayClient
	// clienset for the gate/v1/v1 APIs
	Filters() FilterClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := SchemeBuilder.AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the gate/v1/v1 APIs
func (c *clientSet) CaCertificates() CaCertificateClient {
	return NewCaCertificateClient(c.client)
}

// clienset for the gate/v1/v1 APIs
func (c *clientSet) SslCertificates() SslCertificateClient {
	return NewSslCertificateClient(c.client)
}

// clienset for the gate/v1/v1 APIs
func (c *clientSet) Upstreams() UpstreamClient {
	return NewUpstreamClient(c.client)
}

// clienset for the gate/v1/v1 APIs
func (c *clientSet) Gateways() GatewayClient {
	return NewGatewayClient(c.client)
}

// clienset for the gate/v1/v1 APIs
func (c *clientSet) Filters() FilterClient {
	return NewFilterClient(c.client)
}

// Reader knows how to read and list CaCertificates.
type CaCertificateReader interface {
	// Get retrieves a CaCertificate for the given object key
	GetCaCertificate(ctx context.Context, key client.ObjectKey) (*CaCertificate, error)

	// List retrieves list of CaCertificates for a given namespace and list options.
	ListCaCertificate(ctx context.Context, opts ...client.ListOption) (*CaCertificateList, error)
}

// CaCertificateTransitionFunction instructs the CaCertificateWriter how to transition between an existing
// CaCertificate object and a desired on an Upsert
type CaCertificateTransitionFunction func(existing, desired *CaCertificate) error

// Writer knows how to create, delete, and update CaCertificates.
type CaCertificateWriter interface {
	// Create saves the CaCertificate object.
	CreateCaCertificate(ctx context.Context, obj *CaCertificate, opts ...client.CreateOption) error

	// Delete deletes the CaCertificate object.
	DeleteCaCertificate(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given CaCertificate object.
	UpdateCaCertificate(ctx context.Context, obj *CaCertificate, opts ...client.UpdateOption) error

	// Patch patches the given CaCertificate object.
	PatchCaCertificate(ctx context.Context, obj *CaCertificate, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all CaCertificate objects matching the given options.
	DeleteAllOfCaCertificate(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the CaCertificate object.
	UpsertCaCertificate(ctx context.Context, obj *CaCertificate, transitionFuncs ...CaCertificateTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a CaCertificate object.
type CaCertificateStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given CaCertificate object.
	UpdateCaCertificateStatus(ctx context.Context, obj *CaCertificate, opts ...client.UpdateOption) error

	// Patch patches the given CaCertificate object's subresource.
	PatchCaCertificateStatus(ctx context.Context, obj *CaCertificate, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on CaCertificates.
type CaCertificateClient interface {
	CaCertificateReader
	CaCertificateWriter
	CaCertificateStatusWriter
}

type caCertificateClient struct {
	client client.Client
}

func NewCaCertificateClient(client client.Client) *caCertificateClient {
	return &caCertificateClient{client: client}
}

func (c *caCertificateClient) GetCaCertificate(ctx context.Context, key client.ObjectKey) (*CaCertificate, error) {
	obj := &CaCertificate{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *caCertificateClient) ListCaCertificate(ctx context.Context, opts ...client.ListOption) (*CaCertificateList, error) {
	list := &CaCertificateList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *caCertificateClient) CreateCaCertificate(ctx context.Context, obj *CaCertificate, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *caCertificateClient) DeleteCaCertificate(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &CaCertificate{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *caCertificateClient) UpdateCaCertificate(ctx context.Context, obj *CaCertificate, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *caCertificateClient) PatchCaCertificate(ctx context.Context, obj *CaCertificate, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *caCertificateClient) DeleteAllOfCaCertificate(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &CaCertificate{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *caCertificateClient) UpsertCaCertificate(ctx context.Context, obj *CaCertificate, transitionFuncs ...CaCertificateTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*CaCertificate), desired.(*CaCertificate)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *caCertificateClient) UpdateCaCertificateStatus(ctx context.Context, obj *CaCertificate, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *caCertificateClient) PatchCaCertificateStatus(ctx context.Context, obj *CaCertificate, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides CaCertificateClients for multiple clusters.
type MulticlusterCaCertificateClient interface {
	// Cluster returns a CaCertificateClient for the given cluster
	Cluster(cluster string) (CaCertificateClient, error)
}

type multiclusterCaCertificateClient struct {
	client multicluster.Client
}

func NewMulticlusterCaCertificateClient(client multicluster.Client) MulticlusterCaCertificateClient {
	return &multiclusterCaCertificateClient{client: client}
}

func (m *multiclusterCaCertificateClient) Cluster(cluster string) (CaCertificateClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewCaCertificateClient(client), nil
}

// Reader knows how to read and list SslCertificates.
type SslCertificateReader interface {
	// Get retrieves a SslCertificate for the given object key
	GetSslCertificate(ctx context.Context, key client.ObjectKey) (*SslCertificate, error)

	// List retrieves list of SslCertificates for a given namespace and list options.
	ListSslCertificate(ctx context.Context, opts ...client.ListOption) (*SslCertificateList, error)
}

// SslCertificateTransitionFunction instructs the SslCertificateWriter how to transition between an existing
// SslCertificate object and a desired on an Upsert
type SslCertificateTransitionFunction func(existing, desired *SslCertificate) error

// Writer knows how to create, delete, and update SslCertificates.
type SslCertificateWriter interface {
	// Create saves the SslCertificate object.
	CreateSslCertificate(ctx context.Context, obj *SslCertificate, opts ...client.CreateOption) error

	// Delete deletes the SslCertificate object.
	DeleteSslCertificate(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given SslCertificate object.
	UpdateSslCertificate(ctx context.Context, obj *SslCertificate, opts ...client.UpdateOption) error

	// Patch patches the given SslCertificate object.
	PatchSslCertificate(ctx context.Context, obj *SslCertificate, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all SslCertificate objects matching the given options.
	DeleteAllOfSslCertificate(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the SslCertificate object.
	UpsertSslCertificate(ctx context.Context, obj *SslCertificate, transitionFuncs ...SslCertificateTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a SslCertificate object.
type SslCertificateStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given SslCertificate object.
	UpdateSslCertificateStatus(ctx context.Context, obj *SslCertificate, opts ...client.UpdateOption) error

	// Patch patches the given SslCertificate object's subresource.
	PatchSslCertificateStatus(ctx context.Context, obj *SslCertificate, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on SslCertificates.
type SslCertificateClient interface {
	SslCertificateReader
	SslCertificateWriter
	SslCertificateStatusWriter
}

type sslCertificateClient struct {
	client client.Client
}

func NewSslCertificateClient(client client.Client) *sslCertificateClient {
	return &sslCertificateClient{client: client}
}

func (c *sslCertificateClient) GetSslCertificate(ctx context.Context, key client.ObjectKey) (*SslCertificate, error) {
	obj := &SslCertificate{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *sslCertificateClient) ListSslCertificate(ctx context.Context, opts ...client.ListOption) (*SslCertificateList, error) {
	list := &SslCertificateList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *sslCertificateClient) CreateSslCertificate(ctx context.Context, obj *SslCertificate, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *sslCertificateClient) DeleteSslCertificate(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &SslCertificate{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *sslCertificateClient) UpdateSslCertificate(ctx context.Context, obj *SslCertificate, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *sslCertificateClient) PatchSslCertificate(ctx context.Context, obj *SslCertificate, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *sslCertificateClient) DeleteAllOfSslCertificate(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &SslCertificate{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *sslCertificateClient) UpsertSslCertificate(ctx context.Context, obj *SslCertificate, transitionFuncs ...SslCertificateTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*SslCertificate), desired.(*SslCertificate)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *sslCertificateClient) UpdateSslCertificateStatus(ctx context.Context, obj *SslCertificate, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *sslCertificateClient) PatchSslCertificateStatus(ctx context.Context, obj *SslCertificate, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides SslCertificateClients for multiple clusters.
type MulticlusterSslCertificateClient interface {
	// Cluster returns a SslCertificateClient for the given cluster
	Cluster(cluster string) (SslCertificateClient, error)
}

type multiclusterSslCertificateClient struct {
	client multicluster.Client
}

func NewMulticlusterSslCertificateClient(client multicluster.Client) MulticlusterSslCertificateClient {
	return &multiclusterSslCertificateClient{client: client}
}

func (m *multiclusterSslCertificateClient) Cluster(cluster string) (SslCertificateClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewSslCertificateClient(client), nil
}

// Reader knows how to read and list Upstreams.
type UpstreamReader interface {
	// Get retrieves a Upstream for the given object key
	GetUpstream(ctx context.Context, key client.ObjectKey) (*Upstream, error)

	// List retrieves list of Upstreams for a given namespace and list options.
	ListUpstream(ctx context.Context, opts ...client.ListOption) (*UpstreamList, error)
}

// UpstreamTransitionFunction instructs the UpstreamWriter how to transition between an existing
// Upstream object and a desired on an Upsert
type UpstreamTransitionFunction func(existing, desired *Upstream) error

// Writer knows how to create, delete, and update Upstreams.
type UpstreamWriter interface {
	// Create saves the Upstream object.
	CreateUpstream(ctx context.Context, obj *Upstream, opts ...client.CreateOption) error

	// Delete deletes the Upstream object.
	DeleteUpstream(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Upstream object.
	UpdateUpstream(ctx context.Context, obj *Upstream, opts ...client.UpdateOption) error

	// Patch patches the given Upstream object.
	PatchUpstream(ctx context.Context, obj *Upstream, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Upstream objects matching the given options.
	DeleteAllOfUpstream(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Upstream object.
	UpsertUpstream(ctx context.Context, obj *Upstream, transitionFuncs ...UpstreamTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Upstream object.
type UpstreamStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Upstream object.
	UpdateUpstreamStatus(ctx context.Context, obj *Upstream, opts ...client.UpdateOption) error

	// Patch patches the given Upstream object's subresource.
	PatchUpstreamStatus(ctx context.Context, obj *Upstream, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Upstreams.
type UpstreamClient interface {
	UpstreamReader
	UpstreamWriter
	UpstreamStatusWriter
}

type upstreamClient struct {
	client client.Client
}

func NewUpstreamClient(client client.Client) *upstreamClient {
	return &upstreamClient{client: client}
}

func (c *upstreamClient) GetUpstream(ctx context.Context, key client.ObjectKey) (*Upstream, error) {
	obj := &Upstream{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *upstreamClient) ListUpstream(ctx context.Context, opts ...client.ListOption) (*UpstreamList, error) {
	list := &UpstreamList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *upstreamClient) CreateUpstream(ctx context.Context, obj *Upstream, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *upstreamClient) DeleteUpstream(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Upstream{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *upstreamClient) UpdateUpstream(ctx context.Context, obj *Upstream, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *upstreamClient) PatchUpstream(ctx context.Context, obj *Upstream, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *upstreamClient) DeleteAllOfUpstream(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Upstream{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *upstreamClient) UpsertUpstream(ctx context.Context, obj *Upstream, transitionFuncs ...UpstreamTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*Upstream), desired.(*Upstream)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *upstreamClient) UpdateUpstreamStatus(ctx context.Context, obj *Upstream, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *upstreamClient) PatchUpstreamStatus(ctx context.Context, obj *Upstream, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides UpstreamClients for multiple clusters.
type MulticlusterUpstreamClient interface {
	// Cluster returns a UpstreamClient for the given cluster
	Cluster(cluster string) (UpstreamClient, error)
}

type multiclusterUpstreamClient struct {
	client multicluster.Client
}

func NewMulticlusterUpstreamClient(client multicluster.Client) MulticlusterUpstreamClient {
	return &multiclusterUpstreamClient{client: client}
}

func (m *multiclusterUpstreamClient) Cluster(cluster string) (UpstreamClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewUpstreamClient(client), nil
}

// Reader knows how to read and list Gateways.
type GatewayReader interface {
	// Get retrieves a Gateway for the given object key
	GetGateway(ctx context.Context, key client.ObjectKey) (*Gateway, error)

	// List retrieves list of Gateways for a given namespace and list options.
	ListGateway(ctx context.Context, opts ...client.ListOption) (*GatewayList, error)
}

// GatewayTransitionFunction instructs the GatewayWriter how to transition between an existing
// Gateway object and a desired on an Upsert
type GatewayTransitionFunction func(existing, desired *Gateway) error

// Writer knows how to create, delete, and update Gateways.
type GatewayWriter interface {
	// Create saves the Gateway object.
	CreateGateway(ctx context.Context, obj *Gateway, opts ...client.CreateOption) error

	// Delete deletes the Gateway object.
	DeleteGateway(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Gateway object.
	UpdateGateway(ctx context.Context, obj *Gateway, opts ...client.UpdateOption) error

	// Patch patches the given Gateway object.
	PatchGateway(ctx context.Context, obj *Gateway, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Gateway objects matching the given options.
	DeleteAllOfGateway(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Gateway object.
	UpsertGateway(ctx context.Context, obj *Gateway, transitionFuncs ...GatewayTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Gateway object.
type GatewayStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Gateway object.
	UpdateGatewayStatus(ctx context.Context, obj *Gateway, opts ...client.UpdateOption) error

	// Patch patches the given Gateway object's subresource.
	PatchGatewayStatus(ctx context.Context, obj *Gateway, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Gateways.
type GatewayClient interface {
	GatewayReader
	GatewayWriter
	GatewayStatusWriter
}

type gatewayClient struct {
	client client.Client
}

func NewGatewayClient(client client.Client) *gatewayClient {
	return &gatewayClient{client: client}
}

func (c *gatewayClient) GetGateway(ctx context.Context, key client.ObjectKey) (*Gateway, error) {
	obj := &Gateway{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *gatewayClient) ListGateway(ctx context.Context, opts ...client.ListOption) (*GatewayList, error) {
	list := &GatewayList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *gatewayClient) CreateGateway(ctx context.Context, obj *Gateway, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *gatewayClient) DeleteGateway(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Gateway{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *gatewayClient) UpdateGateway(ctx context.Context, obj *Gateway, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *gatewayClient) PatchGateway(ctx context.Context, obj *Gateway, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *gatewayClient) DeleteAllOfGateway(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Gateway{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *gatewayClient) UpsertGateway(ctx context.Context, obj *Gateway, transitionFuncs ...GatewayTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*Gateway), desired.(*Gateway)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *gatewayClient) UpdateGatewayStatus(ctx context.Context, obj *Gateway, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *gatewayClient) PatchGatewayStatus(ctx context.Context, obj *Gateway, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides GatewayClients for multiple clusters.
type MulticlusterGatewayClient interface {
	// Cluster returns a GatewayClient for the given cluster
	Cluster(cluster string) (GatewayClient, error)
}

type multiclusterGatewayClient struct {
	client multicluster.Client
}

func NewMulticlusterGatewayClient(client multicluster.Client) MulticlusterGatewayClient {
	return &multiclusterGatewayClient{client: client}
}

func (m *multiclusterGatewayClient) Cluster(cluster string) (GatewayClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewGatewayClient(client), nil
}

// Reader knows how to read and list Filters.
type FilterReader interface {
	// Get retrieves a Filter for the given object key
	GetFilter(ctx context.Context, key client.ObjectKey) (*Filter, error)

	// List retrieves list of Filters for a given namespace and list options.
	ListFilter(ctx context.Context, opts ...client.ListOption) (*FilterList, error)
}

// FilterTransitionFunction instructs the FilterWriter how to transition between an existing
// Filter object and a desired on an Upsert
type FilterTransitionFunction func(existing, desired *Filter) error

// Writer knows how to create, delete, and update Filters.
type FilterWriter interface {
	// Create saves the Filter object.
	CreateFilter(ctx context.Context, obj *Filter, opts ...client.CreateOption) error

	// Delete deletes the Filter object.
	DeleteFilter(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Filter object.
	UpdateFilter(ctx context.Context, obj *Filter, opts ...client.UpdateOption) error

	// Patch patches the given Filter object.
	PatchFilter(ctx context.Context, obj *Filter, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Filter objects matching the given options.
	DeleteAllOfFilter(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Filter object.
	UpsertFilter(ctx context.Context, obj *Filter, transitionFuncs ...FilterTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Filter object.
type FilterStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Filter object.
	UpdateFilterStatus(ctx context.Context, obj *Filter, opts ...client.UpdateOption) error

	// Patch patches the given Filter object's subresource.
	PatchFilterStatus(ctx context.Context, obj *Filter, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Filters.
type FilterClient interface {
	FilterReader
	FilterWriter
	FilterStatusWriter
}

type filterClient struct {
	client client.Client
}

func NewFilterClient(client client.Client) *filterClient {
	return &filterClient{client: client}
}

func (c *filterClient) GetFilter(ctx context.Context, key client.ObjectKey) (*Filter, error) {
	obj := &Filter{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *filterClient) ListFilter(ctx context.Context, opts ...client.ListOption) (*FilterList, error) {
	list := &FilterList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *filterClient) CreateFilter(ctx context.Context, obj *Filter, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *filterClient) DeleteFilter(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Filter{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *filterClient) UpdateFilter(ctx context.Context, obj *Filter, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *filterClient) PatchFilter(ctx context.Context, obj *Filter, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *filterClient) DeleteAllOfFilter(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Filter{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *filterClient) UpsertFilter(ctx context.Context, obj *Filter, transitionFuncs ...FilterTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*Filter), desired.(*Filter)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *filterClient) UpdateFilterStatus(ctx context.Context, obj *Filter, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *filterClient) PatchFilterStatus(ctx context.Context, obj *Filter, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides FilterClients for multiple clusters.
type MulticlusterFilterClient interface {
	// Cluster returns a FilterClient for the given cluster
	Cluster(cluster string) (FilterClient, error)
}

type multiclusterFilterClient struct {
	client multicluster.Client
}

func NewMulticlusterFilterClient(client multicluster.Client) MulticlusterFilterClient {
	return &multiclusterFilterClient{client: client}
}

func (m *multiclusterFilterClient) Cluster(cluster string) (FilterClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewFilterClient(client), nil
}
