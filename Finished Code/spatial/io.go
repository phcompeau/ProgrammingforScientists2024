package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadBoardFromFile should open the given file and read the initial
// values for the field. The first line of the file will contain
// two space-separated integers saying how many rows and columns
// the field should have:
//    10 15
// each subsequent line will consist of a string of Cs and Ds, which
// are the initial strategies for the cells:
//    CCCCCCDDDCCCCCC
func ReadBoardFromFile(filename string) GameBoard {
	in, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error: couldn't open the file!")
	}

	defer in.Close()

	// we will store the lines of the file as a slice of strings
	lines := make([]string, 0)

	//create a scanner object and read in the lines one at a time
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//Parse out the data in the first line of the file
	var params []string = strings.Split(lines[0], " ")

	rows, err1 := strconv.Atoi(params[0])
	columns, err2 := strconv.Atoi(params[1])

	if err1 != nil || err2 != nil {
		fmt.Println("Error: row and column must not be numbers.")
		os.Exit(1)
	}

	// Initialize the game board
	g := CreateBoard(rows, columns)

	// Parse the remaining lines of the data and enter them into the cells

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			g[i][j].strategy = lines[i+1][j : j+1] //make sure that it has type string
			g[i][j].score = 0.0                    // default behavior, redundant
		}
	}
	return g
}
