package mermaid

import (
	"testing"

	"github.com/src2crs/algovis/graphdrawing"
)

func TestDrawNode_invisible_nolabel(t *testing.T) {
	node := graphdrawing.NodeInfo{
		Name:    "root",
		Label:   "",
		Visible: false,
	}
	actual := DrawNode(node)
	expected := "root:::hidden"
	if actual != expected {
		t.Errorf("DrawNode() is %s, but expected %s", actual, expected)
	}

}

func TestDrawNode_invisible_label(t *testing.T) {
	node := graphdrawing.NodeInfo{
		Name:    "root",
		Label:   "foo",
		Visible: false,
	}
	actual := DrawNode(node)
	expected := "root[foo]:::hidden"
	if actual != expected {
		t.Errorf("DrawNode() is %s, but expected %s", actual, expected)
	}

}

func TestDrawNode_visible_nolabel(t *testing.T) {
	node := graphdrawing.NodeInfo{
		Name:    "root",
		Label:   "",
		Visible: true,
	}
	actual := DrawNode(node)
	expected := "root"
	if actual != expected {
		t.Errorf("DrawNode() is %s, but expected %s", actual, expected)
	}

}

func TestDrawNode_visible_label(t *testing.T) {
	node := graphdrawing.NodeInfo{
		Name:    "root",
		Label:   "foo",
		Visible: true,
	}
	actual := DrawNode(node)
	expected := "root[foo]"
	if actual != expected {
		t.Errorf("DrawNode() is %s, but expected %s", actual, expected)
	}
}
