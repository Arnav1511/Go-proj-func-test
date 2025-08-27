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
