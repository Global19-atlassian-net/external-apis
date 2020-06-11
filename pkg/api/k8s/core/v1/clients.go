// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the /v1 APIs
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

// clienset for the /v1 APIs
type Clientset interface {
	// clienset for the v1/v1 APIs
	Secrets() SecretClient
	// clienset for the v1/v1 APIs
	ServiceAccounts() ServiceAccountClient
	// clienset for the v1/v1 APIs
	ConfigMaps() ConfigMapClient
	// clienset for the v1/v1 APIs
	Services() ServiceClient
	// clienset for the v1/v1 APIs
	Pods() PodClient
	// clienset for the v1/v1 APIs
	Namespaces() NamespaceClient
	// clienset for the v1/v1 APIs
	Nodes() NodeClient
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

// clienset for the v1/v1 APIs
func (c *clientSet) Secrets() SecretClient {
	return NewSecretClient(c.client)
}

// clienset for the v1/v1 APIs
func (c *clientSet) ServiceAccounts() ServiceAccountClient {
	return NewServiceAccountClient(c.client)
}

// clienset for the v1/v1 APIs
func (c *clientSet) ConfigMaps() ConfigMapClient {
	return NewConfigMapClient(c.client)
}

// clienset for the v1/v1 APIs
func (c *clientSet) Services() ServiceClient {
	return NewServiceClient(c.client)
}

// clienset for the v1/v1 APIs
func (c *clientSet) Pods() PodClient {
	return NewPodClient(c.client)
}

// clienset for the v1/v1 APIs
func (c *clientSet) Namespaces() NamespaceClient {
	return NewNamespaceClient(c.client)
}

// clienset for the v1/v1 APIs
func (c *clientSet) Nodes() NodeClient {
	return NewNodeClient(c.client)
}

// Reader knows how to read and list Secrets.
type SecretReader interface {
	// Get retrieves a Secret for the given object key
	GetSecret(ctx context.Context, key client.ObjectKey) (*v1.Secret, error)

	// List retrieves list of Secrets for a given namespace and list options.
	ListSecret(ctx context.Context, opts ...client.ListOption) (*v1.SecretList, error)
}

// SecretTransitionFunction instructs the SecretWriter how to transition between an existing
// Secret object and a desired on an Upsert
type SecretTransitionFunction func(existing, desired *v1.Secret) error

// Writer knows how to create, delete, and update Secrets.
type SecretWriter interface {
	// Create saves the Secret object.
	CreateSecret(ctx context.Context, obj *v1.Secret, opts ...client.CreateOption) error

	// Delete deletes the Secret object.
	DeleteSecret(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Secret object.
	UpdateSecret(ctx context.Context, obj *v1.Secret, opts ...client.UpdateOption) error

	// Patch patches the given Secret object.
	PatchSecret(ctx context.Context, obj *v1.Secret, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Secret objects matching the given options.
	DeleteAllOfSecret(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Secret object.
	UpsertSecret(ctx context.Context, obj *v1.Secret, transitionFuncs ...SecretTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Secret object.
type SecretStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Secret object.
	UpdateSecretStatus(ctx context.Context, obj *v1.Secret, opts ...client.UpdateOption) error

	// Patch patches the given Secret object's subresource.
	PatchSecretStatus(ctx context.Context, obj *v1.Secret, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Secrets.
type SecretClient interface {
	SecretReader
	SecretWriter
	SecretStatusWriter
}

type secretClient struct {
	client client.Client
}

func NewSecretClient(client client.Client) *secretClient {
	return &secretClient{client: client}
}

func (c *secretClient) GetSecret(ctx context.Context, key client.ObjectKey) (*v1.Secret, error) {
	obj := &Secret{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *secretClient) ListSecret(ctx context.Context, opts ...client.ListOption) (*v1.SecretList, error) {
	list := &SecretList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *secretClient) CreateSecret(ctx context.Context, obj *v1.Secret, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *secretClient) DeleteSecret(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Secret{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *secretClient) UpdateSecret(ctx context.Context, obj *v1.Secret, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *secretClient) PatchSecret(ctx context.Context, obj *v1.Secret, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *secretClient) DeleteAllOfSecret(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Secret{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *secretClient) UpsertSecret(ctx context.Context, obj *v1.Secret, transitionFuncs ...SecretTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*v1.Secret), desired.(*v1.Secret)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *secretClient) UpdateSecretStatus(ctx context.Context, obj *v1.Secret, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *secretClient) PatchSecretStatus(ctx context.Context, obj *v1.Secret, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list ServiceAccounts.
type ServiceAccountReader interface {
	// Get retrieves a ServiceAccount for the given object key
	GetServiceAccount(ctx context.Context, key client.ObjectKey) (*v1.ServiceAccount, error)

	// List retrieves list of ServiceAccounts for a given namespace and list options.
	ListServiceAccount(ctx context.Context, opts ...client.ListOption) (*v1.ServiceAccountList, error)
}

// ServiceAccountTransitionFunction instructs the ServiceAccountWriter how to transition between an existing
// ServiceAccount object and a desired on an Upsert
type ServiceAccountTransitionFunction func(existing, desired *v1.ServiceAccount) error

// Writer knows how to create, delete, and update ServiceAccounts.
type ServiceAccountWriter interface {
	// Create saves the ServiceAccount object.
	CreateServiceAccount(ctx context.Context, obj *v1.ServiceAccount, opts ...client.CreateOption) error

	// Delete deletes the ServiceAccount object.
	DeleteServiceAccount(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given ServiceAccount object.
	UpdateServiceAccount(ctx context.Context, obj *v1.ServiceAccount, opts ...client.UpdateOption) error

	// Patch patches the given ServiceAccount object.
	PatchServiceAccount(ctx context.Context, obj *v1.ServiceAccount, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ServiceAccount objects matching the given options.
	DeleteAllOfServiceAccount(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the ServiceAccount object.
	UpsertServiceAccount(ctx context.Context, obj *v1.ServiceAccount, transitionFuncs ...ServiceAccountTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a ServiceAccount object.
type ServiceAccountStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ServiceAccount object.
	UpdateServiceAccountStatus(ctx context.Context, obj *v1.ServiceAccount, opts ...client.UpdateOption) error

	// Patch patches the given ServiceAccount object's subresource.
	PatchServiceAccountStatus(ctx context.Context, obj *v1.ServiceAccount, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on ServiceAccounts.
type ServiceAccountClient interface {
	ServiceAccountReader
	ServiceAccountWriter
	ServiceAccountStatusWriter
}

type serviceAccountClient struct {
	client client.Client
}

func NewServiceAccountClient(client client.Client) *serviceAccountClient {
	return &serviceAccountClient{client: client}
}

func (c *serviceAccountClient) GetServiceAccount(ctx context.Context, key client.ObjectKey) (*v1.ServiceAccount, error) {
	obj := &ServiceAccount{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *serviceAccountClient) ListServiceAccount(ctx context.Context, opts ...client.ListOption) (*v1.ServiceAccountList, error) {
	list := &ServiceAccountList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *serviceAccountClient) CreateServiceAccount(ctx context.Context, obj *v1.ServiceAccount, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *serviceAccountClient) DeleteServiceAccount(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &ServiceAccount{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *serviceAccountClient) UpdateServiceAccount(ctx context.Context, obj *v1.ServiceAccount, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *serviceAccountClient) PatchServiceAccount(ctx context.Context, obj *v1.ServiceAccount, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *serviceAccountClient) DeleteAllOfServiceAccount(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &ServiceAccount{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *serviceAccountClient) UpsertServiceAccount(ctx context.Context, obj *v1.ServiceAccount, transitionFuncs ...ServiceAccountTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*v1.ServiceAccount), desired.(*v1.ServiceAccount)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *serviceAccountClient) UpdateServiceAccountStatus(ctx context.Context, obj *v1.ServiceAccount, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *serviceAccountClient) PatchServiceAccountStatus(ctx context.Context, obj *v1.ServiceAccount, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list ConfigMaps.
type ConfigMapReader interface {
	// Get retrieves a ConfigMap for the given object key
	GetConfigMap(ctx context.Context, key client.ObjectKey) (*v1.ConfigMap, error)

	// List retrieves list of ConfigMaps for a given namespace and list options.
	ListConfigMap(ctx context.Context, opts ...client.ListOption) (*v1.ConfigMapList, error)
}

// ConfigMapTransitionFunction instructs the ConfigMapWriter how to transition between an existing
// ConfigMap object and a desired on an Upsert
type ConfigMapTransitionFunction func(existing, desired *v1.ConfigMap) error

// Writer knows how to create, delete, and update ConfigMaps.
type ConfigMapWriter interface {
	// Create saves the ConfigMap object.
	CreateConfigMap(ctx context.Context, obj *v1.ConfigMap, opts ...client.CreateOption) error

	// Delete deletes the ConfigMap object.
	DeleteConfigMap(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given ConfigMap object.
	UpdateConfigMap(ctx context.Context, obj *v1.ConfigMap, opts ...client.UpdateOption) error

	// Patch patches the given ConfigMap object.
	PatchConfigMap(ctx context.Context, obj *v1.ConfigMap, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ConfigMap objects matching the given options.
	DeleteAllOfConfigMap(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the ConfigMap object.
	UpsertConfigMap(ctx context.Context, obj *v1.ConfigMap, transitionFuncs ...ConfigMapTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a ConfigMap object.
type ConfigMapStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ConfigMap object.
	UpdateConfigMapStatus(ctx context.Context, obj *v1.ConfigMap, opts ...client.UpdateOption) error

	// Patch patches the given ConfigMap object's subresource.
	PatchConfigMapStatus(ctx context.Context, obj *v1.ConfigMap, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on ConfigMaps.
type ConfigMapClient interface {
	ConfigMapReader
	ConfigMapWriter
	ConfigMapStatusWriter
}

type configMapClient struct {
	client client.Client
}

func NewConfigMapClient(client client.Client) *configMapClient {
	return &configMapClient{client: client}
}

func (c *configMapClient) GetConfigMap(ctx context.Context, key client.ObjectKey) (*v1.ConfigMap, error) {
	obj := &ConfigMap{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *configMapClient) ListConfigMap(ctx context.Context, opts ...client.ListOption) (*v1.ConfigMapList, error) {
	list := &ConfigMapList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *configMapClient) CreateConfigMap(ctx context.Context, obj *v1.ConfigMap, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *configMapClient) DeleteConfigMap(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &ConfigMap{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *configMapClient) UpdateConfigMap(ctx context.Context, obj *v1.ConfigMap, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *configMapClient) PatchConfigMap(ctx context.Context, obj *v1.ConfigMap, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *configMapClient) DeleteAllOfConfigMap(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &ConfigMap{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *configMapClient) UpsertConfigMap(ctx context.Context, obj *v1.ConfigMap, transitionFuncs ...ConfigMapTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*v1.ConfigMap), desired.(*v1.ConfigMap)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *configMapClient) UpdateConfigMapStatus(ctx context.Context, obj *v1.ConfigMap, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *configMapClient) PatchConfigMapStatus(ctx context.Context, obj *v1.ConfigMap, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list Services.
type ServiceReader interface {
	// Get retrieves a Service for the given object key
	GetService(ctx context.Context, key client.ObjectKey) (*v1.Service, error)

	// List retrieves list of Services for a given namespace and list options.
	ListService(ctx context.Context, opts ...client.ListOption) (*v1.ServiceList, error)
}

// ServiceTransitionFunction instructs the ServiceWriter how to transition between an existing
// Service object and a desired on an Upsert
type ServiceTransitionFunction func(existing, desired *v1.Service) error

// Writer knows how to create, delete, and update Services.
type ServiceWriter interface {
	// Create saves the Service object.
	CreateService(ctx context.Context, obj *v1.Service, opts ...client.CreateOption) error

	// Delete deletes the Service object.
	DeleteService(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Service object.
	UpdateService(ctx context.Context, obj *v1.Service, opts ...client.UpdateOption) error

	// Patch patches the given Service object.
	PatchService(ctx context.Context, obj *v1.Service, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Service objects matching the given options.
	DeleteAllOfService(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Service object.
	UpsertService(ctx context.Context, obj *v1.Service, transitionFuncs ...ServiceTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Service object.
type ServiceStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Service object.
	UpdateServiceStatus(ctx context.Context, obj *v1.Service, opts ...client.UpdateOption) error

	// Patch patches the given Service object's subresource.
	PatchServiceStatus(ctx context.Context, obj *v1.Service, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Services.
type ServiceClient interface {
	ServiceReader
	ServiceWriter
	ServiceStatusWriter
}

type serviceClient struct {
	client client.Client
}

func NewServiceClient(client client.Client) *serviceClient {
	return &serviceClient{client: client}
}

func (c *serviceClient) GetService(ctx context.Context, key client.ObjectKey) (*v1.Service, error) {
	obj := &Service{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *serviceClient) ListService(ctx context.Context, opts ...client.ListOption) (*v1.ServiceList, error) {
	list := &ServiceList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *serviceClient) CreateService(ctx context.Context, obj *v1.Service, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *serviceClient) DeleteService(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Service{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *serviceClient) UpdateService(ctx context.Context, obj *v1.Service, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *serviceClient) PatchService(ctx context.Context, obj *v1.Service, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *serviceClient) DeleteAllOfService(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Service{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *serviceClient) UpsertService(ctx context.Context, obj *v1.Service, transitionFuncs ...ServiceTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*v1.Service), desired.(*v1.Service)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *serviceClient) UpdateServiceStatus(ctx context.Context, obj *v1.Service, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *serviceClient) PatchServiceStatus(ctx context.Context, obj *v1.Service, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list Pods.
type PodReader interface {
	// Get retrieves a Pod for the given object key
	GetPod(ctx context.Context, key client.ObjectKey) (*v1.Pod, error)

	// List retrieves list of Pods for a given namespace and list options.
	ListPod(ctx context.Context, opts ...client.ListOption) (*v1.PodList, error)
}

// PodTransitionFunction instructs the PodWriter how to transition between an existing
// Pod object and a desired on an Upsert
type PodTransitionFunction func(existing, desired *v1.Pod) error

// Writer knows how to create, delete, and update Pods.
type PodWriter interface {
	// Create saves the Pod object.
	CreatePod(ctx context.Context, obj *v1.Pod, opts ...client.CreateOption) error

	// Delete deletes the Pod object.
	DeletePod(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Pod object.
	UpdatePod(ctx context.Context, obj *v1.Pod, opts ...client.UpdateOption) error

	// Patch patches the given Pod object.
	PatchPod(ctx context.Context, obj *v1.Pod, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Pod objects matching the given options.
	DeleteAllOfPod(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Pod object.
	UpsertPod(ctx context.Context, obj *v1.Pod, transitionFuncs ...PodTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Pod object.
type PodStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Pod object.
	UpdatePodStatus(ctx context.Context, obj *v1.Pod, opts ...client.UpdateOption) error

	// Patch patches the given Pod object's subresource.
	PatchPodStatus(ctx context.Context, obj *v1.Pod, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Pods.
type PodClient interface {
	PodReader
	PodWriter
	PodStatusWriter
}

type podClient struct {
	client client.Client
}

func NewPodClient(client client.Client) *podClient {
	return &podClient{client: client}
}

func (c *podClient) GetPod(ctx context.Context, key client.ObjectKey) (*v1.Pod, error) {
	obj := &Pod{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *podClient) ListPod(ctx context.Context, opts ...client.ListOption) (*v1.PodList, error) {
	list := &PodList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *podClient) CreatePod(ctx context.Context, obj *v1.Pod, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *podClient) DeletePod(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Pod{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *podClient) UpdatePod(ctx context.Context, obj *v1.Pod, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *podClient) PatchPod(ctx context.Context, obj *v1.Pod, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *podClient) DeleteAllOfPod(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Pod{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *podClient) UpsertPod(ctx context.Context, obj *v1.Pod, transitionFuncs ...PodTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*v1.Pod), desired.(*v1.Pod)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *podClient) UpdatePodStatus(ctx context.Context, obj *v1.Pod, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *podClient) PatchPodStatus(ctx context.Context, obj *v1.Pod, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list Namespaces.
type NamespaceReader interface {
	// Get retrieves a Namespace for the given object key
	GetNamespace(ctx context.Context, name string) (*v1.Namespace, error)

	// List retrieves list of Namespaces for a given namespace and list options.
	ListNamespace(ctx context.Context, opts ...client.ListOption) (*v1.NamespaceList, error)
}

// NamespaceTransitionFunction instructs the NamespaceWriter how to transition between an existing
// Namespace object and a desired on an Upsert
type NamespaceTransitionFunction func(existing, desired *v1.Namespace) error

// Writer knows how to create, delete, and update Namespaces.
type NamespaceWriter interface {
	// Create saves the Namespace object.
	CreateNamespace(ctx context.Context, obj *v1.Namespace, opts ...client.CreateOption) error

	// Delete deletes the Namespace object.
	DeleteNamespace(ctx context.Context, name string, opts ...client.DeleteOption) error

	// Update updates the given Namespace object.
	UpdateNamespace(ctx context.Context, obj *v1.Namespace, opts ...client.UpdateOption) error

	// Patch patches the given Namespace object.
	PatchNamespace(ctx context.Context, obj *v1.Namespace, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Namespace objects matching the given options.
	DeleteAllOfNamespace(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Namespace object.
	UpsertNamespace(ctx context.Context, obj *v1.Namespace, transitionFuncs ...NamespaceTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Namespace object.
type NamespaceStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Namespace object.
	UpdateNamespaceStatus(ctx context.Context, obj *v1.Namespace, opts ...client.UpdateOption) error

	// Patch patches the given Namespace object's subresource.
	PatchNamespaceStatus(ctx context.Context, obj *v1.Namespace, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Namespaces.
type NamespaceClient interface {
	NamespaceReader
	NamespaceWriter
	NamespaceStatusWriter
}

type namespaceClient struct {
	client client.Client
}

func NewNamespaceClient(client client.Client) *namespaceClient {
	return &namespaceClient{client: client}
}

func (c *namespaceClient) GetNamespace(ctx context.Context, name string) (*v1.Namespace, error) {
	obj := &Namespace{}
	key := client.ObjectKey{
		Name: name,
	}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *namespaceClient) ListNamespace(ctx context.Context, opts ...client.ListOption) (*v1.NamespaceList, error) {
	list := &NamespaceList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *namespaceClient) CreateNamespace(ctx context.Context, obj *v1.Namespace, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *namespaceClient) DeleteNamespace(ctx context.Context, name string, opts ...client.DeleteOption) error {
	obj := &Namespace{}
	obj.SetName(name)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *namespaceClient) UpdateNamespace(ctx context.Context, obj *v1.Namespace, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *namespaceClient) PatchNamespace(ctx context.Context, obj *v1.Namespace, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *namespaceClient) DeleteAllOfNamespace(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Namespace{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *namespaceClient) UpsertNamespace(ctx context.Context, obj *v1.Namespace, transitionFuncs ...NamespaceTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*v1.Namespace), desired.(*v1.Namespace)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *namespaceClient) UpdateNamespaceStatus(ctx context.Context, obj *v1.Namespace, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *namespaceClient) PatchNamespaceStatus(ctx context.Context, obj *v1.Namespace, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list Nodes.
type NodeReader interface {
	// Get retrieves a Node for the given object key
	GetNode(ctx context.Context, name string) (*v1.Node, error)

	// List retrieves list of Nodes for a given namespace and list options.
	ListNode(ctx context.Context, opts ...client.ListOption) (*v1.NodeList, error)
}

// NodeTransitionFunction instructs the NodeWriter how to transition between an existing
// Node object and a desired on an Upsert
type NodeTransitionFunction func(existing, desired *v1.Node) error

// Writer knows how to create, delete, and update Nodes.
type NodeWriter interface {
	// Create saves the Node object.
	CreateNode(ctx context.Context, obj *v1.Node, opts ...client.CreateOption) error

	// Delete deletes the Node object.
	DeleteNode(ctx context.Context, name string, opts ...client.DeleteOption) error

	// Update updates the given Node object.
	UpdateNode(ctx context.Context, obj *v1.Node, opts ...client.UpdateOption) error

	// Patch patches the given Node object.
	PatchNode(ctx context.Context, obj *v1.Node, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Node objects matching the given options.
	DeleteAllOfNode(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the Node object.
	UpsertNode(ctx context.Context, obj *v1.Node, transitionFuncs ...NodeTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a Node object.
type NodeStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Node object.
	UpdateNodeStatus(ctx context.Context, obj *v1.Node, opts ...client.UpdateOption) error

	// Patch patches the given Node object's subresource.
	PatchNodeStatus(ctx context.Context, obj *v1.Node, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Nodes.
type NodeClient interface {
	NodeReader
	NodeWriter
	NodeStatusWriter
}

type nodeClient struct {
	client client.Client
}

func NewNodeClient(client client.Client) *nodeClient {
	return &nodeClient{client: client}
}

func (c *nodeClient) GetNode(ctx context.Context, name string) (*v1.Node, error) {
	obj := &Node{}
	key := client.ObjectKey{
		Name: name,
	}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *nodeClient) ListNode(ctx context.Context, opts ...client.ListOption) (*v1.NodeList, error) {
	list := &NodeList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *nodeClient) CreateNode(ctx context.Context, obj *v1.Node, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *nodeClient) DeleteNode(ctx context.Context, name string, opts ...client.DeleteOption) error {
	obj := &Node{}
	obj.SetName(name)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *nodeClient) UpdateNode(ctx context.Context, obj *v1.Node, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *nodeClient) PatchNode(ctx context.Context, obj *v1.Node, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *nodeClient) DeleteAllOfNode(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Node{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *nodeClient) UpsertNode(ctx context.Context, obj *v1.Node, transitionFuncs ...NodeTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*v1.Node), desired.(*v1.Node)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *nodeClient) UpdateNodeStatus(ctx context.Context, obj *v1.Node, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *nodeClient) PatchNodeStatus(ctx context.Context, obj *v1.Node, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}