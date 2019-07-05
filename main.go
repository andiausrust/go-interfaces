package main

import "fmt"

type englishBot struct {

}
type spanishBot struct {

}

type greeting interface {
	getGreeting()
}

func main() {

	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)
}

func printGreeting(g greeting){
	g.getGreeting()
}

func (englishBot) getGreeting() {
	fmt.Println("hello")
}

func (spanishBot) getGreeting() {
	fmt.Println("ola")
}
