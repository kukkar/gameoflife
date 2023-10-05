package main

import (
	"fmt"
	"gameoflife/gameoflife"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	grid := gameoflife.InitializeMatrix()

	for generation := 1; ; generation++ {
		fmt.Printf("Generation %d:\n", generation)
		gameoflife.PrintGrid(grid)
		time.Sleep(1000 * time.Millisecond)

		grid = gameoflife.GameOfLifeRules(grid)
	}
}
