package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input = `1
2
3

4
5
6

7
8

9
10
11
12`
	lines, err := parseLinesToIntGroups(input)
	if err != nil {
		log.Fatal(err)
	}

	sums := sumGroups(lines)
	sort.Ints(sums)
	maxCal := sums[len(sums)-1]
	fmt.Printf("Part 1: %v\n", maxCal)

	top3Cal := sums[len(sums)-3:]
	fmt.Printf("Part 2: %v\n", sum(top3Cal))
}

func parseLinesToIntGroups(lines string) ([][]int, error) {
	var groups = make([][]int, 0)
	var group = make([]int, 0)
	for _, l := range strings.Split(lines, "\n") {
		l = strings.TrimSpace(l)
		if l == "" {
			groups = append(groups, group)
			group = make([]int, 0)
			continue
		}

		i, err := strconv.ParseInt(l, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("Could not parse line `%v` to an int", l)
		}
		group = append(group, int(i))
	}
	return groups, nil
}

func sumGroups(groups [][]int) []int {
	sums := make([]int, 0)
	for _, group := range groups {
		sums = append(sums, sum(group))
	}
	return sums
}

func sum(group []int) int {
	s := 0
	for _, i := range group {
		s += i
	}
	return s
}
