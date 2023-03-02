package main

import (
	"fmt"
	"net/http"
	"os"
)
	

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	/* resp is response object */

	fmt.Println(resp)
}