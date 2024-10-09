package main

import (
	"fmt"
	"strconv"
)

// UPGMA takes a distance matrix and a collection of species names as input.
// It returns a Tree (an array of nodes) resulting from applying
// UPGMA to this dataset.
func UPGMA(mtx DistanceMatrix, speciesNames []string) Tree {
	AssertSquareMatrix(mtx)
	AssertSameNumberSpecies(mtx, speciesNames)

	// we now know that len(mtx) and len(speciesNames) are equal
	numLeaves := len(mtx)
	t := InitializeTree(speciesNames)
	clusters := t.InitializeClusters() // will point clusters to leaves of the tree

	for p := numLeaves; p < 2*numLeaves-1; p++ {
		// p will represent the current internal node's index in tree
		// in each step of the for loop will be one step of UPGMA
		row, col, minVal := FindMinElement(mtx)

		// min value immediately gives me the age of current node
		t[p].Age = minVal / 2.0

		//next, set children of t[P]
		t[p].Child1 = clusters[row]
		t[p].Child2 = clusters[col]

		// we have now set the fields of t[p]

		// next, update the matrix

		//grab the sizes of the two clusters we are joining together
		clusterSize1 := CountLeaves(t[p].Child1)
		clusterSize2 := CountLeaves(t[p].Child2)

		// first, add a row and column corresponding to new cluster
		mtx = AddRowCol(row, col, clusterSize1, clusterSize2, mtx)
		mtx = DeleteRowCol(mtx, row, col)

		// finally, we clean up clusters
		//add current node to end of our clusters
		clusters = append(clusters, t[p])
		clusters = DeleteClusters(clusters, row, col)
	}
	return t
}

func AssertSquareMatrix(mtx DistanceMatrix) {
	numRows := len(mtx)
	// check that number of elements in each row is equal to numRows
	for r := 0; r < numRows; r++ {
		if len(mtx[r]) != numRows {
			fmt.Println("Row", r, "of matrix has length", len(mtx[r]), "and matrix has", numRows, "rows.")
			panic("Error!")
		}
	}
}

func AssertSameNumberSpecies(mtx DistanceMatrix, speciesNames []string) {
	if len(mtx) != len(speciesNames) {
		panic("Error: Number of rows of matrix don't match number of species.")
	}
}

// AddRowCol takes a distance matrix DistanceMatrix, two cluster sizes,
// and a row/col index (NOTE: col > row).
// It returns the matrix corresponding to "gluing" together clusters[row] and clusters[col]
// forming a new row/col of the matrix for the new cluster, computing
// distances to other elements of the matrix weighted according to the sizes
// of clusters[row] and clusters[col].
func AddRowCol(row, col, clusterSize1, clusterSize2 int, mtx DistanceMatrix) DistanceMatrix {
	numRows := len(mtx)
	newRow := make([]float64, numRows+1)

	// range over new row and set its values based on our weighted average
	// don't need to set the final element because it is zero by default, representing distance from cluster to itself
	for r := 0; r < len(newRow)-1; r++ {
		// only set a value of the row if it's not at index row or column
		if r != row && r != col {
			// set the value: average distance from element r to clusters corresponding to row and col
			newRow[r] = (float64(clusterSize1)*mtx[r][row] + float64(clusterSize2)*mtx[r][col]) / float64(clusterSize1+clusterSize2)
		}
	}

	// I have made my row, so I just need to add it to the matrix
	mtx = append(mtx, newRow)

	// also add the values to the end of each row of mtx ...
	for c := 0; c < numRows; c++ {
		mtx[c] = append(mtx[c], newRow[c])
	}

	return mtx
}

// CountLeaves takes a non-nil pointer to a Node object and returns
// the number of leaves in the tree rooted at the node. It returns 1 at a leaf.
func CountLeaves(v *Node) int {
	//base case: we are at a leaf, return 1
	if v.Child1 == nil && v.Child2 == nil {
		return 1
	}
	// what if only one is nil?
	if v.Child1 == nil {
		return CountLeaves(v.Child2) // we know this is OK since we made it here
	}
	if v.Child2 == nil {
		return CountLeaves(v.Child1) // same reasoning
	}
	//if I make it here, the node has two children that are non-nil
	return CountLeaves(v.Child1) + CountLeaves(v.Child2)
}

// DeleteClusters takes a slice of Node objects corresponding to "clusters"
// along with two index parameters. It deletes the two nodes from the
// slice at these two indices.
func DeleteClusters(clusters []*Node, row, col int) []*Node {
	//first, get rid of clusters[col]
	clusters = append(clusters[:col], clusters[col+1:]...)

	// next, get rid of clusters[row]
	clusters = append(clusters[:row], clusters[row+1:]...)

	return clusters
}

// DeleteRowCol takes a distance matrix along with two indices.
// It returns the matrix after deleting the row and column corresponding
// to each of the indices.
// NOTE: you should assume that row < col.
func DeleteRowCol(mtx DistanceMatrix, row, col int) DistanceMatrix {
	// first, let's delete the rows that we don't want any more

	mtx = append(mtx[:col], mtx[col+1:]...)
	mtx = append(mtx[:row], mtx[row+1:]...)

	// to get rid of columns, iterate over all rows and delete elements at indices row and col
	for r := range mtx {
		mtx[r] = append(mtx[r][:col], mtx[r][col+1:]...)
		mtx[r] = append(mtx[r][:row], mtx[r][row+1:]...)
	}

	return mtx
}

// FindMinElement takes a distance matrix as input.
// It returns a pair (row, col, val) where (row, col) corresponds to the minimum
// value of the matrix, and val is the minimum value.
// NOTE: you should have that row < col.
func FindMinElement(mtx DistanceMatrix) (int, int, float64) {
	if len(mtx) <= 1 || len(mtx[0]) <= 1 {
		panic("One row or one column!")
	}

	// can now assume that matrix is at least 2 x 2
	row := 0
	col := 1
	minVal := mtx[row][col]

	// range over matrix, and see if we can do better than minVal.
	for i := 0; i < len(mtx)-1; i++ {
		// start column ranging at i + 1
		for j := i + 1; j < len(mtx[i]); j++ {
			// do we have a winner?
			if mtx[i][j] < minVal {
				// update all three variables
				minVal = mtx[i][j]
				row = i
				col = j
				// col will still always be > row.
			}
		}
	}
	return row, col, minVal
}

// InitializeTree takes a slice of n species names as input.
// It returns a rooted binary tree with with 2n-1 total nodes, where
// the leaves are the first n and are labeled by the associated species names.
func InitializeTree(speciesNames []string) Tree {
	var t Tree

	// we haven't made slice
	// how long should the slice be? 2 * len(speciesNames) - 1
	numLeaves := len(speciesNames)
	t = make([]*Node, 2*numLeaves-1)

	// all pointers are currently nil
	// if I try to set anything, Go will get angry

	// we should create nodes and point the pointers in slice to my nodes
	for i := range t {
		// create a node
		var vx Node

		// set its fields
		if i < numLeaves {
			// at a leaf
			// set name of the leaf
			vx.Label = speciesNames[i]
		} else {
			vx.Label = "Ancestor Species: " + strconv.Itoa(i)
		}

		// I don't need to set any ages because .....
		// for internal nodes, it gets calculated by UPGMA
		// for leaves, it gets initialized as zero

		// point the current node of t to vx
		t[i] = &vx
	}

	return t
}

// InitializeClusters is a tree method that returns a slice of
// pointers to the leaves of the tree.
func (t Tree) InitializeClusters() []*Node {
	// the tree has 2n-1 total nodes
	// I want the first n nodes of the tree, corresponding to the leaves
	numNodes := len(t)
	numLeaves := (numNodes + 1) / 2

	clusters := make([]*Node, numLeaves)

	for i := range clusters {
		clusters[i] = t[i] // copy pointer of t[i] into clusters[i]
	}

	return clusters
}
