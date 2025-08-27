package main

import (
	"net/http"
)

func main() {
	links := []string{
		"https://www.youtube.com",
		"https://www.google.com",
		"https://www.arnavranjan.com",
	}

	for _, links := range links {
		checkLink(links)

	}
}

func checkLink(links string) {
	_, err := http.Get(links)
	if err != nil {
		println(links, "might be down!")
		return
	}
	println(links, "is up!")
}

// The above is not the best approach for this so some sort of a parallel approach would be better.
