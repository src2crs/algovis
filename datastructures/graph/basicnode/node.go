package basicnode

// Node is a generic node in a graph.
// It implements the NodeInfo interface providing an id, a label and a visibility.
type Node struct {
	id      string // the node's id
	label   string // the node's label
	visible bool   // whether the node is visible
}

// New creates a new node with the given id, label and visibility.
func New(id string, label string, visible bool) *Node {
	return &Node{id, label, visible}
}

// Id returns the node's id.
func (n *Node) Id() string {
	return n.id
}

// Label returns the node's label.
func (n *Node) Label() string {
	return n.label
}

// Visible returns whether the node is visible.
func (n *Node) Visible() bool {
	return n.visible
}
