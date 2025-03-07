// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// HostEndpointLister helps list HostEndpoints.
// All objects returned here must be treated as read-only.
type HostEndpointLister interface {
	// List lists all HostEndpoints in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.HostEndpoint, err error)
	// Get retrieves the HostEndpoint from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.HostEndpoint, error)
	HostEndpointListerExpansion
}

// hostEndpointLister implements the HostEndpointLister interface.
type hostEndpointLister struct {
	listers.ResourceIndexer[*v3.HostEndpoint]
}

// NewHostEndpointLister returns a new HostEndpointLister.
func NewHostEndpointLister(indexer cache.Indexer) HostEndpointLister {
	return &hostEndpointLister{listers.New[*v3.HostEndpoint](indexer, v3.Resource("hostendpoint"))}
}
