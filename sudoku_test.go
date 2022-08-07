package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertNoZeroes(t *testing.T, tab [][]int) (solved bool) {
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

func TestSudokuFillLastNumber(t *testing.T) {
	table = [][]int{
		{1, 2, 3, 4},
		{3, 4, 2, 1},
		{2, 1, 4, 3},
		{4, 3, 1, 0},
	}
	setDimensions()
	sudokuSolver()
	assertNoZeroes(t, <-solutions)
}

func TestSudokuSolve(t *testing.T) {
	table = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	setDimensions()
	go sudokuSolver()
	assertNoZeroes(t, <-solutions)
}

func TestNoTable(t *testing.T) {
	table = [][]int{}
	err := validateTable()
	assert.NotNil(t, err)
}

func TestMismatchedTable(t *testing.T) {
	table = [][]int{{0, 0}}
	err := validateTable()
	assert.NotNil(t, err)
}

func TestImperfectSquareTable(t *testing.T) {
	table = [][]int{
		{0, 0},
		{0, 0},
	}
	err := validateTable()
	assert.NotNil(t, err)
}

func TestValidTable(t *testing.T) {
	table = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	err := validateTable()
	assert.Nil(t, err)
}
