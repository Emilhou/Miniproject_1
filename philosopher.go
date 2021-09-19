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
	leftFork, rightFork                            *Fork
}

func CreatePhilosopher(name string, forkLeft, forkLeftIn, forkRight, forkRightIn, queryIN, queryOUT chan string, leftFork, rightFork *Fork) *Philosopher {
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
	p.leftFork = leftFork
	p.rightFork = rightFork
	p.numEaten = 0
	fmt.Println(name + " is created")
	return &p
}

func dine(p *Philosopher) {
	go startQueryPhilosopher(p)
	for {
		<-p.inputLeft
		p.outputLeft <- "I picked up my left fork!"
		//p.leftFork.mutex.Lock()
		<-p.inputRight
		p.outputRight <- "I picked up my right fork!"
		//p.rightFork.mutex.Lock()
		p.isEating = true
		//fmt.Println(p.name + " is eating")
		p.numEaten++
		sleepRandomSeconds()
		//p.leftFork.mutex.Unlock()
		//p.rightFork.mutex.Unlock()
		p.isEating = false
		//fmt.Println(p.name + " is done eating and has eaten " + strconv.Itoa(p.numEaten))

	}
}

func startQueryPhilosopher(p *Philosopher) {
	for {
		<-p.queryIN
		p.queryOUT <- fmt.Sprintf("%s has eaten %d times", p.name, p.numEaten)
	}
}
