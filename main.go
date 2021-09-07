package main
import "fmt"

func main() {
	phil := CreatePhilosopher(1)
	fmt.Println(phil.isThinking)
	fmt.Println(phil.numEaten)
	fmt.Println(phil.name)
}
