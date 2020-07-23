package cmd

import (
	"fmt" //apiGroup, APIResourcelist
	"strings"
	_ "context"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func findApiResource(name string) ([]apiResource ,error) {
	
	var resources []apiResource//types.go
	
	//https://godoc.org/k8s.io/client-go/discovery#CachedDiscoveryInterface
	klog.V(3).Info("fetching Resources lists from Kubernetes server")
	apiResourceLists, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		return resources, fmt.Errorf("unable to fetch api resource list: %w")
	}
	
	// Get matching API Resource from the given name
	resources = getResourceFromList(name, apiResourceLists)
	klog.V(3).Info(resources)
	
	/*
	Donot Disambiguate. Have all api version, if one doesnt work... try another!
	For example, some replicasets may be extensions/v1beta1, some may be at core/v1

	if len(resources) > 1 {
		resources = disAmbiguate(resources) 
		if len(resources) > 1 {
			var groups []schema.GroupVersion
			for _, resource := range(resources) {
				group 	:= resource.groupVersion()
				groups 	= append(groups, group)
			}
			return resources, fmt.Errorf("multiple matches found for %v, matching groups: %v . Please diambiguate the kind name.", name, groups) 
		}
		
	}
	*/
	if len(resources) == 0 {
		return resources, fmt.Errorf("no matches found for kind %v", name)
	}
	return resources, nil
}

func findObjectResource( resources []apiResource, objectName string) (*unstructured.Unstructured, error) {
		// get resource

		//https://godoc.org/k8s.io/client-go/dynamic#Interface
		//https://godoc.org/k8s.io/client-go/dynamic#NamespaceableResourceInterface
		namespace 	:= getNamespace()
		objectFound := false
		var err error
		var object *unstructured.Unstructured
		for _, resource := range(resources) {
			groupVersionResource 	:= resource.groupVersionResource()
			resourceInterface 		:= dynamicInterface.Resource(resource.groupVersionResource()).Namespace(namespace)
			
			klog.V(4).Infof("trying - group-version-resource : %v", )
			
			object, err = resourceInterface.Get("", v1.GetOptions{})
			if err != nil {
				klog.V(3).Infof("could not find %v in groupVersionResource: %v",objectName, groupVersionResource)
				continue
			}
			objectFound = true
		}
		//https://godoc.org/k8s.io/client-go/dynamic#ResourceInterface
		if !objectFound {
			return object, fmt.Errorf("unable to find %v in any of api/group version")
		}
		klog.V(3).Infof("successfully obtained object: %v", object)
		return object, nil
}

func disAmbiguate(resources []apiResource) []apiResource {
	
	var unAmbigousResources []apiResource
	name := resources[0].getName()
	name = strings.ToLower(name) 
	switch(name) {
	case "replicasets": for _,resource := range(resources) {
		klog.V(4).Infof("disambiguation: comparing: %v, apps",resource.getAPIVersion(),  )
		if resource.getAPIVersion() ==  "apps" {
			return append(unAmbigousResources,resource)
		}
	}
	case "deployments": for _,resource := range(resources) {
		if resource.getAPIVersion() ==  "apps" {
			return append(unAmbigousResources,resource)
		}
	}
	case "daemonsets": for _,resource := range(resources) {
		if resource.getAPIVersion() ==  "apps" {
			return append(unAmbigousResources,resource)
		}
	}
	case "statefulsets": for _,resource := range(resources) {
		if resource.getAPIVersion() ==  "apps" {
			return append(unAmbigousResources,resource)
		}
	}
	case "jobs": for _,resource := range(resources) {
		if resource.getAPIVersion() ==  "batch" {
			return append(unAmbigousResources,resource)
		}
	}
	}
	return resources
}