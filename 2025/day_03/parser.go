package main

import (
	"bufio"
	"os"
)

type BatteryBank struct {
	Batteries []int32
}

func ParseInput(filename string) ([]BatteryBank, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var banks []BatteryBank
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var batteries []int32
		for _, r := range line {
			if r >= '0' && r <= '9' {
				batteries = append(batteries, int32(r-'0'))
			}
		}
		banks = append(banks, BatteryBank{Batteries: batteries})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return banks, nil
}
