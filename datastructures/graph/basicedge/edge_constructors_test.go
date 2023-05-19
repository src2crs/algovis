package basicedge

import (
	"testing"

	"github.com/src2crs/algovis/datastructures/graph/basicnode"
)

// TestEdge_New_visibleNodes tests the New function with two visible nodes.
// It checks whether the returned edge is visible and has the correct source and target.
func TestEdge_New_visibleNodes(t *testing.T) {
	source := basicnode.New("source", "source node", true)
	target := basicnode.New("target", "target node", true)

	edge := New(source, target)

	if edge.Source() != source {
		t.Errorf("Expected edge's source to be %v, but was %v", source, edge.Source())
	}
	if edge.Target() != target {
		t.Errorf("Expected edge's target to be %v, but was %v", target, edge.Target())
	}
	if edge.Visible() != true {
		t.Errorf("Expected edge's visibility to be %t, but was %t", true, edge.Visible())
	}
}

// TestEdge_New_invisibleNodes tests the New function with two invisible nodes.
// It checks whether the returned edge is invisible and has the correct source and target.
func TestEdge_New_invisibleNodes(t *testing.T) {
	source := basicnode.New("source", "source node", false)
	target := basicnode.New("target", "target node", false)

	edge := New(source, target)

	if edge.Source() != source {
		t.Errorf("Expected edge's source to be %v, but was %v", source, edge.Source())
	}
	if edge.Target() != target {
		t.Errorf("Expected edge's target to be %v, but was %v", target, edge.Target())
	}
	if edge.Visible() != false {
		t.Errorf("Expected edge's visibility to be %t, but was %t", false, edge.Visible())
	}
}

// TestEdge_New_sourceInvisible tests the New function with an invisible source node.
// It checks whether the returned edge is invisible and has the correct source and target.
func TestEdge_New_sourceInvisible(t *testing.T) {
	source := basicnode.New("source", "source node", false)
	target := basicnode.New("target", "target node", true)

	edge := New(source, target)

	if edge.Source() != source {
		t.Errorf("Expected edge's source to be %v, but was %v", source, edge.Source())
	}
	if edge.Target() != target {
		t.Errorf("Expected edge's target to be %v, but was %v", target, edge.Target())
	}
	if edge.Visible() != false {
		t.Errorf("Expected edge's visibility to be %t, but was %t", false, edge.Visible())
	}
}

// TestEdge_New_targetInvisible tests the New function with an invisible target node.
// It checks whether the returned edge is invisible and has the correct source and target.
func TestEdge_New_targetInvisible(t *testing.T) {
	source := basicnode.New("source", "source node", true)
	target := basicnode.New("target", "target node", false)

	edge := New(source, target)

	if edge.Source() != source {
		t.Errorf("Expected edge's source to be %v, but was %v", source, edge.Source())
	}
	if edge.Target() != target {
		t.Errorf("Expected edge's target to be %v, but was %v", target, edge.Target())
	}
	if edge.Visible() != false {
		t.Errorf("Expected edge's visibility to be %t, but was %t", false, edge.Visible())
	}
}
