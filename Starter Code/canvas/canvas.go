//Thanks to Carl Kingsford and Hannah Kim for producing
//the canvas.go file, which makes drawing objects on a
//virtual canvas much easier by communicating with
//drawing code in Go from code.google.com.
//Thanks to John D. Cox for helping to optimize this code for most of our purposes.

package canvas

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/llgcode/draw2d/draw2dimg"
)

type Canvas struct {
	gc     *draw2dimg.GraphicContext
	img    image.Image
	width  int // both width and height are in pixels
	height int
}

func (c *Canvas) GetImage() image.Image {
	return c.img
}

// Create a new canvas
func CreateNewCanvas(w, h int) Canvas {
	i := image.NewRGBA(image.Rect(0, 0, w, h))
	gc := draw2dimg.NewGraphicContext(i)

	gc.SetStrokeColor(image.Black)
	gc.SetFillColor(image.White)
	// fill the background
	gc.Clear()
	gc.SetFillColor(image.Black)

	return Canvas{gc, i, w, h}
}

// Create a new Paletted canvas
func CreateNewPalettedCanvas(w, h int, cp color.Palette) Canvas {
	if cp == nil {
		cp = palette.WebSafe
	}
	i := image.NewPaletted(image.Rect(0, 0, w, h), cp)
	gc := draw2dimg.NewGraphicContext(i)

	gc.SetStrokeColor(image.Black)
	gc.SetFillColor(image.White)
	// fill the background
	gc.Clear()
	gc.SetFillColor(image.Black)

	return Canvas{gc, i, w, h}
}

// Create a new color
func MakeColor(r, g, b uint8) color.Color {
	return &color.RGBA{r, g, b, 255}
}

// Move the current point to (x,y)
func (c *Canvas) MoveTo(x, y float64) {
	c.gc.MoveTo(x, y)
}

// Draw a line from the current point to (x,y), and set the current point to (x,y)
func (c *Canvas) LineTo(x, y float64) {
	c.gc.LineTo(x, y)
}

// Draw an arc from the current point to (x, y)
// Can be used to easily draw a circle or an ellipse
func (c *Canvas) ArcTo(x, y, radiusX, radiusY, degStart, degEnd float64) {
	c.gc.ArcTo(x, y, radiusX, radiusY, degStart, degEnd)
}

// Set the line color
func (c *Canvas) SetStrokeColor(col color.Color) {
	c.gc.SetStrokeColor(col)
}

// Set the fill color
func (c *Canvas) SetFillColor(col color.Color) {
	c.gc.SetFillColor(col)
}

// Set the line width
func (c *Canvas) SetLineWidth(w float64) {
	c.gc.SetLineWidth(w)
}

// Actually draw the lines you've set up with LineTo
func (c *Canvas) Stroke() {
	c.gc.Stroke()
}

// Fill the area inside the lines you've set up with LineTo
func (c *Canvas) FillStroke() {
	c.gc.FillStroke()
}

// Fill the area inside the lines you've set up with LineTo, but don't
// draw the lines
func (c *Canvas) Fill() {
	c.gc.Fill()
}

// Fill the whole canvas with the fill color
func (c *Canvas) Clear() {
	c.gc.Clear()
}

// Fill the given rectangle with the fill color
func (c *Canvas) ClearRect(x1, y1, x2, y2 int) {
	c.gc.ClearRect(x1, y1, x2, y2)
}

// Draws an empty circle
// Fill the given circle with the fill color
// Stroke() each time to avoid connected circles
func (c *Canvas) Circle(cx, cy, r float64) {
	c.gc.ArcTo(cx, cy, r, r, 0, -math.Pi*2)
	c.gc.Close()
}

// Draws an empty ellipse
// Fill the given ellipse with the fill color
// Stroke() each time to avoid connected ellipses
func (c *Canvas) Ellipse(cx, cy, rx, ry float64) {
	c.gc.ArcTo(cx, cy, rx, ry, 0, -math.Pi*2)
	c.gc.Close()
}

// Save the current canvas to a PNG file
func (c *Canvas) SaveToPNG(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, c.img)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filename)
}

// Return the width of the canvas
func (c *Canvas) Width() int {
	return c.width
}

// Return the height of the canvas
func (c *Canvas) Height() int {
	return c.height
}
