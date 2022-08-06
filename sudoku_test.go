package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertSolved(t *testing.T, tab [][]int) (solved bool) {
	solved = true
	for r, row := range tab {
		if !assert.NotContains(t, row, 0) {
			t.Logf("row: %d, %v", r, row)
			solved = false
			break
		}
	}
	return
}

func TestSudoku(t *testing.T) {
	table = [][]int{
		{1, 2, 3, 4},
		{3, 4, 2, 1},
		{2, 1, 4, 3},
		{4, 3, 1, 0},
	}
	sudokuSolver()
	solution := <-solutions
	assertSolved(t, solution)
}
