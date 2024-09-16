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
		oldAcceleration, oldVelocity := b.acceleration, b.velocity // OK :)
		newUniverse.bodies[i].acceleration = UpdateAcceleration(currentUniverse, b)
		newUniverse.bodies[i].velocity = UpdateVelocity(b, oldAcceleration, time)
		newUniverse.bodies[i].position = UpdatePosition(b, oldAcceleration, oldVelocity, time)
	}

	return newUniverse
}

// UpdateVelocity
// Input: Body object b, previous acceleration and velocity as OrderedPair objects, and a float time.
// Output: Updated velocity vector of b as an OrderedPair according to physics nerd velocity update equations.
func UpdateVelocity(b Body, oldAcceleration OrderedPair, time float64) OrderedPair {
	var currentVelocity OrderedPair // starts at (0, 0)

	// apply equations. Note that b's acceleration has already been updated :)
	currentVelocity.x = b.velocity.x + 0.5*(b.acceleration.x+oldAcceleration.x)*time
	currentVelocity.y = b.velocity.y + 0.5*(b.acceleration.y+oldAcceleration.y)*time

	return currentVelocity
}

// UpdatePosition
// Input: Body object b, previous acceleration and velocity as OrderedPair objects, and a float time.
// Output: Updated position of b as an OrderedPair according to physics nerd update equations.
func UpdatePosition(b Body, oldAcceleration, oldVelocity OrderedPair, time float64) OrderedPair {
	var pos OrderedPair

	pos.x = b.position.x + oldVelocity.x*time + 0.5*oldAcceleration.x*time*time
	pos.y = b.position.y + oldVelocity.y*time + 0.5*oldAcceleration.y*time*time

	return pos
}

// UpdateAcceleration
// Input: currentUniverse Universe object and a Body object b.
// Output: The acceleration of b in the next generation after computing the net force of gravity acting on b over all bodies in currentUniverse.
func UpdateAcceleration(currentUniverse Universe, b Body) OrderedPair {
	var accel OrderedPair

	// get the net force vector (Ordered Pair)
	force := ComputeNetForce(currentUniverse, b)

	// thank u Newton for telling us F = m * a or a = F/m
	accel.x = force.x / b.mass
	accel.y = force.y / b.mass

	return accel
}

// ComputeNetForce
// Input: currentUniverse Universe object and a Body object b.
// Output: The net force of gravity acting on b over all bodies in currentUniverse, as an OrderedPair.
func ComputeNetForce(currentUniverse Universe, b Body) OrderedPair {
	var netForce OrderedPair //starts at (0, 0)

	for i := range currentUniverse.bodies {
		// I only want to compute force if currentbody is not b
		if currentUniverse.bodies[i] != b { // this is OK :)
			G := currentUniverse.gravitationalConstant
			currentForce := ComputeForce(b, currentUniverse.bodies[i], G)

			// add the contribution of b to the net force, componentwise
			netForce.x += currentForce.x
			netForce.y += currentForce.y
		}
	}

	return netForce
}

// ComputeForce
// Input: Two Body objects b and b2, along with a float G.
// Output: OrderedPair object representing the force of gravity acting on b according to b2, using G as the gravitational constant.
func ComputeForce(b, b2 Body, G float64) OrderedPair {
	var force OrderedPair

	d := Distance(b.position, b2.position)

	if d == 0.0 { // objects are the same or somehow occupy identical position, set force = (0, 0)
		return force
	}

	F := G * b.mass * b2.mass / (d * d)

	deltaX := b2.position.x - b.position.x
	deltaY := b2.position.y - b.position.y

	force.x = F * deltaX / d
	force.y = F * deltaY / d

	return force
}

// CopyUniverse
// Input: Universe object currentUniverse
// Output: Another Universe object newUniverse resulting from copying over all the fields in currentUniverse.
func CopyUniverse(currentUniverse Universe) Universe {
	var newUniverse Universe

	newUniverse.gravitationalConstant = currentUniverse.gravitationalConstant
	newUniverse.width = currentUniverse.width
	// newUniverse.bodies = currentUniverse.bodies // never do this with slices and think the whole slice will get copied over
	// instead, make a new slice of Body objects and copy them over
	numBodies := len(currentUniverse.bodies)
	newUniverse.bodies = make([]Body, numBodies)

	// then we want to range over all these bodies and copy their fields over to the new body
	for i := range newUniverse.bodies {
		//newUniverse.bodies[i] = currentUniverse.bodies[i] //BAD
		newUniverse.bodies[i] = CopyBody(currentUniverse.bodies[i])
	}

	return newUniverse
}

// CopyBody
// Input: a Body object b
// Output: a new Body object b2 whose fields are identical to b
func CopyBody(b Body) Body {
	var b2 Body

	// now for tedious copying
	b2.name = b.name
	b2.mass = b.mass
	b2.radius = b.radius
	b2.red = b.red
	b2.green = b.green
	b2.blue = b.blue

	//copy over ordered pairs too
	b2.position.x = b.position.x
	b2.position.y = b.position.y
	b2.velocity.x = b.velocity.x
	b2.velocity.y = b.velocity.y
	b2.acceleration.x = b.acceleration.x
	b2.acceleration.y = b.acceleration.y

	return b2
}

// Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
