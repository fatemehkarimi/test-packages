package matew

import (
	"golang.org/x/exp/constraints"
)

// Add returns the sum of two int32 number
func Add[T constraints.Float | constraints.Integer](a T, b T) T {
	return a + b
}
