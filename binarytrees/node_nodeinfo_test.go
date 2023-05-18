package binarytrees

import (
	"testing"

	"github.com/src2crs/algovis/graphdrawing"
)

// TestNode_NodeInfo_root_empty tests whether the NodeInfo() method
// returns the correct NodeInfo for an empty root node without a label.
func TestNode_NodeInfo_root_empty_nolabel(t *testing.T) {
	root := NewRootNode()
	actualinfo := root.NodeInfo()
	expectedinfo := graphdrawing.NodeInfo{
		Name:    "root",
		Label:   "",
		Visible: false,
	}
	if actualinfo != expectedinfo {
		t.Errorf("root.NodeInfo() is %#v, but expected %#v", actualinfo, expectedinfo)
	}
}

// TestNode_NodeInfo_root_nonempty tests whether the NodeInfo() method
// returns the correct NodeInfo for a non-empty root node without a label.
func TestNode_NodeInfo_root_nonempty_nolabel(t *testing.T) {
	root := NewRootNode()
	root.CreateChild("L")
	actualinfo := root.NodeInfo()
	expectedinfo := graphdrawing.NodeInfo{
		Name:    "root",
		Label:   "",
		Visible: true,
	}
	if actualinfo != expectedinfo {
		t.Errorf("root.NodeInfo() is %#v, but expected %#v", actualinfo, expectedinfo)
	}
}

// TestNode_NodeInfo_root_empty_label tests whether the NodeInfo() method
// returns the correct NodeInfo for an empty root node with a label.
func TestNode_NodeInfo_root_empty_label(t *testing.T) {
	root := NewRootNode()
	root.SetLabel("foo")
	actualinfo := root.NodeInfo()
	expectedinfo := graphdrawing.NodeInfo{
		Name:    "root",
		Label:   "foo",
		Visible: false,
	}
	if actualinfo != expectedinfo {
		t.Errorf("root.NodeInfo() is %#v, but expected %#v", actualinfo, expectedinfo)
	}
}

// TestNode_NodeInfo_root_nonempty_label tests whether the NodeInfo() method
// returns the correct NodeInfo for a non-empty root node with a label.
func TestNode_NodeInfo_root_nonempty_label(t *testing.T) {
	root := NewRootNode()
	root.CreateChild("L")
	root.SetLabel("foo")
	actualinfo := root.NodeInfo()
	expectedinfo := graphdrawing.NodeInfo{
		Name:    "root",
		Label:   "foo",
		Visible: true,
	}
	if actualinfo != expectedinfo {
		t.Errorf("root.NodeInfo() is %#v, but expected %#v", actualinfo, expectedinfo)
	}
}
