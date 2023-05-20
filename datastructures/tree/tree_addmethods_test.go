package tree

import (
	"testing"

	"github.com/src2crs/algovis/datastructures/graph"
)

// TestTree_AddNodeAtPath tests the AddNodeAtPath function.
// It creates a new tree and adds a root node.
// It checks whether the root node exists.
func TestTree_AddNodeAtPath(t *testing.T) {
	t1 := New()
	root := t1.AddNodeAtPath("root")

	actualnodes := t1.Nodes()
	expectednodes := []graph.NodeInfo{root}

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
		t.Errorf("Expected tree's nodes to be %v, but was %v", expectednodes, actualnodes)
	}
}

// TestTree_AddNodeAtPath_intermediate tests the AddNodeAtPath function.
// It creates a new tree and adds a node at a position where
// intermediate nodes and edges have to be created.
// It checks whether the intermediate nodes and edges exist after adding the node.
func TestTree_AddNodeAtPath_intermediate(t *testing.T) {
	t1 := New()
	t1.AddNodeAtPath("root")
	t1.AddNodeAtPath("0.0.1.0")

	actualnodes := t1.Nodes()
	actualids := []string{}
	for _, node := range actualnodes {
		actualids = append(actualids, node.Id())
	}

	expectednodeids := []string{"root", "0", "0.0", "0.0.1", "0.0.1.0"}

	ok := true
	if len(actualids) != len(expectednodeids) {
		ok = false
	} else {
		for i := range actualids {
			if actualids[i] != expectednodeids[i] {
				ok = false
				break
			}
		}
	}
	if !ok {
		t.Errorf("Expected tree's node ids to be %v, but was %v", expectednodeids, actualids)
	}

	actualedges := t1.Edges()
	actualedgeids := []struct{ from, to string }{}
	for _, edge := range actualedges {
		actualedgeids = append(actualedgeids, struct{ from, to string }{edge.Source().Id(), edge.Target().Id()})
	}
	expectededgeids := []struct{ from, to string }{
		{"root", "0"}, {"0", "0.0"}, {"0.0", "0.0.1"}, {"0.0.1", "0.0.1.0"},
	}

	ok = true
	if len(actualedgeids) != len(expectededgeids) {
		ok = false
	} else {
		for i := range actualedgeids {
			if actualedgeids[i] != expectededgeids[i] {
				ok = false
				break
			}
		}
	}
	if !ok {
		t.Errorf("Expected tree's edge ids to be %v, but was %v", expectededgeids, actualedgeids)
	}
}

// TestTree_AddNodeAtPath_emptypath tests whether the AddNodeAtPath function
// creates a root node when the path is empty.
func TestTree_AddNodeAtPath_emptypath(t *testing.T) {
	t1 := New()
	t1.AddNodeAtPath("")

	actualnodes := t1.Nodes()
	actualids := []string{}
	for _, node := range actualnodes {
		actualids = append(actualids, node.Id())
	}

	expectednodeids := []string{"root"}

	ok := true
	if len(actualids) != len(expectednodeids) {
		ok = false
	} else {
		for i := range actualids {
			if actualids[i] != expectednodeids[i] {
				ok = false
				break
			}
		}
	}
	if !ok {
		t.Errorf("Expected tree's node ids to be %v, but was %v", expectednodeids, actualids)
	}
}

// TestTree_AddEdgeBetweenIds_emptypath tests whether the AddEdgeBetweenIds function
// treats the empty path as the root node.
func TestTree_AddEdgeBetweenIds_emptypath(t *testing.T) {
	t1 := New()
	t1.AddNodeAtPath("root")
	t1.AddEdgeBetweenIds("", "0")

	actualnodes := t1.Nodes()
	actualnodeids := []string{}
	for _, node := range actualnodes {
		actualnodeids = append(actualnodeids, node.Id())
	}
	expectednodeids := []string{"root", "0"}

	ok := true
	if len(actualnodeids) != len(expectednodeids) {
		ok = false
	} else {
		for i := range actualnodeids {
			if actualnodeids[i] != expectednodeids[i] {
				ok = false
				break
			}
		}
	}
	if !ok {
		t.Errorf("Expected tree's node ids to be %v, but was %v", expectednodeids, actualnodeids)
	}
}

// TestTree_AddEdgeBetweenIds_noduplicates tests whether the AddEdgeBetweenIds function
// adds an edge only if it does not exist yet.
func TestTree_AddEdgeBetweenIds_noduplicates(t *testing.T) {
	t1 := New()
	t1.AddNodeAtPath("0")

	t1.AddEdgeBetweenIds("", "0")

	actualedges := t1.Edges()

	if len(actualedges) != 1 {
		t.Errorf("Expected tree to have 1 edge, but was %v", actualedges)
	}

}
