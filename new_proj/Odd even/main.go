package main

//create a slice of ints from 0 through 10
func main() {
	// Create a slice of integers from 0 to 10
	numbers := make([]int, 11)
	for i := 0; i <= 10; i++ {
		numbers[i] = i
	}

	//check to see if the number is even or odd and then print it
	for _, number := range numbers {
		if number%2 == 0 {
			println(number, "is even")
		} else {
			println(number, "is odd")
		}
	}
}
