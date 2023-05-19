package basicgraph

import (
	"testing"

	"github.com/src2crs/algovis/datastructures/graph"
	"github.com/src2crs/algovis/datastructures/graph/basicedge"
	"github.com/src2crs/algovis/datastructures/graph/basicnode"
)

// TestGraph_AddNode tests the AddNode function.
// It checks whether, after adding a node, the graph contains the node.
func TestGraph_AddNode(t *testing.T) {
	g := New()
	node := basicnode.New("node1", "node 1", true)

	g.AddNode(node)

	actualnodes := g.Nodes()
	expectednodes := []graph.NodeInfo{node}

	ok := true
	if len(actualnodes) != len(expectednodes) {
		ok = false
	} else {
		for i := range actualnodes {
			if actualnodes[i] != expectednodes[i] {
				ok = false
				break
			}
		}
	}
	if !ok {
		t.Errorf("Expected graph's nodes to be %v, but was %v", expectednodes, actualnodes)
	}
}

// TestGraph_AddNode_Duplicate tests the AddNode function.
// It checks whether, after adding a node twice, the graph contains the node only once.
func TestGraph_AddNode_Duplicate(t *testing.T) {
	g := New()
	node := basicnode.New("node1", "node 1", true)

	g.AddNode(node)
	g.AddNode(node)

	actualnodes := g.Nodes()
	expectednodes := []graph.NodeInfo{node}

	ok := true
	if len(actualnodes) != len(expectednodes) {
		ok = false
	} else {
		for i := range actualnodes {
			if actualnodes[i] != expectednodes[i] {
				ok = false
				break
			}
		}
	}
	if !ok {
		t.Errorf("Expected graph's nodes to be %v, but was %v", expectednodes, actualnodes)
	}
}

// TestGraph_AddNodeWithId tests the AddNodeWithId function.
// It checks whether, after adding a node, the graph contains the node.
func TestGraph_AddNodeWithId(t *testing.T) {
	g := New()
	node1 := g.AddNodeWithId("node1")

	actualnodes := g.Nodes()
	expectednodes := []graph.NodeInfo{node1}

	ok := true
	if len(actualnodes) != len(expectednodes) {
		ok = false
	} else {
		for i := range actualnodes {
			if actualnodes[i] != expectednodes[i] {
				ok = false
				break
			}
		}
	}
	if !ok {
		t.Errorf("Expected graph's nodes to be %v, but was %v", expectednodes, actualnodes)
	}
}

// TestGraph_AddNodeWithId_Duplicate tests the AddNodeWithId function.
// It checks whether, after adding a node twice, the graph contains the node only once.
func TestGraph_AddNodeWithId_Duplicate(t *testing.T) {
	g := New()
	node1 := g.AddNodeWithId("node1")
	g.AddNodeWithId("node1")

	actualnodes := g.Nodes()
	expectednodes := []graph.NodeInfo{node1}

	ok := true
	if len(actualnodes) != len(expectednodes) {
		ok = false
	} else {
		for i := range actualnodes {
			if actualnodes[i] != expectednodes[i] {
				ok = false
				break
			}
		}
	}
	if !ok {
		t.Errorf("Expected graph's nodes to be %v, but was %v", expectednodes, actualnodes)
	}
}

// TestGraph_AddInvisibleNodeWithId tests the AddInvisibleNodeWithId function.
// It checks whether, after adding a node, the graph contains the node.
func TestGraph_AddInvisibleNodeWithId(t *testing.T) {
	g := New()
	node1 := g.AddInvisibleNodeWithId("node1")

	actualnodes := g.Nodes()
	expectednodes := []graph.NodeInfo{node1}

	ok := true
	if len(actualnodes) != len(expectednodes) {
		ok = false
	} else {
		for i := range actualnodes {
			if actualnodes[i] != expectednodes[i] {
				ok = false
				break
			}
		}
	}
	if !ok {
		t.Errorf("Expected graph's nodes to be %v, but was %v", expectednodes, actualnodes)
	}
}

// TestGraph_AddInvisibleNodeWithId_Duplicate tests the AddInvisibleNodeWithId function.
// It checks whether, after adding a node twice, the graph contains the node only once.
func TestGraph_AddInvisibleNodeWithId_Duplicate(t *testing.T) {
	g := New()
	node1 := g.AddInvisibleNodeWithId("node1")
	g.AddInvisibleNodeWithId("node1")

	actualnodes := g.Nodes()
	expectednodes := []graph.NodeInfo{node1}

	ok := true
	if len(actualnodes) != len(expectednodes) {
		ok = false
	} else {
		for i := range actualnodes {
			if actualnodes[i] != expectednodes[i] {
				ok = false
				break
			}
		}
	}
	if !ok {
		t.Errorf("Expected graph's nodes to be %v, but was %v", expectednodes, actualnodes)
	}
}

// TestGraph_AddNodes tests the AddNodes function.
// It checks whether, after adding nodes, the graph contains the nodes.
func TestGraph_AddNodes(t *testing.T) {
	g := New()
	node1 := basicnode.New("node1", "node 1", true)
	node2 := basicnode.New("node2", "node 2", true)

	g.AddNodes(node1, node2)

	actualnodes := g.Nodes()
	expectednodes := []graph.NodeInfo{node1, node2}

	ok := true
	if len(actualnodes) != len(expectednodes) {
		ok = false
	} else {
		for i := range actualnodes {
			if actualnodes[i] != expectednodes[i] {
				ok = false
				break
			}
		}
	}
	if !ok {
		t.Errorf("Expected graph's nodes to be %v, but was %v", expectednodes, actualnodes)
	}
}

// TestGraph_AddEdge tests the AddEdge function.
// It checks whether, after adding an edge, the graph contains the edge.
func TestGraph_AddEdge(t *testing.T) {
	g := New()
	node1 := basicnode.New("source", "source node", true)
	node2 := basicnode.New("target", "target node", true)
	edge := basicedge.New(node1, node2)

	g.AddEdge(edge)

	actualedges := g.Edges()
	expectededges := []graph.EdgeInfo{edge}

	ok := true
	if len(actualedges) != len(expectededges) {
		ok = false
	} else {
		for i := range actualedges {
			if actualedges[i] != expectededges[i] {
				ok = false
				break
			}
		}
	}
	if !ok {
		t.Errorf("Expected graph's edges to be %v, but was %v", expectededges, actualedges)
	}
}

// TestGraph_AddEdge_CreatesNodes tests the AddEdge function.
// It checks whether, after adding an edge, the graph contains the edge's source and target nodes
// even if they were not part of the graph before.
func TestGraph_AddEdge_CreatesNodes(t *testing.T) {
	g := New()
	node1 := basicnode.New("source", "source node", true)
	node2 := basicnode.New("target", "target node", true)
	edge := basicedge.New(node1, node2)

	g.AddEdge(edge)

	actualnodes := g.Nodes()
	expectednodes := []graph.NodeInfo{node1, node2}

	ok := true
	if len(actualnodes) != len(expectednodes) {
		ok = false
	} else {
		for i := range actualnodes {
			if actualnodes[i] != expectednodes[i] {
				ok = false
				break
			}
		}
	}
	if !ok {
		t.Errorf("Expected graph's nodes to be %v, but was %v", expectednodes, actualnodes)
	}
}

// TestGraph_AddEdgeBetweenIds tests the AddEdgeBetweenIds function.
// It checks whether, after adding an edge to a graph with two nodes,
// the graph contains the edge and the two nodes.
func TestGraph_AddEdgeBetweenIds(t *testing.T) {
	g := New()
	node1 := g.AddNode(basicnode.New("node1", "node 1", true))
	node2 := g.AddNode(basicnode.New("node2", "node 2", true))

	g.AddEdgeBetweenIds("node1", "node2")

	actualedges := g.Edges()

	ok := true
	if len(actualedges) != 1 {
		ok = false
	} else {
		actualedge := actualedges[0]
		if actualedge.Source() != node1 || actualedge.Target() != node2 {
			ok = false
		}
	}
	if !ok {
		t.Errorf("Expected only one edge between %v and %v, but got %v", node1, node2, actualedges)
	}
}

// TestGraph_AddEdgeBetween tests the AddEdgeBetween function.
// It checks whether, after adding an edge between two nodes, the graph contains the edge.
func TestGraph_AddEdgeBetween(t *testing.T) {
	graph := New()
	source := basicnode.New("source", "source node", true)
	target := basicnode.New("target", "target node", true)

	graph.AddEdgeBetween(source, target)

	actualedges := graph.Edges()

	ok := true
	if len(actualedges) != 1 {
		ok = false
	} else {
		actualedge := actualedges[0]
		if actualedge.Source() != source || actualedge.Target() != target {
			ok = false
		}
	}
	if !ok {
		t.Errorf("Expected only one edge between %v and %v, but got %v", source, target, actualedges)
	}
}
