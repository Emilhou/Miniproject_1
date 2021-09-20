package main

import "fmt"

type Fork struct {
	input, queryIN, queryOUT      chan string
	output     chan bool
	isOccupied bool
	numUsed    int
	name       int
}

func CreateFork(name int, queryIN, queryOUT chan string) *Fork {
	f := Fork{name: name}
	f.input = make(chan string)
	f.output = make(chan bool)
	f.queryIN = queryIN
	f.queryOUT = queryOUT
	f.isOccupied = false
	f.numUsed = 0
	return &f
}

func Work(f *Fork) {
	go StartRecievingQueriesFork(f)
	for {
		select {
		case <-f.input:
			f.numUsed++
			f.isOccupied = true;
			f.output <- f.isOccupied
		}
	}
}

func StartRecievingQueriesFork(f *Fork) {
	for {
		select {
		case <-f.queryIN:
			var output = fmt.Sprintf("Fork %d has been used %d times, and is currently", f.name, f.numUsed)
			if f.isOccupied {
				output += " occupied."
			} else {
				output += " free to use."
			}
			f.queryOUT <- output
		}
	}
}
