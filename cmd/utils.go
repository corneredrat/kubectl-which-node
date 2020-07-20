package cmd

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func getNameList(apiResourceLists []*v1.APIResourceList) string{
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
			append(names,[]string{singularName, pluralName})
			append(names,apiResource.ShortNames)
		}
	}
	return string
}