package cmd

import (
	"fmt"
	"github.com/gosuri/uitable"
)

func printPodsAndNodes(podNodeMap map[string]string) {
	table := uitable.New()
	table.AddRow("POD","NODE")
	for pod, node := range (podNodeMap) {
		table.AddRow(pod, node)
	}
	fmt.Println(table)
}