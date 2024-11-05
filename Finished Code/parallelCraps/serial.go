package main

import (
	"math/rand"
)

func RollDie() int {
	return rand.Intn(6) + 1
}

// SumTwoDice takes no inputs and returns the sum of two simulated six-sided dice.
func SumTwoDice() int {
	return RollDie() + RollDie()
}

// math/rand has three built-in functions we will use a lot:
// 1. rand.Int: pseudorandom integer
// 2. rand.Float64: pseudorandom decimal in [0, 1)
// 3. rand.Intn: pseudorandom integer between 0 and n-1, inclusively

// PlayCrapsOnce takes no input parameters and returns true or false depending on outcome of a single simulated game of craps.
func PlayCrapsOnce() bool {
	firstRoll := SumTwoDice()
	if firstRoll == 7 || firstRoll == 11 {
		return true // winner!
	} else if firstRoll == 2 || firstRoll == 3 || firstRoll == 12 {
		return false // loser!
	} else { //roll again until we hit a 7 or our original roll
		for true {
			newRoll := SumTwoDice()
			if newRoll == firstRoll {
				// winner! :)
				return true
			} else if newRoll == 7 {
				//loser :(
				return false
			}
		}
	}
	// Go often likes default values at end of function
	panic("We shouldn't be here.")
	return false
}

// ComputeHouseEdge takes an integer numTrials and returns an estimate of the house edge of craps (or whatever binary game) played over numTrials simulated games.
func ComputeHouseEdge(numTrials int) float64 {
	// we use count to keep track of money won/lost
	count := 0
	for i := 0; i < numTrials; i++ {
		var outcome bool
		outcome = PlayCrapsOnce()
		// did we win or lose?
		if outcome == true { // win
			count++
		} else {
			// lost
			count--
		}
	}
	// we want to return the average won/lost
	return float64(count) / float64(numTrials)
}
