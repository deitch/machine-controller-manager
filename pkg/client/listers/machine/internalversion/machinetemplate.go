// This file was automatically generated by lister-gen

package internalversion

import (
	machine "github.com/gardener/node-controller-manager/pkg/apis/machine"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MachineTemplateLister helps list MachineTemplates.
type MachineTemplateLister interface {
	// List lists all MachineTemplates in the indexer.
	List(selector labels.Selector) (ret []*machine.MachineTemplate, err error)
	// MachineTemplates returns an object that can list and get MachineTemplates.
	MachineTemplates(namespace string) MachineTemplateNamespaceLister
	MachineTemplateListerExpansion
}

// machineTemplateLister implements the MachineTemplateLister interface.
type machineTemplateLister struct {
	indexer cache.Indexer
}

// NewMachineTemplateLister returns a new MachineTemplateLister.
func NewMachineTemplateLister(indexer cache.Indexer) MachineTemplateLister {
	return &machineTemplateLister{indexer: indexer}
}

// List lists all MachineTemplates in the indexer.
func (s *machineTemplateLister) List(selector labels.Selector) (ret []*machine.MachineTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.MachineTemplate))
	})
	return ret, err
}

// MachineTemplates returns an object that can list and get MachineTemplates.
func (s *machineTemplateLister) MachineTemplates(namespace string) MachineTemplateNamespaceLister {
	return machineTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MachineTemplateNamespaceLister helps list and get MachineTemplates.
type MachineTemplateNamespaceLister interface {
	// List lists all MachineTemplates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*machine.MachineTemplate, err error)
	// Get retrieves the MachineTemplate from the indexer for a given namespace and name.
	Get(name string) (*machine.MachineTemplate, error)
	MachineTemplateNamespaceListerExpansion
}

// machineTemplateNamespaceLister implements the MachineTemplateNamespaceLister
// interface.
type machineTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MachineTemplates in the indexer for a given namespace.
func (s machineTemplateNamespaceLister) List(selector labels.Selector) (ret []*machine.MachineTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.MachineTemplate))
	})
	return ret, err
}

// Get retrieves the MachineTemplate from the indexer for a given namespace and name.
func (s machineTemplateNamespaceLister) Get(name string) (*machine.MachineTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(machine.Resource("machinetemplate"), name)
	}
	return obj.(*machine.MachineTemplate), nil
}
