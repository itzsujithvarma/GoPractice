package main

import "fmt"

type EnglishBot struct{}
type SpanishBot struct{}

type bot interface {
	getGreeting() string
}

func (EnglishBot) getGreeting() string {
	return "Hello!"
}

func (SpanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}
	printGreeting(sb)
	printGreeting(eb)
}
