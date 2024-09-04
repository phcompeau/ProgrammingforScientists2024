package main

import (
	"canvas"
	"image"
)

func BoardsToImages(boards []GameBoard, cellWidth int) []image.Image {
	imageList := make([]image.Image, len(boards))
	for i := range boards {
		imageList[i] = boards[i].BoardToImage(cellWidth)
	}
	return imageList
}

// BoardToImage converts a GameBoard to an image, in which
// each cell has a cell width given by a parameter
func (g GameBoard) BoardToImage(cellWidth int) image.Image {
	//Parse out the # of rows and columns in the field
	rows := len(g)
	columns := len(g[0])

	height := rows * cellWidth
	width := columns * cellWidth
	c := canvas.CreateNewCanvas(width, height)

	darkGray := canvas.MakeColor(60, 60, 60)
	blue := canvas.MakeColor(7, 30, 230)
	red := canvas.MakeColor(239, 71, 111)

	//set the entire board as black

	c.SetFillColor(darkGray)
	c.ClearRect(0, 0, height, width)
	c.Clear()

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if g[i][j].strategy == "C" {
				//set fill color to blue
				c.SetFillColor(blue)
			} else if g[i][j].strategy == "D" {
				c.SetFillColor(red)
			}

			x := j * cellWidth
			y := i * cellWidth

			scalingFactor := 0.8 // to make circle smaller

			c.Circle(float64(x), float64(y), scalingFactor*float64(cellWidth)/2)
			c.Fill()

			/*
				//draw the rectangle
				x1, y1 := cellWidth*j, cellWidth*i
				x2, y2 := cellWidth*(j+1), cellWidth*(i+1)
				c.ClearRect(x1, y1, x2, y2)
			*/

			//what if we want to draw circles instead?
			/*
				x1, y1 := cellWidth*j, cellWidth*i
				radius := float64(cellWidth) / 2.0
				c.Circle(float64(x1), float64(y1), float64(radius))
			*/

			c.Fill()

		}
	}
	return c.GetImage()
}
