// Package path provides path utilities for trees.
// A path is a string that describes the path from the root node to a node in the tree.
// The path is a sequence of path segments separated by dots.
// For example, "0.1.0" means the left child of the right child of the left child of the root.
// For the root node, either the empty string "" or the special path string "root" can be used.
package path

import (
	"fmt"
	"strings"
)

// IsValid expects a path string.
// Returns true if the path is valid.
// A path is valid if it consists of one or more path segments separated by dots.
func IsValid(path string) bool {
	if path == "" || path == "root" {
		return true
	}
	pathSegments := strings.Split(path, ".")
	for _, pathSegment := range pathSegments {
		if !PathSegmentIsValid(pathSegment) {
			return false
		}
	}
	return true
}

// PathSegmentIsValid expects a path segment string.
// Returns true if the path segment is valid.
// A path segment is valid if it consists of one or more digits.
func PathSegmentIsValid(pathSegment string) bool {
	if pathSegment == "" {
		return false
	}
	for _, c := range pathSegment {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// Child expects a path string and a single path segment as a string.
// Returns the path of the child node at the given index.
// Returns \"invalid\" if the path or the path segment is invalid.
func Child(path, index string) string {
	if !IsValid(path) || !PathSegmentIsValid(index) {
		return "invalid"
	}
	if IsRoot(path) {
		return index
	}
	return path + "." + index
}

// Children expects a path string and an integer n.
// Returns the first n child paths of the node at the given path.
// Returns an empty slice if the path is invalid or n is zero or negative.
func Children(path string, n int) []string {
	if !IsValid(path) || n <= 0 {
		return []string{}
	}
	children := []string{}
	for i := 0; i < n; i++ {
		children = append(children, Child(path, fmt.Sprintf("%d", i)))
	}
	return children
}

// NormalizeRoot expects a path string and returns a path string, where the root node is identified by \"root\".
// No other checks are performed. In particular, the function does not check whether the path is valid.
func NormalizeRoot(path string) string {
	if path == "" {
		return "root"
	}
	return path
}

// IsRoot expects a path string and returns true if the path identifies the root node.
// No other checks are performed. In particular, the function does not check whether the path is valid.
func IsRoot(path string) bool {
	return NormalizeRoot(path) == "root"
}
