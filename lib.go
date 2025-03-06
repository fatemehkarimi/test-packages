package matew

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// Add returns the sum of two int32 number
func Add[T Number](a T, b T) T {
	return a + b
}
