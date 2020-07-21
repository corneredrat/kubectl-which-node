package cmd

import (
	"fmt" //apiGroup, APIResourcelist

	"k8s.io/klog"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	
	if len(resources) > 1 {
		var groups []string
		for _, resource := range(resources) {
			groups := append(groups, getGroupVersion(resource))
		}
		return resources, fmt.Errorf("multiple matches found for %v, matching groups: %v . Please diambiguate the kind name.", name, groups) 
	}
	if len(resources) == 0 {
		return resources, fmt.Errorf("no matches found for kind %v", name)
	}
	return resources, nil
}

