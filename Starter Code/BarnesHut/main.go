package main

import (
	"fmt"
	"gifhelper"
)

func main() {

	// the following sample parameters may be helpful for the "collide" command
	// all units are in SI (meters, kg, etc.)
	// but feel free to change the positions of the galaxies.

	g0 := InitializeGalaxy(500, 4e21, 7e22, 2e22)
	g1 := InitializeGalaxy(500, 4e21, 3e22, 7e22)

	// you probably want to apply a "push" function at this point to these galaxies to move
	// them toward each other to collide.
	// be careful: if you push them too fast, they'll just fly through each other.
	// too slow and the black holes at the center collide and hilarity ensues.

	width := 1.0e23
	galaxies := []Galaxy{g0, g1}

	initialUniverse := InitializeUniverse(galaxies, width)

	// now evolve the universe: feel free to adjust the following parameters.
	numGens := 100000
	time := 2e14
	theta := 0.5

	timePoints := BarnesHut(initialUniverse, numGens, time, theta)

	fmt.Println("Simulation run. Now drawing images.")
	canvasWidth := 1000
	frequency := 1000
	scalingFactor := 1e11 // a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse
	imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

	fmt.Println("Images drawn. Now generating GIF.")
	gifhelper.ImagesToGIF(imageList, "galaxy")
	fmt.Println("GIF drawn.")
}
