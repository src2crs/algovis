package basicnode

import (
	"testing"

	"github.com/src2crs/algovis/datastructures/graph"
)

// TestNode_interface tests whether the Node struct implements the Node interface.
// This test should never fail, but it should not compile if the Node struct does not implement the Node interface.
func TestNode_interface(t *testing.T) {
	var _ graph.NodeInfo = &Node{}
}
