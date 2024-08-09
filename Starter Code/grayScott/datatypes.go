package main

//Cell contains two attributes corresponding to
//the concentration of predator (0-th element) and prey (1-th element) in the cell
type Cell [2]float64

//Board is a two-dimensional slice of Cells
type Board [][]Cell
