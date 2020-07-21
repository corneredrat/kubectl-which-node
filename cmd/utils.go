package cmd

import (
	"strings"
	"k8s.io/klog"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Get names of all available api resources in the kubernetes server
func getNameList(apiResourceLists []*v1.APIResourceList) []string{
	var names []string
	for _, apiResourceList := range(apiResourceLists) {
		klog.V(4).Infof("parsing over resource list: group - %v, version - %v",apiResourceList.GroupVersion,apiResourceList.APIVersion)
		for _, apiResource := range(apiResourceList.APIResources) {
			pluralName		:= strings.ToLower(apiResource.Name)
			singularName	:= ""
			if apiResource.SingularName == "" {
				singularName = strings.ToLower(apiResource.Kind)
			} else {
				singularName = strings.ToLower(apiResource.SingularName)
			}
			names = append(names,singularName)
			names = append(names,pluralName)
			names = append(names,apiResource.ShortNames...)
		}
	}
	return names
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
