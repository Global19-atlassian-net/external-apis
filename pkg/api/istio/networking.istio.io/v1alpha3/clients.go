// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1alpha3

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	networking_istio_io_v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the networking.istio.io/v1alpha3 APIs
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

// clienset for the networking.istio.io/v1alpha3 APIs
type Clientset interface {
	// clienset for the networking.istio.io/v1alpha3/v1alpha3 APIs
	DestinationRules() DestinationRuleClient
	// clienset for the networking.istio.io/v1alpha3/v1alpha3 APIs
	EnvoyFilters() EnvoyFilterClient
	// clienset for the networking.istio.io/v1alpha3/v1alpha3 APIs
	Gateways() GatewayClient
	// clienset for the networking.istio.io/v1alpha3/v1alpha3 APIs
	ServiceEntries() ServiceEntryClient
	// clienset for the networking.istio.io/v1alpha3/v1alpha3 APIs
	VirtualServices() VirtualServiceClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := AddToScheme(scheme); err != nil {
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

// clienset for the networking.istio.io/v1alpha3/v1alpha3 APIs
func (c *clientSet) DestinationRules() DestinationRuleClient {
	return NewDestinationRuleClient(c.client)
}

// clienset for the networking.istio.io/v1alpha3/v1alpha3 APIs
func (c *clientSet) EnvoyFilters() EnvoyFilterClient {
	return NewEnvoyFilterClient(c.client)
}

// clienset for the networking.istio.io/v1alpha3/v1alpha3 APIs
func (c *clientSet) Gateways() GatewayClient {
	return NewGatewayClient(c.client)
}

// clienset for the networking.istio.io/v1alpha3/v1alpha3 APIs
func (c *clientSet) ServiceEntries() ServiceEntryClient {
	return NewServiceEntryClient(c.client)
}

// clienset for the networking.istio.io/v1alpha3/v1alpha3 APIs
func (c *clientSet) VirtualServices() VirtualServiceClient {
	return NewVirtualServiceClient(c.client)
}

// Reader knows how to read and list DestinationRules.
type DestinationRuleReader interface {
	// Get retrieves a DestinationRule for the given object key
	GetDestinationRule(ctx context.Context, key client.ObjectKey) (*networking_istio_io_v1alpha3.DestinationRule, error)

	// List retrieves list of DestinationRules for a given namespace and list options.
	ListDestinationRule(ctx context.Context, opts ...client.ListOption) (*networking_istio_io_v1alpha3.DestinationRuleList, error)
}

// DestinationRuleTransitionFunction instructs the DestinationRuleWriter how to transition between an existing
// DestinationRule object and a desired on an Upsert
type DestinationRuleTransitionFunction func(existing, desired *networking_istio_io_v1alpha3.DestinationRule) error

// Writer knows how to create, delete, and update DestinationRules.
type DestinationRuleWriter interface {
	// Create saves the DestinationRule object.
	CreateDestinationRule(ctx context.Context, obj *networking_istio_io_v1alpha3.DestinationRule, opts ...client.CreateOption) error

	// Delete deletes the DestinationRule object.
	DeleteDestinationRule(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given DestinationRule object.
	UpdateDestinationRule(ctx context.Context, obj *networking_istio_io_v1alpha3.DestinationRule, opts ...client.UpdateOption) error

	// Patch patches the given DestinationRule object.
	PatchDestinationRule(ctx context.Context, obj *networking_istio_io_v1alpha3.DestinationRule, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all DestinationRule objects matching the given options.
	DeleteAllOfDestinationRule(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the DestinationRule object.
	UpsertDestinationRule(ctx context.Context, obj *networking_istio_io_v1alpha3.DestinationRule, transitionFuncs ...DestinationRuleTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a DestinationRule object.
type DestinationRuleStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given DestinationRule object.
	UpdateDestinationRuleStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.DestinationRule, opts ...client.UpdateOption) error

	// Patch patches the given DestinationRule object's subresource.
	PatchDestinationRuleStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.DestinationRule, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on DestinationRules.
type DestinationRuleClient interface {
	DestinationRuleReader
	DestinationRuleWriter
	DestinationRuleStatusWriter
}

type destinationRuleClient struct {
	client client.Client
}

func NewDestinationRuleClient(client client.Client) *destinationRuleClient {
	return &destinationRuleClient{client: client}
}

func (c *destinationRuleClient) GetDestinationRule(ctx context.Context, key client.ObjectKey) (*networking_istio_io_v1alpha3.DestinationRule, error) {
	obj := &DestinationRule{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *destinationRuleClient) ListDestinationRule(ctx context.Context, opts ...client.ListOption) (*networking_istio_io_v1alpha3.DestinationRuleList, error) {
	list := &DestinationRuleList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *destinationRuleClient) CreateDestinationRule(ctx context.Context, obj *networking_istio_io_v1alpha3.DestinationRule, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *destinationRuleClient) DeleteDestinationRule(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &DestinationRule{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *destinationRuleClient) UpdateDestinationRule(ctx context.Context, obj *networking_istio_io_v1alpha3.DestinationRule, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *destinationRuleClient) PatchDestinationRule(ctx context.Context, obj *networking_istio_io_v1alpha3.DestinationRule, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *destinationRuleClient) DeleteAllOfDestinationRule(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &DestinationRule{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *destinationRuleClient) UpsertDestinationRule(ctx context.Context, obj *networking_istio_io_v1alpha3.DestinationRule, transitionFuncs ...DestinationRuleTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*networking_istio_io_v1alpha3.DestinationRule), desired.(*networking_istio_io_v1alpha3.DestinationRule)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *destinationRuleClient) UpdateDestinationRuleStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.DestinationRule, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *destinationRuleClient) PatchDestinationRuleStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.DestinationRule, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list EnvoyFilters.
type EnvoyFilterReader interface {
	// Get retrieves a EnvoyFilter for the given object key
	GetEnvoyFilter(ctx context.Context, key client.ObjectKey) (*networking_istio_io_v1alpha3.EnvoyFilter, error)

	// List retrieves list of EnvoyFilters for a given namespace and list options.
	ListEnvoyFilter(ctx context.Context, opts ...client.ListOption) (*networking_istio_io_v1alpha3.EnvoyFilterList, error)
}

// EnvoyFilterTransitionFunction instructs the EnvoyFilterWriter how to transition between an existing
// EnvoyFilter object and a desired on an Upsert
type EnvoyFilterTransitionFunction func(existing, desired *networking_istio_io_v1alpha3.EnvoyFilter) error

// Writer knows how to create, delete, and update EnvoyFilters.
type EnvoyFilterWriter interface {
	// Create saves the EnvoyFilter object.
	CreateEnvoyFilter(ctx context.Context, obj *networking_istio_io_v1alpha3.EnvoyFilter, opts ...client.CreateOption) error

	// Delete deletes the EnvoyFilter object.
	DeleteEnvoyFilter(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given EnvoyFilter object.
	UpdateEnvoyFilter(ctx context.Context, obj *networking_istio_io_v1alpha3.EnvoyFilter, opts ...client.UpdateOption) error

	// Patch patches the given EnvoyFilter object.
	PatchEnvoyFilter(ctx context.Context, obj *networking_istio_io_v1alpha3.EnvoyFilter, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all EnvoyFilter objects matching the given options.
	DeleteAllOfEnvoyFilter(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the EnvoyFilter object.
	UpsertEnvoyFilter(ctx context.Context, obj *networking_istio_io_v1alpha3.EnvoyFilter, transitionFuncs ...EnvoyFilterTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a EnvoyFilter object.
type EnvoyFilterStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given EnvoyFilter object.
	UpdateEnvoyFilterStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.EnvoyFilter, opts ...client.UpdateOption) error

	// Patch patches the given EnvoyFilter object's subresource.
	PatchEnvoyFilterStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.EnvoyFilter, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on EnvoyFilters.
type EnvoyFilterClient interface {
	EnvoyFilterReader
	EnvoyFilterWriter
	EnvoyFilterStatusWriter
}

type envoyFilterClient struct {
	client client.Client
}

func NewEnvoyFilterClient(client client.Client) *envoyFilterClient {
	return &envoyFilterClient{client: client}
}

func (c *envoyFilterClient) GetEnvoyFilter(ctx context.Context, key client.ObjectKey) (*networking_istio_io_v1alpha3.EnvoyFilter, error) {
	obj := &EnvoyFilter{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *envoyFilterClient) ListEnvoyFilter(ctx context.Context, opts ...client.ListOption) (*networking_istio_io_v1alpha3.EnvoyFilterList, error) {
	list := &EnvoyFilterList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *envoyFilterClient) CreateEnvoyFilter(ctx context.Context, obj *networking_istio_io_v1alpha3.EnvoyFilter, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *envoyFilterClient) DeleteEnvoyFilter(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &EnvoyFilter{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *envoyFilterClient) UpdateEnvoyFilter(ctx context.Context, obj *networking_istio_io_v1alpha3.EnvoyFilter, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *envoyFilterClient) PatchEnvoyFilter(ctx context.Context, obj *networking_istio_io_v1alpha3.EnvoyFilter, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *envoyFilterClient) DeleteAllOfEnvoyFilter(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &EnvoyFilter{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *envoyFilterClient) UpsertEnvoyFilter(ctx context.Context, obj *networking_istio_io_v1alpha3.EnvoyFilter, transitionFuncs ...EnvoyFilterTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*networking_istio_io_v1alpha3.EnvoyFilter), desired.(*networking_istio_io_v1alpha3.EnvoyFilter)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *envoyFilterClient) UpdateEnvoyFilterStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.EnvoyFilter, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *envoyFilterClient) PatchEnvoyFilterStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.EnvoyFilter, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list Gateways.
type GatewayReader interface {
	// Get retrieves a Gateway for the given object key
	GetGateway(ctx context.Context, key client.ObjectKey) (*networking_istio_io_v1alpha3.Gateway, error)

	// List retrieves list of Gateways for a given namespace and list options.
	ListGateway(ctx context.Context, opts ...client.ListOption) (*networking_istio_io_v1alpha3.GatewayList, error)
}

// GatewayTransitionFunction instructs the GatewayWriter how to transition between an existing
// Gateway object and a desired on an Upsert
type GatewayTransitionFunction func(existing, desired *networking_istio_io_v1alpha3.Gateway) error

// Writer knows how to create, delete, and update Gateways.
type GatewayWriter interface {
	// Create saves the Gateway object.
	CreateGateway(ctx context.Context, obj *networking_istio_io_v1alpha3.Gateway, opts ...client.CreateOption) error

	// Delete deletes the Gateway object.
	DeleteGateway(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Gateway object.
	UpdateGateway(ctx context.Context, obj *networking_istio_io_v1alpha3.Gateway, opts ...client.UpdateOption) error

	// Patch patches the given Gateway object.
	PatchGateway(ctx context.Context, obj *networking_istio_io_v1alpha3.Gateway, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Gateway objects matching the given options.
	DeleteAllOfGateway(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Gateway object.
	UpsertGateway(ctx context.Context, obj *networking_istio_io_v1alpha3.Gateway, transitionFuncs ...GatewayTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Gateway object.
type GatewayStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Gateway object.
	UpdateGatewayStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.Gateway, opts ...client.UpdateOption) error

	// Patch patches the given Gateway object's subresource.
	PatchGatewayStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.Gateway, patch client.Patch, opts ...client.PatchOption) error
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

func (c *gatewayClient) GetGateway(ctx context.Context, key client.ObjectKey) (*networking_istio_io_v1alpha3.Gateway, error) {
	obj := &Gateway{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *gatewayClient) ListGateway(ctx context.Context, opts ...client.ListOption) (*networking_istio_io_v1alpha3.GatewayList, error) {
	list := &GatewayList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *gatewayClient) CreateGateway(ctx context.Context, obj *networking_istio_io_v1alpha3.Gateway, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *gatewayClient) DeleteGateway(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Gateway{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *gatewayClient) UpdateGateway(ctx context.Context, obj *networking_istio_io_v1alpha3.Gateway, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *gatewayClient) PatchGateway(ctx context.Context, obj *networking_istio_io_v1alpha3.Gateway, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *gatewayClient) DeleteAllOfGateway(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Gateway{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *gatewayClient) UpsertGateway(ctx context.Context, obj *networking_istio_io_v1alpha3.Gateway, transitionFuncs ...GatewayTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*networking_istio_io_v1alpha3.Gateway), desired.(*networking_istio_io_v1alpha3.Gateway)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *gatewayClient) UpdateGatewayStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.Gateway, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *gatewayClient) PatchGatewayStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.Gateway, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list ServiceEntrys.
type ServiceEntryReader interface {
	// Get retrieves a ServiceEntry for the given object key
	GetServiceEntry(ctx context.Context, key client.ObjectKey) (*networking_istio_io_v1alpha3.ServiceEntry, error)

	// List retrieves list of ServiceEntrys for a given namespace and list options.
	ListServiceEntry(ctx context.Context, opts ...client.ListOption) (*networking_istio_io_v1alpha3.ServiceEntryList, error)
}

// ServiceEntryTransitionFunction instructs the ServiceEntryWriter how to transition between an existing
// ServiceEntry object and a desired on an Upsert
type ServiceEntryTransitionFunction func(existing, desired *networking_istio_io_v1alpha3.ServiceEntry) error

// Writer knows how to create, delete, and update ServiceEntrys.
type ServiceEntryWriter interface {
	// Create saves the ServiceEntry object.
	CreateServiceEntry(ctx context.Context, obj *networking_istio_io_v1alpha3.ServiceEntry, opts ...client.CreateOption) error

	// Delete deletes the ServiceEntry object.
	DeleteServiceEntry(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given ServiceEntry object.
	UpdateServiceEntry(ctx context.Context, obj *networking_istio_io_v1alpha3.ServiceEntry, opts ...client.UpdateOption) error

	// Patch patches the given ServiceEntry object.
	PatchServiceEntry(ctx context.Context, obj *networking_istio_io_v1alpha3.ServiceEntry, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ServiceEntry objects matching the given options.
	DeleteAllOfServiceEntry(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the ServiceEntry object.
	UpsertServiceEntry(ctx context.Context, obj *networking_istio_io_v1alpha3.ServiceEntry, transitionFuncs ...ServiceEntryTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a ServiceEntry object.
type ServiceEntryStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ServiceEntry object.
	UpdateServiceEntryStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.ServiceEntry, opts ...client.UpdateOption) error

	// Patch patches the given ServiceEntry object's subresource.
	PatchServiceEntryStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.ServiceEntry, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on ServiceEntrys.
type ServiceEntryClient interface {
	ServiceEntryReader
	ServiceEntryWriter
	ServiceEntryStatusWriter
}

type serviceEntryClient struct {
	client client.Client
}

func NewServiceEntryClient(client client.Client) *serviceEntryClient {
	return &serviceEntryClient{client: client}
}

func (c *serviceEntryClient) GetServiceEntry(ctx context.Context, key client.ObjectKey) (*networking_istio_io_v1alpha3.ServiceEntry, error) {
	obj := &ServiceEntry{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *serviceEntryClient) ListServiceEntry(ctx context.Context, opts ...client.ListOption) (*networking_istio_io_v1alpha3.ServiceEntryList, error) {
	list := &ServiceEntryList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *serviceEntryClient) CreateServiceEntry(ctx context.Context, obj *networking_istio_io_v1alpha3.ServiceEntry, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *serviceEntryClient) DeleteServiceEntry(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &ServiceEntry{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *serviceEntryClient) UpdateServiceEntry(ctx context.Context, obj *networking_istio_io_v1alpha3.ServiceEntry, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *serviceEntryClient) PatchServiceEntry(ctx context.Context, obj *networking_istio_io_v1alpha3.ServiceEntry, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *serviceEntryClient) DeleteAllOfServiceEntry(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &ServiceEntry{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *serviceEntryClient) UpsertServiceEntry(ctx context.Context, obj *networking_istio_io_v1alpha3.ServiceEntry, transitionFuncs ...ServiceEntryTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*networking_istio_io_v1alpha3.ServiceEntry), desired.(*networking_istio_io_v1alpha3.ServiceEntry)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *serviceEntryClient) UpdateServiceEntryStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.ServiceEntry, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *serviceEntryClient) PatchServiceEntryStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.ServiceEntry, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list VirtualServices.
type VirtualServiceReader interface {
	// Get retrieves a VirtualService for the given object key
	GetVirtualService(ctx context.Context, key client.ObjectKey) (*networking_istio_io_v1alpha3.VirtualService, error)

	// List retrieves list of VirtualServices for a given namespace and list options.
	ListVirtualService(ctx context.Context, opts ...client.ListOption) (*networking_istio_io_v1alpha3.VirtualServiceList, error)
}

// VirtualServiceTransitionFunction instructs the VirtualServiceWriter how to transition between an existing
// VirtualService object and a desired on an Upsert
type VirtualServiceTransitionFunction func(existing, desired *networking_istio_io_v1alpha3.VirtualService) error

// Writer knows how to create, delete, and update VirtualServices.
type VirtualServiceWriter interface {
	// Create saves the VirtualService object.
	CreateVirtualService(ctx context.Context, obj *networking_istio_io_v1alpha3.VirtualService, opts ...client.CreateOption) error

	// Delete deletes the VirtualService object.
	DeleteVirtualService(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given VirtualService object.
	UpdateVirtualService(ctx context.Context, obj *networking_istio_io_v1alpha3.VirtualService, opts ...client.UpdateOption) error

	// Patch patches the given VirtualService object.
	PatchVirtualService(ctx context.Context, obj *networking_istio_io_v1alpha3.VirtualService, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all VirtualService objects matching the given options.
	DeleteAllOfVirtualService(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the VirtualService object.
	UpsertVirtualService(ctx context.Context, obj *networking_istio_io_v1alpha3.VirtualService, transitionFuncs ...VirtualServiceTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a VirtualService object.
type VirtualServiceStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given VirtualService object.
	UpdateVirtualServiceStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.VirtualService, opts ...client.UpdateOption) error

	// Patch patches the given VirtualService object's subresource.
	PatchVirtualServiceStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.VirtualService, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on VirtualServices.
type VirtualServiceClient interface {
	VirtualServiceReader
	VirtualServiceWriter
	VirtualServiceStatusWriter
}

type virtualServiceClient struct {
	client client.Client
}

func NewVirtualServiceClient(client client.Client) *virtualServiceClient {
	return &virtualServiceClient{client: client}
}

func (c *virtualServiceClient) GetVirtualService(ctx context.Context, key client.ObjectKey) (*networking_istio_io_v1alpha3.VirtualService, error) {
	obj := &VirtualService{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *virtualServiceClient) ListVirtualService(ctx context.Context, opts ...client.ListOption) (*networking_istio_io_v1alpha3.VirtualServiceList, error) {
	list := &VirtualServiceList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *virtualServiceClient) CreateVirtualService(ctx context.Context, obj *networking_istio_io_v1alpha3.VirtualService, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *virtualServiceClient) DeleteVirtualService(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &VirtualService{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *virtualServiceClient) UpdateVirtualService(ctx context.Context, obj *networking_istio_io_v1alpha3.VirtualService, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *virtualServiceClient) PatchVirtualService(ctx context.Context, obj *networking_istio_io_v1alpha3.VirtualService, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *virtualServiceClient) DeleteAllOfVirtualService(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &VirtualService{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *virtualServiceClient) UpsertVirtualService(ctx context.Context, obj *networking_istio_io_v1alpha3.VirtualService, transitionFuncs ...VirtualServiceTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*networking_istio_io_v1alpha3.VirtualService), desired.(*networking_istio_io_v1alpha3.VirtualService)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *virtualServiceClient) UpdateVirtualServiceStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.VirtualService, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *virtualServiceClient) PatchVirtualServiceStatus(ctx context.Context, obj *networking_istio_io_v1alpha3.VirtualService, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}