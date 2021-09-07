package main

type Philosopher struct {
	input      chan string
	output     chan string
	isEating   bool
	isThinking bool
	numEaten   int
	name       int
}

func CreatePhilosopher(name int) *Philosopher {
	p := Philosopher{name: name}
	p.input = make(chan string)
	p.output = make(chan string)
	p.isEating = false
	p.isThinking = true
	p.numEaten = 0
	return &p
}
