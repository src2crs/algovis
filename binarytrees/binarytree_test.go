package binarytrees

import "testing"

// TestBinaryTree_NewBinaryTree_IsEmpty tests whether a newly created binary tree is empty.
func TestBinaryTree_NewBinaryTree_IsEmpty(t *testing.T) {
	tree := NewBinaryTree()
	if !tree.IsEmpty() {
		t.Errorf("NewBinaryTree() is not empty")
	}
}

// TestBinaryTree_NewBinaryTree_Root tests whether the Root() method of a newly created binary tree
// returns the correct root node.
func TestBinaryTree_NewBinaryTree_Root(t *testing.T) {
	tree := NewBinaryTree()
	root := tree.Root()
	if root != tree.root {
		t.Errorf("Root() returned %v, but expected %v", root, tree.root)
	}
}

// TestBinaryTree_CreateNode_IsEmpty tests whether a binary tree is not empty after creating a node.
func TestBinaryTree_CreateNode_IsEmpty(t *testing.T) {
	tree := NewBinaryTree()
	tree.CreateNode("L")
	if tree.IsEmpty() {
		t.Errorf("CreateNode() did not make binary tree non-empty")
	}
}

// TestBinaryTree_CreateNode_Path tests whether nodes are created at the correct positions.
func TestBinaryTree_CreateNode_Path(t *testing.T) {
	tree := NewBinaryTree()
	tree.CreateNode("L")
	tree.CreateNode("RLR")

	// The following tree structure is expected:
	//   root
	//   +-- L (empty)
	//   +-- R
	//       +-- L
	//       |   +-- L (empty)
	//       |   +-- R (empty)
	//       +-- R (empty)
	root := tree.root
	L := root.Left // (empty)
	R := root.Right
	RL := R.Left
	RR := R.Right   // (empty)
	RLL := RL.Left  // (empty)
	RLR := RL.Right // (empty)

	// Check node "root"
	if root.Path != "" {
		t.Errorf("root  has unexpected path string %s.", root.Path)
	}
	if root.IsEmpty() {
		t.Errorf("root is empty.")
	}

	// Check node "L"
	if L.Path != "L" {
		t.Errorf("Node \"L\" has unexpected path string %s.", L.Path)
	}
	if !L.IsEmpty() {
		t.Errorf("Node \"L\" is not empty.")
	}

	// Check node "R"
	if R.Path != "R" {
		t.Errorf("Node \"R\"  has unexpected path string %s.", R.Path)
	}
	if R.IsEmpty() {
		t.Errorf("Node \"R\" is empty.")
	}

	// Check node "RL"
	if RL.Path != "RL" {
		t.Errorf("Node \"RL\" has unexpected path string %s.", RL.Path)
	}
	if RL.IsEmpty() {
		t.Errorf("Node \"RL\" is empty.")
	}

	// Check node "RR"
	if RR.Path != "RR" {
		t.Errorf("Node \"RR\"  has unexpected path string %s.", RR.Path)
	}
	if !RR.IsEmpty() {
		t.Errorf("Node \"RR\" is not empty.")
	}

	// Check node "RLL"
	if RLL.Path != "RLL" {
		t.Errorf("Node \"RLL\" has unexpected path string %s.", RLL.Path)
	}
	if !RLL.IsEmpty() {
		t.Errorf("Node \"RLL\" is not empty.")
	}

	// Check node "RLR"
	if RLR.Path != "RLR" {
		t.Errorf("Node \"RLR\" has unexpected path string %s.", RLR.Path)
	}
	if !RLR.IsEmpty() {
		t.Errorf("Node \"RLR\" is not empty.")
	}
}

// TestBinaryTree_CreateLabelledNode_Label tests whether nodes are created with the correct labels.
func TestBinaryTree_CreateLabelledNode_Label(t *testing.T) {
	tree := NewBinaryTree()
	tree.CreateLabelledNode("LL", "foo")

	// The following tree structure is expected:
	//   root
	//    +-- L
	//    |   +-- L (label: "foo") (empty)
	//    |   +-- R (empty)
	//    +-- R (empty)
	root := tree.root
	L := root.Left
	LL := L.Left    // (empty)
	LR := L.Right   // (empty)
	R := root.Right // (empty)

	// Check node "root"
	if root.Path != "" {
		t.Errorf("root  has unexpected path string %s.", root.Path)
	}
	if root.IsEmpty() {
		t.Errorf("root is empty.")
	}

	// Check node "L"
	if L.Path != "L" {
		t.Errorf("Node \"L\" has unexpected path string %s.", L.Path)
	}
	if L.IsEmpty() {
		t.Errorf("Node \"L\" is empty.")
	}

	// Check node "LL"
	if LL.Path != "LL" {
		t.Errorf("Node \"LL\" has unexpected path string %s.", LL.Path)
	}
	if !LL.IsEmpty() {
		t.Errorf("Node \"LL\" is not empty.")
	}
	if LL.Label() != "foo" {
		t.Errorf("Node \"LL\" has unexpected label %s.", LL.Label())
	}

	// Check node "LR"
	if LR.Path != "LR" {
		t.Errorf("Node \"LR\" has unexpected path string %s.", LR.Path)
	}
	if !LR.IsEmpty() {
		t.Errorf("Node \"LR\" is not empty.")
	}

	// Check node "R"
	if R.Path != "R" {
		t.Errorf("Node \"R\"  has unexpected path string %s.", R.Path)
	}
	if !R.IsEmpty() {
		t.Errorf("Node \"R\" is not empty.")
	}
}
