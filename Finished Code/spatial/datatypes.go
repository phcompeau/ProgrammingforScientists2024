package main

// Cell contains a strategy (as a string) and a score (as a decimal).
type Cell struct {
	strategy string  //represents "C" or "D" corresponding to the type of prisoner in the cell
	score    float64 //represents the score of the cell based on the prisoner's relationship with neighboring cells
}

// GameBoard is a 2D slice of Cell objects representing our board.
type GameBoard [][]Cell
