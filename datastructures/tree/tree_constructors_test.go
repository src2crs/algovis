package tree

import "testing"

// TestTree_New tests the New function.
// It creates a new tree and checks whether it is empty.
func TestTree_New(t *testing.T) {
	t1 := New()
	if !t1.IsEmpty() {
		t.Error("New tree is not empty")
	}
}
