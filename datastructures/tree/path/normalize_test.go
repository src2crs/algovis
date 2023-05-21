package path

import "testing"

// TestNormalizeRoot tests the NormalizeRoot function.
// It checks whether the function returns the correct result for several test cases.
func TestNormalizeRoot(t *testing.T) {
	testcases := []struct {
		path     string
		expected string
	}{
		{"", "root"},
		{"root", "root"},
		{"0", "0"},
		{"0.0", "0.0"},
		{"0.0.0", "0.0.0"},
		{"0.1.0", "0.1.0"},
		{"a", "a"},
		{"a.b", "a.b"},
	}

	for _, testcase := range testcases {
		actual := NormalizeRoot(testcase.path)
		if actual != testcase.expected {
			t.Errorf("Expected NormalizeRoot(%v) to be %v, but was %v", testcase.path, testcase.expected, actual)
		}
	}
}
