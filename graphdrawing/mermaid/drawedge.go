package mermaid

import (
	"fmt"

	"github.com/src2crs/algovis/graphdrawing"
)

func DrawEdge(edgeinfo graphdrawing.EdgeInfo) string {
	template := "%s %s %s"
	edgestring := "---"
	if !edgeinfo.Visible {
		edgestring = "~~~"
	}
	return fmt.Sprintf(template, edgeinfo.Source, edgestring, edgeinfo.Target)
}
