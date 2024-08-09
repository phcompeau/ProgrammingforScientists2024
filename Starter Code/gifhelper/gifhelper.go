//special thanks to John Cox, who optimized this code to run faster and generate smaller GIFs.

package gifhelper

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
)

//ImagesToGIF() takes a slice of images and uses them to generate an animated GIF
// with the name "filename.out.gif" where filename is an input parameter.
func ImagesToGIF(imglist []image.Image, filename string) {

	// get ready to write images to files
	w, err := os.Create(filename + ".out.gif")

	if err != nil {
		fmt.Println("Sorry: couldn't create the file!")
		os.Exit(1)
	}

	defer w.Close()
	var g gif.GIF
	g.Delay = make([]int, len(imglist))
	g.Image = make([]*image.Paletted, len(imglist))
	g.LoopCount = 10

	for i := range imglist {
		g.Image[i] = ImageToPaletted(imglist[i])
		g.Delay[i] = 1
	}

	gif.EncodeAll(w, &g)
}

// ImageToPaletted converts an image to an image.Paletted with 256 colors.
// It is used by a subroutine by process() to generate an animated GIF.
func ImageToPalettedVersion1(img image.Image) *image.Paletted {
	pm, ok := img.(*image.Paletted)
	if !ok {
		b := img.Bounds()
		pm = image.NewPaletted(b, palette.WebSafe)
		draw.Draw(pm, pm.Bounds(), img, image.Point{}, draw.Src)
	}
	return pm
}

var mapOfColorIndices map[color.Color]uint8

func init() {
	mapOfColorIndices = make(map[color.Color]uint8)
}

func ImageToPaletted(img image.Image) *image.Paletted {
        pm, ok := img.(*image.Paletted)
        if !ok {
                b := img.Bounds()
                pm = image.NewPaletted(b, palette.WebSafe)
                var prevC color.Color = nil
                var idx uint8
                var ok bool
                for y := b.Min.Y; y < b.Max.Y; y++ {
                        for x := b.Min.X; x < b.Max.X; x++ {
                                c := img.At(x, y)
                                if c != prevC {
                                        if idx, ok = mapOfColorIndices[c]; !ok {
                                                idx = uint8(pm.Palette.Index(c))
                                                mapOfColorIndices[c] = idx
                                        }
                                        prevC = c
                                }
                                i := pm.PixOffset(x, y)
                                pm.Pix[i] = idx
                        }
                }
        }
        return pm
}
