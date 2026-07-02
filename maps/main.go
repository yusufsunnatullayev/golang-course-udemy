package main

import "fmt"

func main() {
	// 2) way to declare map
	// var colors map[string]string

	// 3) way to declare map
	// colors := make(map[int]string)

	// colors[10] = "#fff"

	// delete(colors, 10)

	// 1) way to declare map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"white": "#fff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
