package cmd

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type apiResource struct {
	resource  		v1.APIResource
	group			string
	apiVersion 		string
	name			string
}

// https://godoc.org/k8s.io/apimachinery/pkg/runtime/schema#GroupVersion
func (r *apiResource) groupVersion() schema.GroupVersion {
	var groupVersion schema.GroupVersion
	groupVersion.Group 		= r.group
	groupVersion.Version 	= r.apiVersion
	return  groupVersion
}

func makeAPIResource(resourceList *v1.APIResourceList, resource v1.APIResource) apiResource {
	var apiResourceElement apiResource
	apiResourceElement.resource = resource
	apiResourceElement.group	= getGroupVersionFromMetadata(resourceList)
	apiResourceElement.apiVersion	= getAPIVersionFromMetadata(resourceList)
	apiResourceElement.name		= resource.Name
	return apiResourceElement
} 

func  (r *apiResource) getGroupVersion() string {
	return  r.group
}

func  (r *apiResource) getAPIVersion() string {
	return  r.apiVersion
}

func  (r *apiResource) getName() string {
	return  r.name
}

