package main

import (
	"errors"
	"fmt"
)

func main() {
	var printvalue string = "Hello, World!"
	printMe(printvalue)

	var numerator int = 10
	var denominator int = 2
	var result, remainder, err = intdivision(numerator, denominator)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	switch remainder {
	case 0:
		fmt.Printf("the division was exact: %d / %d = %d\n", numerator, denominator, result)
	case 1, 2:
		fmt.Printf("the division was close")
	default:
		fmt.Printf("the division was not close")
	}
}

func printMe(printvalue string) {
	fmt.Println(printvalue)
}

func intdivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("division by zero so we cannot divide")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
