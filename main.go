package main

import (
	"fmt"
	"math/rand"
	"time"
)

var queryForkIN chan string
var queryForkOUT chan string
var queryPhiloIN chan string
var queryPhiloOUT chan string

func main() {
	createTable()
	for {
		time.Sleep(5 * time.Second)
		query()
	}
}

func createTable() {
	queryForkIN = make(chan string, 5)
	queryForkOUT = make(chan string, 5)
	queryPhiloIN = make(chan string, 5)
	queryPhiloOUT = make(chan string, 5)

	//fmt.Println("this is the start")
	forkZero := CreateFork("0", queryForkIN, queryForkOUT)
	forkOne := CreateFork("1", queryForkIN, queryForkOUT)
	forkTwo := CreateFork("2", queryForkIN, queryForkOUT)
	forkThree := CreateFork("3", queryForkIN, queryForkOUT)
	forkFour := CreateFork("4", queryForkIN, queryForkOUT)

	demokritus := CreatePhilosopher("Demokritus", forkZero.outputRight, forkZero.inputRight, forkFour.outputLeft, forkFour.inputLeft, queryPhiloIN, queryPhiloOUT, forkZero, forkFour)
	sokrates := CreatePhilosopher("Sokrates", forkOne.outputRight, forkOne.inputRight, forkZero.outputLeft, forkZero.inputLeft, queryPhiloIN, queryPhiloOUT, forkOne, forkZero)
	platon := CreatePhilosopher("Platon", forkTwo.outputRight, forkTwo.inputRight, forkOne.outputLeft, forkOne.inputLeft, queryPhiloIN, queryPhiloOUT, forkTwo, forkOne)
	pythagoras := CreatePhilosopher("Pythagoras", forkThree.outputRight, forkThree.inputRight, forkTwo.outputLeft, forkTwo.inputLeft, queryPhiloIN, queryPhiloOUT, forkThree, forkTwo)
	aristoteles := CreatePhilosopher("Aristoteles", forkThree.outputLeft, forkThree.inputLeft, forkFour.outputRight, forkFour.inputRight, queryPhiloIN, queryPhiloOUT, forkThree, forkFour) // Left and right is swapped for this philosopher

	go work(forkZero)
	//time.Sleep(1 * time.Second)
	go work(forkOne)
	//time.Sleep(1 * time.Second)
	go work(forkTwo)
	//time.Sleep(1 * time.Second)
	go work(forkThree)
	//time.Sleep(1 * time.Second)
	go work(forkFour)
	time.Sleep(5 * time.Second)
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

	//fmt.Println("this is the end")
}

func sleepRandomSeconds() {
	time.Sleep(time.Duration(rand.Intn(10)))
}

func query() {
	var numQueries = 5
	fmt.Println()
	for i := 0; i < numQueries; i++ {
		queryForkIN <- "Answer me!"
		queryPhiloIN <- "Answer me!"
	}
	fmt.Println("****************************************************************************")
	for i := 0; i < numQueries; i++ {
		fmt.Println(<-queryPhiloOUT)
	}
	fmt.Println("----------------------------------------------------------------------------")
	for i := 0; i < numQueries; i++ {
		fmt.Println(<-queryForkOUT)
	}
	fmt.Println("****************************************************************************")
	fmt.Println()
}
