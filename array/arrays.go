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
	// we can print out the length and capacity of the slice

	fmt.Println("Length:", len(intSlice))
	fmt.Println("Capacity:", cap(intSlice))

	intSlice = append(intSlice, 12, 13, 14)
	fmt.Println(intSlice)

	// we can print out the length and capacity of the slice
	fmt.Println("Length:", len(intSlice))
	fmt.Println("Capacity:", cap(intSlice))

	//we can also append a slice to another slice so let me show you that as well below :)
	intSlice2 := []int32{15, 16, 17}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	//I found another way to make a slice from an array
	//using the make function
	//this is useful when you want to create a slice with a specific length and capacity
	//but you don't want to use the array syntax
	//so we can use the make function to create a slice from an array
	//and then copy the elements from the array to the slice
	intSlice3 := make([]int32, 2)
	copy(intSlice3, intArr[1:3])
	fmt.Println(intSlice3)

	var intSlice4 []int32 = make([]int32, 2, 5)
	copy(intSlice4, intArr[1:3])
	fmt.Println(intSlice4)
}
