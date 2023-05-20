package basicgraph

import (
	"github.com/src2crs/algovis/datastructures/graph"
	"github.com/src2crs/algovis/datastructures/graph/basicedge"
	"github.com/src2crs/algovis/datastructures/graph/basicnode"
)

// Graph is a struct that represents a graph.
// It implements the GraphInfo interface providing lists of nodes and edges.
type Graph struct {
	nodes []graph.NodeInfo // the graph's nodes
	edges []graph.EdgeInfo // the graph's edges
}

// New creates a new empty graph.
func New() *Graph {
	return &Graph{nodes: []graph.NodeInfo{}, edges: []graph.EdgeInfo{}}
}

// Nodes returns the graph's nodes.
func (g *Graph) Nodes() []graph.NodeInfo {
	return g.nodes
}

// Edges returns the graph's edges.
func (g *Graph) Edges() []graph.EdgeInfo {
	return g.edges
}

// Size returns the number of nodes in the graph.
func (g *Graph) Size() int {
	return len(g.Nodes())
}

// IsEmpty returns true if the graph is empty.
func (g *Graph) IsEmpty() bool {
	return g.Size() == 0
}

// AddNode adds the given node to the graph.
// If a node with the same id already exists, the function has no effect.
// Returns the node as a graph.NodeInfo.
func (g *Graph) AddNode(node graph.NodeInfo) graph.NodeInfo {
	if !g.HasNodeWithId(node.Id()) {
		g.nodes = append(g.nodes, node)
	}
	return node
}

// AddNodeWithId adds a node with the given id to the graph.
// If a node with the same id already exists, the function has no effect.
// Returns the node as a graph.NodeInfo.
func (g *Graph) AddNodeWithId(id string) graph.NodeInfo {
	return g.AddNode(basicnode.New(id, id, true))
}

// AddInvisibleNodeWithId adds an invisible node with the given id to the graph.
// If a node with the same id already exists, the function has no effect.
// Returns the node as a graph.NodeInfo.
func (g *Graph) AddInvisibleNodeWithId(id string) graph.NodeInfo {
	return g.AddNode(basicnode.New(id, id, false))
}

// AddNodes adds the given nodes to the graph.
func (g *Graph) AddNodes(nodes ...graph.NodeInfo) {
	for _, node := range nodes {
		g.AddNode(node)
	}
}

// AddEdge adds the given edge to the graph.
// If the edge's source or target node is not part of the graph, they are added.
// Does not check whether the edge is already part of the graph.
// I.e. the graph may contain multiple edges between the same nodes.
// Returns the edge as a graph.EdgeInfo.
func (g *Graph) AddEdge(edge graph.EdgeInfo) graph.EdgeInfo {
	g.AddNode(edge.Source())
	g.AddNode(edge.Target())
	g.edges = append(g.edges, edge)
	return edge
}

// AddEdgeBetween adds an edge between the given source and target nodes to the graph.
func (g *Graph) AddEdgeBetween(source, target graph.NodeInfo) {
	g.AddEdge(basicedge.New(source, target))
}

// AddEdgeBetweenIds adds an edge between the nodes with the given ids to the graph.
// If source or target node do not exist, they are created.
// Does not check whether the edge is already part of the graph.
// I.e. the graph may contain multiple edges between the same nodes.
// Returns the edge as a graph.EdgeInfo.
func (g *Graph) AddEdgeBetweenIds(sourceid string, targetid string) graph.EdgeInfo {
	source := g.GetOrCreateNodeWithId(sourceid)
	target := g.GetOrCreateNodeWithId(targetid)
	return g.AddEdge(basicedge.New(source, target))
}

// GetNodeWithId returns the node with the given id.
// If no such node exists, nil is returned.
func (g *Graph) GetNodeWithId(id string) graph.NodeInfo {
	for _, node := range g.Nodes() {
		if node.Id() == id {
			return node
		}
	}
	return nil
}

// GetOrCreateNodeWithId returns the node with the given id.
// If no such node exists, a new node with the given id is created and returned.
func (g *Graph) GetOrCreateNodeWithId(id string) graph.NodeInfo {
	if !g.HasNodeWithId(id) {
		g.AddNode(basicnode.New(id, id, true))
	}
	return g.GetNodeWithId(id)
}

// HasNodeWithId returns whether the graph contains a node with the given id.
func (g *Graph) HasNodeWithId(id string) bool {
	return g.GetNodeWithId(id) != nil
}

/*
TODO:
- Build a binary tree on top of the graph that ensures the tree structure and
  creates intermediate nodes and edges if necessary when adding nodes.
  - This can be a struct that embeds a graph and provides methods for adding nodes.
  - No methods for adding edges should be exposed.
  - Path strings like "0.1.0" or "LRL" can be used to specify the path to the node.
    - Are path strings like "0.1.0" a good idea? Can they be used in Tikz, Mermaid, etc.?
	- Better use slices of ints or bools and create the ids via the Id() method of the nodes?
	- What else could be used for n-ary trees?

Optimizations:
- Use a map for storing the nodes?
*/
