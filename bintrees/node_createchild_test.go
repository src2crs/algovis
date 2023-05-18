package bintrees

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

// TestNode_CreateChild_Overwrite tests whether a child node is overwritten when a child node already exists on the given side.
func TestNode_CreateChild_Overwrite(t *testing.T) {
	e := NewRootNode()
	e.CreateChild("L")
	c := e.CreateChild("L")
	if e.Left != c {
		t.Errorf("CreateChild() did not overwrite existing child node")
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
