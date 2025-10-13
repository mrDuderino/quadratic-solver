package main

import (
	"fmt"
	"quadratic-solver/solver"
)

func main() {
	fmt.Println(solver.Solve(1, 2, 1+1e-15))
}
