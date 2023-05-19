# AlgoVis

AlgoVis is a visualisation tool for algorithms and data structures.

## Binary Tree Visualiser

* Done: data types for binary trees with nodes that carry a path string and optionally a label string.
* Done: code that generates a mermaid graph description from a binary tree.

### TODO / Ideas

* Create a generalised data type for graphs consisting of:
  * Nodes with names, labels, and visibility
  * Edges with source, target, and visibility (derived)
* Create a generalised visualiser for graphs that generates a mermaid graph description.
* Refactor the binary tree data type to export the generalised graph data type.
* Refactor the binary tree data type to simplify its code.
  * Can it be changed to use the generalised graph data type?

### TODO later

* Create a command line tool that allows to generate a mermaid graph description from a list of node paths.
* Extend the binary tree data type to allow for more than two children per node.
* Create a visualiser for binary trees that generates a LaTeX / Tikz representation.
  * This may rely on the graph being a tree.
