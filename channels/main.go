package main

import (
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.youtube.com",
		"https://www.google.com",
		// "https://www.arnavranjan.com",
	}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)

	}
	for l := range c {
		time.Sleep(time.Second)
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		println(link, "might be down!")
		c <- link
		return
	}
	println(link, "is up!")
	c <- link
}

// need to create channels above to make the print happen.
