package main

//this is where we will put functions that correspond only to the parallel simulation.

// DiffuseParallel is a Board method that takes as input an integer numProcs.
// It updates the board by diffusing each particle one time step, dividing the work over numProcs workers.
func (b *Board) DiffuseParallel(numProcs int) {
	numParticles := len(b.particles)

	finished := make(chan bool, numProcs)

	// split the work over numProcs processors
	for i := 0; i < numProcs; i++ {
		// each processor needs about the same number of particles

		//figure out which start and end indices of a this worker is getting assigned
		chunkSize := numParticles / numProcs
		startIndex := i * chunkSize
		endIndex := startIndex + chunkSize

		if i == numProcs-1 {
			// if we're in the final subslice, we need to make sure to extend it all the way to the end of a
			endIndex = numParticles
		}

		go DiffuseOneProc(b.particles[startIndex:endIndex], finished)
	}
	// we can't end the function until all goroutines are finished

	// now we need a separate for loop to grab from channel
	for i := 0; i < numProcs; i++ {
		<-finished
	}

	// now the function can end because receiving will block until I've received numProcs times
}

func DiffuseOneProc(particles []*Particle, finished chan bool) {
	// all we have to do is range over the particles and take a random step with each one
	for _, p := range particles {
		p.RandStep()
	}

	// hey, I'm done
	finished <- true
}
