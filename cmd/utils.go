package cmd

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func getNameList(apiResourceLists []*v1.APIResourceList) []string{
	var names []string
	for _, apiResourceList := range(apiResourceLists) {
		for _, apiResource := range(apiResourceList.APIResources) {
			pluralName		:= apiResource.Name
			singularName	:= ""
			if apiResource.SingularName == "" {
				singularName = apiResource.Kind
			} else {
				singularName = apiResource.SingularName
			}
			names = append(names,singularName)
			names = append(names,pluralName)
			names = append(names,apiResource.ShortNames...)
		}
	}
	return names
}

func exists(key, list) bool {
	found := false
	for _, k := range(list) {
		if key == k {
			found = true
		}
	}
	return found
}