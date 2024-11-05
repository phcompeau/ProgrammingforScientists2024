package main

// ComputeHouseEdgeMultiproc takes an integer numTrials as well as an integer numProcs and returns an estimate of the house edge of craps (or whatever binary game) played over numTrials simulated games, distributed over numProcs processors.
func ComputeHouseEdgeMultiproc(numTrials, numProcs int) float64 {
	count := 0

	// make a channel to hold win totals
	c := make(chan int, numProcs)

	// play the game in parallel over numProcs processors
	for i := 0; i < numProcs; i++ {
		// how many times does each goroutine play the game?
		// standard behavior: numTrials/numProcs
		if i < numProcs-1 {
			go TotalWinOneProc(numTrials/numProcs, c)
		} else {
			//final processor should get remainder of the division as well
			go TotalWinOneProc(numTrials/numProcs+numTrials%numProcs, c)
		}

	}

	// grab all values from channel
	for i := 0; i < numProcs; i++ {
		count += <-c
	}

	return float64(count) / float64(numTrials)
}

// TotalWinOneProc
// Input: numTrials as an integer, and an integer channel
// Output: doesn't return anything, but it enters the total amount won in numTrials games of simulated craps into the channel
func TotalWinOneProc(numTrials int, c chan int) {
	count := 0 // keeps track of winnings (or more like losings)

	for i := 0; i < numTrials; i++ {
		// don't parallelize this! Use serial code that we already have
		outcome := PlayCrapsOnce()
		if outcome {
			count++ // won
		} else {
			count-- // lost
		}
	}

	c <- count
}
