package main

import (
	"fmt"
	"math/big"
)

// HugeFibMod returns the last digit of the nth Fibonacci number
func HugeFibMod(n, m int) int {
	if n <= 1 {
		return n
	}
	prev := big.NewInt(0)
	curr := big.NewInt(1)
	for i := 0; i < n-1; i++ {
		prev.Add(prev, curr)
		prev, curr = curr, prev
	}
	last := curr.Mod(curr, big.NewInt(int64(m)))
	fmt.Println(last)
	return int(last.Int64())
}
