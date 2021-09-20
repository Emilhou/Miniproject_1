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
		left := <-p.leftFork.output
		right := <-p.rightFork.output
		if(!left && !right){
			p.isEating = true;
			p.leftFork.isOccupied = true
			p.rightFork.isOccupied = true
			p.leftFork.numUsed++
			p.rightFork.numUsed++
		}else {
			mutex.Unlock()
			continue
		}

		mutex.Unlock()

		SleepRandomSeconds()
		
		if(p.isEating){
			mutex.Lock()
			p.numEaten++
			p.isEating = false
			p.leftFork.isOccupied = false
			p.rightFork.isOccupied = false
			mutex.Unlock()
		}

		SleepRandomSeconds()
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
