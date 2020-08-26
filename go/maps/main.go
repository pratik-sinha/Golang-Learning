package main

import "fmt"

//maps all keys are of the same type and all values must be
//the same type (keys and values can be of different types)
//reference type and we can iterate over them
func main() {
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// delete(colors, "white")

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf545",
		"white": "#ffffff",
	}
	printMap(colors)
	// fmt.Println(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println(color, hex)
	}
}
