package cmd

import (
	"fmt" //apiGroup, APIResourcelist

	"k8s.io/klog"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func findApiResourceNames()  ([]string ,error) {
	//https://godoc.org/k8s.io/client-go/discovery#CachedDiscoveryInterface
	klog.V(3).Info("fetching Resources lists from Kubernetes server")
	apiResourceLists, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		return []string{}, fmt.Errorf("unable to fetch api resource list: %w")
	}
	return getNameList(apiResourceLists), nil
}

