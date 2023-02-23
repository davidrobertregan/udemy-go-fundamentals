package main /* what does this mean? 
	- package == project == workspace
	- collection of common source files
	- all files that belong to package main must have "package main" at the top
*/

import "fmt"
/* 
	- gives our package main access to the package 'fmt'
	- fmt is standard library package (format)
	- standard library (fmt, math, debug) or other reusable packages authored by other devs
	- golang.org/pkg
	- a lot of learning go is learning the standard library packages
*/

func main() {
	fmt.Println("Hello World")
}
/* 
	- similar to js
*/

/* how do we run the code?
	go run main.go (builds and executes)
	go build (compiles all files)
*/