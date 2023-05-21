package tree

import (
	"github.com/src2crs/algovis/datastructures/graph"
	"github.com/src2crs/algovis/datastructures/graph/basicgraph"
	"github.com/src2crs/algovis/datastructures/tree/path"
)

// Tree is a tree data structure.
// It is implemented on top of a basic graph.
// The nodes are identified by a string of integers separated by dots.
// For example, "0.1.0" means the left child of the right child of the left child of the root.
// For the root node, the tree uses the special path string "root".
// However, the root node can also be identified by the empty string "".
type Tree struct {
	*basicgraph.Graph
}

// New creates a new empty tree.
func New() *Tree {
	return &Tree{basicgraph.New()}
}

// HasChildAtPos expects a path string and a single path segment as a string.
// Returns true if the node with the given path exists and has a child at the given index.
func (t *Tree) HasChildAtPos(p, index string) bool {
	return t.HasNodeWithId(path.Child(p, index))
}

// AddNodeAtPath adds a node at the given path.
// All nodes and edges along the path are created if they don't exist.
// Returns the node as a graph.NodeInfo.
// If the target node already exists, the function has no effect and the existing node is returned.
func (t *Tree) AddNodeAtPath(p string) graph.NodeInfo {
	p = path.NormalizeRoot(p)
	for _, edge := range path.Edges(p) {
		t.AddEdgeBetweenIds(edge.From, edge.To)
	}
	return t.GetOrCreateNodeWithId(p)
}

// AddEdgeBetweenIds adds an edge between the nodes at the given paths.
// It reuses the AddEdgeBetweenIds function to create the nodes if they don't exist,
// but it handles the case that the root node may be specified as "" or "root".
// Furthermor, it will not create duplicate edges.
func (t *Tree) AddEdgeBetweenIds(sourceid, targetid string) {
	sourceid = path.NormalizeRoot(sourceid)
	if !t.HasNodeWithId(targetid) {
		t.Graph.AddEdgeBetweenIds(sourceid, targetid)
	}
}
