package main

import "fmt"

// Slice is an array that can grow or shrink
// Array is fixed lenght list of things

func main() {
	// Slice of String
	cards := []string{"Spade", "Diamonds"}

	// Appending Slice
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	// Loop through Slice
	for index, card := range cards {
		fmt.Println(index, card)
	}

	// Manual Loop
	for i := 0; i < len(cards); i++ {
		fmt.Println(i, cards[i])
	}

	// Slice Range
	// Get Card with index [0 to 1]
	sliced := cards[0:2]
	fmt.Println(sliced)

	// Get Card with index [0 until before n]
	sliceNew := cards[:2]
	fmt.Println(sliceNew)

	// Get Card with index [2 until n]
	sliceNewNew := cards[2:]
	fmt.Println(sliceNewNew)

}
