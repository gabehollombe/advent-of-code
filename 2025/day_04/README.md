# Day 04: Paper Roll Accessibility

## Problem Summary

Given a 2D grid with paper rolls (`@`) and empty spaces (`.`), determine how many paper rolls can be accessed by a forklift. A paper roll is accessible if it has fewer than 4 paper rolls in its 8 adjacent positions.

## Example

Input:
```
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
```

Expected output: **13 accessible rolls**

## Code Structure

- **`parser.go`**: Contains the `Cell`, `Room` structs and `ParseInput` function
  - `Cell`: Represents a grid position with coordinates, empty status, and neighbor links
  - `Room`: Contains all cells in the warehouse and provides helper methods
  - `Room.FindCell(x, y)`: Finds a cell by its coordinates
  - `ParseInput`: Reads input file and creates a Room with all cells and neighbor relationships
  
- **`parser_test.go`**: Tests for parsing functionality
  - Verifies correct number of cells parsed
  - Verifies correct count of empty vs paper roll cells
  - Verifies neighbor relationships (corners have 3, edges have 5, center has 8)

- **`solver.go`**: Contains the main solving logic
  - `Cell.IsAccessible()`: Determines if a paper roll can be accessed (< 4 neighbors with paper)
  - `CountAccessibleRolls(room)`: Counts total accessible rolls in the room

- **`solver_test.go`**: Tests for solving functionality
  - Tests the example case (should return 13)
  - Tests specific cells that should/shouldn't be accessible

- **`main.go`**: Entry point that reads input and prints the result

## Running Tests

```bash
cd day_04
go test -v
```

## Running the Solution

```bash
cd day_04
go run .
```

## Implementation Notes

The parsing is fully implemented. You need to implement:

1. ✅ `ParseInput()` - **DONE** - Parses the grid and creates a Room with cells and neighbor links
2. ✅ `Room.FindCell(x, y)` - **DONE** - Finds a cell by coordinates
3. `Cell.IsAccessible()` - Check if a paper roll has < 4 paper roll neighbors
4. `CountAccessibleRolls(room)` - Count all accessible rolls in the room