package main

import (
	"fmt"
	"log"
)

func main() {
	banks, err := ParseInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	var bestJoltagesPerBank []int32
	for _, bank := range banks {
		bestJoltagesPerBank = append(bestJoltagesPerBank, GetBestJoltage(bank))
	}

	var sum int32
	for _, joltage := range bestJoltagesPerBank {
		sum += joltage
	}

	fmt.Printf("Sum of best joltages: %d\n", sum)

	var bestJoltagesPerBankGeneric []int64
	for _, bank := range banks {
		bestJoltagesPerBankGeneric = append(bestJoltagesPerBankGeneric, GetBestJoltageGeneric(bank, 12))
	}

	var sumGeneric int64
	for _, joltage := range bestJoltagesPerBankGeneric {
		sumGeneric += joltage
	}

	fmt.Printf("Sum of best joltages (generic with 12): %d\n", sumGeneric)
}
