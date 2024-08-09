package main

import (
	"canvas"
	"fmt"
	"image"
)

//AnimateSystem takes a slice of Universe objects along with a canvas width
//parameter and a frequency parameter.
//Every frequency steps, it generates a slice of images corresponding to drawing each Universe
//on a canvasWidth x canvasWidth canvas.
//A scaling factor is a final input that is used to scale the stars big enough to see them.
func AnimateSystem(timePoints []*Universe, canvasWidth, frequency int, scalingFactor float64) []image.Image {
	images := make([]image.Image, 0)

	if len(timePoints) == 0 {
		panic("Error: no Universe objects present in AnimateSystem.")
	}

	// for every universe, draw to canvas and grab the image
	for i := range timePoints {
		if i%frequency == 0 {
			fmt.Println(i)
			images = append(images, timePoints[i].DrawToCanvas(canvasWidth, scalingFactor))
		}
	}

	return images
}

//DrawToCanvas generates the image corresponding to a canvas after drawing a Universe
//object's bodies on a square canvas that is canvasWidth pixels x canvasWidth pixels.
//A scaling factor is needed to make the stars big enough to see them.
func (u *Universe) DrawToCanvas(canvasWidth int, scalingFactor float64) image.Image {
	if u == nil {
		panic("Can't Draw a nil Universe.")
	}

	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// range over all the bodies and draw them.
	for _, b := range u.stars {
		c.SetFillColor(canvas.MakeColor(b.red, b.green, b.blue))
		cx := (b.position.x / u.width) * float64(canvasWidth)
		cy := (b.position.y / u.width) * float64(canvasWidth)
		r := scalingFactor * (b.radius / u.width) * float64(canvasWidth)
		c.Circle(cx, cy, r)
		c.Fill()
	}
	// we want to return an image!
	return c.GetImage()
}
