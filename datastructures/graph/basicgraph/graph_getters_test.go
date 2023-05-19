package basicgraph

import (
	"testing"

	"github.com/src2crs/algovis/datastructures/graph/basicnode"
)

// TestNode_GetNodeWithId tests the GetNodeById method.
// It creates a graph with some nodes and tests whether the correct nodes are returned.
func TestNode_GetNodeWithId(t *testing.T) {
	g := New()
	node1 := basicnode.New("1", "1", true)
	node2 := basicnode.New("2", "2", true)
	node3 := basicnode.New("3", "3", true)
	g.AddNodes(node1, node2, node3)

	if g.GetNodeWithId("1") != node1 {
		t.Errorf("Expected node with id 1, but got %v", g.GetNodeWithId("1"))
	}
	if g.GetNodeWithId("2") != node2 {
		t.Errorf("Expected node with id 2, but got %v", g.GetNodeWithId("2"))
	}
	if g.GetNodeWithId("3") != node3 {
		t.Errorf("Expected node with id 3, but got %v", g.GetNodeWithId("3"))
	}
	if g.GetNodeWithId("4") != nil {
		t.Errorf("Expected nil, but got %v", g.GetNodeWithId("4"))
	}
}

// TestNode_GetOrCreateNodeWithId tests the GetOrCreateNodeWithId method.
// It creates a graph with some nodes and those, but also missing nodes.
// It tests whether all returned nodes are part of the graph.
func TestNode_GetOrCreateNodeWithId(t *testing.T) {
	g := New()
	node1 := g.AddNode(basicnode.New("1", "1", true))
	node2 := g.AddNode(basicnode.New("2", "2", true))

	// Test existing nodes
	got1 := g.GetOrCreateNodeWithId("1")
	if got1 != node1 {
		t.Errorf("Expected node with id 1, but got %v", got1)
	}
	got2 := g.GetOrCreateNodeWithId("2")
	if got2 != node2 {
		t.Errorf("Expected node with id 2, but got %v", got2)
	}

	// Test if GetOrCreateNodeWithId creates a missing node.
	node3 := g.GetOrCreateNodeWithId("3")
	if node3 == nil {
		t.Errorf("Expected node with id 3, but got nil")
	}
	got3 := g.GetNodeWithId("3")
	if got3 != node3 {
		t.Errorf("Expected node with id 3, but got %v", got3)
	}
}

// TestNode_HasNodeWithId tests the HasNodeWithId method.
// It creates a graph with some nodes and tests whether the method returns
// the correct result for known and unknown ids.
func TestNode_HasNodeWithId(t *testing.T) {
	g := New()
	node1 := basicnode.New("1", "1", true)
	node2 := basicnode.New("2", "2", true)
	node3 := basicnode.New("3", "3", true)
	g.AddNodes(node1, node2, node3)

	if !g.HasNodeWithId("1") {
		t.Errorf("Expected true, but got false")
	}
	if !g.HasNodeWithId("2") {
		t.Errorf("Expected true, but got false")
	}
	if !g.HasNodeWithId("3") {
		t.Errorf("Expected true, but got false")
	}
	if g.HasNodeWithId("4") {
		t.Errorf("Expected false, but got true")
	}
}
