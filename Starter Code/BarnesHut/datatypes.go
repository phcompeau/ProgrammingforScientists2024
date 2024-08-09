package main

const G = 6.67408e-11 // gravitational constant -- don't change this!

const solarMass = 1.989e30 // mass of sun -- don't change this!

const blackHoleMass = 8e36 // mass of black hole -- don't change!

// Universe contains a slice of pointers to stars and a width parameter.
// We conceptualize the universe as a square -- stars may go outside the universe
// but the width dictates relative distances when drawing the universe.
type Universe struct {
	stars []*Star
	width float64
}

// Galaxy is a potentially useful object holding a list of star positions
type Galaxy []*Star

// Star is analogous to the "Body" object from the jupiter simulations.
type Star struct {
	position, velocity, acceleration OrderedPair
	mass                             float64
	radius                           float64
	red, blue, green                 uint8
}

// OrderedPair represents a point or vector.
type OrderedPair struct {
	x float64
	y float64
}

// QuadTree simply contains a pointer to the root.
// Another way of doing this would be type QuadTree *Node
type QuadTree struct {
	root *Node
}

// Node object contains a slice of children (this could just as easily be an array of length 4).
// A node refers to a star. Sometimes, the star will be a "dummy" star, sometimes it is a star in the
// universe, and sometimes it is nil. Every internal node points to a dummy star.
type Node struct {
	children []*Node
	star     *Star
	sector   Quadrant
}

// Quadrant is an object representing a sub-square within a larger universe.
type Quadrant struct {
	x     float64 //bottom left corner x coordinate
	y     float64 //bottom left corner y coordinate
	width float64
}
