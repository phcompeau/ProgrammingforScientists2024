package main

import (
	"gifhelper"
	"os"
	"strconv"
)

func main() {
	filename := os.Args[1]

	b, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		os.Exit(1)
	}

	steps, err := strconv.Atoi(os.Args[3])
	if err != nil {
		os.Exit(1)
	}

	cellWidth, err := strconv.Atoi(os.Args[4])
	if err != nil {
		os.Exit(1)
	}

	initialBoard := ReadBoardFromFile(filename)

	boards := initialBoard.Evolve(steps, b)

	// generate the GIF
	imageList := BoardsToImages(boards, cellWidth)
	gifhelper.ImagesToGIF(imageList, "prisoners")
}
