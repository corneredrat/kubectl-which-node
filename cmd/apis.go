package cmd

import (
	"fmt" //apiGroup, APIResourcelist

	"k8s.io/klog"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func findApiResources()  ([]*v1.APIResourceList ,error) {
	// https://godoc.org/k8s.io/client-go/discovery#DiscoveryInterface
	apiResourceList, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		return apiResourceList, fmt.Errorf("unable to fetch api resource list: %w")
	}

	for _, apiResource := range(Name) {
		klog.Info("%v", apiResourceList.String())
	}
	
	return apiResourceList, nil
}

