package quadratic_solver

import "testing"

func TestSolveNoRoots(t *testing.T) {
	roots := Solve(1, 0, 1)
	if len(roots) != 0 {
		t.Errorf("Ожидалось отсутствие корней, но получили %v корней", roots)
	}
}
