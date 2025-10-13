package solver

import "math"

const epsilon = 1e-10

func Solve(a, b, c float64) []float64 {
	if a == 0 {
		panic("Коэффициент 'a' не может быть равен 0")
	}

	discriminant := b*b - 4*a*c

	if discriminant < -epsilon {
		return []float64{} // нет действительных корней
	}

	if math.Abs(discriminant) <= epsilon {
		root := -b / (2 * a) // один корень
		return []float64{root}
	}

	// два различных корня
	sqrtD := math.Sqrt(discriminant)
	root1 := (-b + sqrtD) / (2 * a)
	root2 := (-b - sqrtD) / (2 * a)

	return []float64{root1, root2}
}
