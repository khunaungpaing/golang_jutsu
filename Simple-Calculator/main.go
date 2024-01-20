package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	fmt.Println("Welcome to Simple Calculator")
	fmt.Println("Choose the operation you want to perform")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Exit")
	var choice int
	fmt.Scanln(&choice)

	clearScreen()
	switch choice {
	case 1:
		fmt.Print("Enter counts of the numbers you want to add : ")
		var n int
		fmt.Scanln(&n)
		var a []float32
		for i := 0; i < n; i++ {
			var num float32
			fmt.Scanln(&num)
			a = append(a, num)
		}
		fmt.Println("Sum is", Sum(a...))
	case 2:
		fmt.Print("Enter first number : ")
		var a float32
		fmt.Scanln(&a)
		fmt.Print("Enter second number : ")
		var b float32
		fmt.Scanln(&b)
		fmt.Println("Difference is", Sub(a, b))
	case 3:
		fmt.Print("Enter counts of the numbers you want to multiply : ")
		var n int
		fmt.Scanln(&n)
		var a []float32
		for i := 0; i < n; i++ {
			var num float32
			fmt.Scanln(&num)
			a = append(a, num)
		}
		fmt.Println("Product is", Mul(a...))
	case 4:
		fmt.Print("Enter first number : ")
		var a float32
		fmt.Scanln(&a)
		fmt.Print("Enter second number : ")
		var b float32
		fmt.Scanln(&b)
		res, err := Div(a, b)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Quotient is", res)
		}
	case 5:
		fmt.Println("Thank you for using Simple Calculator")
	default:
		fmt.Println("Invalid choice")
	}
}
func Sum(a ...float32) float32 {
	var sum float32
	for _, v := range a {
		sum += v
	}
	return sum
}

func Sub(a float32, b float32) float32 {
	return a - b
}

func Mul(a ...float32) float32 {
	var mul float32 = 1
	for _, v := range a {
		mul *= v
	}
	return mul
}

func Div(a float32, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
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
