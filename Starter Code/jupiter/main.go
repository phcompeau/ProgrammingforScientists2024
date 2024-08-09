package main

import (
	"math"
)

//G is the gravitational constant in the gravitational force equation.  It is declared as a "global" constant that can be accessed by all functions.
const G = 6.67408e-11

func main() {
	// declaring objects
	var jupiter, io, europa, ganymede, callisto Body

	jupiter.name = "Jupiter"
	io.name = "Io"
	europa.name = "Europa"
	ganymede.name = "Ganymede"
	callisto.name = "Callisto"

	jupiter.red, jupiter.green, jupiter.blue = 223, 227, 202
	io.red, io.green, io.blue = 249, 249, 165
	europa.red, europa.green, europa.blue = 132, 83, 52
	ganymede.red, ganymede.green, ganymede.blue = 76, 0, 153
	callisto.red, callisto.green, callisto.blue = 0, 153, 76

	jupiter.mass = 1.898 * math.Pow(10, 27)
	io.mass = 8.9319 * math.Pow(10, 22)
	europa.mass = 4.7998 * math.Pow(10, 22)
	ganymede.mass = 1.4819 * math.Pow(10, 23)
	callisto.mass = 1.0759 * math.Pow(10, 23)

	jupiter.radius = 71000000
	io.radius = 1821000
	europa.radius = 1569000
	ganymede.radius = 2631000
	callisto.radius = 2410000

	jupiter.position.x, jupiter.position.y = 2000000000, 2000000000
	io.position.x, io.position.y = 2000000000-421600000, 2000000000
	europa.position.x, europa.position.y = 2000000000, 2000000000+670900000
	ganymede.position.x, ganymede.position.y = 2000000000+1070400000, 2000000000
	callisto.position.x, callisto.position.y = 2000000000, 2000000000-1882700000

	jupiter.velocity.x, jupiter.velocity.y = 0, 0
	io.velocity.x, io.velocity.y = 0, -17320
	europa.velocity.x, europa.velocity.y = -13740, 0
	ganymede.velocity.x, ganymede.velocity.y = 0, 10870
	callisto.velocity.x, callisto.velocity.y = 8200, 0

	// declaring universe and setting its fields.
	var jupiterSystem Universe
	jupiterSystem.width = 4000000000
	jupiterSystem.bodies = []Body{jupiter, io, europa, ganymede, callisto}

	// now we need to implement the system
}
