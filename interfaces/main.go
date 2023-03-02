package main

import "fmt"

/* interfaces...
	- to all types out there! there's a new type called bot! Do you qualify?
	- wanna be in the club? You MUST have a method called getGreeting() that returns a string
	- if you're in, you get a new method called printGreeting()... congrats and welcome!
*/

type bot interface {
	getGreeting() string /* must list arg types too if they exist! */
}

type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
