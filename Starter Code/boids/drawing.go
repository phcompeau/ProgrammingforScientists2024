package main

import (
	"canvas"
	"image"
	"math"
)

// Config contains customizable parameters for the animation
type Config struct {
	CanvasWidth     int
	BoidSize        float64
	BoidColor       Color
	BackgroundColor Color
}

// Color represents an RGB color with an optional alpha component
type Color struct {
	R, G, B, A uint8
}

// AnimateSystem takes a collection of Sky objects along with a configuration.
// It generates a slice of images corresponding to drawing every frequency-th Sky on the canvas.
func AnimateSystem(timePoints []Sky, config Config, drawingFrequency int) []image.Image {
	var images []image.Image

	for i, sky := range timePoints {
		if i%drawingFrequency == 0 {
			img := DrawToCanvas(sky, config)
			images = append(images, img)
		}
	}

	return images
}

func DrawToCanvas(currentSky Sky, config Config) image.Image {
	c := canvas.CreateNewCanvas(config.CanvasWidth, config.CanvasWidth)

	// Set background color
	c.SetFillColor(canvas.MakeColor(config.BackgroundColor.R, config.BackgroundColor.G, config.BackgroundColor.B))
	c.ClearRect(0, 0, config.CanvasWidth, config.CanvasWidth)
	c.Fill()

	for _, b := range currentSky.boids {
		// Draw the boid
		DrawBoid(&c, b, config, currentSky.width)
	}

	return c.GetImage()
}

// DrawBoid draws the boid on the canvas
func DrawBoid(c *canvas.Canvas, b Boid, config Config, skyWidth float64) {
	// Compute triangle points for the boid
	point1, point2, point3 := ComputeTrianglePoints(b.position, b.velocity)

	// Draw the boid's triangle
	c.SetFillColor(canvas.MakeColor(config.BoidColor.R, config.BoidColor.G, config.BoidColor.B))
	c.MoveTo((point1.x/skyWidth)*float64(config.CanvasWidth), (point1.y/skyWidth)*float64(config.CanvasWidth))
	c.LineTo((point2.x/skyWidth)*float64(config.CanvasWidth), (point2.y/skyWidth)*float64(config.CanvasWidth))
	c.LineTo((point3.x/skyWidth)*float64(config.CanvasWidth), (point3.y/skyWidth)*float64(config.CanvasWidth))
	c.LineTo((point1.x/skyWidth)*float64(config.CanvasWidth), (point1.y/skyWidth)*float64(config.CanvasWidth))
	c.Fill()

	// Draw triangle outline
	c.SetStrokeColor(canvas.MakeColor(0, 0, 0))
	c.MoveTo((point1.x/skyWidth)*float64(config.CanvasWidth), (point1.y/skyWidth)*float64(config.CanvasWidth))
	c.LineTo((point2.x/skyWidth)*float64(config.CanvasWidth), (point2.y/skyWidth)*float64(config.CanvasWidth))
	c.LineTo((point3.x/skyWidth)*float64(config.CanvasWidth), (point3.y/skyWidth)*float64(config.CanvasWidth))
	c.LineTo((point1.x/skyWidth)*float64(config.CanvasWidth), (point1.y/skyWidth)*float64(config.CanvasWidth))
	c.Stroke()
}

// ComputeTrianglePoints calculates the three points of a triangle representing a boid
func ComputeTrianglePoints(position OrderedPair, velocity OrderedPair) (OrderedPair, OrderedPair, OrderedPair) {
	direction := math.Atan2(velocity.y, velocity.x)

	point1 := OrderedPair{
		x: position.x + 80*math.Cos(direction),
		y: position.y + 80*math.Sin(direction),
	}
	point2 := OrderedPair{
		x: position.x + 30*math.Cos(direction+2*math.Pi/3),
		y: position.y + 30*math.Sin(direction+2*math.Pi/3),
	}
	point3 := OrderedPair{
		x: position.x + 30*math.Cos(direction+4*math.Pi/3),
		y: position.y + 30*math.Sin(direction+4*math.Pi/3),
	}

	return point1, point2, point3
}
