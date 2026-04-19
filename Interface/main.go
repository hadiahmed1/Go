package main

import "fmt"

type bot interface {
	getGreeting() string
}

type enBot struct{}
type esBot struct{}

func main() {
	es := esBot{}
	en := enBot{}

	printGreeting(es)
	printGreeting(en)
}

func (enBot) getGreeting() string {
	return "Hi, how r you?"
}

func (esBot) getGreeting() string {
	return "Ola, como esta?"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}