package main

import "fmt"

func main() {
	var intSlice []int
	for i := 0; i <= 9; i++ {
		intSlice = append(intSlice, i)

		if intSlice[i]%2 == 0 {
			fmt.Println("Even:", intSlice[i])
		} else {
			fmt.Println("Odd:", intSlice[i])
		}
	}
}
