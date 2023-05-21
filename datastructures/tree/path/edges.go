package path

import "strings"

// Edge represents an edge in a tree.
type Edge struct {
	From string
	To   string
}

// Edges expects a path string.
// Returns the edges along the path.
// The edges are returned as a slice of Edge structs.
// The edges are returned in the order they appear in the path from the root to the node.
// If the path is invalid, the function returns an empty slice.
func Edges(path string) []Edge {
	if !IsValid(path) || IsRoot(path) {
		return []Edge{}
	}

	edges := []Edge{}
	pathSegments := strings.Split(path, ".")
	currentPath := "root"
	for _, pathSegment := range pathSegments {
		nextPath := Child(currentPath, pathSegment)
		edges = append(edges, Edge{currentPath, nextPath})
		currentPath = nextPath
	}
	return edges
}
