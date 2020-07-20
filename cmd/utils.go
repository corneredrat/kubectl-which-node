package cmd

import (
	"k8s.io/klog"
	"strings"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Get names of all available api resources in the kubernetes server
func getNameList(apiResourceLists []*v1.APIResourceList) []string{
	var names []string
	for _, apiResourceList := range(apiResourceLists) {
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
func exists(key interface{}, list ...interface{}) bool {
	found := false
	for _, k := range(list[0]) {
		
		if key == k {
			found = true
		}
	}
	return found
}