package main

import (
	"fmt"
	"math"
	"quadratic-solver/solver"
)

func main() {
	testCases := []struct {
		a, b, c float64
		desc    string
	}{
		{1, 0, 1, "x^2 + 1 = 0 (нет корней)"},
		{1, 0, -1, "x^2 - 1 = 0 (два корня)"},
		{1, 2, 1, "x^2 + 2x + 1 = 0 (один корень)"},
		{0, 2, 1, "0*x^2 + 2x + 1 = 0 (ошибка)"},
	}

	for _, tc := range testCases {
		fmt.Printf("\nУравнение: %s\n", tc.desc)
		roots, err := solver.Solve(tc.a, tc.b, tc.c)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Printf("Корни: %v\n", roots)
		}
	}

	// тест специальных значений
	fmt.Printf("\nСпециальные значения (NaN):\n")
	_, err := solver.Solve(math.NaN(), 1, 1)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}
}
