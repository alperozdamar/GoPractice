package main

import "fmt"

func main() {
	colors := map[int]string{
		1: "#ff0000",
		2: "#4bf745",
	}

	colors[10] = "#ffffff"
	fmt.Println(colors)
	delete(colors, 10)
	//fmt.Println(colors)
	printMap(colors)

	myMap := make(map[string]string)
	fmt.Println(myMap)
}

func printMap(c map[int]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}
