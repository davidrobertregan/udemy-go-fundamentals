package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com", 
		"http://facebook.com", 
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	for _, l := range links {
		go checkStatus(l)
	}
}

func checkStatus(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		return
	}

	fmt.Println(link, "is up")
}