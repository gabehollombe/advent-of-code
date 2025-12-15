package main

import (
	"bufio"
	"fmt"
	"os"
)

// Cell represents a position on the grid
type Cell struct {
	X         int
	Y         int
	IsEmpty   bool
	Neighbors []*Cell
}

// IsAccessible returns true if the cell can be accessed by a forklift
// A roll of paper (@) can be accessed if there are fewer than 4 rolls
// in the eight adjacent positions
func (c *Cell) IsAccessible() bool {
	// Empty cells are not accessible (they don't contain paper rolls)
	if c.IsEmpty {
		return false
	}

	occupiedNeighborsCount := 0
	for _, neighbor := range c.Neighbors {
		if !neighbor.IsEmpty {
			occupiedNeighborsCount++
		}
	}
	return occupiedNeighborsCount < 4
}

// Room represents the warehouse containing all cells
type Room struct {
	Cells []*Cell
}

// FindCell returns the cell at the given coordinates, or nil if not found
func (r *Room) FindCell(x, y int) *Cell {
	for _, cell := range r.Cells {
		if cell.X == x && cell.Y == y {
			return cell
		}
	}
	return nil
}

// ParseInput reads the input file and returns a Room containing all cells
// Each cell is linked to its adjacent neighbors (up to 8)
func ParseInput(filename string) (*Room, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// First pass: read the grid and create cells
	var grid [][]*Cell
	scanner := bufio.NewScanner(file)
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		row := make([]*Cell, len(line))
		for x, char := range line {
			cell := &Cell{
				X:         x,
				Y:         y,
				IsEmpty:   char == '.',
				Neighbors: make([]*Cell, 0, 8),
			}
			row[x] = cell
		}
		grid = append(grid, row)
		y++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Second pass: link neighbors
	height := len(grid)
	if height == 0 {
		return &Room{Cells: []*Cell{}}, nil
	}
	width := len(grid[0])

	// Define the 8 directions (dx, dy)
	directions := [][2]int{
		{-1, -1}, {0, -1}, {1, -1}, // top-left, top, top-right
		{-1, 0}, {1, 0}, // left, right
		{-1, 1}, {0, 1}, {1, 1}, // bottom-left, bottom, bottom-right
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := grid[y][x]

			// Check all 8 directions
			for _, dir := range directions {
				nx, ny := x+dir[0], y+dir[1]

				// Check if neighbor is within bounds
				if nx >= 0 && nx < width && ny >= 0 && ny < height {
					neighbor := grid[ny][nx]
					cell.Neighbors = append(cell.Neighbors, neighbor)
				}
			}
		}
	}

	// Flatten grid into a single slice
	var cells []*Cell
	for _, row := range grid {
		cells = append(cells, row...)
	}

	return &Room{Cells: cells}, nil
}
