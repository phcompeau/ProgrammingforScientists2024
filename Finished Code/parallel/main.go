package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	//IntroToConcurrency()
	//GoRoutinesCantReturnVals()

	// channels will help us communicate across goroutines
	c := make(chan string)

	// this channel is synchronous, which means that the send to the channel and the receive from the channel are coordinated to occur at the same moment in time

	// c <- "Hello!" // this places a word into the channel

	go SayHi(c)

	// when you have a synchronous channel, adding to the channel *blocks*, which means that anything in this function beneath this line of code does not execute

	// to receive from the channel, use <- c
	fmt.Println(<-c)
	//receiving from a synchronous channel blocks too
}

func SayHi(c chan string) {
	c <- "Hello!"
	// this blocks, but what does it block? Rest of this function
	fmt.Println("I'm at the end of Say Hi().")
}

func GoRoutinesCantReturnVals() {
	fmt.Println("Let's compute 40 factorial by splitting it into two tasks.")

	/*
		n := 40

		product1 := go Perm(1, n/2+1)
		product2 := go Perm(n/2+1, n)
		fmt.Println(product1 * product2) // concern is that print happens before functions even start, let alone finish. Go won't allow this
	*/

	//Goroutines are barred from returning values

	// what we need is a method of communication across Goroutines for things like "here is a value" or "I am done"
}

func IntroToConcurrency() {
	fmt.Println("Parallelism and concurrency.")

	fmt.Println("This computer has", runtime.NumCPU(), "total cores available.")

	n := 100000000

	start := time.Now()
	Factorial(n)
	elapsed := time.Since(start)
	log.Printf("Using multiple processors took %s", elapsed)

	// By default, Go has access to all cores available and tries to use all of them.

	// I can change the number of processors that Go has access to.

	runtime.GOMAXPROCS(1) // this forces Go to only use one processor

	// let's do the same thing as before, but with Go only having access to one processor

	start2 := time.Now()
	Factorial(n)
	elapsed2 := time.Since(start2)
	log.Printf("Using one processor took %s", elapsed2)

	// adding the word "go" before a function starts this function as a separate ****concurrent**** process.
	go PrintFactorials(10)

	//time.Sleep(time.Second) // incorporates a little delay on func main() process to let function call catch up

	fmt.Println("I am still here.")

	fmt.Println("Program exiting normally.")

	// If I don't set num processors to 1, and I use goroutines, then I have not just concurrency but parallelism :)
}

// Perm computes n permuted with k from the prep materials
func Perm(k, n int) int {
	p := 1

	for i := k; i < n; i++ {
		p *= i
	}

	return p
}

func PrintFactorials(n int) {
	fmt.Println("Entering PrintFactorials function.")
	p := 1
	for i := 1; i <= n; i++ {
		fmt.Println(p)
		p *= i
	}
	fmt.Println("Exiting PrintFactorials function.")
}

func Factorial(n int) int {
	prod := 1
	for i := 2; i <= n; i++ {
		prod *= i
	}
	return prod
}
