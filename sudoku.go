package main

import (
	"errors"
	"fmt"
	"math"
)

var (
	dimension, boxDimension int
	solutions               = make(chan [][]int, 2)
	table                   = [][]int{

		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		// {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},

		// {0, 0, 0, 0},
		// {0, 0, 0, 0},
		// {0, 0, 0, 0},
		// {0, 0, 0, 0},

		// {0},
	}
)

func main() {
	setDimensions()
	if err := validateTable(); err != nil {
		fmt.Println(err)
		return
	}
	go sudokuSolver()
	another := ""
	for solution := range solutions {
		for _, row := range solution {
			fmt.Println(row)
		}
		fmt.Print("Show another solution [y/n]? ")
		fmt.Scanln(&another)
		if another != "y" {
			break
		}
		if len(solutions) == 0 {
			close(solutions)
		}
	}
}

func setDimensions() {
	dimension = len(table)
	boxDimension = int(math.Sqrt(float64(dimension)))
}

func validateTable() error {
	if len(table) < 1 {
		return errors.New("no data")
	}
	for _, row := range table {
		if len(row) != len(table) {
			return errors.New("size mismatch")
		}
	}
	dimensionSqrt := math.Sqrt(float64(len(table)))
	if math.Pow(math.Floor(dimensionSqrt), 2) != float64(len(table)) {
		return errors.New("not a perfect square")
	}
	return nil
}

func sudokuSolver() {
	for row := range table {
		for col := range table[row] {
			if table[row][col] == 0 {
				for potentialValue := 1; potentialValue <= len(table); potentialValue++ {
					if isValidValue(row, col, potentialValue) {
						table[row][col] = potentialValue
						sudokuSolver()
						table[row][col] = 0
					}
				}
				return
			}
		}
	}
	solutions <- copySolution()
}

func copySolution() (solution [][]int) {
	solution = make([][]int, dimension)
	for i, v := range table {
		solution[i] = make([]int, dimension)
		copy(solution[i], v)
	}
	return
}

func isValidValue(row, col, potentialValue int) bool {
	for colCheck := range table[row] {
		if table[row][colCheck] == potentialValue {
			return false
		}
	}
	for rowCheck := range table {
		if table[rowCheck][col] == potentialValue {
			return false
		}
	}
	boxColStart := (col / boxDimension) * boxDimension
	boxRowStart := (row / boxDimension) * boxDimension
	for boxRow := 0; boxRow < boxDimension; boxRow++ {
		for boxCol := 0; boxCol < boxDimension; boxCol++ {
			if table[boxRowStart+boxRow][boxColStart+boxCol] == potentialValue {
				return false
			}
		}
	}
	return true
}
