package solver

type Solver interface {
	Solve(a, b, c float64) ([]float64, error)
}

type MockSolver struct {
	SolveFunc func(a, b, c float64) ([]float64, error)
}

func (m *MockSolver) Solve(a, b, c float64) ([]float64, error) {
	if m.SolveFunc != nil {
		return m.SolveFunc(a, b, c)
	}
	return []float64{}, nil
}
