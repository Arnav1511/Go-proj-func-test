package main

import (
	"net/http"
)

func main() {
	links := []string{
		"https://www.youtube.com/watch?v=YS4e4q9oBaU",
		"https://www.youtube.com/watch?v=HxaD_trXwRE",
		"https://www.youtube.com/watch?v=Z1Yd7upQsXY",
		"https://www.youtube.com/watch?v=8pTEmbeENF4",
		"https://www.youtube.com/watch?v=2LhoCfjm8R4",
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
