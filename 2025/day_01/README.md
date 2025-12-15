# Day 01: Virtual Safe Dial

## Problem Summary

A virtual safe has a dial with numbers 0-99 arranged in a circle. Starting at position 50, we need to follow a sequence of rotations (Left or Right with a distance) and count how many times the dial points at 0 after any rotation.

## Solution

**Answer: 1139**

## Implementation

The solution is implemented in Go using Test-Driven Development (TDD) principles.

### File Structure

```
day_01/
├── dial.go           # Core dial rotation logic
├── dial_test.go      # Unit tests for dial operations
├── parser.go         # Rotation parsing from input
├── parser_test.go    # Parser unit tests
├── solver.go         # Main solving logic
├── solver_test.go    # Integration tests
├── main.go           # Entry point
├── input.txt         # Puzzle input
└── README.md         # This file
```

### Key Components

1. **Data Structures**
   - [`Direction`](dial.go:5): Enum for Left ('L') or Right ('R')
   - [`Rotation`](parser.go:11): Represents a rotation instruction
   - [`Dial`](dial.go:9): Tracks the current position (0-99)

2. **Core Functions**
   - [`ParseRotation()`](parser.go:17): Parses input lines like "L68" or "R48"
   - [`Dial.Rotate()`](dial.go:15): Applies rotation with wraparound handling
   - [`CountZeroPositions()`](solver.go:3): Main solver that counts zeros

### Algorithm

The rotation logic handles wraparound correctly:

```go
// Right rotation
position = (position + distance) % 100

// Left rotation (add 100 to handle negative numbers)
position = (position - distance + 100) % 100
```

### Testing

All tests pass, including:
- Unit tests for parsing, rotation, and zero detection
- Integration test with the provided example (expected: 3, got: 3)
- Edge cases for wraparound behavior

Run tests:
```bash
cd day_01
go test -v
```

Run solution:
```bash
cd day_01
go run .
```

### Example Walkthrough

From the problem:
```
Start: 50
L68  → 82
L30  → 52
R48  → 0  ✓ (count: 1)
L5   → 95
R60  → 55
L55  → 0  ✓ (count: 2)
L1   → 99
L99  → 0  ✓ (count: 3)
R14  → 14
L82  → 32

Total zeros: 3
```

This matches our test expectations perfectly.