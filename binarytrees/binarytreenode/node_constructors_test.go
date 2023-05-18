package binarytreenode

import "testing"

// TestNode_NewNode_Path tests whether a newly created node has the correct path string.
func TestNode_NewNode_Path(t *testing.T) {
	e := NewNode("LLR")
	if e.Path != "LLR" {
		t.Errorf("NewNode() has path string %s, but expected %s", e.Path, "LLR")
	}
}

// TestNode_NewRootNode_Path tests whether a newly created root node has the correct path string.
func TestNode_NewRootNode_Path(t *testing.T) {
	e := NewRootNode()
	if e.Path != "" {
		t.Errorf("NewRootNode() has path string %s, but expected %s", e.Path, "")
	}
}

// TestNode_NewNode_IsEmpty tests whether a newly created node is empty.
func TestNode_NewNode_IsEmpty(t *testing.T) {
	e := NewNode("LLR")
	if !e.IsEmpty() {
		t.Errorf("NewNode() is not empty")
	}
}

// TestNode_Label_root checks the label of a root node.
// If no label is set, the empty string is expected.
// If a label is set, the label is expected.
func TestNode_Label_root(t *testing.T) {
	root := NewRootNode()
	if root.Label() != "" {
		t.Errorf("root.Label() returned %s, but expected empty string", root.Label())
	}
	root.SetLabel("foo")
	if root.Label() != "foo" {
		t.Errorf("root.Label() returned %s, but expected %s", root.Label(), "foo")
	}
}

// TestNode_Label_nonroot checks the label of a non-root node.
// If no label is set, the path string is expected.
// If a label is set, the label is expected.
func TestNode_Label_nonroot(t *testing.T) {
	root := NewRootNode()
	l := root.CreateChild("L")
	if l.Label() != "L" {
		t.Errorf("l.Label() returned %s, but expected %s", l.Label(), "L")
	}
	l.SetLabel("foo")
	if l.Label() != "foo" {
		t.Errorf("l.Label() returned %s, but expected %s", l.Label(), "foo")
	}
}

// TestNode_Name_root checks the name of a root node.
// The name of a root node is "root".
func TestNode_Name_root(t *testing.T) {
	root := NewRootNode()
	if root.Name() != "root" {
		t.Errorf("root.Name() returned %s, but expected %s", root.Name(), "root")
	}
}

// TestNode_Name_nonroot checks the name of a non-root node.
// The name of a non-root node is the path string.
func TestNode_Name_nonroot(t *testing.T) {
	root := NewRootNode()
	l := root.CreateChild("L")
	if l.Name() != "L" {
		t.Errorf("l.Name() returned %s, but expected %s", l.Name(), "L")
	}
}
