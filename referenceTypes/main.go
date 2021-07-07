package main

import "fmt"

func main() {

	//When you create slice Go internally creates 2 seperate data structures.
	//1-) Slice(pointer,capacity,length)
	//2-) Array(Actual Array contains data)
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
