package basicedge

import (
	"testing"

	"github.com/src2crs/algovis/datastructures/graph"
)

// TestEdge_interface tests whether the Edge struct implements the Edge interface.
// This test should never fail, but it should not compile if the Edge struct does not implement the Edge interface.
func TestEdge_interface(t *testing.T) {
	var _ graph.EdgeInfo = &Edge{}
}
