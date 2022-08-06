package main

import (
	"errors"
	"fmt"
	"math"
)

var (
	solutions = make(chan [][]int, 2)
	table     = [][]int{

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

		// {0, 0},
		// {0, 0},

		// {0},
	}
)

func main() {
	if err := validInput(); err != nil {
		fmt.Println(err)
		return
	}
	go sudokuSolver()
	more := ""
	for solution := range solutions {
		for _, row := range solution {
			fmt.Println(row)
		}
		fmt.Print("Show another solution [y/n]? ")
		fmt.Scanln(&more)
		if more != "y" {
			break
		}
		if len(solutions) == 0 {
			close(solutions)
		}
	}
}

func validInput() error {
	dimension := len(table)
	if dimension < 1 {
		return errors.New("no data")
	}
	for _, row := range table {
		if len(row) != dimension {
			return errors.New("size mismatch")
		}
	}
	dimensionSqrt := math.Sqrt(float64(dimension))
	if math.Pow(dimensionSqrt, 2) != float64(dimension) {
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
	solution = make([][]int, len(table))
	for i, v := range table {
		solution[i] = make([]int, len(v))
		copy(solution[i], v)
	}
	return
}

func isValidValue(row, col, potentialValue int) (isValid bool) {
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
	dimension := int(math.Sqrt(float64(len(table))))
	boxColStart := (col / dimension) * dimension
	boxRowStart := (row / dimension) * dimension
	for boxRow := 0; boxRow < dimension; boxRow++ {
		for boxCol := 0; boxCol < dimension; boxCol++ {
			if table[boxRowStart+boxRow][boxColStart+boxCol] == potentialValue {
				return false
			}
		}
	}
	return true
}
