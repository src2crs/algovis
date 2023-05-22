# Package `datastructures`

This package contains interfaces and implementations of various graph like data structures.

## Overview of types and interfaces (not necessarily complete)

### Interfaces `GraphInfo`, `NodeInfo`, `EdgeInfo` in the `graph` package

These interfaces are used to store information about nodes and edges in a graph.
They are implememted by various concrete implementations.
They are used by the Drawing algorithms to query information about the graph.

### Type `basicgraph.Graph` in the `graph` package

A basic graph implementation that stores the graph as lists of `NodeInfo` and `EdgeInfo`.
Provides methods to add nodes and edges, and to query the graph.

### Type `Tree` in the `tree` package

A tree implementation that uses a `basicgraph.Graph` to store the tree.
Provides methods to add nodes, and to query the tree.
The methods here also ensure that the requred edges are added to the graph.

The package also contains a sub package `path` that provides functions to create
and validate paths for trees.

## TODO

* Add a way to influence the layout by adding invisible nodes when drawing a tree with mermaid.
  * A function or method `DrawTree` that takes a minimum number of children for each node,
    filling them up with invisible nodes if necessary?
  * Helper functions in `Tree` or `path` that list missing children for a node?
