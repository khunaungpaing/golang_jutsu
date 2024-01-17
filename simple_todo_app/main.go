package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	todoLists := []string{"hi", "hello"}

	for {

		fmt.Println("Todo Lists :")
		for i, todo := range todoLists {
			fmt.Printf("%d. %v\n", i+1, todo)
		}

		fmt.Printf("\n\nOptions : \n1. Add task\n3.Exit choice: ")
		var choice uint8
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			todoLists = append(todoLists, scanner.Text())
		case 3:
			os.Exit(0)
		}
	}

}
