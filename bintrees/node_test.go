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
