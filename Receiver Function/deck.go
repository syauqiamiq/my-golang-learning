package main

import (
	"fmt"
	"math/rand"
)

type deck []string

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func (d deck) shuffle() {
	for index := range d {
		newPos := rand.Intn(len(d) - 1)
		d[index], d[newPos] = d[newPos], d[index]
	}
}
