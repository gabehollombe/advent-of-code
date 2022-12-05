package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Section struct {
	start int64
	end   int64
}

type AssignmentGroup []Section

type OverlapResult struct {
	fullOverlaps    []AssignmentGroup
	partialOverlaps []AssignmentGroup
}

func newSection(s string) Section {
	parts := strings.Split(s, "-")
	start, _ := strconv.ParseInt(parts[0], 10, 32)
	end, _ := strconv.ParseInt(parts[1], 10, 32)
	return Section{start: start, end: end}
}

func main() {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	assignmentGroups := newAssignmentGroups(lines)
	overlapResults := overlapResults(assignmentGroups)

	part1 := len(overlapResults.fullOverlaps)
	part2 := len(overlapResults.partialOverlaps)

	fmt.Printf("Day 1: %v\n", part1)
	fmt.Printf("Day 2: %v\n", part2)
}

func newAssignmentGroups(lines []string) []AssignmentGroup {
	groups := make([]AssignmentGroup, 0)
	for _, l := range lines {
		parts := strings.Split(string(l), ",")
		group := []Section{newSection(parts[0]), newSection(parts[1])}
		groups = append(groups, group)
	}
	return groups
}

func sectionsFullyOverlap(s1 Section, s2 Section) bool {
	return (s2.start >= s1.start && s2.end <= s1.end) || (s1.start >= s2.start && s1.end <= s2.end)
}

func sectionsOverlapAtAll(s1 Section, s2 Section) bool {
	return (s1.start <= s2.start && s2.start <= s1.end) ||
		(s2.start <= s1.start && s1.start <= s2.end)
}

func overlapResults(groups []AssignmentGroup) OverlapResult {
	fullOverlaps := make([]AssignmentGroup, 0)
	partialOverlaps := make([]AssignmentGroup, 0)

	for _, group := range groups {
		//fmt.Printf("Checking Groups: %v\n", groups)
		lastCompFullyOverlaps := false
		lastCompPartiallyOverlaps := false
		for i := 0; i < len(group)-1; i++ {
			if sectionsFullyOverlap(group[i], group[i+1]) {
				lastCompFullyOverlaps = true
			}
			if sectionsOverlapAtAll(group[i], group[i+1]) {
				lastCompPartiallyOverlaps = true
			}
		}
		if lastCompFullyOverlaps {
			//fmt.Printf("OVERLAP FOUND: %v\n", group)
			fullOverlaps = append(fullOverlaps, group)
		}
		if lastCompPartiallyOverlaps {
			//fmt.Printf("OVERLAP FOUND: %v\n", group)
			partialOverlaps = append(partialOverlaps, group)
		}
	}

	return OverlapResult{fullOverlaps: fullOverlaps, partialOverlaps: partialOverlaps}
}
