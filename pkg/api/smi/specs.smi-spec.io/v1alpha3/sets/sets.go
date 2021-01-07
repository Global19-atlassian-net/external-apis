// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha3sets

import (
	specs_smi_spec_io_v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type HTTPRouteGroupSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) bool) []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup
	// Return the Set as a map of key to resource.
	Map() map[string]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup
	// Insert a resource into the set.
	Insert(hTTPRouteGroup ...*specs_smi_spec_io_v1alpha3.HTTPRouteGroup)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(hTTPRouteGroupSet HTTPRouteGroupSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(hTTPRouteGroup ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(hTTPRouteGroup ezkube.ResourceId)
	// Return the union with the provided set
	Union(set HTTPRouteGroupSet) HTTPRouteGroupSet
	// Return the difference with the provided set
	Difference(set HTTPRouteGroupSet) HTTPRouteGroupSet
	// Return the intersection with the provided set
	Intersection(set HTTPRouteGroupSet) HTTPRouteGroupSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*specs_smi_spec_io_v1alpha3.HTTPRouteGroup, error)
	// Get the length of the set
	Length() int
}

func makeGenericHTTPRouteGroupSet(hTTPRouteGroupList []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range hTTPRouteGroupList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type hTTPRouteGroupSet struct {
	set sksets.ResourceSet
}

func NewHTTPRouteGroupSet(hTTPRouteGroupList ...*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) HTTPRouteGroupSet {
	return &hTTPRouteGroupSet{set: makeGenericHTTPRouteGroupSet(hTTPRouteGroupList)}
}

func NewHTTPRouteGroupSetFromList(hTTPRouteGroupList *specs_smi_spec_io_v1alpha3.HTTPRouteGroupList) HTTPRouteGroupSet {
	list := make([]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup, 0, len(hTTPRouteGroupList.Items))
	for idx := range hTTPRouteGroupList.Items {
		list = append(list, &hTTPRouteGroupList.Items[idx])
	}
	return &hTTPRouteGroupSet{set: makeGenericHTTPRouteGroupSet(list)}
}

func (s *hTTPRouteGroupSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.set.Keys()
}

func (s *hTTPRouteGroupSet) List(filterResource ...func(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup) bool) []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup))
		})
	}

	var hTTPRouteGroupList []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup
	for _, obj := range s.set.List(genericFilters...) {
		hTTPRouteGroupList = append(hTTPRouteGroupList, obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup))
	}
	return hTTPRouteGroupList
}

func (s *hTTPRouteGroupSet) Map() map[string]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup {
	if s == nil {
		return nil
	}

	newMap := map[string]*specs_smi_spec_io_v1alpha3.HTTPRouteGroup{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup)
	}
	return newMap
}

func (s *hTTPRouteGroupSet) Insert(
	hTTPRouteGroupList ...*specs_smi_spec_io_v1alpha3.HTTPRouteGroup,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range hTTPRouteGroupList {
		s.set.Insert(obj)
	}
}

func (s *hTTPRouteGroupSet) Has(hTTPRouteGroup ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.set.Has(hTTPRouteGroup)
}

func (s *hTTPRouteGroupSet) Equal(
	hTTPRouteGroupSet HTTPRouteGroupSet,
) bool {
	if s == nil {
		return hTTPRouteGroupSet == nil
	}
	return s.set.Equal(makeGenericHTTPRouteGroupSet(hTTPRouteGroupSet.List()))
}

func (s *hTTPRouteGroupSet) Delete(HTTPRouteGroup ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.set.Delete(HTTPRouteGroup)
}

func (s *hTTPRouteGroupSet) Union(set HTTPRouteGroupSet) HTTPRouteGroupSet {
	if s == nil {
		return set
	}
	return NewHTTPRouteGroupSet(append(s.List(), set.List()...)...)
}

func (s *hTTPRouteGroupSet) Difference(set HTTPRouteGroupSet) HTTPRouteGroupSet {
	if s == nil {
		return set
	}
	newSet := s.set.Difference(makeGenericHTTPRouteGroupSet(set.List()))
	return &hTTPRouteGroupSet{set: newSet}
}

func (s *hTTPRouteGroupSet) Intersection(set HTTPRouteGroupSet) HTTPRouteGroupSet {
	if s == nil {
		return nil
	}
	newSet := s.set.Intersection(makeGenericHTTPRouteGroupSet(set.List()))
	var hTTPRouteGroupList []*specs_smi_spec_io_v1alpha3.HTTPRouteGroup
	for _, obj := range newSet.List() {
		hTTPRouteGroupList = append(hTTPRouteGroupList, obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup))
	}
	return NewHTTPRouteGroupSet(hTTPRouteGroupList...)
}

func (s *hTTPRouteGroupSet) Find(id ezkube.ResourceId) (*specs_smi_spec_io_v1alpha3.HTTPRouteGroup, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find HTTPRouteGroup %v", sksets.Key(id))
	}
	obj, err := s.set.Find(&specs_smi_spec_io_v1alpha3.HTTPRouteGroup{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*specs_smi_spec_io_v1alpha3.HTTPRouteGroup), nil
}

func (s *hTTPRouteGroupSet) Length() int {
	if s == nil {
		return 0
	}
	return s.set.Length()
}