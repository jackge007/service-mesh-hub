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
func NewTlsSecret(namespace, name string) *TlsSecret {
	return &TlsSecret{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (r *TlsSecret) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *TlsSecret) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.RootCert,
		r.CertChain,
		r.CaCert,
		r.CaKey,
	)
}

type TlsSecretList []*TlsSecret
type TlssecretsByNamespace map[string]TlsSecretList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list TlsSecretList) Find(namespace, name string) (*TlsSecret, error) {
	for _, tlsSecret := range list {
		if tlsSecret.Metadata.Name == name {
			if namespace == "" || tlsSecret.Metadata.Namespace == namespace {
				return tlsSecret, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find tlsSecret %v.%v", namespace, name)
}

func (list TlsSecretList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, tlsSecret := range list {
		ress = append(ress, tlsSecret)
	}
	return ress
}

func (list TlsSecretList) Names() []string {
	var names []string
	for _, tlsSecret := range list {
		names = append(names, tlsSecret.Metadata.Name)
	}
	return names
}

func (list TlsSecretList) NamespacesDotNames() []string {
	var names []string
	for _, tlsSecret := range list {
		names = append(names, tlsSecret.Metadata.Namespace+"."+tlsSecret.Metadata.Name)
	}
	return names
}

func (list TlsSecretList) Sort() TlsSecretList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Metadata.Less(list[j].Metadata)
	})
	return list
}

func (list TlsSecretList) Clone() TlsSecretList {
	var tlsSecretList TlsSecretList
	for _, tlsSecret := range list {
		tlsSecretList = append(tlsSecretList, proto.Clone(tlsSecret).(*TlsSecret))
	}
	return tlsSecretList
}

func (list TlsSecretList) Each(f func(element *TlsSecret)) {
	for _, tlsSecret := range list {
		f(tlsSecret)
	}
}

func (list TlsSecretList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *TlsSecret) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

func (byNamespace TlssecretsByNamespace) Add(tlsSecret ...*TlsSecret) {
	for _, item := range tlsSecret {
		byNamespace[item.Metadata.Namespace] = append(byNamespace[item.Metadata.Namespace], item)
	}
}

func (byNamespace TlssecretsByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace TlssecretsByNamespace) List() TlsSecretList {
	var list TlsSecretList
	for _, tlsSecretList := range byNamespace {
		list = append(list, tlsSecretList...)
	}
	return list.Sort()
}

func (byNamespace TlssecretsByNamespace) Clone() TlssecretsByNamespace {
	cloned := make(TlssecretsByNamespace)
	for ns, list := range byNamespace {
		cloned[ns] = list.Clone()
	}
	return cloned
}

var _ resources.Resource = &TlsSecret{}

// Kubernetes Adapter for TlsSecret

func (o *TlsSecret) GetObjectKind() schema.ObjectKind {
	t := TlsSecretCrd.TypeMeta()
	return &t
}

func (o *TlsSecret) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*TlsSecret)
}

var TlsSecretCrd = crd.NewCrd("supergloo.solo.io",
	"tlssecrets",
	"supergloo.solo.io",
	"v1",
	"TlsSecret",
	"ts",
	false,
	&TlsSecret{})
