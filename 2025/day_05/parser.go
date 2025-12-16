package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FreshRange struct {
	Start int
	End   int
}

type Database struct {
	Ranges    []FreshRange
	Available []int
}

func (d *Database) IsFresh(id int) bool {
	for _, r := range d.Ranges {
		if id >= r.Start && id <= r.End {
			return true
		}
	}
	return false
}

func (d *Database) AddRange(r FreshRange) {
	// Merge the new range with all overlapping existing ranges
	newRanges := []FreshRange{}

	for _, existing := range d.Ranges {
		if r.Start <= existing.End && r.End >= existing.Start {
			// Ranges overlap - merge them
			if existing.Start < r.Start {
				r.Start = existing.Start
			}
			if existing.End > r.End {
				r.End = existing.End
			}
		} else {
			// No overlap - keep the existing range
			newRanges = append(newRanges, existing)
		}
	}

	// Add the merged (or new) range
	newRanges = append(newRanges, r)
	d.Ranges = newRanges
}

func (d *Database) CountTotalPossibleFresh() int {
	count := 0
	for _, r := range d.Ranges {
		count += r.End - r.Start + 1
	}
	return count
}

func ParseInput(filename string) (*Database, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	db := &Database{
		Ranges:    []FreshRange{},
		Available: []int{},
	}

	parsingRanges := true
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Empty line separates ranges from available IDs
		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			// Parse range
			freshRange, err := parseRange(line)
			if err != nil {
				return nil, err
			}
			db.AddRange(freshRange)
			// db.Ranges = append(db.Ranges, freshRange)
		} else {
			// Parse available ID
			id, err := strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("invalid available ID: %s", line)
			}
			db.Available = append(db.Available, id)
		}
	}

	return db, nil
}

func parseRange(line string) (FreshRange, error) {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return FreshRange{}, fmt.Errorf("invalid range format: %s", line)
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return FreshRange{}, fmt.Errorf("invalid start value: %s", parts[0])
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return FreshRange{}, fmt.Errorf("invalid end value: %s", parts[1])
	}

	return FreshRange{Start: start, End: end}, nil
}
