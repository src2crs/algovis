package binarytreenode

import (
	"testing"

	"github.com/src2crs/algovis/graphdrawing"
)

// TestNode_EdgeInfo_source_empty_target_nil
// Creates an empty root node.
// Calls the EdgeInfo() method on this root with a nil target node.
// Expects a panic.
func TestNode_EdgeInfo_source_empty_target_nil(t *testing.T) {
	root := NewRootNode()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("root.EdgeInfo() did not panic for empty source node")
		}
	}()
	root.EdgeInfo(nil)
}

// TestNode_EdgeInfo_source_empty_target_empty
// Creates an empty root node.
// Calls the EdgeInfo() method on this root with itself as the target node.
// Expects an EdgeInfo with:
// - Source: "root"
// - Target: "root"
// - Visible: false
func TestNode_EdgeInfo_source_empty_target_empty(t *testing.T) {
	root := NewRootNode()
	actualinfo := root.EdgeInfo(root)
	expectedinfo := graphdrawing.EdgeInfo{
		Source:  "root",
		Target:  "root",
		Visible: false,
	}
	if actualinfo != expectedinfo {
		t.Errorf("root.EdgeInfo() is %#v, but expected %#v", actualinfo, expectedinfo)
	}
}

// TestNode_EdgeInfo_source_nonempty_target_empty
// Creates a root node.
// Creates a left child node.
// Calls the EdgeInfo() method on this root with the child as the target node.
// Expects an EdgeInfo with:
// - Source: "root"
// - Target: "L"
// - Visible: false (because the child is empty)
func TestNode_EdgeInfo_source_nonempty_target_empty(t *testing.T) {
	root := NewRootNode()
	root.CreateChild("L")
	actualinfo := root.EdgeInfo(root.Left)
	expectedinfo := graphdrawing.EdgeInfo{
		Source:  "root",
		Target:  "L",
		Visible: false,
	}
	if actualinfo != expectedinfo {
		t.Errorf("root.EdgeInfo() is %#v, but expected %#v", actualinfo, expectedinfo)
	}
}

// TestNode_EdgeInfo_source_nonempty_target_nonempty
// Creates a root node.
// Creates a left child node.
// Creates a left-left grandchild node.
// Calls the EdgeInfo() method on the root with its left child as the target node.
// Expects an EdgeInfo with:
// - Source: "root"
// - Target: "L"
// - Visible: true (because the child is non-empty)
func TestNode_EdgeInfo_source_nonempty_target_nonempty(t *testing.T) {
	root := NewRootNode()
	root.CreateChild("L")
	root.Left.CreateChild("L")
	actualinfo := root.EdgeInfo(root.Left)
	expectedinfo := graphdrawing.EdgeInfo{
		Source:  "root",
		Target:  "L",
		Visible: true,
	}
	if actualinfo != expectedinfo {
		t.Errorf("root.EdgeInfo() is %#v, but expected %#v", actualinfo, expectedinfo)
	}
}

// TestNode_EdgeInfo_source_empty_target_nonempty
// Creates a root node.
// Creates a left child node.
// Calls the EdgeInfo() method on the child with the root as the target node.
// Expects an EdgeInfo with:
// - Source: "L"
// - Target: "root"
// - Visible: false (because the child (source) is empty)
func TestNode_EdgeInfo_source_empty_target_nonempty(t *testing.T) {
	root := NewRootNode()
	root.CreateChild("L")
	actualinfo := root.Left.EdgeInfo(root)
	expectedinfo := graphdrawing.EdgeInfo{
		Source:  "L",
		Target:  "root",
		Visible: false,
	}
	if actualinfo != expectedinfo {
		t.Errorf("root.Left.EdgeInfo() is %#v, but expected %#v", actualinfo, expectedinfo)
	}
}
