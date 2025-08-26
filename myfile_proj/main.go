package main

import (
	"fmt"
	"io"
	"os"
)

//create a program that read the contents of a text file and prints it to the terminal

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
