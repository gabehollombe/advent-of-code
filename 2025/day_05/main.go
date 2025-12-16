package main

import (
	"fmt"
	"log"
)

func main() {
	db, err := ParseInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	freshCount := CountFreshIngredients(db)

	fmt.Printf("Number of fresh ingredients: %d\n", freshCount)

	fmt.Printf("Number of total possible fresh ingredients: %d\n", db.CountTotalPossibleFresh())
}
