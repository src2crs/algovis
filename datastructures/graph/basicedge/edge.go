package basicedge

import "github.com/src2crs/algovis/datastructures/graph"

// Edge is an edge in a graph.
// It implements the EdgeInfo interface providing source and target nodes and a visibility.
// An edge is visible if both its source and target nodes are visible.
type Edge struct {
	source graph.NodeInfo // the edge's source node
	target graph.NodeInfo // the edge's target node
}

// New creates a new edge with the given source node, target node and visibility.
func New(source, target graph.NodeInfo) *Edge {
	return &Edge{source, target}
}

// Source returns the edge's source node.
func (e *Edge) Source() graph.NodeInfo {
	return e.source
}

// Target returns the edge's target node.
func (e *Edge) Target() graph.NodeInfo {
	return e.target
}

// Visible returns whether the edge is visible.
// An edge is visible if both its source and target nodes are visible.
func (e *Edge) Visible() bool {
	return e.Source().Visible() && e.Target().Visible()
}
