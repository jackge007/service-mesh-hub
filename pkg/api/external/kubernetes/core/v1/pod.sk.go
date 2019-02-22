// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO: modify as needed to populate additional fields
func NewPod(namespace, name string) *Pod {
	return &Pod{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (r *Pod) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *Pod) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.Spec,
	)
}

type PodList []*Pod
type PodsByNamespace map[string]PodList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list PodList) Find(namespace, name string) (*Pod, error) {
	for _, pod := range list {
		if pod.Metadata.Name == name {
			if namespace == "" || pod.Metadata.Namespace == namespace {
				return pod, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find pod %v.%v", namespace, name)
}

func (list PodList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, pod := range list {
		ress = append(ress, pod)
	}
	return ress
}

func (list PodList) Names() []string {
	var names []string
	for _, pod := range list {
		names = append(names, pod.Metadata.Name)
	}
	return names
}

func (list PodList) NamespacesDotNames() []string {
	var names []string
	for _, pod := range list {
		names = append(names, pod.Metadata.Namespace+"."+pod.Metadata.Name)
	}
	return names
}

func (list PodList) Sort() PodList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Metadata.Less(list[j].Metadata)
	})
	return list
}

func (list PodList) Clone() PodList {
	var podList PodList
	for _, pod := range list {
		podList = append(podList, proto.Clone(pod).(*Pod))
	}
	return podList
}

func (list PodList) Each(f func(element *Pod)) {
	for _, pod := range list {
		f(pod)
	}
}

func (list PodList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Pod) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

func (byNamespace PodsByNamespace) Add(pod ...*Pod) {
	for _, item := range pod {
		byNamespace[item.Metadata.Namespace] = append(byNamespace[item.Metadata.Namespace], item)
	}
}

func (byNamespace PodsByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace PodsByNamespace) List() PodList {
	var list PodList
	for _, podList := range byNamespace {
		list = append(list, podList...)
	}
	return list.Sort()
}

func (byNamespace PodsByNamespace) Clone() PodsByNamespace {
	cloned := make(PodsByNamespace)
	for ns, list := range byNamespace {
		cloned[ns] = list.Clone()
	}
	return cloned
}

var _ resources.Resource = &Pod{}

// Kubernetes Adapter for Pod

func (o *Pod) GetObjectKind() schema.ObjectKind {
	t := PodCrd.TypeMeta()
	return &t
}

func (o *Pod) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Pod)
}

var PodCrd = crd.NewCrd("core.kubernetes.io",
	"pods",
	"core.kubernetes.io",
	"v1",
	"Pod",
	"pod",
	false,
	&Pod{})
