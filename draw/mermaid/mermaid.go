package mermaid

import (
	"fmt"
	"strings"

	"github.com/src2crs/algovis/datastructures/graph"
)

// Mermaid returns the graph in Mermaid format.
func Mermaid(g graph.GraphInfo) string {
	lines := []string{"graph TD"}

	for _, node := range g.Nodes() {
		nodeline := fmt.Sprintf("%s[%s]", node.Id(), node.Label())
		if !node.Visible() {
			nodeline += ":::hidden"
		}
		lines = append(lines, nodeline)
	}

	for _, edge := range g.Edges() {
		edgestring := "---"
		if !edge.Visible() {
			edgestring = "~~~"
		}
		lines = append(lines, fmt.Sprintf("%s %s %s", edge.Source().Id(), edgestring, edge.Target().Id()))
	}

	return strings.Join(lines, "\n  ")
}
