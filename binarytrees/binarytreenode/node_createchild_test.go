package binarytreenode

import "testing"

// TestNode_CreateChild_Left tests whether a child node is created on the left side with the correct path string.
func TestNode_CreateChild_Left(t *testing.T) {
	e := NewRootNode()
	c := e.CreateChild("L")
	if c.Path != "L" {
		t.Errorf("CreateChild() has path string %s, but expected %s", c.Path, "L")
	}
}

// TestNode_CreateChild_Right tests whether a child node is created on the right side with the correct path string.
func TestNode_CreateChild_Right(t *testing.T) {
	e := NewRootNode()
	c := e.CreateChild("R")
	if c.Path != "R" {
		t.Errorf("CreateChild() has path string %s, but expected %s", c.Path, "R")
	}
}

// TestNode_CreateChild_InvalidDirection tests whether a panic occurs when an invalid direction is given.
func TestNode_CreateChild_InvalidDirection(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("CreateChild() did not panic")
		}
	}()
	e := NewRootNode()
	e.CreateChild("X")
}

// TestNode_CreateChild creates two child nodes and tests whether both are the root's left child.
func TestNode_CreateChild_overwrite_nonempty_panic(t *testing.T) {
	root := NewRootNode()
	l1 := root.CreateChild("L")
	if l1 != root.Left {
		t.Errorf("CreateChild() did not create node at correct position")
	}
	l2 := root.CreateChild("L")
	if l1 != l2 {
		t.Errorf("CreateChild() did not return the same node when creating the same non-empty child twice")
	}
}

// TestNode_CreateChild_IsEmpty tests whether after creating a new node,
// the parent node is not empty anymore, but the child node is empty.
func TestNode_CreateChild_IsEmpty(t *testing.T) {
	e := NewRootNode()
	e.CreateChild("L")
	if e.IsEmpty() {
		t.Errorf("CreateChild() did not make parent node non-empty")
	}
	if !e.Left.IsEmpty() {
		t.Errorf("CreateChild() did not make child node empty")
	}
}

// TestNode_CreateNodeAtSubPos tests whether sub-nodes are created at the correct position.
func TestNode_CreateNodeAtSubPos(t *testing.T) {
	root := NewRootNode()
	l := root.CreateNodeAtSubPos("L")
	ll := root.CreateNodeAtSubPos("LL")
	lr := root.CreateNodeAtSubPos("LR")
	r := root.CreateNodeAtSubPos("R")
	rl := root.CreateNodeAtSubPos("RL")
	rr := root.CreateNodeAtSubPos("RR")

	if l != root.Left {
		t.Errorf("CreateNodeAtSubPos() did not create node at correct position")
	}
	if ll != root.Left.Left {
		t.Errorf("CreateNodeAtSubPos() did not create node at correct position")
	}
	if lr != root.Left.Right {
		t.Errorf("CreateNodeAtSubPos() did not create node at correct position")
	}
	if r != root.Right {
		t.Errorf("CreateNodeAtSubPos() did not create node at correct position")
	}
	if rl != root.Right.Left {
		t.Errorf("CreateNodeAtSubPos() did not create node at correct position")
	}
	if rr != root.Right.Right {
		t.Errorf("CreateNodeAtSubPos() did not create node at correct position")
	}
}
