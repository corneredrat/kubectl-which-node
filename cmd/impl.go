package cmd

import (
	"k8s.io/klog"
	"fmt"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func findAndPrintNodes(kind string, object string) error {

	klog.V(3).Infof("kind:%v object:%v", kind, object)
	klog.V(3).Info("finding API resources")
	
	// Check available "Kinds"
	apiResources, err := findApiResource(kind)
	// Check if kind requests exist
    if err != nil {
		return fmt.Errorf("Could not find API resource: %w",err)
	}
	klog.V(3).Infof("found kind %v",kind)
	
	// Get Object resource, like replicaset, or a deployment or a daemonset or a pod that is specified as input
	objectResource, err := findObjectResource(apiResources,object) //api.go
	if err != nil {
		return fmt.Errorf("Could not find API resources: %w",err)
	}
	klog.V(3).Infof("obtained api resource: %w", objectResource)
	
	// Get all (nodes <-> pods) mappings under given objectResource
	_ , err := findPodAndNode(objectResource)
	
	if err != nil {
		return fmt.Errorf("unable to obtain node info: %w",err)
	}
	return nil
}
