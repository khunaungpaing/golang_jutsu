package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	for {
		clearScreen()
		fmt.Println("Welcome to Number Guessing Game!")
		fmt.Println("1. Start Game\n2. Exit Game")
		fmt.Print("Enter your choice:")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			selectLevel()
		case 2:
			fmt.Println("Exiting Game...")
			clearScreen()
			os.Exit(0)
		}

	}

}

func selectLevel() {
	clearScreen()
	fmt.Println("Select level:")
	fmt.Println("1. Easy")
	fmt.Println("2. Medium")
	fmt.Println("3. Hard")
	fmt.Print("Enter your choice: ")
	var level int
	fmt.Scanln(&level)
	clearScreen()
	playGame(level)
}

func playGame(level int) {
	var max int
	switch level {
	case 1:
		max = 10
	case 2:
		max = 100
	case 3:
		max = 1000
	default:
		fmt.Println("Invalid choice...")
		selectLevel()
	}

	secretNumber := generateRandomNumber(max)
	fmt.Print("Guess a number between 1 and ", max, " : ")

	var guess int
	var i int
	for i = 0; i < 5; i++ {
		fmt.Scanln(&guess)
		if guess < 1 || guess > max {
			fmt.Println("Invalid guess. Try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is too low")
		} else if guess > secretNumber {
			fmt.Println("Your guess is too high")
		} else {
			fmt.Println("Congratulations! You won!")
			fmt.Println("You took", i+1, "tries")
			break
		}
	}
	fmt.Println("The secret number was", secretNumber)
	if i == 0 {
		fmt.Println("Game Over!")
	}

	fmt.Println("\n1.Play again\n2.Exit")
	fmt.Print("Enter your choice: ")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		selectLevel()
	case 2:
		fmt.Println("Exiting Game...")
		clearScreen()
		os.Exit(0)
	}
}

func generateRandomNumber(max int) int {
	b := make([]byte, 1)
	rand.Read(b)
	return int(b[0])%max + 1
}

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported operating system")
	}
}
