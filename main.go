package main

import (
    "fmt"
    "errors"
)

func main(){
    var printvalue string = "Hello, World!"
    printMe(printvalue)

    var numerator int = 10
    var denominator int = 3
    var result, remainder, err = intdivision(numerator, denominator)
    if err!=nil {
        fmt.Printf(err.Error())
    }else if remainder == 0{
        fmt.Printf("Amaze")
    }
    fmt.Printf("the result of the integer division is %v with remainder %v\n", result, remainder)
}

func printMe(printvalue string){
    fmt.Println(printvalue)
}

func intdivision(numerator int, denominator int) (int, int, error){
    var err error
    if denominator == 0 {
        err = errors.New("division by zero so we cannot divide mate!")
        return 0, 0, err
    }
    var result int = numerator/denominator
    var remainder int = numerator%denominator
    return result, remainder , err
}