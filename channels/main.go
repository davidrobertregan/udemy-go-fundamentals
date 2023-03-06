package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com", 
		"http://facebook.com", 
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	for l := range c { /* same as below, but better syntax for future devs */
		go func(link string) { /* function literal, the same as JS anonymous function */
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

	// for  {  /* infinite loop */
	// 	go checkLink(<-c, c)
	// }

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}