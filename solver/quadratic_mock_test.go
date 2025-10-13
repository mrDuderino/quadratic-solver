package solver

import "testing"

func TestSolverWithMock(t *testing.T) {
	mockSolver := &MockSolver{
		SolveFunc: func(a, b, c float64) ([]float64, error) {
			if a == 1 && b == 0 && c == -1 {
				return []float64{1, -1}, nil
			}
			return []float64{}, nil
		},
	}

	roots, err := mockSolver.Solve(1, 0, -1)
	if err != nil {
		t.Errorf("Неожиданная ошибка: %v", err)
	}

	if len(roots) != 2 {
		t.Errorf("Мок вернул неожиданные значения: %v", roots)
	}

	if !((roots[0] == 1 && roots[1] == -1) || (roots[0] == -1 && roots[1] == 1)) {
		t.Errorf("Мок вернул неожиданные значения: %v", roots)
	}
}
