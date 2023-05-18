package binarytrees

import "testing"

// TestNode_Traverse_complete tests the Node.Traverse() method.
// It uses a traversal function that appends each node to a slice.
// Then it checks if the slice contains the expected nodes in the expected order.
func TestNode_Traverse_complete(t *testing.T) {
	// Create a binary tree with the following structure:
	//   root
	//   ├── L
	//   │   ├── LL
	//   │   │   ├── LLL
	//   │   │   └── LLR
	//   │   └── LR
	//   | 	 	 ├── LRL
	//   | 	 	 ├── LRR
	//   └── R
	//       └── RL
	//       └── RR
	//
	root := NewRootNode()
	l := root.CreateChild("L")
	ll := l.CreateChild("L")
	lll := ll.CreateChild("L")
	llr := ll.CreateChild("R")
	lr := l.CreateChild("R")
	lrl := lr.CreateChild("L")
	lrr := lr.CreateChild("R")
	r := root.CreateChild("R")
	rl := r.CreateChild("L")
	rr := r.CreateChild("R")

	// Create a slice to store the nodes in the order they are visited.
	var visited []*Node

	// Create a traversal function that appends each node to the slice.
	traversalFunc := func(n *Node) bool {
		visited = append(visited, n)
		return true
	}

	// Traverse the tree.
	root.Traverse(traversalFunc)

	// Check result.
	ok := true
	expected := []*Node{root, l, ll, lll, llr, lr, lrl, lrr, r, rl, rr}
	if len(visited) != len(expected) {
		ok = false
	} else {
		for i := range visited {
			if visited[i] != expected[i] {
				ok = false
			}
		}
	}
	if !ok {
		actualNames := make([]string, len(visited))
		for i := range visited {
			actualNames[i] = visited[i].Name()
		}
		expectedNames := make([]string, len(expected))
		for i := range expected {
			expectedNames[i] = expected[i].Name()
		}
		t.Errorf(
			"Expected and actual traversed nodes do not match:\n  expected: %v\n  actual:   %v",
			expectedNames,
			actualNames,
		)
	}
}

// TestNode_Traverse_incomplete tests the Node.Traverse() method.
// It uses a traversal function that appends only the first two nodes to a slice.
// Then it checks if the slice contains the first two nodes in the expected order.
func TestNode_Traverse_incomplete(t *testing.T) {
	// Create a binary tree with the following structure:
	//   root
	//   ├── L
	//   │   ├── LL
	//   │   │   ├── LLL
	//   │   │   └── LLR
	//   │   └── LR
	//   | 	 	 ├── LRL
	//   | 	 	 ├── LRR
	//   └── R
	//       └── RL
	//       └── RR
	//
	root := NewRootNode()
	l := root.CreateChild("L")
	ll := l.CreateChild("L")
	ll.CreateChild("L")
	ll.CreateChild("R")
	lr := l.CreateChild("R")
	lr.CreateChild("L")
	lr.CreateChild("R")
	r := root.CreateChild("R")
	r.CreateChild("L")
	r.CreateChild("R")

	// Create a slice to store the nodes in the order they are visited.
	var visited []*Node

	// Create a traversal function that appends only the first three nodes to the slice.
	traversalFunc := func(n *Node) bool {
		if len(visited) >= 3 {
			return false
		}
		visited = append(visited, n)
		return true
	}

	// Traverse the tree.
	root.Traverse(traversalFunc)

	// Check result.
	ok := true
	expected := []*Node{root, l, ll}
	if len(visited) != len(expected) {
		ok = false
	} else {
		for i := range visited {
			if visited[i] != expected[i] {
				ok = false
			}
		}
	}
	if !ok {
		actualNames := make([]string, len(visited))
		for i := range visited {
			actualNames[i] = visited[i].Name()
		}
		expectedNames := make([]string, len(expected))
		for i := range expected {
			expectedNames[i] = expected[i].Name()
		}
		t.Errorf(
			"Expected and actual traversed nodes do not match:\n  expected: %v\n  actual:   %v",
			expectedNames,
			actualNames,
		)
	}
}
