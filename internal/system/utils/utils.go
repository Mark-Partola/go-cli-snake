package utils

import "math/rand/v2"

func Remainder(a, b int) int {
	return ((a % b) + b) % b
}

func FastRand(limit int) int {
	return rand.IntN(limit)
}
