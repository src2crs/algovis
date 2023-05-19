package graph

// EdgeInfo is an interface for an edge in a graph that is used for drawing the graph.
// An edge connects a source node with a target node.
// An edge can be visible or invisible.
// Invisible edges are not drawn, but they may be of importance for the graph's layout.
type EdgeInfo interface {
	Source() NodeInfo // Source returns the edge's source node.
	Target() NodeInfo // Target returns the edge's target node.
	Visible() bool    // Visible returns whether the edge is visible.
}
