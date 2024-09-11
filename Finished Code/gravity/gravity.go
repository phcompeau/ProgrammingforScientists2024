package main

import (
	"math"
)

//let's place our gravity simulation functions here.

// SimulateGravity
// Input: an initial Universe object, a number of generations, and a float time.
// Output: a slice of numGens + 1 Universes resulting from simulating gravity over numGens generations, where the time interval between generations is specified by time.
func SimulateGravity(initialUniverse Universe, numGens int, time float64) []Universe {
	timePoints := make([]Universe, numGens+1)

	timePoints[0] = initialUniverse
	// range from 1 to numGens, and call UpdateUniverse to set timePoints[i] based on updating previous generation
	for i := 1; i < numGens+1; i++ {
		timePoints[i] = UpdateUniverse(timePoints[i-1], time)
	}

	return timePoints
}

// UpdateUniverse
// Input: a Universe object and a float time.
// Output: a Universe object resulting from a single step according to the gravity simulation, using a time interval specified by time.
func UpdateUniverse(currentUniverse Universe, time float64) Universe {
	newUniverse := CopyUniverse(currentUniverse)

	// range and update every body in universe
	for i, b := range newUniverse.bodies {
		oldAcceleration, oldVelocity := b.acceleration, b.velocity
		newUniverse.bodies[i].acceleration = UpdateAcceleration(currentUniverse, b)
		newUniverse.bodies[i].velocity = UpdateVelocity(b, oldAcceleration, oldVelocity, time)
		newUniverse.bodies[i].position = UpdatePosition(b, oldAcceleration, oldVelocity, time)
	}

	return newUniverse
}

// Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
