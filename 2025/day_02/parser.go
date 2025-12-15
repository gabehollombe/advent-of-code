package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

// ParseInput reads the input file and returns a list of ranges
func ParseInput(filename string) ([]Range, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	content := strings.TrimSpace(string(data))
	return ParseRanges(content)
}

// ParseRanges parses a string of comma-separated ranges
func ParseRanges(input string) ([]Range, error) {
	if input == "" {
		return []Range{}, nil
	}

	parts := strings.Split(input, ",")
	ranges := make([]Range, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		dashIndex := strings.Index(part, "-")
		if dashIndex == -1 {
			return nil, fmt.Errorf("invalid range format: %s (missing dash)", part)
		}

		startStr := part[:dashIndex]
		endStr := part[dashIndex+1:]

		if startStr == "" {
			return nil, fmt.Errorf("invalid range format: %s (missing start)", part)
		}
		if endStr == "" {
			return nil, fmt.Errorf("invalid range format: %s (missing end)", part)
		}

		start, err := strconv.Atoi(startStr)
		if err != nil {
			return nil, fmt.Errorf("invalid start value in range %s: %w", part, err)
		}

		end, err := strconv.Atoi(endStr)
		if err != nil {
			return nil, fmt.Errorf("invalid end value in range %s: %w", part, err)
		}

		ranges = append(ranges, Range{Start: start, End: end})
	}

	return ranges, nil
}
