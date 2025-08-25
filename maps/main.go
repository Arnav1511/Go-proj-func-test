package main

import "fmt"

func main() {
	colours := map[string]string{
		"red":   " #ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}
	delete(colours, "green")
	printMap(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Printf("Hex code for %s is %s\n", colour, hex)
	}
}
