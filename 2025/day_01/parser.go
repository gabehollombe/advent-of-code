package main

import (
	"fmt"
	"strconv"
)

type Direction rune

const (
	Left  Direction = 'L'
	Right Direction = 'R'
)

type Rotation struct {
	Direction Direction
	Distance  int
}

func ParseRotation(line string) (Rotation, error) {
	if len(line) < 2 {
		return Rotation{}, fmt.Errorf("invalid rotation: %s", line)
	}

	direction := Direction(line[0])
	if direction != Left && direction != Right {
		return Rotation{}, fmt.Errorf("invalid direction: %c", direction)
	}

	distance, err := strconv.Atoi(line[1:])
	if err != nil {
		return Rotation{}, fmt.Errorf("invalid distance: %s", line[1:])
	}

	return Rotation{
		Direction: direction,
		Distance:  distance,
	}, nil
}