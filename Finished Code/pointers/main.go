package main

import "fmt"

type Node struct {
	age  float64
	name string
}

type Tree struct {
	nodes []Node
	label string
}

func Equality() {
	var v1, v2 Node
	v1.age = 24
	v2.age = 24
	v1.name = "Bro"
	v2.name = "Bro"

	if v1 == v2 {
		fmt.Println("Same")
	} else {
		fmt.Println("Not same")
	}

	/*
		// Go doesn't even allow this
		a := []int{1, 5, 2, 1, 3}
		b := []int{1, 5, 2, 1, 3}
		if a == b {
			fmt.Println("Same!")
		} else {
			fmt.Println("Not same.")
		}
	*/
}

func ShallowAndDeepCopy() {
	var t Tree
	var v1, v2 Node
	t.nodes = []Node{v1, v2}
	t.label = "Hi my name is t"

	fmt.Println("t is", t)

	s := t // the key to understanding all this is this line

	// this creates a "shallow copy" of t
	// basic fields (integers, strings, floats) get copied over
	// slices/maps are pointers and so the POINTERS get copied but
	// (this is a good thing)
	// Go doesn't go in and duplicate all the data in the underlying array

	s.label = "My name is s"

	//if we don't like shallow copy, just create an array
	s.nodes = make([]Node, 2)
	// let's set the nodes of s too
	s.nodes[0].name = "Fred"
	s.nodes[1].name = "Mercury"
	s.nodes[0].age = 83.3
	s.nodes[1].age = 42347239047.47389

	fmt.Println("s is", s)
	fmt.Println("t is", t)

	// this is why I wrote CopyUniverse() in gravity functions
}

func Bro() {
	var v1, v2 Node
	v1.name = "Hi"
	v1.age = 68.2

	v2 = v1 // fields of v1 copied into v2
	//fmt.Println(v2.name, v2.age)

	v2.name = "Yo"
	v2.age = 13.7

	fmt.Println(v1.name, v1.age)

}

func ChangeFirst(list []int) {
	// list is a copy of b
	// when we access list[0], it's identical to b[0]
	list[0] = 1
	//list dies :(
}

func main() {
	fmt.Println("More pointers ha")

	//SliceHell()
	//MoreSliceHell()
	//Bro()
	//ShallowAndDeepCopy()
	Equality()
}

// Delete
// Input: a slice of integers a and an index
// Output: the updated slice after removing a[index]
func Delete(a []int, index int) []int {
	// copy of c was created here called a
	// it points to same location in array and has length 5
	a = append(a[:index], a[index+1:]...)
	fmt.Println("a is", a)

	//a dies UNLESS we return it
	return a
}

func MoreSliceHell() {
	c := make([]int, 5)
	// c is two things: a pointer to an array position, and a length

	for i := range c {
		c[i] = 2*i + 1
	}

	fmt.Println("c is", c)

	c = Delete(c, 2) // this will get rid of middle element of c

	// c should now be [1 3 7 9]

	fmt.Println("c is now", c)
}

func SliceHell() {
	b := make([]int, 5) // bunch of default zeroes

	// some array behind the scenes gets created

	// b is a pointer to this array

	ChangeFirst(b)

	b[4] = 3

	fmt.Println(b)

	// a slice is a pointer to an array!

	var c []int

	// what value does a have? nil

	c = b

	// pass the address of b to a

	c[2] = -489032784723894

	fmt.Println(c)

	fmt.Println(b)

	a := make([]int, 10, 20)

	for i := range a {
		a[i] = -i - 1
	}

	fmt.Println("a has length", len(a), "and is", a)

	q := a[8:15]

	fmt.Println("q is", q)

	fmt.Println(q[6])

	q[0] = 237489723894

	fmt.Println("a is now", a)
}
