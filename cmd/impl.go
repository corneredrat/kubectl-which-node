package cmd

import (
	"k8s.io/klog"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func findNodes(kind string, object string) error {

	klog.V(3).Infof("kind:%v object:%v", kind, object)
	klog.V(3).Info("finding API resources")
	
	// Check available "Kinds"
	apiResources, err := findApiResource(kind)
	// Check if kind requests exist
    if err != nil {
		return fmt.Errorf("Could not find API resources: %w",err)
	}
	klog.V(3).Infof("found kind %v",kind)
	
	objectResource, err := findObjectResource(apiResources[0],object) //api.go


	return nil
}
