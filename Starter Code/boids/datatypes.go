package main

// OrderedPair contains two float64 fields corresponding to
// the x and y coordinates of a point or vector in two-dimensional space.
type OrderedPair struct {
	x, y float64
}

// Boid represents our "bird" object. It contains two
// OrderedPair fields corresponding to its position, velocity, and acceleration.
type Boid struct {
	position, velocity, acceleration OrderedPair
}

// Sky represents a single time point of the simulation.
// It contains a width parameter indicating the boundary of the sky, as well as a slice of Boid objects.
// It also contains the system parameters.
type Sky struct {
	width                                             float64
	boids                                             []Boid
	proximity                                         float64 // used to determine if boids are close enough for forces to apply
	separationFactor, alignmentFactor, cohesionFactor float64 //multiply by each respective force
	maxBoidSpeed                                      float64 //fastest speed that a boid can fly
}
