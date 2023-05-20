package tree

import "testing"

// TestTree_HasChildAtPos tests the HasChildAtPos function.
// It creates a new tree and adds some nodes.
// It tests the HasChildAtPos function with each of the nodes.
func TestTree_HasChildAtPos(t *testing.T) {
	t1 := New()

	// Create nodes for the following tree:
	// root
	//   +- 0
	//   |  +- 0
	//   |  +- 1
	//   +- 1
	//   |  +- 1
	//   +- 2
	//      +- 0
	//      +- 2
	//         +- 1

	root := "root"
	n_0 := "0"
	n_1 := "1"
	n_2 := "2"
	n_0_0 := "0.0"
	n_0_1 := "0.1"
	n_1_1 := "1.1"
	n_2_0 := "2.0"
	n_2_2 := "2.2"
	n_2_2_1 := "2.2.1"

	// Add nodes to the tree.
	// Use graph.AddNodeWithId instead of tree.AddNodeAtPath because the tested
	// function is graph.HasChildAtPos and may be used by tree.AddNodeAtPath.
	for _, node := range []string{
		root, n_0, n_1, n_2, n_0_0, n_0_1, n_1_1, n_2_0, n_2_2, n_2_2_1,
	} {
		t1.AddNodeWithId(node)
	}

	// Test nodes where true is expected
	for _, test := range []struct {
		path  string
		index string
	}{
		{root, "0"}, {root, "1"}, {root, "2"},
		{n_0, "0"}, {n_0, "1"},
		{n_1, "1"},
		{n_2, "0"}, {n_2, "2"},
		{n_2_2, "1"},
	} {
		if !t1.HasChildAtPos(test.path, test.index) {
			t.Errorf("Expected tree to have child at position %s of node %s", test.index, test.path)
		}
	}

	// Test nodes where false is expected
	for _, test := range []struct {
		path  string
		index string
	}{
		{root, "3"},
		{n_0, "2"},
		{n_1, "0"}, {n_1, "2"},
		{n_2, "1"}, {n_2, "3"},
		{n_0_0, "0"}, {n_0_0, "1"},
		{n_0_1, "0"}, {n_0_1, "1"},
		{n_1_1, "0"}, {n_1_1, "1"},
		{n_2_0, "0"}, {n_2_0, "1"},
		{n_2_2, "0"},
		{n_2_2_1, "0"}, {n_2_2_1, "1"}, {n_2_2_1, "2"},
	} {
		if t1.HasChildAtPos(test.path, test.index) {
			t.Errorf("Expected tree not to have child at position %s of node %s", test.index, test.path)
		}
	}
}

// TestTree_HasChildAtPos_root_emptypath tests whether the HasChildAtPos function
// returns true for the root node when the path is empty.
func TestTree_HasChildAtPos_root_emptypath(t *testing.T) {
	t1 := New()
	t1.AddNodeWithId("root")
	t1.AddNodeWithId("0")

	if !t1.HasChildAtPos("", "0") {
		t.Errorf("Expected tree to have child at position 0 of node root")
	}
}

// TestTree_HasChildAtPos_root_nonexistingpath tests whether the HasChildAtPos function
// returns false for a node that doesn't exist.
func TestTree_HasChildAtPos_root_nonexistingpath(t *testing.T) {
	t1 := New()
	t1.AddNodeWithId("root")
	t1.AddNodeWithId("0")

	if t1.HasChildAtPos("fakepath", "0") {
		t.Errorf("Expected tree not to have child at position 0 of node \"fakepath\"")
	}
}
