package main

import "fmt"

func main() {
	colours := map[string]string{
		"red":   "s#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	delete(colours, "green")
	fmt.Println(colours)
}
