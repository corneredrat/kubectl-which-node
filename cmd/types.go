package cmd

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type apiResource struct {
	resource  		v1.APIResource
	group			string
	apiVersion 		string
	name			string
}

// https://godoc.org/k8s.io/apimachinery/pkg/runtime/schema#GroupVersion
func  (r *apiResource) getGroupVersion() string {
	return  struct{
		Group   string
		Version string
	} {
		Group: 		r.group
		Version:	r.apiVersion
	}
}

func makeAPIResource(resourceList v1.APIResourceList,resource v1.APIResource) apiResource {
	var apiResourceElement apiResource
	apiResourceElement.resource := resource
	apiResourceElement.group	:= resourceList.getGroupVersion()
	apiResourceElement.version	:= resourceList.getAPIVersion()
	apiResourceElement.name		:= resource.Name
} 

func  (r *apiResource) getGroupVersion() string {
	return  r.groupVersion
}

func  (r *apiResource) getAPIVersion() string {
	return  r.apiVersion
}

func  (r *apiResource) getName() string {
	return  r.name
}

