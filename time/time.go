package main

import (
	"fmt"
	"time"
)

func main() {

	//testing the go speed by hacing a long int of 100000 value wihout a preset capacity
	//and then with a preset capacity
	//this is to show the difference in performance when using slices with and without preallocation
	//the first one will take longer because it has to resize the slice multiple times
	//the second one will be faster because it has a preset capacity and doesn't need to resize
	//so we can see how much faster it is to use preallocation
	var n int = 100000
	var testslice = []int{}
	var testslice2 = make([]int, 0, n)

	fmt.Println("Starting loop to fill slice with 100000 integers...%v\n", timeLoop(testslice, n))
	fmt.Println("Starting loop to fill slice using preallocation...%v\n", timeLoop(testslice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)

	}
	return time.Since(t0)
}
