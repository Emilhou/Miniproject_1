package main

import (
	"fmt"
)

type Philosopher struct {
	queryIN             chan string
	queryOUT            chan string
	leftFork, rightFork *Fork
	isEating            bool
	numEaten            int
	name                int
}

func CreatePhilosopher(name int, leftFork, rightFork *Fork, queryIN, queryOUT chan string) *Philosopher {
	p := Philosopher{name: name}
	p.queryIN = queryIN
	p.queryOUT = queryOUT
	p.leftFork = leftFork
	p.rightFork = rightFork
	p.isEating = false
	p.numEaten = 0
	return &p
}

func Dine(p *Philosopher) {
	go StartRecievingQueriesPhilosopher(p)
	for {
		mutex.Lock()

		p.leftFork.input <- "This fork is in use now"
		p.rightFork.input <- "This fork is in use now"
		<-p.leftFork.output
		<-p.rightFork.output

		p.isEating = true
		p.numEaten++

		SleepRandomSeconds()

		p.isEating = false
		p.leftFork.isOccupied = false
		p.rightFork.isOccupied = false

		mutex.Unlock()
	}
}

func StartRecievingQueriesPhilosopher(p *Philosopher) {
	for {
		select {
		case <-p.queryIN:
			var output = fmt.Sprintf("Philosopher %d has eaten %d times, and is currently", p.name, p.numEaten)
			if p.isEating {
				output += " eating."
			} else {
				output += " thinking."
			}
			p.queryOUT <- output
		}
	}
}
