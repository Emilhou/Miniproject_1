package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex
var queryPhiloIN = make(chan string, 5)
var queryPhiloOUT = make(chan string, 5)
var queryForkIN = make(chan string, 5)
var queryForkOUT = make(chan string, 5)

func main() {
	CreateTable()
	for {
		time.Sleep(5 * time.Second)
		Query()
	}
}

func CreateTable() {
	forkZero := CreateFork(0, queryForkIN, queryForkOUT)
	forkOne := CreateFork(1, queryForkIN, queryForkOUT)
	forkTwo := CreateFork(2, queryForkIN, queryForkOUT)
	forkThree := CreateFork(3, queryForkIN, queryForkOUT)
	forkFour := CreateFork(4, queryForkIN, queryForkOUT)

	philoZero := CreatePhilosopher(0, forkZero, forkFour, queryPhiloIN, queryPhiloOUT)
	philoOne := CreatePhilosopher(1, forkOne, forkZero, queryPhiloIN, queryPhiloOUT)
	philoTwo := CreatePhilosopher(2, forkTwo, forkOne, queryPhiloIN, queryPhiloOUT)
	philoThree := CreatePhilosopher(3, forkThree, forkTwo, queryPhiloIN, queryPhiloOUT)
	philoFour := CreatePhilosopher(4, forkThree, forkFour, queryPhiloIN, queryPhiloOUT)

	go Work(forkZero)
	go Work(forkOne)
	go Work(forkTwo)
	go Work(forkThree)
	go Work(forkFour)

	go Dine(philoZero)
	go Dine(philoOne)
	go Dine(philoTwo)
	go Dine(philoThree)
	go Dine(philoFour)

}

func SleepRandomMilliSeconds() {
	rand.Seed(time.Now().UnixNano())
	var n = rand.Intn(5)

	time.Sleep(time.Duration(n) * time.Millisecond)
}

func Query() {
	var QueriesToEachChannel = 5

	for i := 0; i < QueriesToEachChannel; i++ {
		queryForkIN <- "Answer me!"
		queryPhiloIN <- "Answer me!"
	}

	fmt.Println()
	fmt.Println("_________________________________________________________________")
	for i := 0; i < QueriesToEachChannel; i++ {
		x := <-queryPhiloOUT
		fmt.Println(x)
	}
	fmt.Println("-----------------------------------------------------------------")
	for i := 0; i < QueriesToEachChannel; i++ {
		x := <-queryForkOUT
		fmt.Println(x)
	}
	fmt.Println("_________________________________________________________________")
	fmt.Println()
}
