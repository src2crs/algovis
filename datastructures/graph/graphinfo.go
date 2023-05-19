package graph

// GraphInfo is an interface for a graph that is used for drawing the graph.
// A graph consists of nodes and edges.
type GraphInfo interface {
	Nodes() []NodeInfo // Nodes returns the graph's nodes.
	Edges() []EdgeInfo // Edges returns the graph's edges.
}
