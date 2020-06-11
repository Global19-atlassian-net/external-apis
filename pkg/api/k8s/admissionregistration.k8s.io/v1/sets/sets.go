// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1sets

import (
	admissionregistration_k8s_io_v1 "k8s.io/api/admissionregistration/v1"

	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

type ValidatingWebhookConfigurationSet interface {
	Keys() sets.String
	List() []*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration
	Map() map[string]*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration
	Insert(validatingWebhookConfiguration ...*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration)
	Equal(validatingWebhookConfigurationSet ValidatingWebhookConfigurationSet) bool
	Has(validatingWebhookConfiguration *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) bool
	Delete(validatingWebhookConfiguration *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration)
	Union(set ValidatingWebhookConfigurationSet) ValidatingWebhookConfigurationSet
	Difference(set ValidatingWebhookConfigurationSet) ValidatingWebhookConfigurationSet
	Intersection(set ValidatingWebhookConfigurationSet) ValidatingWebhookConfigurationSet
}

func makeGenericValidatingWebhookConfigurationSet(validatingWebhookConfigurationList []*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range validatingWebhookConfigurationList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type validatingWebhookConfigurationSet struct {
	set sksets.ResourceSet
}

func NewValidatingWebhookConfigurationSet(validatingWebhookConfigurationList ...*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) ValidatingWebhookConfigurationSet {
	return &validatingWebhookConfigurationSet{set: makeGenericValidatingWebhookConfigurationSet(validatingWebhookConfigurationList)}
}

func (s validatingWebhookConfigurationSet) Keys() sets.String {
	return s.set.Keys()
}

func (s validatingWebhookConfigurationSet) List() []*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration {
	var validatingWebhookConfigurationList []*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration
	for _, obj := range s.set.List() {
		validatingWebhookConfigurationList = append(validatingWebhookConfigurationList, obj.(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration))
	}
	return validatingWebhookConfigurationList
}

func (s validatingWebhookConfigurationSet) Map() map[string]*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration {
	newMap := map[string]*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration)
	}
	return newMap
}

func (s validatingWebhookConfigurationSet) Insert(
	validatingWebhookConfigurationList ...*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration,
) {
	for _, obj := range validatingWebhookConfigurationList {
		s.set.Insert(obj)
	}
}

func (s validatingWebhookConfigurationSet) Has(validatingWebhookConfiguration *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) bool {
	return s.set.Has(validatingWebhookConfiguration)
}

func (s validatingWebhookConfigurationSet) Equal(
	validatingWebhookConfigurationSet ValidatingWebhookConfigurationSet,
) bool {
	return s.set.Equal(makeGenericValidatingWebhookConfigurationSet(validatingWebhookConfigurationSet.List()))
}

func (s validatingWebhookConfigurationSet) Delete(ValidatingWebhookConfiguration *admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration) {
	s.set.Delete(ValidatingWebhookConfiguration)
}

func (s validatingWebhookConfigurationSet) Union(set ValidatingWebhookConfigurationSet) ValidatingWebhookConfigurationSet {
	return NewValidatingWebhookConfigurationSet(append(s.List(), set.List()...)...)
}

func (s validatingWebhookConfigurationSet) Difference(set ValidatingWebhookConfigurationSet) ValidatingWebhookConfigurationSet {
	newSet := s.set.Difference(makeGenericValidatingWebhookConfigurationSet(set.List()))
	return validatingWebhookConfigurationSet{set: newSet}
}

func (s validatingWebhookConfigurationSet) Intersection(set ValidatingWebhookConfigurationSet) ValidatingWebhookConfigurationSet {
	newSet := s.set.Intersection(makeGenericValidatingWebhookConfigurationSet(set.List()))
	var validatingWebhookConfigurationList []*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration
	for _, obj := range newSet.List() {
		validatingWebhookConfigurationList = append(validatingWebhookConfigurationList, obj.(*admissionregistration_k8s_io_v1.ValidatingWebhookConfiguration))
	}
	return NewValidatingWebhookConfigurationSet(validatingWebhookConfigurationList...)
}