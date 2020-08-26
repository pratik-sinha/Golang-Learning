package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

//if a type exists in the program that has this function
//then its also a member of bot type and so it can be
//used with bot types functions
type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//if we want to omit value we can leave it out
func (englishBot) getGreeting() string {
	return "hello there"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
