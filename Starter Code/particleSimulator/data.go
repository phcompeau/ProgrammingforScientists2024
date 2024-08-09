package main

// Particle represents each individual particle in our diffusion simulation.
type Particle struct {
	position         OrderedPair
	name             string
	radius           float64
	diffusionRate    float64 // length of single step
	red, green, blue uint8   // color function of type?
}

// Board represents the visible part of the simulation.
type Board struct {
	width, height float64
	particles     []*Particle
}

// OrderedPair is an object that represents a point or vector in two-dimensional space.
type OrderedPair struct {
	x, y float64
}
