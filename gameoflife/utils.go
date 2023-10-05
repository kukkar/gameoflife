package gameoflife

import (
	"fmt"
	"math/rand"
)

func countLiveNeighbors(grid [][]bool, x, y int) int {
	count := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i >= 0 && i < rows && j >= 0 && j < columns && !(i == x && j == y) && grid[i][j] {
				count++
			}
		}
	}
	return count
}

func InitializeMatrix() [][]bool {
	grid := make([][]bool, rows)
	for i := range grid {
		grid[i] = make([]bool, columns)
		for j := range grid[i] {
			// Randomly initialize the grid with live or dead cells
			grid[i][j] = rand.Intn(2) == 1
		}
	}
	return grid
}

func PrintGrid(grid [][]bool) {
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				fmt.Print("■")
			} else {
				fmt.Print("□")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
