package main

import (
	"fmt"
	"os"
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

	//let's take CLAs: initial universe file, numGens, time, canvas width, drawing frequency

}
