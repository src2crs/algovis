package tree

import "github.com/src2crs/algovis/datastructures/graph"

// This file contains the methods that implement the Graph interface.

// Nodes returns all nodes of the tree.
func (t *Tree) Nodes() []graph.NodeInfo {
	return t.g.Nodes()
}

// Edges returns all edges of the tree.
func (t *Tree) Edges() []graph.EdgeInfo {
	return t.g.Edges()
}

// Size returns the number of nodes in the tree.
func (t *Tree) Size() int {
	return t.g.Size()
}

// IsEmpty returns true if the tree is empty.
func (t *Tree) IsEmpty() bool {
	return t.g.IsEmpty()
}
