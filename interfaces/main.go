package main

import "fmt"

//In go interfaces are not generic types. Other languages have generic types.
//Interfaces are implicit. We don't manually have to say that our custom type satisfies some interface.
//Interfaces are a contract to help us manage types. If our custom type's implementation of a function is broken then interfaces won't help us.
//Interfaces are tough. Understand how to read interfaces in the standard lib. Writing your own interfaces is tough and requires experience.
type bot interface { //interfaceType : we can't create a value directly from interface type.
	getGreeting() string
}

type englishBot struct{} //concreteType : englishBot, map, int, string : You can directly create a value from concrete type
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//VERY custom logic for generating an English greeting.
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	//VERY custom logic for generating an Spanish greeting.
	return "Hola!"
}
