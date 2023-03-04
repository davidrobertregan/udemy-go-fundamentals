package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	name := os.Args[1]
	f, err := os.Open(name)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}