package main

//GameBetween computes the contribution of Cell c2 to the score of Cell c1, assuming
import (
	"fmt"
	"os"
)

//that c2 is in the neighborhood of c1.
func GameBetween(c1, c2 Cell, b float64) float64 {
	if c1.strategy == "C" && c2.strategy == "C" {
		return 1.0
	} else if c1.strategy == "C" && c2.strategy == "D" {
		return 0.0
	} else if c1.strategy == "D" && c2.strategy == "C" {
		return b
	} else if c1.strategy == "D" && c2.strategy == "D" {
		return 0.0
	} else {
		fmt.Println("Error: Not all strategies are C or D.")
		os.Exit(1)
	}
	return -1.0
}

// updateScores goes through every cell, and plays the Prisoner's dilemma game
// with each of its in-field nieghbors (including itself). It updates the
// score of each cell to be the sum of that cell's winnings from the game.
func (g2 GameBoard) UpdateScores(g1 GameBoard, b float64) {
	//parse out the dimensions of the board
	rows := len(g2)
	columns := len(g2[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			for k := i - 1; k <= i+1; k++ {
				for l := j - 1; l <= j+1; l++ {
					if InField(k, l, rows, columns) && (i != k || j != l) {
						g2[i][j].score += GameBetween(g1[i][j], g1[k][l], b)
					}
				}
			}
		}
	}
}

// updateStrategies updates the strategies in g2 based on the scores in g2 and the
// strategies from g1 that these scores were computed from.
func (g2 GameBoard) UpdateStrategies(g1 GameBoard) {

	// parse out the # of rows and columns of matrix
	rows := len(g2)
	columns := len(g2[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			max := -1.0
			for k := i - 1; k <= i+1; k++ {
				for l := j - 1; l <= j+1; l++ {
					if InField(k, l, rows, columns) {
						if g2[k][l].score > max {
							max = g2[k][l].score
							g2[i][j].strategy = g1[k][l].strategy
						}
					}
				}
			}
		}
	}
}

// CreateBoard should create a new gameboard g of the ysize rows and xsize columns,
// so that g[r][c] gives the Cell at position (r,c).
func CreateBoard(rows, columns int) GameBoard {
	g := make(GameBoard, rows)
	// need to make all slices inside the cells
	for i := range g {
		g[i] = make([]Cell, columns)
	}
	return g
}

// InField returns true iff (row,col) is a valid cell in the field
func InField(i, j, rows, columns int) bool {
	return i >= 0 && i < rows && j >= 0 && j < columns
}

// Evolve() takes an intial field and evolves it for steps according to the game
// rule. At each step, it should call "updateScores()" and the updateStrategies
func (initialBoard GameBoard) Evolve(steps int, b float64) []GameBoard {
	boards := make([]GameBoard, steps+1)
	boards[0] = initialBoard
	for i := 1; i <= steps; i++ {
		boards[i] = Update(boards[i-1], b)
	}
	return boards
}

func Update(g1 GameBoard, b float64) GameBoard {
	g2 := CreateBoard(len(g1), len(g1[0]))

	g2.UpdateScores(g1, b)
	g2.UpdateStrategies(g1)

	return g2
}
