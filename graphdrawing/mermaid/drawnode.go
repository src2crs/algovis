package mermaid

import (
	"fmt"

	"github.com/src2crs/algovis/graphdrawing"
)

// DrawNode returns the mermaid representation of a node.
func DrawNode(node graphdrawing.NodeInfo) string {
	template := "%s%s%s"
	labelString := ""
	if node.Label != "" {
		labelString = fmt.Sprintf("[%s]", node.Label)
	}
	visibleTag := ""
	if !node.Visible {
		visibleTag = ":::hidden"
	}
	return fmt.Sprintf(template, node.Name, labelString, visibleTag)
}
