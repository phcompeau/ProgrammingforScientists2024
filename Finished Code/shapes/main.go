// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Rectangle struct {
	width, height float64
	x1, y1        float64
	rotation      float64
}

type Circle struct {
	x1, y1 float64
	radius float64
}

func (c *Circle) Scale(f float64) {
	c.radius *= f
}

func (r *Rectangle) Scale(f float64) {
	r.width *= f
	r.height *= f
}

func (c *Circle) Area() float64 {
	return 3.0 * c.radius * c.radius
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Translate(a, b float64) {
	// add a to x1 and b to y1
	r.x1 += a
	r.y1 += b
}

func (c *Circle) Translate(a, b float64) {
	// add a to x1 and b to y1
	c.x1 += a
	c.y1 += b
	fmt.Println("c's center is at", c.x1, c.y1)
}

func main() {
	var r Rectangle
	var myCirc Circle

	// set center
	myCirc.x1, myCirc.y1 = 1.0, 3.0

	myCirc.radius = 2.0
	r.width = 3.0
	r.height = 5.0

	fmt.Println("myCirc's Area is", myCirc.Area())
	fmt.Println("r's Area is", r.Area())

	// let's move myCirc
	//myCirc = myCirc.Translate(-2.1, 4.7)
	// new coordinates: -1.1, 7.7
	fmt.Println("New coordinates of myCirc:", myCirc.x1, myCirc.y1)

	var b int = -14

	// where is b stored?

	var a *int // a has type "pointer to an integer" (address of some integer)
	// a starts with the default value nil
	fmt.Println("b is", b)
	//fmt.Println(a)

	//I can point a at the location of b
	a = &b

	fmt.Println("a is", a)

	*a = 4 // changing b without even knowing its name
	// * here is called a "dereference" -- go in and unlock the door

	fmt.Println("b is now", b)

	var pointerToCirc *Circle // pointerToC = nil
	//well, let's point this thing to my circle

	pointerToCirc = &myCirc
	// I can change myCirc's fields bahahaha
	(*pointerToCirc).x1 = -1237891237890.090
	(*pointerToCirc).y1 = -32840387489237489273.8908

	fmt.Println("circle center:", myCirc.x1, myCirc.y1)

	//Go is nice because I don't need pointer dereferencing
	pointerToCirc.x1 = 0.0
	pointerToCirc.y1 = 0.0

	fmt.Println("circle center:", myCirc.x1, myCirc.y1)

	//let's move this thing one more time
	pointerToCirc.Translate(2.0, 4.0)
	fmt.Println("circle center:", myCirc.x1, myCirc.y1)

	myCirc.Translate(1.3, 4.5)
	fmt.Println("circle center:", myCirc.x1, myCirc.y1)

	myCirc.radius = 3.0
	myCirc.Scale(2.0)
	fmt.Println("new radius:", myCirc.radius)

	//final point: using "new"
	d := new(Circle) // creates a Circle behind the scenes, and returns a pointer to it called d
	d.x1 = 3.0
	d.y1 = -2.45
	d.radius = 5.0
	fmt.Println("Area of d is", d.Area())
	fmt.Println(d)
}
