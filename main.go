package main

import (
	"fmt"
	"time"
)

func main() {
	createTable()
	for {
		continue
	}
}

func createTable() {
	fmt.Println("this is the start")
	forkZero := CreateFork("0")
	forkOne := CreateFork("1")
	forkTwo := CreateFork("2")
	forkThree := CreateFork("3")
	forkFour := CreateFork("4")

	demokritus := CreatePhilosopher("Demokritus", forkZero.outputRight, forkZero.inputRight, forkFour.outputLeft, forkFour.inputLeft)
	sokrates := CreatePhilosopher("Sokrates", forkOne.outputRight, forkOne.inputRight, forkZero.outputLeft, forkZero.inputLeft)
	platon := CreatePhilosopher("Platon", forkTwo.outputRight, forkTwo.inputRight, forkOne.outputLeft, forkOne.inputLeft)
	pythagoras := CreatePhilosopher("Pythagoras", forkThree.outputRight, forkThree.inputRight, forkTwo.outputLeft, forkTwo.inputLeft)
	aristoteles := CreatePhilosopher("Aristoteles", forkThree.outputLeft, forkThree.inputLeft, forkFour.outputRight, forkFour.inputRight) // Left and right is swapped for this philosopher

	go work(forkZero)
	//time.Sleep(1 * time.Second)
	go work(forkOne)
	//time.Sleep(1 * time.Second)
	go work(forkTwo)
	//time.Sleep(1 * time.Second)
	go work(forkThree)
	//time.Sleep(1 * time.Second)
	go work(forkFour)
	//time.Sleep(1 * time.Second)
	go dine(demokritus)
	//time.Sleep(1 * time.Second)
	go dine(sokrates)
	//time.Sleep(1 * time.Second)
	go dine(platon)
	//time.Sleep(1 * time.Second)
	go dine(pythagoras)
	//time.Sleep(1 * time.Second)
	go dine(aristoteles)
	//time.Sleep(1 * time.Second)

	fmt.Println("this is the end")
}

func sleepSeconds(i time.Duration) {
	time.Sleep(i * time.Second)
}
