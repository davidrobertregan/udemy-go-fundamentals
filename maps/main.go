package main

import "fmt"

func main() { 
	/* map[keyType]valueType */

	// var colors map[string]string
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"

	// delete(colors, "white")

	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4BF745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%v has hex code of %v\n", color, hex)
	}

	// for key, value := range map { ... }
}