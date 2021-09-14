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
	f.outputRight = make(chan string)
	f.isOccupied = false
	f.numUsed = 0
	fmt.Println(name + "is created")
	return &f
}

func work(f *Fork) {
	for {
		f.outputRight <- f.name + " left pick me up!"
		<-f.inputRight
		f.outputLeft <- f.name + " right pick me up!"
		<-f.inputLeft
	}

}
