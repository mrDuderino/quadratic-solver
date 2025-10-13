package quadraticSolver

import "math"

func Solve(a, b, c float64) []float64 {
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return []float64{}
	}

	if discriminant == 0 {
		root := -b / (2 * a)
		return []float64{root}
	}

	sqrtDiscriminant := math.Sqrt(discriminant)
	root1 := (-b + sqrtDiscriminant) / (2 * a)
	root2 := (-b - sqrtDiscriminant) / (2 * a)

	return []float64{root1, root2}
}
