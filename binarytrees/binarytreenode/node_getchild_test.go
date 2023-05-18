package binarytreenode

import "testing"

// TestNode_GetChild_Left tests whether the correct child node is returned on the left side.
func TestNode_GetChild_Left(t *testing.T) {
	e := NewRootNode()
	e.CreateChild("L")
	c := e.GetChild("L")
	if c != e.Left {
		t.Errorf("GetChild() returned child node with address %p, but expected %p", c, e.Left)
	}
}

// TestNode_GetChild_Right tests whether the correct child node is returned on the right side.
func TestNode_GetChild_Right(t *testing.T) {
	e := NewRootNode()
	e.CreateChild("R")
	c := e.GetChild("R")
	if c != e.Right {
		t.Errorf("GetChild() returned child node with address %p, but expected %p", c, e.Right)
	}
}

// TestNode_GetChild_InvalidDirection tests whether a panic occurs when an invalid direction is given.
func TestNode_GetChild_InvalidDirection(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetChild() did not panic")
		}
	}()
	e := NewRootNode()
	e.GetChild("X")
}
