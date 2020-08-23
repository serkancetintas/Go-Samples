package main

import "fmt"

func main() {
	// var colors map[string]string
	colors := map[string]string{
		"red":   "redasd",
		"green": "greenvalue",
		"blue":  "bluevalue",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("key", color, "is", hex)
	}
}
