package binarytrees

import "testing"

// TestNode_NodeInfo_root_empty tests whether the NodeInfo() method
// returns the correct NodeInfo for an empty root node without a label.
func TestNode_NodeInfo_root_empty_nolabel(t *testing.T) {
	root := NewRootNode()
	info := root.NodeInfo()
	if info.Name != "root" {
		t.Errorf("root.NodeInfo().Name is %s, but expected %s", info.Name, "root")
	}
	if info.Label != "" {
		t.Errorf("root.NodeInfo().Label is %s, but expected %s", info.Label, "")
	}
	if info.Visible {
		t.Errorf("root.NodeInfo().Visible is %t, but expected %t", info.Visible, false)
	}
}

// TestNode_NodeInfo_root_nonempty tests whether the NodeInfo() method
// returns the correct NodeInfo for a non-empty root node without a label.
func TestNode_NodeInfo_root_nonempty_nolabel(t *testing.T) {
	root := NewRootNode()
	root.CreateChild("L")
	info := root.NodeInfo()
	if info.Name != "root" {
		t.Errorf("root.NodeInfo().Name is %s, but expected %s", info.Name, "root")
	}
	if info.Label != "" {
		t.Errorf("root.NodeInfo().Label is %s, but expected %s", info.Label, "")
	}
	if !info.Visible {
		t.Errorf("root.NodeInfo().Visible is %t, but expected %t", info.Visible, true)
	}
}

// TestNode_NodeInfo_root_empty_label tests whether the NodeInfo() method
// returns the correct NodeInfo for an empty root node with a label.
func TestNode_NodeInfo_root_empty_label(t *testing.T) {
	root := NewRootNode()
	root.SetLabel("foo")
	info := root.NodeInfo()
	if info.Name != "root" {
		t.Errorf("root.NodeInfo().Name is %s, but expected %s", info.Name, "root")
	}
	if info.Label != "foo" {
		t.Errorf("root.NodeInfo().Label is %s, but expected %s", info.Label, "foo")
	}
	if info.Visible {
		t.Errorf("root.NodeInfo().Visible is %t, but expected %t", info.Visible, false)
	}
}

// TestNode_NodeInfo_root_nonempty_label tests whether the NodeInfo() method
// returns the correct NodeInfo for a non-empty root node with a label.
func TestNode_NodeInfo_root_nonempty_label(t *testing.T) {
	root := NewRootNode()
	root.CreateChild("L")
	root.SetLabel("foo")
	info := root.NodeInfo()
	if info.Name != "root" {
		t.Errorf("root.NodeInfo().Name is %s, but expected %s", info.Name, "root")
	}
	if info.Label != "foo" {
		t.Errorf("root.NodeInfo().Label is %s, but expected %s", info.Label, "foo")
	}
	if !info.Visible {
		t.Errorf("root.NodeInfo().Visible is %t, but expected %t", info.Visible, true)
	}
}
