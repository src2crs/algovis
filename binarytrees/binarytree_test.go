package binarytrees

import "testing"

// TestBinaryTree_NewBinaryTree_IsEmpty tests whether a newly created binary tree is empty.
func TestBinaryTree_NewBinaryTree_IsEmpty(t *testing.T) {
	tree := NewBinaryTree()
	if !tree.IsEmpty() {
		t.Errorf("NewBinaryTree() is not empty")
	}
}
