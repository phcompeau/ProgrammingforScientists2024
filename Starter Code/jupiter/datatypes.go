package main

//data setup.

type Body struct {
	name                             string
	mass, radius                     float64
	position, velocity, acceleration OrderedPair
	red, green, blue                 uint8
}

type OrderedPair struct {
	x, y float64
}

type Universe struct {
	bodies []Body
	width  float64
}
