package main

import (
	"fmt"
	"os"

	"cpcmd"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: cp <source> <destination>")
		os.Exit(1)
	}

	source := os.Args[1]
	destination := os.Args[2]

	err := cp.CopyFile(source, destination)
	if err != nil {
		fmt.Printf("Error copying file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("File %s copied to %s successfully.\n", source, destination)
}
