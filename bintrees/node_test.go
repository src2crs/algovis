package bintrees

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
