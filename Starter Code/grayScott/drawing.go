package main

import (
	"canvas"
	"image"

	"gonum.org/v1/plot/palette/moreland"
)

//DrawBoards takes a slice of Board objects as input along with a cellWidth and n parameter.
//It returns a slice of images corresponding to drawing every nth board to a file,
//where each cell is cellWidth x cellWidth pixels.
func DrawBoards(boards []Board, cellWidth, n int) []image.Image {
	imageList := make([]image.Image, 0)

	// range over boards and if divisible by n, draw board and add to our list
	for i := range boards {
		if i%n == 0 {
			imageList = append(imageList, DrawBoard(boards[i], cellWidth))
		}
	}

	return imageList
}

//DrawBoard takes a Board objects as input along with a cellWidth and n parameter.
//It returns an image corresponding to drawing every nth board to a file,
//where each cell is cellWidth x cellWidth pixels.
func DrawBoard(b Board, cellWidth int) image.Image {
	// need to know how many pixels wide and tall to make our image

	height := len(b) * cellWidth
	width := len(b[0]) * cellWidth

	// think of a canvas as a PowerPoint slide that we draw on
	c := canvas.CreateNewCanvas(width, height)

	// canvas will start as black, so we should fill in colored squares

	for i := range b {
		for j := range b[i] {
			prey := b[i][j][0]
			predator := b[i][j][1]

			// we will color each cell according to a color map.

			val := predator / (predator + prey)

			//colorMap := palette.Reverse(moreland.Kindlmann()) // on white background
			//colorMap := moreland.Kindlmann() // on black background
			colorMap := moreland.SmoothBlueRed() // red-blue a la RNA seq

			//set min and max value of color map
			colorMap.SetMin(0)
			colorMap.SetMax(1)

			//find the color associated with the value predator / (predator + prey)
			color, err := colorMap.At(val)

			if err != nil {
				panic("Error converting color!")
			}

			// draw a rectangle in right place with this color
			c.SetFillColor(color)

			x := i * cellWidth
			y := j * cellWidth
			c.ClearRect(x, y, x+cellWidth, y+cellWidth)
			c.Fill()
		}
	}

	// canvas has an image field that we should return
	return c.GetImage()
}
