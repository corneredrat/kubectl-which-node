package cmd

import (
	"fmt" //apiGroup, APIResourcelist
	"strings"
	"k8s.io/klog"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
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
	if len(resources) == 0 {
		return resources, fmt.Errorf("no matches found for kind %v", name)
	}
	return resources, nil
}

func disAmbiguate(resources []apiResource) []apiResource {
	
	var unAmbigousResources []apiResource
	name := resources[0].getName()
	name = strings.ToLower(name) 
	switch(name) {
	case "replicasets": for resource := range(resources) {
		if resource.getGroupVersion() ==  "apps" {
			return append(unAmbigousResources,resource)
		}
	}
	case "deployments": for resource := range(resources) {
		if resource.getGroupVersion() ==  "apps" {
			return append(unAmbigousResources,resource)
		}
	}
	case "daemonsets": for resource := range(resources) {
		if resource.getGroupVersion() ==  "apps" {
			return append(unAmbigousResources,resource)
		}
	}
	case "statefulsets": for resource := range(resources) {
		if resource.getGroupVersion() ==  "apps" {
			return append(unAmbigousResources,resource)
		}
	}
	case "jobs": for resource := range(resources) {
		if resource.getGroupVersion() ==  "batch" {
			return append(unAmbigousResources,resource)
		}
	}
	}
}