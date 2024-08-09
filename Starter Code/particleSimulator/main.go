package main

import (
	"fmt"
	"gifhelper"
)

func main() {
	fmt.Println("Particle simulator.")

	fmt.Println("Generating random particles and initializing board.")

	numParticles := 100
	boardWidth := 1000.0
	boardHeight := 1000.0
	particleRadius := 5.0
	diffusionRate := 1.0

	//assumption: all particles are white

	random := true // make true if we want to scatter across board

	initialBoard := InitializeBoard(boardWidth, boardHeight, numParticles, particleRadius, diffusionRate, random)

	fmt.Println("Running simulation in serial.")

	numSteps := 2000

	boards := UpdateBoards(initialBoard, numSteps)

	fmt.Println("Simulation run. Animating system.")
	canvasWidth := 300
	frequency := 10
	images := AnimateSystem(boards, canvasWidth, frequency)

	fmt.Println("Images drawn. Generating GIF.")

	outFileName := "diffusion"
	gifhelper.ImagesToGIF(images, outFileName)
}
