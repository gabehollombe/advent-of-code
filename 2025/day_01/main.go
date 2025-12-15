package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	rotations, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	part1 := CountZeroPositions(rotations)
	part2 := CountZeroClicks(rotations)
	
	fmt.Printf("Part 1 - Password (final positions at 0): %d\n", part1)
	fmt.Printf("Part 2 - Password (all clicks through 0): %d\n", part2)
}

func readInput(filename string) ([]Rotation, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rotations []Rotation
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		
		rotation, err := ParseRotation(line)
		if err != nil {
			return nil, fmt.Errorf("error parsing line '%s': %w", line, err)
		}
		
		rotations = append(rotations, rotation)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rotations, nil
}