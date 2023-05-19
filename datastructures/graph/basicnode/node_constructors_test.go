package basicnode

import "testing"

// TestBasicNode_New tests the New function.
// It checks whether the returned node has the correct id, label and visibility.
func TestBasicNode_New(t *testing.T) {
	id := "node1"
	label := "node 1"
	visible := true

	node := New(id, label, visible)

	if node.Id() != id {
		t.Errorf("Expected node's id to be %s, but was %s", id, node.Id())
	}
	if node.Label() != label {
		t.Errorf("Expected node's label to be %s, but was %s", label, node.Label())
	}
	if node.Visible() != visible {
		t.Errorf("Expected node's visibility to be %t, but was %t", visible, node.Visible())
	}
}
