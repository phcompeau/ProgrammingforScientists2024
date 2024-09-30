package main

import "fmt"

func ChangeFirst(list []int) {
	// list is a copy of b
	// when we access list[0], it's identical to b[0]
	list[0] = 1
	//list dies :(
}

// Delete
// Input: a slice of integers a and an index
// Output: the updated slice after removing a[index]
func Delete(a []int, index int) {
	a = append(a[:index], a[index+1:]...)
}

func main() {
	fmt.Println("More pointers ha")

	//SliceHell()
	MoreSliceHell()
}

func MoreSliceHell() {
	c := make([]int, 5)
	for i := range c {
		c[i] = 2*i + 1
	}

	fmt.Println("c is", c)

	Delete(c, 2) // this will get rid of middle element of c

	// c should now be [1 3 7 9]

	fmt.Println(c)
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
