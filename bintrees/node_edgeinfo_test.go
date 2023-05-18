package bintrees

import "testing"

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
	info := root.EdgeInfo(root)
	if info.Source != "root" {
		t.Errorf("root.EdgeInfo().Source is %s, but expected %s", info.Source, "root")
	}
	if info.Target != "root" {
		t.Errorf("root.EdgeInfo().Target is %s, but expected %s", info.Target, "root")
	}
	if info.Visible {
		t.Errorf("root.EdgeInfo().Visible is %t, but expected %t", info.Visible, false)
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
	info := root.EdgeInfo(root.Left)
	if info.Source != "root" {
		t.Errorf("root.EdgeInfo().Source is %s, but expected %s", info.Source, "root")
	}
	if info.Target != "L" {
		t.Errorf("root.EdgeInfo().Target is %s, but expected %s", info.Target, "L")
	}
	if info.Visible {
		t.Errorf("root.EdgeInfo().Visible is %t, but expected %t", info.Visible, false)
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
	info := root.EdgeInfo(root.Left)
	if info.Source != "root" {
		t.Errorf("root.EdgeInfo().Source is %s, but expected %s", info.Source, "root")
	}
	if info.Target != "L" {
		t.Errorf("root.EdgeInfo().Target is %s, but expected %s", info.Target, "L")
	}
	if !info.Visible {
		t.Errorf("root.EdgeInfo().Visible is %t, but expected %t", info.Visible, true)
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
	info := root.Left.EdgeInfo(root)
	if info.Source != "L" {
		t.Errorf("root.Left.EdgeInfo().Source is %s, but expected %s", info.Source, "L")
	}
	if info.Target != "root" {
		t.Errorf("root.Left.EdgeInfo().Target is %s, but expected %s", info.Target, "root")
	}
	if info.Visible {
		t.Errorf("root.Left.EdgeInfo().Visible is %t, but expected %t", info.Visible, false)
	}
}
