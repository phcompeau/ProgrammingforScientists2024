package main

// UPGMA takes a distance matrix and a collection of species names as input.
// It returns a Tree (an array of nodes) resulting from applying
// UPGMA to this dataset.

// AddRowCol takes a distance matrix Given a DistanceMatrix, a slice of current clusters,
// and a row/col index (NOTE: col > row).
// It returns the matrix corresponding to "gluing" together clusters[row] and clusters[col]
// forming a new row/col of the matrix for the new cluster, computing
// distances to other elements of the matrix weighted according to the sizes
// of clusters[row] and clusters[col].

// DeleteClusters takes a slice of Node objects corresponding to "clusters"
// along with two index parameters. It deletes the two nodes from the
// slice at these two indices.

// DeleteRowCol takes a distance matrix along with two indices.
// It returns the matrix after deleting the row and column corresponding
// to each of the indices.
// NOTE: you should assume that row < col.

// FindMinElement takes a distance matrix as input.
// It returns a pair (row, col, val) where (row, col) corresponds to the minimum
// value of the matrix, and val is the minimum value.
// NOTE: you should have that row < col.

// InitializeClusters takes a tree and returns a slice of
// pointers to the leaves of the tree.

// InitializeTree takes a slice of n species names as input.
// It returns a rooted binary tree with with 2n-1 total nodes, where
// the leaves are the first n and have the associated species names.
