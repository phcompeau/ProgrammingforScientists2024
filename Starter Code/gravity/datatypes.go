package main

//data setup.

type Body struct {
	name                             string
	mass, radius                     float64
	position, velocity, acceleration OrderedPair
	red, green, blue                 uint8 // values between 0 and 255, inclusively
}

type OrderedPair struct {
	x, y float64
}

type Universe struct {
	bodies                []Body
	width                 float64 // we will draw the universe as square
	gravitationalConstant float64 // represents the gravitational constant in this system
}
