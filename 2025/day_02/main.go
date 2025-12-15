package main

import (
	"fmt"
	"log"
)

func main() {
	ranges, err := ParseInput("input.txt")
	if err != nil {
		log.Fatalf("Failed to parse input: %v", err)
	}

	sum := SumInvalidIds(ranges)
	fmt.Printf("Part One - Sum of invalid IDs: %d\n", sum)

	sumPartTwo := SumInvalidIdsPartTwo(ranges)
	fmt.Printf("Part Two - Sum of invalid IDs: %d\n", sumPartTwo)
}
