package tree

import (
	"fmt"
	"strings"

	"github.com/src2crs/algovis/datastructures/graph"
	"github.com/src2crs/algovis/datastructures/graph/basicgraph"
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
func (t *Tree) HasChildAtPos(path, index string) bool {
	if path == "" {
		path = "root"
	}
	node := t.GetNodeWithId(path)
	if node == nil {
		return false
	}
	if path == "root" {
		return t.HasNodeWithId(index)
	}
	return t.HasNodeWithId(fmt.Sprintf("%s.%s", path, index))
}

// childPath expects a path string and a single path segment as a string.
// Returns the path of the child node at the given index.
// TODO: Provide a set of functions to properly check and handle invalid paths.
func childPath(path, index string) string {
	if path == "" || path == "root" {
		return index
	}
	return fmt.Sprintf("%s.%s", path, index)
}

// AddNodeAtPath adds a node at the given path.
// All nodes and edges along the path are created if they don't exist.
// Returns the node as a graph.NodeInfo.
// If the target node already exists, the function has no effect and the existing node is returned.
func (t *Tree) AddNodeAtPath(path string) graph.NodeInfo {
	if path == "" || path == "root" {
		return t.AddNodeWithId("root")
	}

	pathsegments := strings.Split(path, ".")
	currentpath := "root"
	for _, dir := range pathsegments {
		nextpath := childPath(currentpath, dir)
		if !t.HasNodeWithId(nextpath) {
			t.AddEdgeBetweenIds(currentpath, nextpath)
		}
		currentpath = nextpath
	}
	return t.GetOrCreateNodeWithId(currentpath)
}

// AddEdgeBetweenIds adds an edge between the nodes at the given paths.
// It reuses the AddEdgeBetweenIds function to create the nodes if they don't exist,
// but it handles the case that the root node may be specified as "" or "root".
// Furthermor, it will not create duplicate edges.
func (t *Tree) AddEdgeBetweenIds(sourceid, targetid string) {
	if sourceid == "" {
		sourceid = "root"
	}
	if !t.HasNodeWithId(targetid) {
		t.Graph.AddEdgeBetweenIds(sourceid, targetid)
	}
}
