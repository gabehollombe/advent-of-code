package main

import "testing"

func TestGetBestJoltage(t *testing.T) {
	testCases := []struct {
		name     string
		bank     BatteryBank
		expected int32
	}{
		{
			name:     "Example 1",
			bank:     BatteryBank{Batteries: []int32{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}},
			expected: 98,
		},
		{
			name:     "Example 2",
			bank:     BatteryBank{Batteries: []int32{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}},
			expected: 89,
		},
		{
			name:     "Example 3",
			bank:     BatteryBank{Batteries: []int32{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}},
			expected: 78,
		},
		{
			name:     "Example 4",
			bank:     BatteryBank{Batteries: []int32{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}},
			expected: 92,
		},
		{
			name:     "Small Bank",
			bank:     BatteryBank{Batteries: []int32{1, 2, 3, 4, 5}},
			expected: 45, // 4 and 5
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetBestJoltage(tc.bank)
			if got != tc.expected {
				t.Errorf("GetBestJoltage() = %d, want %d", got, tc.expected)
			}
		})
	}
}

func TestGetBestJoltagePartTwo(t *testing.T) {
	testCases := []struct {
		name     string
		bank     BatteryBank
		expected int64
	}{
		{
			name:     "Example 1 - 987654321111111",
			bank:     BatteryBank{Batteries: []int32{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}},
			expected: 987654321111,
		},
		{
			name:     "Example 2 - 811111111111119",
			bank:     BatteryBank{Batteries: []int32{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}},
			expected: 811111111119,
		},
		{
			name:     "Example 3 - 234234234234278",
			bank:     BatteryBank{Batteries: []int32{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}},
			expected: 434234234278,
		},
		{
			name:     "Example 4 - 818181911112111",
			bank:     BatteryBank{Batteries: []int32{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}},
			expected: 888911112111,
		},
		{
			name:     "All same digits",
			bank:     BatteryBank{Batteries: []int32{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}},
			expected: 555555555555,
		},
		{
			name:     "Exactly 12 batteries",
			bank:     BatteryBank{Batteries: []int32{9, 8, 7, 6, 5, 4, 3, 2, 1, 9, 8, 7}},
			expected: 987654321987,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetBestJoltagePartTwo(tc.bank)
			if got != tc.expected {
				t.Errorf("GetBestJoltagePartTwo() = %d, want %d", got, tc.expected)
			}
		})
	}
}

func BenchmarkGetBestJoltageGeneric(b *testing.B) {
	banks, err := ParseInput("input.txt")
	if err != nil {
		b.Fatalf("Error reading input: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, bank := range banks {
			GetBestJoltageGeneric(bank, 12)
		}
	}
}

// func BenchmarkGetBestJoltageGenericVsPartTwo(b *testing.B) {
// 	banks, err := ParseInput("input.txt")
// 	if err != nil {
// 		b.Fatalf("Error reading input: %v", err)
// 	}

// 	b.Run("Generic_12", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			for _, bank := range banks {
// 				GetBestJoltageGeneric(bank, 12)
// 			}
// 		}
// 	})

// 	b.Run("PartTwo", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			for _, bank := range banks {
// 				GetBestJoltagePartTwo(bank)
// 			}
// 		}
// 	})
// }
