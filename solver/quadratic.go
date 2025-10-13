package solver

import (
	"fmt"
	"math"
)

const epsilon = 1e-10

func Solve(a, b, c float64) ([]float64, error) {
	// проверка на специальные значения
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return nil, fmt.Errorf("коэффициенты не могут быть равны NaN")
	}

	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return nil, fmt.Errorf("коэффициенты не могут быть бесконечностью")
	}

	if math.Abs(a) <= epsilon {
		return nil, fmt.Errorf("коэффициент 'a' не может быть равен 0")
	}

	discriminant := b*b - 4*a*c

	if discriminant < -epsilon {
		return []float64{}, nil // нет действительных корней
	}

	if math.Abs(discriminant) <= epsilon {
		root := -b / (2 * a) // один корень
		return []float64{root}, nil
	}

	// два различных корня
	sqrtD := math.Sqrt(discriminant)
	root1 := (-b + sqrtD) / (2 * a)
	root2 := (-b - sqrtD) / (2 * a)

	return []float64{root1, root2}, nil
}
