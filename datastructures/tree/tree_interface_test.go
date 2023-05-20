package tree

import (
	"testing"

	"github.com/src2crs/algovis/datastructures/graph"
)

// TestTree__graph_interface tests whether the Tree struct implements the Graph interface.
// This test should never fail, but it should not compile if the Tree struct does not implement the Graph interface.
func TestTree_interface(t *testing.T) {
	var _ graph.GraphInfo = &Tree{}
}
