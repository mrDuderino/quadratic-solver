package solver

import (
	"math"
	"testing"
)

func TestSolveNoRoots(t *testing.T) {
	roots, err := Solve(1, 0, 1)

	if err != nil {
		t.Errorf("Получили ошибку: %v", err)
		return
	}

	if len(roots) != 0 {
		t.Errorf("Ожидалось отсутствие корней, но получили %v корней", roots)
	}
}

func TestSolveTwoRoots(t *testing.T) {
	roots, err := Solve(1, 0, -1)

	if err != nil {
		t.Errorf("Получили ошибку: %v", err)
		return
	}

	if len(roots) != 2 {
		t.Errorf("Ожидалось 2 корня но получили %v корней", roots)
		return
	}

	expectedResult := []float64{1, -1}
	if !((roots[0] == expectedResult[0] && roots[1] == expectedResult[1]) ||
		(roots[0] == expectedResult[1] && roots[1] == expectedResult[0])) {
		t.Errorf("Ожидаемый результат: %v, но получили: %v", expectedResult, roots)
	}
}

func TestSolveOneRoot(t *testing.T) {
	roots, err := Solve(1, 2, 1)

	if err != nil {
		t.Errorf("Получили ошибку: %v", err)
		return
	}
	if len(roots) != 1 {
		t.Errorf("Ожидался 1 корень, но получили %v корней", roots)
		return
	}

	expectedResult := []float64{-1}
	if math.Abs(roots[0]-expectedResult[0]) > 1e-10 {
		t.Errorf("Ожидался корень %v, но получили: %v", expectedResult, roots)
	}
}

func TestSolveZeroCoeffA(t *testing.T) {
	_, err := Solve(0, 2, 1)
	if err == nil {
		t.Errorf("Ожидаемая ошибка для a = 0")
	}
}

func TestSolveDiscriminantIsCloseToZero(t *testing.T) {
	roots, err := Solve(1, 2, 1+1e-15)

	if err != nil {
		t.Errorf("Получили ошибку: %v", err)
	}

	if len(roots) != 1 {
		t.Errorf("Ожидался 1 корень для почти нулевого дискриминанта, но получили %v корней", roots)
	}
}

func TestSolveNonNumericalCoeffs(t *testing.T) {
	testCases := []struct {
		name        string
		a, b, c     float64
		shouldError bool
	}{
		{"NaN a", math.NaN(), 1, 1, true},
		{"NaN b", 1, math.NaN(), 1, true},
		{"NaN c", 1, 1, math.NaN(), true},
		{"Positive Infinity a", math.Inf(1), 1, 1, true},
		{"Negative Infinity a", math.Inf(-1), 1, 1, true},
		{"Positive Infinity b", 1, math.Inf(1), 1, true},
		{"Negative Infinity b", 1, math.Inf(-1), 1, true},
		{"Positive Infinity c", 1, 1, math.Inf(1), true},
		{"Negative Infinity c", 1, 1, math.Inf(-1), true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Solve(tc.a, tc.b, tc.c)
			if tc.shouldError && err == nil {
				t.Errorf("Ожидалась ошибка для %s", tc.name)
			}
			if !tc.shouldError && err != nil {
				t.Errorf("Неожиданная ошибка для %s: %v", tc.name, err)
			}
		})
	}
}
