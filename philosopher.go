package main

import (
	"fmt"
)

type Philosopher struct {
	inputLeft, outputLeft, inputRight, outputRight chan string
	queryIN, queryOUT                              chan string
	isEating                                       bool
	hasRightFork                                   bool
	hasLeftFork                                    bool
	numEaten                                       int
	name                                           string
}

func CreatePhilosopher(name string, forkLeft, forkLeftIn, forkRight, forkRightIn, queryIN, queryOUT chan string) *Philosopher {
	p := Philosopher{name: name}
	p.inputLeft = forkLeft
	p.outputLeft = forkLeftIn
	p.inputRight = forkRight
	p.outputRight = forkRightIn
	p.queryIN = queryIN
	p.queryOUT = queryOUT
	p.isEating = false
	p.hasLeftFork = false
	p.hasRightFork = false
	p.numEaten = 0
	fmt.Println(name + " is created")
	return &p
}

func dine(p *Philosopher) {
	go startQueryPhilosopher(p)
	for {
		<-p.inputLeft
		p.outputLeft <- "I picked up my left fork!"
		<-p.inputRight
		p.outputRight <- "I picked up my right fork!"
		p.isEating = true
		//fmt.Println(p.name + " is eating")
		p.numEaten++
		sleepSeconds()

		p.isEating = false
		//fmt.Println(p.name + " is done eating and has eaten " + strconv.Itoa(p.numEaten))

	}
}

func startQueryPhilosopher(p *Philosopher) {
	for {
		select {
		case <-p.queryIN:
			p.queryOUT <- fmt.Sprintf("%s has eaten %d times", p.name, p.numEaten)
		}

	}
}
