package main

//this file contains functions shared by the serial and parallel versions of our code.

import (
	"math"
	"math/rand"
	"runtime"
)

// CopyBoard is a Board method that makes a deep copy of a board and returns
// a pointer to it.
func (b *Board) CopyBoard() *Board {
	var newBoard Board

	newBoard.width = b.width
	newBoard.height = b.height
	newBoard.particles = make([]*Particle, len(b.particles))

	for i, p := range b.particles {
		newBoard.particles[i] = p.CopyParticle()
	}

	return &newBoard
}

// CopyParticle is a Particle method that makes a deep copy of a Particle
// and returns a pointer to the new Particle.
func (p *Particle) CopyParticle() *Particle {
	var p2 Particle

	p2 = *p // shallow copy ok because all fields are elementary

	//but if you changed the data structure representing a Particle, this could be horrible!

	return &p2
}

// UpdateBoards takes a pointer to an initial Board object, a number of steps parameter, and a boolean flag isParallel.
// It returns a slice of pointers to Board objects corresponding to simulating diffusion
// over the number of steps given.
// It runs the algorithm serially if isParallel is false and in parallel if isParallel is true.
func UpdateBoards(initialBoard *Board, numSteps int, isParallel bool) []*Board {
	boards := make([]*Board, numSteps+1)
	boards[0] = initialBoard

	for i := 1; i <= numSteps; i++ {
		boards[i] = boards[i-1].UpdateBoard(isParallel)
	}

	return boards
}

// UpdateBoard is a Board method that returns a pointer to a new Board object
// corresponding to a single time step update of the Board.
// It takes a boolean input isParallel.
// It runs the algorithm serially if isParallel is false and in parallel if isParallel is true.
func (b *Board) UpdateBoard(isParallel bool) *Board {
	newBoard := b.CopyBoard()

	newBoard.Diffuse(isParallel)

	return newBoard
}

// Diffuse is a Board method that diffuses each Particle in the Board over a single time step.
// It takes a boolean input isParallel.
// It runs the algorithm serially if isParallel is false and in parallel if isParallel is true.
func (b *Board) Diffuse(isParallel bool) {
	// if isParallel is true, I want the diffusion process to happen in parallel, meaning that I'll need to adjust the existing code.
	if isParallel {
		// I have to do parallel work now
		numProcs := runtime.NumCPU()
		b.DiffuseParallel(numProcs)
	} else { // serial case
		for _, p := range b.particles {
			p.RandStep()
		}
	}
}

// RandStep is a Particle method that moves the Particle by the Particle's diffusion rate
// parameter in a randomly chosen direction.
func (p *Particle) RandStep() {
	stepLength := p.diffusionRate
	angle := rand.Float64() * 2 * math.Pi
	p.position.x += stepLength * math.Cos(angle)
	p.position.y += stepLength * math.Sin(angle)
}

// InitializeBoard takes board parameters and initializes a Board with these parameters
// for a collection of randomly placed particles in the Board.
func InitializeBoard(boardWidth, boardHeight float64, numParticles int, particleRadius float64, diffusionRate float64, random bool) *Board {
	var b Board

	b.width = boardWidth
	b.height = boardHeight

	b.particles = make([]*Particle, numParticles)

	for i := range b.particles {
		var p Particle
		if random {
			p.position.x = rand.Float64() * boardWidth
			p.position.y = rand.Float64() * boardHeight
		} else {
			// default: non-random: assign all to center of board
			p.position.x = boardWidth / 2
			p.position.y = boardHeight / 2
		}
		p.radius = particleRadius
		p.diffusionRate = diffusionRate
		p.red, p.green, p.blue = 255, 255, 255
		b.particles[i] = &p
	}

	return &b
}
