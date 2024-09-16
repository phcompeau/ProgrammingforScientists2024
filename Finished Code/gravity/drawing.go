package main

import (
	"canvas"
	"image"
)

// let's place our drawing functions here.

// AnimateSystem takes a slice of Universe objects along with a canvas width and an integer drawingFrequency.
// It generates a slice of images corresponding to drawing each Universe that is divisible by drawingFrequency on a canvasWidth x canvasWidth canvas (in pixels)
func AnimateSystem(timePoints []Universe, canvasWidth, drawingFrequency int) []image.Image {
	images := make([]image.Image, 0)

	for i, u := range timePoints {
		//only draw the current universe to an image IF i is divisible by drawingFrequency
		if i%drawingFrequency == 0 {
			images = append(images, DrawToCanvas(u, canvasWidth))
		}
	}

	return images
}

// DrawToCanvas generates the image corresponding to a canvas after drawing a Universe
// object's bodies on a square canvas that is canvasWidth pixels x canvasWidth pixels
func DrawToCanvas(u Universe, canvasWidth int) image.Image {
	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a white background
	c.SetFillColor(canvas.MakeColor(255, 255, 255))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// range over all the bodies and draw them.
	for _, b := range u.bodies {
		c.SetFillColor(canvas.MakeColor(b.red, b.green, b.blue))
		cx := (b.position.x / u.width) * float64(canvasWidth)
		cy := (b.position.y / u.width) * float64(canvasWidth)
		r := (b.radius / u.width) * float64(canvasWidth)
		if b.name == "Europa" || b.name == "Io" || b.name == "Ganymede" || b.name == "Callisto" {
			c.Circle(cx, cy, 10.0*r)
		} else {
			c.Circle(cx, cy, r)
		}
		c.Fill()
	}
	// we want to return an image!
	return c.GetImage()
}
