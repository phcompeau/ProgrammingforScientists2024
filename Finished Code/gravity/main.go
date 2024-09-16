package main

import (
	"fmt"
	"gifhelper"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Let's simulate gravity!")

	//let's take command line arguments (CLAs) from the user
	//CLAs get stored in an ARRAY of strings called os.Args
	//this array has length equal to number of arguments given by the user + 1

	//os.Args[0] is the name of the program (./gravity)

	if len(os.Args) != 6 {
		panic("Error: incorrect number of command line arguments.")
	}

	//let's take CLAs: initial universe file, numGens, time, canvas width (in pixels), drawing frequency
	// remember: os.Args[] is an array of strings

	inputFile := "data/" + os.Args[1] + ".txt"

	// let's ensure that the file exists

	initialUniverse, err := ReadUniverse(inputFile)

	Check(err)

	// I wil eventually write the simulation to a beautiful GIF
	outputFile := "output/" + os.Args[1]

	numGens, err2 := strconv.Atoi(os.Args[2])

	Check(err2)

	time, err3 := strconv.ParseFloat(os.Args[3], 64)
	Check(err3)

	canvasWidth, err4 := strconv.Atoi(os.Args[4])
	Check(err4)

	drawingFrequency, err5 := strconv.Atoi(os.Args[5])
	Check(err5)

	if drawingFrequency <= 0 {
		panic("Error: nonpositive number given as drawingFrequency.")
	}

	fmt.Println("Command line arguments read!")

	fmt.Println("Simulating gravity now.")

	timePoints := SimulateGravity(initialUniverse, numGens, time)

	fmt.Println("Simulation run!")

	fmt.Println("Drawing universes.")

	images := AnimateSystem(timePoints, canvasWidth, drawingFrequency)

	fmt.Println("Images drawn!")

	fmt.Println("Now, let's make an animated GIF.")

	gifhelper.ImagesToGIF(images, outputFile)

	fmt.Println("GIF drawn!")

	fmt.Println("Simulation complete.")
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
