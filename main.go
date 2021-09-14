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

	sokrates := CreatePhilosopher("Sokrates", forkOne.outputRight, forkOne.inputRight, forkZero.outputLeft, forkZero.inputLeft)
	platon := CreatePhilosopher("Platon", forkTwo.outputRight, forkTwo.inputRight, forkOne.outputLeft, forkOne.inputLeft)
	pythagoras := CreatePhilosopher("Pythagoras", forkThree.outputRight, forkThree.inputRight, forkTwo.outputLeft, forkTwo.inputLeft)
	aristoteles := CreatePhilosopher("Aristoteles", forkFour.outputRight, forkFour.inputRight, forkThree.outputLeft, forkThree.inputLeft)
	demokritus := CreatePhilosopher("Demokritus", forkFour.outputLeft, forkFour.inputLeft, forkZero.outputRight, forkZero.inputRight) // Left and right is swapped for this philosopher

	go work(forkZero)
	time.Sleep(1 * time.Second)
	go work(forkOne)
	time.Sleep(1 * time.Second)
	go work(forkTwo)
	time.Sleep(1 * time.Second)
	go work(forkThree)
	time.Sleep(1 * time.Second)
	go work(forkFour)
	time.Sleep(1 * time.Second)
	go dine(sokrates)
	time.Sleep(1 * time.Second)
	go dine(platon)
	time.Sleep(1 * time.Second)
	go dine(pythagoras)
	time.Sleep(1 * time.Second)
	go dine(aristoteles)
	time.Sleep(1 * time.Second)
	go dine(demokritus)
	time.Sleep(1 * time.Second)

	fmt.Println("this is the end")
}
