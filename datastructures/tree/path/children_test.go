package path

import "testing"

// TestChild tests the ChildPath function.
// It checks whether the function returns the correct result for several test cases.
func TestChild(t *testing.T) {
	testcases := []struct {
		path     string
		index    string
		expected string
	}{
		{"", "0", "0"},
		{"root", "0", "0"},
		{"0", "0", "0.0"},
		{"0.0", "0", "0.0.0"},
		{"root", "1", "1"},
		{"0", "1", "0.1"},
		{"0.0", "1", "0.0.1"},
		{"1", "0", "1.0"},
		{"1.0", "0", "1.0.0"},
		{"fakepath", "0", "invalid"},
		{"0", "fakepath", "invalid"},
	}

	for _, testcase := range testcases {
		actual := Child(testcase.path, testcase.index)
		if actual != testcase.expected {
			t.Errorf("Expected ChildPath(%v, %v) to be %v, but was %v", testcase.path, testcase.index, testcase.expected, actual)
		}
	}
}

// TestChildren tests the Children function.
// It checks whether the function returns the correct result for several test cases.
func TestChildren(t *testing.T) {
	testcases := []struct {
		path     string
		n        int
		expected []string
	}{
		{"", 0, []string{}},
		{"", 1, []string{"0"}},
		{"", 2, []string{"0", "1"}},
		{"0", 0, []string{}},
		{"0", 1, []string{"0.0"}},
		{"0", 2, []string{"0.0", "0.1"}},
		{"0", 3, []string{"0.0", "0.1", "0.2"}},
		{"0.1", 0, []string{}},
		{"0.1", 1, []string{"0.1.0"}},
		{"0.1", 2, []string{"0.1.0", "0.1.1"}},
	}

	for _, testcase := range testcases {
		actual := Children(testcase.path, testcase.n)
		if !slicesEqual(actual, testcase.expected) {
			t.Errorf("Expected Children(%v, %v) to be %v, but was %v", testcase.path, testcase.n, testcase.expected, actual)
		}
	}
}

// slicesEqual checks whether two slices of strings are equal.
func slicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return true
		}
	}
	return true
}
