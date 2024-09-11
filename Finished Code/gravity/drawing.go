package main

import (
	"canvas"
	"image"
)

// let's place our drawing functions here.

//AnimateSystem takes a slice of Universe objects along with a canvas width
//parameter and generates a slice of images corresponding to drawing each Universe
//on a canvasWidth x canvasWidth canvas

//DrawToCanvas generates the image corresponding to a canvas after drawing a Universe
//object's bodies on a square canvas that is canvasWidth pixels x canvasWidth pixels
func DrawToCanvas(u Universe, canvasWidth int) image.Image {
	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// range over all the bodies and draw them.
	for _, b := range u.bodies {
		c.SetFillColor(canvas.MakeColor(b.red, b.green, b.blue))
		cx := (b.position.x / u.width) * float64(canvasWidth)
		cy := (b.position.y / u.width) * float64(canvasWidth)
		r := (b.radius / u.width) * float64(canvasWidth)
		c.Circle(cx, cy, r)
		c.Fill()
	}
	// we want to return an image!
	return c.GetImage()
}
