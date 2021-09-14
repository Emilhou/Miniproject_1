package main

import "fmt"

type Fork struct {
	inputLeft   chan string
	outputLeft  chan string
	inputRight  chan string
	outputRight chan string
	isOccupied  bool
	numUsed     int
	name        string
}

func CreateFork(name string) *Fork {
	f := Fork{name: name}
	f.inputLeft = make(chan string)
	f.outputLeft = make(chan string)
	f.inputRight = make(chan string)
	f.inputRight = make(chan string)
	f.isOccupied = false
	f.numUsed = 0
	fmt.Println(name + "is created")
	return &f
}

func work(f *Fork) {
	fmt.Println(f.name + " work init")
	f.outputLeft <- f.name + " left pick me up!"
	<-f.inputLeft
	f.outputRight <- f.name + " right pick me up!"
	<-f.inputRight
	fmt.Println(f.name + " done")
}
