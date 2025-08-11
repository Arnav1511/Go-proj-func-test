package main

import (
	"fmt"
)

func main() {
	intArr := [...]int32{1, 5, 7}
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[1])

	var intSlice []int32 = intArr[1:3]
	fmt.Println(intSlice)

	intSlice = append(intSlice, 12, 13, 14)
	fmt.Println(intSlice)
}
