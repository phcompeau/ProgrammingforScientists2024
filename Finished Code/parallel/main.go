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
	//SyncChannels()
	//ParallelFactorial()
	//ParallelSumming()
	BufferedChannels()
}

func BufferedChannels() {
	n := 10

	// a buffered channel can store more than one thing
	// buffered channels are *asynchronous*, which means that
	// sending to the channel does not block
	c := make(chan int, n) // capacity of our channel

	for k := 0; k < n; k++ {
		go Push(k, c)
	}

	// buffered channels are like queues or lines: first in, first out

	for i := 0; i < n; i++ {
		fmt.Println(<-c)
		// receiving from a buffered channel takes the first thing in line out of the channel
		// is there anything in my mailbox? If so, grab it
		// If not, block until there is
	}
}

func Push(k int, c chan int) {
	time.Sleep(time.Second)
	c <- k
	// because c is asynchronous, there is no blocking on the sender side; it's like dropping something in a mailbox for someone to pick up later
}

func ParallelSumming() {
	// create some array a
	n := 10000000000
	a := make([]int, n)
	for i := range a {
		a[i] = 2*i + 1
	}

	numProcs := runtime.NumCPU() // number of workers

	// let's time serial and parallel approaches
	start := time.Now()
	SerialSum(a)
	elapsed := time.Since(start)
	log.Printf("Summing serially took %s", elapsed)

	start2 := time.Now()
	SumMultiProc(a, numProcs)
	elapsed2 := time.Since(start2)
	log.Printf("Summing in parallel took %s", elapsed2)

}

func SerialSum(a []int) int {
	s := 0

	for _, val := range a {
		s += val
	}

	return s
}

// better than hard coding tasks into a constant number of workers, let the # workers vary based on what we have access to

// SumMultiProc
// Input: slice of integers a, numProcs integer
// Output: sum of all elements in a, parallelized over numProcs workers
func SumMultiProc(a []int, numProcs int) int {
	n := len(a)
	s := 0              // stores the sum
	c := make(chan int) // used for exchange of information

	//idea: split the job into numProcs pieces, each of approx equal size, and let numProcs workers compute the sum of one subslice of a

	for i := 0; i < numProcs; i++ {
		//figure out which start and end indices of a this worker is getting assigned
		chunkSize := n / numProcs
		startIndex := i * chunkSize
		endIndex := startIndex + chunkSize

		if i == numProcs-1 {
			// if we're in the final subslice, we need to make sure to extend it all the way to the end of a
			endIndex = n
		}

		go SumOneProc(a[startIndex:endIndex], c)
	}

	// once these are done, grab values from the channel
	for i := 0; i < numProcs; i++ {
		s += <-c
		// this blocks too, which is what you want, otherwise you can't return s
	}

	return s
}

// SumOneProc
// Input: slice of integers a, and an integer channel c
// Output: nothing -- but puts sum of elements of a into c
func SumOneProc(a []int, c chan int) {
	s := 0
	for _, val := range a {
		s += val
	}
	c <- s // only blocks rest of this function

	// but there is nothing down here
}

func ParallelFactorial() {
	fmt.Println("Let's compute 40 factorial by splitting it into two tasks.")

	n := 40

	c := make(chan int)

	go Perm(1, n/2+1, c)
	go Perm(n/2+1, n, c)
	fmt.Println(<-c * <-c) // concern is that print happens before functions even start, let alone finish. Go won't allow this

	//Goroutines are barred from returning values

	// what we need is a method of communication across Goroutines for things like "here is a value" or "I am done"
}

// Perm computes n permuted with k from the prep materials
func Perm(k, n int, c chan int) {
	p := 1

	for i := k; i < n; i++ {
		p *= i
	}

	c <- p
}

func SyncChannels() {

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
