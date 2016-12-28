package main

import "fmt"

func main() {
	fmt.Println(FibLastDigit(10))

}

// FibLastDigit returns the last digit of the nth Fibonacci number
func FibLastDigit(n int) int {
	for a, b := 0, 1; a < n; a, b = a+b, a {
	}
	return 0
}
