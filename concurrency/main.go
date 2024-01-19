package main

import (
	"fmt"
	"sync"
	"time"
)

// Declare a WaitGroup to keep track of tasks
var friends sync.WaitGroup

func main() {
	startTime := time.Now()
	// Add three tasks to the WaitGroup
	friends.Add(3)

	// Friends start their tasks (goroutines)
	go paint(&friends)
	go build(&friends)
	go cook(&friends)

	// Wait for all friends to finish their tasks
	friends.Wait()

	// Celebrate together!
	fmt.Println("Hooray! All friends finished their tasks! ", (time.Now().Second() - startTime.Second()))

}

// Friend 1: Painting
func paint(friends *sync.WaitGroup) {
	defer friends.Done() // Decrement the WaitGroup counter when done

	fmt.Println("Friend 1: Painting started")
	time.Sleep(1 * time.Second) // Simulate painting
	fmt.Println("Friend 1: Painting done!")
}

// Friend 2: Building
func build(friends *sync.WaitGroup) {
	defer friends.Done() // Decrement the WaitGroup counter when done

	fmt.Println("Friend 2: Building started")
	time.Sleep(2 * time.Second) // Simulate building
	fmt.Println("Friend 2: Building done!")
}

// Friend 3: Cooking
func cook(friends *sync.WaitGroup) {
	defer friends.Done() // Decrement the WaitGroup counter when done

	fmt.Println("Friend 3: Cooking started")
	time.Sleep(3 * time.Second) // Simulate cooking
	fmt.Println("Friend 3: Cooking done!")
}
