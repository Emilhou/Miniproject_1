package main

import "fmt"

type Philosopher struct {
	inputLeft    chan string
	outputLeft   chan string
	inputRight   chan string
	outputRight  chan string
	isEating     bool
	hasRightFork bool
	hasLeftFork  bool
	numEaten     int
	name         string
}

func CreatePhilosopher(name string, forkLeft, forkLeftIn, forkRight, forkRightIn chan string) *Philosopher {
	p := Philosopher{name: name}
	p.inputLeft = forkLeft
	p.outputLeft = forkLeftIn
	p.inputRight = forkRight
	p.outputRight = forkRightIn
	p.isEating = false
	p.hasLeftFork = false
	p.hasRightFork = false
	p.numEaten = 0
	fmt.Println(name + "is created")
	return &p
}

func dine(p *Philosopher) {
	for {
		<-p.inputLeft
		p.outputLeft <- "I picked up my left fork!"
		<-p.inputRight
		p.outputRight <- "I picked up my right fork!"
		fmt.Println(p.name + "is eating")
		//layForks(p)

	}
}

func layForks(p *Philosopher) {
	p.outputLeft <- "I´m done with my left fork!"
	p.outputRight <- "I´m done with my right fork!"
}
