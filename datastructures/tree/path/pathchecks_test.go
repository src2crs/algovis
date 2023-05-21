package path

import "testing"

// TestIsValid tests the IsValid function.
// It checks whether the function returns the correct result for several test cases.
func TestIsValid(t *testing.T) {
	testcases := []struct {
		path     string
		expected bool
	}{
		{"", true},
		{"root", true},
		{"0", true},
		{"0.0", true},
		{"0.0.0", true},
		{"0.1.0", true},
		{"0.10.42.2", true},
		{"10.2.0.5", true},
		{"1.2.3.4.5", true},
		{"a", false},
		{"a.b", false},
		{"a.0", false},
		{"0.a", false},
		{"0.0.a", false},
		{"0.0.0.a", false},
	}

	for _, testcase := range testcases {
		actual := IsValid(testcase.path)
		if actual != testcase.expected {
			t.Errorf("Expected IsValid(%v) to be %v, but was %v", testcase.path, testcase.expected, actual)
		}
	}
}

// TestPathSegmentIsValid tests the PathSegmentIsValid function.
// It checks whether the function returns the correct result for several test cases.
func TestPathSegmentIsValid(t *testing.T) {
	testcases := []struct {
		path     string
		expected bool
	}{
		{"", false},
		{"root", false},
		{"0", true},
		{"0.0", false},
		{"0.0.0", false},
		{"0.1.0", false},
	}

	for _, testcase := range testcases {
		actual := PathSegmentIsValid(testcase.path)
		if actual != testcase.expected {
			t.Errorf("Expected PathSegmentIsValid(%v) to be %v, but was %v", testcase.path, testcase.expected, actual)
		}
	}
}

// TestIsRoot tests the IsRoot function.
func TestIsRoot(t *testing.T) {
	testcases := []struct {
		path     string
		expected bool
	}{
		{"", true},
		{"root", true},
		{"0", false},
		{"0.0", false},
		{"0.0.0", false},
		{"0.1.0", false},
		{"a.b.c", false},
		{"abc", false},
	}

	for _, testcase := range testcases {
		actual := IsRoot(testcase.path)
		if actual != testcase.expected {
			t.Errorf("Expected IsRoot(%v) to be %v, but was %v", testcase.path, testcase.expected, actual)
		}
	}

}
