package mermaid

import (
	"testing"

	"github.com/src2crs/algovis/graphdrawing"
)

func TestDrawEdge_invisible(t *testing.T) {
	edgeinfo := graphdrawing.EdgeInfo{
		Source:  "root",
		Target:  "L",
		Visible: false,
	}
	actual := DrawEdge(edgeinfo)
	expected := "root ~~~ L"
	if actual != expected {
		t.Errorf("DrawEdge() is %s, but expected %s", actual, expected)
	}
}

func TestDrawEdge_visible(t *testing.T) {
	edgeinfo := graphdrawing.EdgeInfo{
		Source:  "root",
		Target:  "L",
		Visible: true,
	}
	actual := DrawEdge(edgeinfo)
	expected := "root --- L"
	if actual != expected {
		t.Errorf("DrawEdge() is %s, but expected %s", actual, expected)
	}
}
