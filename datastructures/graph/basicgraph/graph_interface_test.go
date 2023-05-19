package basicgraph

import (
	"testing"

	"github.com/src2crs/algovis/datastructures/graph"
)

// TestGraph_interface tests whether the Graph struct implements the Graph interface.
// This test should never fail, but it should not compile if the Graph struct does not implement the Graph interface.
func TestGraph_interface(t *testing.T) {
	var _ graph.GraphInfo = &Graph{}
}
