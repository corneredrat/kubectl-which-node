package cmd

import (
	"fmt" //apiGroup, APIResourcelist
	"strings"
	_ "context"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/kubernetes/typed/core/v1"
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
	For example, some replicasets may be extensions/v1beta1, some may be at apps/v1

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
		for _, resource := range(resources) {
			groupVersionResource 	:= resource.groupVersionResource()
			resourceInterface 		:= dynamicInterface.Resource(resource.groupVersionResource()).Namespace(namespace)
			
			klog.V(4).Infof("trying - group-version : %v, %v", groupVersionResource.Group, groupVersionResource.Version )
			
			object, err := resourceInterface.Get(objectName, v1.GetOptions{})
			if err != nil {
				klog.V(3).Infof("could not find %v in groupVersionResource: %v",objectName, groupVersionResource)
				continue
			} else {
				klog.V(3).Infof("found resource!")
				return object, nil
			}
		}
		
		//https://godoc.org/k8s.io/client-go/dynamic#ResourceInterface
		var dummy *unstructured.Unstructured
		if !objectFound {
			return dummy, fmt.Errorf("unable to find %v in any of api/group version", objectName)
		}
		klog.V(3).Infof("successfully obtained object: %v", dummy)
		return nil, nil
}

func findPodAndNode(objectResource *unstructured.Unstructured) (map[string]string , error) {
	var podToNodeMap 	map[string]string
	var temp 			map[string]interface{}
	var labels 			map[string]interface{}
	
	if objectResource.GetKind() == "Pod" {
		return getNodeFromPod(objectResource)
	} 
	
	temp 	= objectResource.UnstructuredContent()["spec"].(map[string]interface{})
	temp	= temp["selector"].(map[string]interface{})
	labels 	= temp["matchLabels"].(map[string]interface{})
	klog.V(2).Infof("object : %v",labels)
	return podToNodeMap, nil
}

func getNodeFromPod(podResource *unstructured.Unstructured) (map[string]string, error) {
	var podNodeMap 	map [string]string
	podName			:= podResource.GetName()
	podInterface 	:= v1.CoreV1Client.Pods(getNamespace())
	podObject		:= podInterface(podName,v1.GetOptions{})
	podNodeMap[podName]	= podObject.PodSpec.NodeName
	klog.V(2).Infof("constructed pod-node map: %v",podNodeMap)
	return podNodeMap, nil
}

func disAmbiguate(resources []apiResource) []apiResource {
	
	var unAmbigousResources []apiResource
	name := resources[0].getName()
	name = strings.ToLower(name) 
	switch(name) {
	case "replicasets": for _,resource := range(resources) {
		klog.V(4).Infof("disambiguation: comparing: %v, apps",resource.getAPIVersion() )
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
