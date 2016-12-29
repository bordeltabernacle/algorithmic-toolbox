package main

// NaiveFib returns the nth Fibonacci number
func NaiveFib(n int) int {
	if n <= 1 {
		return n
	}
	return NaiveFib(n-1) + NaiveFib(n-2)
}
