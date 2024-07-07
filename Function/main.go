package main

import "fmt"

func main() {

	// Calling multiple return function
	myString, myNumber := newAnotherCards()
	fmt.Println(myString, myNumber)

	// Calling single return function
	mySecondString := newCards()

	fmt.Println(mySecondString)

}

// Single Return, the Return must be "String"
func newCards() string {
	return "Spades"
}

// Multiple Return, the Return must be "String" and "Int64"
func newAnotherCards() (string, int64) {
	return "Spades", 50
}
