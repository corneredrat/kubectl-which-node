package cmd

import (
	"k8s.io/klog"
)

func findNodes(kind string, object string) error {

	klog.Info("Kind:%v Object:%v", kind, object)
	klog.V(3).Info("Finding API resources")
	_, err = findApiResources()
	if err != nil {
		return fmt.Errorf("Could not find API resources: %w",err)
	}
	return nil
}
