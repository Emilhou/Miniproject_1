package main

type Fork struct {
	input      chan string
	output     chan string
	isOccupied bool
	isFree     bool
	numUsed    int
	name       int
}

func CreateFork(name int) *Fork {
	f := Fork{name: name}
	f.input = make(chan string)
	f.output = make(chan string)
	f.isOccupied = false
	f.isFree = true
	f.numUsed = 0
	return &f
}
