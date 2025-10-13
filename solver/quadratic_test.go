package solver

import (
	"math"
	"testing"
)

func TestSolveNoRoots(t *testing.T) {
	roots := Solve(1, 0, 1)
	if len(roots) != 0 {
		t.Errorf("Ожидалось отсутствие корней, но получили %v корней", roots)
	}
}

func TestSolveTwoRoots(t *testing.T) {
	roots := Solve(1, 0, -1)
	expectedResult := []float64{1, -1}

	if len(roots) != 2 {
		t.Errorf("Ожидалось 2 корня но получили %v корней", roots)
		return
	}

	if !((roots[0] == expectedResult[0] && roots[1] == expectedResult[1]) ||
		(roots[0] == expectedResult[1] && roots[1] == expectedResult[0])) {
		t.Errorf("Ожидаемый результат: %v, но получили: %v", expectedResult, roots)
	}
}

func TestSolveOneRoot(t *testing.T) {
	roots := Solve(1, 2, 1)

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
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Ожидаемая паника для a = 0")
		}
	}()

	Solve(0, 2, 1)
}

func TestSolveDiscriminantIsCloseToZero(t *testing.T) {
	roots := Solve(1, 2, 1+1e-15)

	if len(roots) != 1 {
		t.Errorf("Ожидался 1 корень для почти нулевого дискриминанта, но получили %v корней", roots)
	}
}
