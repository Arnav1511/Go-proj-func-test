package main

import "fmt"

func main(){
    var printvalue string = "Hello, World!"
    printMe(printvalue)

    var numerator int = 10
    var denominator int = 2
    var result, remainder int = intdivision(numerator, denominator)
    fmt.Printf("the result of the integer division is %v with remainder %v\n", result, remainder)
}

func printMe(printvalue string){
    fmt.Println(printvalue)
}

func intdivision(numerator int, denominator int) (int, int){
    var result int = numerator/denominator
    var remainder int = numerator%denominator
    return result, remainder
}