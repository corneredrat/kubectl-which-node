package cmd

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func getNameList(apiResourceLists []*v1.apiResourceList) string{
	var names []string
	for apiResourceList := range(apiResourceLists) {
		for apiResource := range(apiResourceList) {
			pluralName		:= apiResource.Name
			singularName	:= ""
			if apiResource.SingularName == nil {
				singularName = apiResource.Kind
			} else {
				singularName = apiResource.SingularName
			}
			append(names,string{singularName, pluralName},apiResource.ShortNames...)
		}
	}
	return string
}