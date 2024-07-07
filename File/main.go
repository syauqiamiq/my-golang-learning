package main

import (
	"fmt"
	"log"
	"os"
)

func saveToFile(fileName string) error {

	return os.WriteFile(fileName, []byte("ASDFASDF"), 0666)

}

func readFile(fileName string) ([]byte, error) {

	return os.ReadFile(fileName)

}

func main() {

	err := saveToFile("TEST.txt")
	if err != nil {
		log.Fatal("ERROR: SAVING FILE")
	}

	bs, err := readFile("TEST.txt")
	if err != nil {
		log.Fatal("ERROR: READING FILE")
	}

	fmt.Println(string(bs))

}
