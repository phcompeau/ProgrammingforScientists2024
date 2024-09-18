package main

import (
	"canvas"
	"image"
)

const (
	trailFrequency        = 10
	numberOfTrailFrames   = 100
	jupiterMoonMultiplier = 10.0
	trailThicknessFactor  = 0.2
)

func AnimateSystem(timePoints []Universe, canvasWidth, drawingFrequency int) []image.Image {
	images := make([]image.Image, 0)
	trails := make(map[int][]OrderedPair) // Map from body index to its trail of positions

	for i, u := range timePoints {
		// Only draw the frame if the index is divisible by the drawing frequency
		if (i*trailFrequency)%drawingFrequency == 0 {
			// Update trails for all bodies
			for bodyIndex, body := range u.bodies {
				trails[bodyIndex] = append(trails[bodyIndex], body.position)

				// shorten the current trail if it has exceeded the trail capacity
				if len(trails[bodyIndex]) > numberOfTrailFrames*trailFrequency {
					trails[bodyIndex] = trails[bodyIndex][1:] // Limit trail length to MaxTrailLengthFactor
				}
			}
		}
		if i%drawingFrequency == 0 {
			images = append(images, DrawToCanvas(u, canvasWidth, trails, i))
		}
	}

	return images
}

func DrawToCanvas(u Universe, canvasWidth int, trails map[int][]OrderedPair, frameCounter int) image.Image {
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// set canvas to white
	c.SetFillColor(canvas.MakeColor(255, 255, 255))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)

	// Draw trails for all bodies
	DrawTrails(&c, trails, frameCounter, u.width, float64(canvasWidth), u.bodies)

	// Draw the bodies themselves
	for _, b := range u.bodies {
		c.SetFillColor(canvas.MakeColor(b.red, b.green, b.blue))
		centerX := (b.position.x / u.width) * float64(canvasWidth)
		centerY := (b.position.y / u.width) * float64(canvasWidth)
		r := (b.radius / u.width) * float64(canvasWidth)

		if b.name == "Io" || b.name == "Ganymede" || b.name == "Callisto" || b.name == "Europa" {
			c.Circle(centerX, centerY, jupiterMoonMultiplier*r)
		} else {
			c.Circle(centerX, centerY, r)
		}
		c.Fill()
	}

	return c.GetImage()
}

func DrawTrails(c *canvas.Canvas, trails map[int][]OrderedPair, frameCounter int, uWidth, canvasWidth float64, bodies []Body) {
	for bodyIndex, b := range bodies {
		trail := trails[bodyIndex]
		numTrails := len(trail)

		// Adjust line width based on the body's radius
		lineWidth := (b.radius / uWidth) * float64(canvasWidth) * trailThicknessFactor // Adjust multiplier for desired thickness

		// adjust lineWidth if we are a jupiter moon
		if b.name == "Ganymede" || b.name == "Io" || b.name == "Callisto" || b.name == "Europa" {
			lineWidth *= jupiterMoonMultiplier
		}

		c.SetLineWidth(lineWidth)

		// Draw lines between consecutive trail points
		for j := 0; j < numTrails-1; j++ {
			// Calculate fading effect based on the position in the slice
			alpha := 255.0 * float64(j) / float64(numTrails)
			red := uint8((1-alpha/255.0)*255.0 + (alpha/255.0)*float64(b.red))
			green := uint8((1-alpha/255.0)*255 + (alpha/255.0)*float64(b.green))
			blue := uint8((1-alpha/255.0)*255 + (alpha/255.0)*float64(b.blue))

			c.SetStrokeColor(canvas.MakeColor(red, green, blue))

			// Draw a line between consecutive trail points
			startX := (trail[j].x / uWidth) * canvasWidth
			startY := (trail[j].y / uWidth) * canvasWidth
			endX := (trail[j+1].x / uWidth) * canvasWidth
			endY := (trail[j+1].y / uWidth) * canvasWidth

			c.MoveTo(startX, startY)
			c.LineTo(endX, endY)
			c.Stroke()
		}
	}
}
