package path

import "testing"

// TestEdges tests the Edges function.
// It checks whether the function returns the correct result for several test cases.
func TestEdges(t *testing.T) {
	testcases := []struct {
		path     string
		expected []Edge
	}{
		{"", []Edge{}},
		{"root", []Edge{}},
		{"0", []Edge{{"root", "0"}}},
		{"0.0", []Edge{{"root", "0"}, {"0", "0.0"}}},
		{"0.1", []Edge{{"root", "0"}, {"0", "0.1"}}},
		{"0.0.0", []Edge{{"root", "0"}, {"0", "0.0"}, {"0.0", "0.0.0"}}},
		{"0.0.1", []Edge{{"root", "0"}, {"0", "0.0"}, {"0.0", "0.0.1"}}},
		{"0.1.0", []Edge{{"root", "0"}, {"0", "0.1"}, {"0.1", "0.1.0"}}},
	}

	for _, testcase := range testcases {
		actual := Edges(testcase.path)
		if !edgesEqual(actual, testcase.expected) {
			t.Errorf("Expected Edges(%v) to be %v, but was %v", testcase.path, testcase.expected, actual)
		}
	}
}

// edgesEqual checks whether two slices of edges are equal.
func edgesEqual(a, b []Edge) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].From != b[i].From || a[i].To != b[i].To {
			return false
		}
	}
	return true
}
