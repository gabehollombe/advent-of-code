package main

import (
	"strconv"
)

// IsInvalidId checks if an ID is made only of some sequence of digits repeated twice
func IsInvalidId(id int) bool {
	return IsInvalidIdWithReps(id, 2)
}

// IsInvalidIdWithReps checks if an ID is made only of some sequence of digits repeated checkReps times
func IsInvalidIdWithReps(id int, checkReps int) bool {
	// convert the id to a string s
	s := strconv.Itoa(id)

	// Check if the length is divisible by checkReps
	subStrLen := len(s) / checkReps
	if len(s)%checkReps != 0 {
		return false
	}

	// if s is less than 2 characters long return false
	if len(s) < 2 {
		return false
	}

	// make a slice `parts` = split s into strings each subStrLen length long
	parts := make([]string, checkReps)
	for i := 0; i < checkReps; i++ {
		parts[i] = s[i*subStrLen : (i+1)*subStrLen]
	}

	// convert this into a set `partsSet`
	partsSet := make(map[string]bool)
	for _, part := range parts {
		partsSet[part] = true
	}

	// if the size of partsSet == 1 return true else false
	return len(partsSet) == 1
}

// GetInvalidIds returns all invalid IDs within a given range (inclusive)
func GetInvalidIds(r Range) []int {
	var invalidIds []int
	for id := r.Start; id <= r.End; id++ {
		if IsInvalidId(id) {
			invalidIds = append(invalidIds, id)
		}
	}
	return invalidIds
}

func GetInvalidIdsPartTwo(r Range) []int {
	var invalidIds []int
	for id := r.Start; id <= r.End; id++ {
		// Check if the ID is invalid with 2 or more repetitions
		for reps := 2; reps <= len(strconv.Itoa(id)); reps++ {
			if IsInvalidIdWithReps(id, reps) {
				invalidIds = append(invalidIds, id)
				break
			}
		}
	}
	return invalidIds
}

// SumInvalidIds calculates the sum of all invalid IDs across all ranges
func SumInvalidIds(ranges []Range) int {
	sum := 0
	for _, r := range ranges {
		invalidIds := GetInvalidIds(r)
		for _, id := range invalidIds {
			sum += id
		}
	}
	return sum
}

// SumInvalidIdsPartTwo calculates the sum of all invalid IDs (with 2+ repetitions) across all ranges
func SumInvalidIdsPartTwo(ranges []Range) int {
	sum := 0
	for _, r := range ranges {
		invalidIds := GetInvalidIdsPartTwo(r)
		for _, id := range invalidIds {
			sum += id
		}
	}
	return sum
}
