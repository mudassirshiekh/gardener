// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ControllerInstallationLister helps list ControllerInstallations.
// All objects returned here must be treated as read-only.
type ControllerInstallationLister interface {
	// List lists all ControllerInstallations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ControllerInstallation, err error)
	// Get retrieves the ControllerInstallation from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.ControllerInstallation, error)
	ControllerInstallationListerExpansion
}

// controllerInstallationLister implements the ControllerInstallationLister interface.
type controllerInstallationLister struct {
	listers.ResourceIndexer[*v1beta1.ControllerInstallation]
}

// NewControllerInstallationLister returns a new ControllerInstallationLister.
func NewControllerInstallationLister(indexer cache.Indexer) ControllerInstallationLister {
	return &controllerInstallationLister{listers.New[*v1beta1.ControllerInstallation](indexer, v1beta1.Resource("controllerinstallation"))}
}
