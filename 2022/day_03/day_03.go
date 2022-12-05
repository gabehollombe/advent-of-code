package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Returns the intersection of letters between two arrays of strings
func intersection(s1 []string, s2 []string) []string {
	m1 := make(map[rune]bool)
	m2 := make(map[rune]bool)
	for _, s := range s1 {
		for _, r := range s {
			m1[r] = true
		}
	}
	for _, s := range s2 {
		for _, r := range s {
			m2[r] = true
		}
	}
	res := make([]string, 0)
	for s, _ := range m1 {
		if m2[s] {
			res = append(res, string(s))
		}
	}

	return res
}

// Returns priority of item (a=1 ... A=27)
func priority(b rune) int {
	var a rune = 'a'
	var z rune = 'z'
	var A rune = 'A'
	var Z rune = 'Z'
	if b >= a && b <= z {
		return int(b-a) + 1
	}
	if b >= A && b <= Z {
		return int(b-A) + 27
	}

	log.Fatalln("Did not expect to get a byte value outside of a-zA-Z")
	return -1
}

func part1Priority(line string) int {
	// Split rucksack contents into two sets
	itemsPerRucksack := len(line) / 2
	ruck1 := line[:itemsPerRucksack]
	ruck2 := line[itemsPerRucksack:]

	// Find the intersection between both sets (should only be one value)
	commonItem := intersection([]string{ruck1}, []string{ruck2})[0] // should only be 1 common per line
	return priority(rune(commonItem[0]))                            // common item should be a single char so cast to rune
}

func part2Priority(lineGroup [3]string) int {
	// Find the intersection between the first two lines
	i := intersection([]string{lineGroup[0]}, []string{lineGroup[1]})
	// and intersect that with the last line
	i = intersection(i, []string{lineGroup[2]})
	// We should only have one item remaining
	commonItem := i[0]
	return priority(rune(commonItem[0]))
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	part1PrioritiesSum := 0
	part2PrioritiesSum := 0
	lineNum := 0
	lineGroup := [3]string{}
	for scanner.Scan() {
		line := scanner.Text()
		part1PrioritiesSum += part1Priority(line)

		if lineNum != 0 && lineNum%3 == 0 {
			part2PrioritiesSum += part2Priority(lineGroup)
		}
		lineGroup[lineNum%3] = line
		lineNum++
	}
	part2PrioritiesSum += part2Priority(lineGroup)

	fmt.Printf("Day 1: %v\n", part1PrioritiesSum)
	fmt.Printf("Day 2: %v\n", part2PrioritiesSum)
}
