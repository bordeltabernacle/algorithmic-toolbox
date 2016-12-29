package main

// EuclidGCD computes the greatest common divisor of two non-negative integers
func EuclidGCD(a int, b int) int {
	if b == 0 {
		return a
	}
	return EuclidGCD(b, a%b)
}
