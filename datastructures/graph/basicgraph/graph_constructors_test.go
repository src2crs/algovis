package basicgraph

import "testing"

// TestGraph_New_empty tests the New function.
// It checks whether the returned graph has no nodes and edges.
func TestGraph_New_empty(t *testing.T) {
	graph := New()

	if !graph.IsEmpty() {
		t.Errorf("Expected graph's nodes to be empty, but was %v", graph.Nodes())
	}
	if len(graph.Edges()) != 0 {
		t.Errorf("Expected graph's edges to be empty, but was %v", graph.Edges())
	}
}
