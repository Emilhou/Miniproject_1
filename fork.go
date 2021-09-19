package main

import "fmt"

type Fork struct {
	inputLeft, outputLeft, inputRight, outputRight chan string
	queryIN, queryOUT                              chan string
	isOccupied                                     bool
	numUsed                                        int
	name                                           string
}

func CreateFork(name string, queryIN, queryOUT chan string) *Fork {
	f := Fork{name: name}
	f.inputLeft = make(chan string)
	f.outputLeft = make(chan string)
	f.inputRight = make(chan string)
	f.outputRight = make(chan string)
	f.queryIN = queryIN
	f.queryOUT = queryOUT
	f.isOccupied = false
	f.numUsed = 0
	fmt.Println(name + " is created")
	return &f
}

func work(f *Fork) {
	go startQueryFork(f)
	for {
		f.outputRight <- f.name + " left pick me up!"
		<-f.inputRight
		f.outputLeft <- f.name + " right pick me up!"
		<-f.inputLeft
		f.numUsed++
	}

}

func startQueryFork(f *Fork) {
	for {
		select {
		case <-f.queryIN:
			f.queryOUT <- fmt.Sprintf("Fork %s has been used %d times", f.name, f.numUsed)
		}
	}
}
