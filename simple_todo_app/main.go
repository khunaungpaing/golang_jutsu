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

		fmt.Printf("Options : \n1. Add task\nchoice: ")
		var choice uint8
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			todoLists = append(todoLists, scanner.Text())
		}
	}

}
