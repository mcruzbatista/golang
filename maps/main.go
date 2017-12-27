package main

import "fmt"

func main() {

	//Create a empty map
	//var colors map[string]string

	//Create a empty map
	//colors := make(map[string]string)

	//Create a map with some values
	colors := map[string]string{
		"red":"#ff0000",
		"green":"#4bf745",
		"white":"#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string){

	for color,hex := range c{
		fmt.Println("Hex code for", color, "is", hex)
	}
}
