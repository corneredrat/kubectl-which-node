package cmd

import (
	"k8s.io/klog"
	"fmt"
)

func findNodes(kind string, object string) error {

	klog.Infof("kind:%v object:%v", kind, object)
	klog.V(3).Info("finding API resources")
	
	// Check available "Kinds"
	availableResourceNames, err := findApiResourceNames()
	// Check if kind requests exist
	if !exists(kind, availableResourceNames) {
		return fmt.Errorf("kind %v is not available in the server.",kind)
	}
	klog.V(3).Infof("found kind %v",kind)

	if err != nil {
		return fmt.Errorf("Could not find API resources: %w",err)
	}
	return nil
}
