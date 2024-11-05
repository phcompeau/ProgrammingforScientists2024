package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Craps!")
	numTrials := 10000000
	numProcs := runtime.NumCPU()

	start := time.Now()
	ComputeHouseEdge(numTrials)
	elapsed := time.Since(start)
	fmt.Printf("Running serially took %s", elapsed)
	fmt.Println()

	start2 := time.Now()
	ComputeHouseEdgeMultiproc(numTrials, numProcs)
	elapsed2 := time.Since(start2)
	fmt.Printf("Running in parallel took %s", elapsed2)

}
