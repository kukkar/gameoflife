package gameoflife

func GameOfLifeRules(grid [][]bool) [][]bool {
	newGrid := make([][]bool, rows)
	for i := range newGrid {
		newGrid[i] = make([]bool, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			alive := grid[i][j]
			neighbors := countLiveNeighbors(grid, i, j)

			if alive && (neighbors < 2 || neighbors > 3) {
				newGrid[i][j] = false
			} else if !alive && neighbors == 3 {
				newGrid[i][j] = true
			} else {
				newGrid[i][j] = alive
			}
		}
	}
	return newGrid
}
