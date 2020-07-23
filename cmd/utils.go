package cmd

import (
	"strings"
	"k8s.io/klog"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Get names of all available api resources in the kubernetes server
func getResourceFromList(name string, apiResourceLists []*v1.APIResourceList) []apiResource{
	var r     []apiResource
	
	for _, apiResourceList := range(apiResourceLists) {
		
		klog.V(4).Infof("parsing over resource list: group - %v, version - %v",apiResourceList.GroupVersion,apiResourceList.APIVersion)
		for _, apiResourceElement := range(apiResourceList.APIResources) {
			var names []string
			pluralName		:= strings.ToLower(apiResourceElement.Name)
			singularName	:= ""
			if apiResourceElement.SingularName == "" {
				singularName = strings.ToLower(apiResourceElement.Kind)
			} else {
				singularName = strings.ToLower(apiResourceElement.SingularName)
			}
			names = append(names,singularName)
			names = append(names,pluralName)
			names = append(names,apiResourceElement.ShortNames...)

			if stringExists(name, names) {
				klog.V(3).Infof("found match name:%v, range:%v",name,names)
				resource := makeAPIResource(apiResourceList, apiResourceElement)
				r 		 = append(r, resource )
			}
		}
	}
	return r
}

// A small function that checks if 
func stringExists(key string,list []string) bool {
	found := false
	for _, k := range(list) {
		
		if key == k {
			found = true
		}
	}
	return found
}

// Get namespace
// Shamelessly stolen from https://github.com/ahmetb/kubectl-tree/blob/master/cmd/kubectl-tree/namespace.go
func getNamespace() string {
	if v := *configFlags.Namespace; v != "" {
		return v
	}
	clientConfig := configFlags.ToRawKubeConfigLoader()
	defaultNamespace, _, err := clientConfig.Namespace()
	if err != nil {
		defaultNamespace = "default"
	}
	return defaultNamespace
}

func makeAPIResource(resourceList *v1.APIResourceList, resource v1.APIResource) apiResource {
	var apiResourceElement apiResource
	apiResourceElement.resource 	= resource
	apiResourceElement.group		= getGroupVersionFromMetadata(*resourceList)		// utils.go
	apiResourceElement.apiVersion	= getAPIVersionFromMetadata(*resourceList)		// utils.go
	apiResourceElement.name			= resource.Name
	return apiResourceElement
} 

func getGroupVersionFromMetadata(resource v1.APIResourceList) string {
	if resource.GroupVersion == "v1" {
		return "core"
	} else {
		groupAPIVersion := resource.GroupVersion
		return strings.Split(groupAPIVersion, "/")[0]
	}
	
}

func  getAPIVersionFromMetadata(resource v1.APIResourceList) string {
	if resource.GroupVersion == "v1" {
		return "v1"
	} else {
		groupAPIVersion := resource.GroupVersion
		return strings.Split(groupAPIVersion, "/")[1]
	}
	
}

