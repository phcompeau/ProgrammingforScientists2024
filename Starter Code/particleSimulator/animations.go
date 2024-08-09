package main

import (
	"canvas"
	"image"
)

// AnimateSystem takes a slice of Board objects, a canvas width, and a frequency integer.
// It returns a slice of images corresponding to drawing each board that is divisible by frequency
// to a canvas that is width x width.
func AnimateSystem(boards []*Board, canvasWidth int, frequency int) []image.Image {
	images := make([]image.Image, 0, len(boards))
	for i, b := range boards {
		if i%frequency == 0 {
			images = append(images, b.DrawToCanvas(canvasWidth))
		}
	}
	return images
}

// DrawToCanvas is a Board method that takes a canvas width parameter.
// It draws the Board to a Canvas scaled with given width and returns the image corresponding
// to this canvas.
func (b *Board) DrawToCanvas(canvasWidth int) image.Image {
	aspectRatio := b.height / b.width
	canvasHeight := int(float64(canvasWidth) * aspectRatio)
	c := canvas.CreateNewCanvas(canvasWidth, canvasHeight)

	// first, make a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	for _, p := range b.particles {
		// make a circle at p's position with the appropriate width
		scalingFactor := float64(canvasHeight) / b.height

		c.SetFillColor(canvas.MakeColor(p.red, p.green, p.blue))

		c.Circle(p.position.x*scalingFactor, p.position.y*scalingFactor, p.radius*scalingFactor)

		c.Fill()
	}
	return c.GetImage()
}
