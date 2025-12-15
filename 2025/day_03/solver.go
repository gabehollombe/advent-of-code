package main

// GetBestJoltageGeneric selects numBatteries from the bank to form the largest possible number
func GetBestJoltageGeneric(bank BatteryBank, numBatteries int) int64 {
	if numBatteries <= 0 || numBatteries > len(bank.Batteries) {
		return 0
	}

	selectedIndexes := selectBestIndexes(bank.Batteries, 0, numBatteries)

	// Compose the selected batteries into a single number
	var result int64 = 0
	for _, idx := range selectedIndexes {
		result = result*10 + int64(bank.Batteries[idx])
	}

	return result
}

// selectBestIndexes recursively selects the best battery indexes
// startIdx: the earliest index we can consider
// remaining: how many more batteries we need to select
func selectBestIndexes(batteries []int32, startIdx int, remaining int) []int {
	if remaining == 0 {
		return []int{}
	}

	// Try to find the highest value battery that leaves enough batteries for the rest
	for targetValue := int32(9); targetValue >= 0; targetValue-- {
		// Look for the earliest index of targetValue starting from startIdx
		for i := startIdx; i < len(batteries); i++ {
			if batteries[i] == targetValue {
				// Check if there are enough batteries left after this one
				batteriesLeft := len(batteries) - i - 1
				if batteriesLeft >= remaining-1 {
					// Select this index and recursively select the rest
					restIndexes := selectBestIndexes(batteries, i+1, remaining-1)
					return append([]int{i}, restIndexes...)
				}
			}
		}
	}

	// Fallback: just take the first 'remaining' batteries from startIdx
	result := make([]int, remaining)
	for i := 0; i < remaining; i++ {
		result[i] = startIdx + i
	}
	return result
}

func GetBestJoltage(bank BatteryBank) int32 {
	return int32(GetBestJoltageGeneric(bank, 2))
}

func GetBestJoltagePartTwo(bank BatteryBank) int64 {
	return GetBestJoltageGeneric(bank, 12)
}
