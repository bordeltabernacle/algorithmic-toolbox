package main

import "math/big"

// HugeFibMod returns the nth Fibonacci number modulo m
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
	mod := curr.Mod(curr, big.NewInt(int64(m)))
	return int(mod.Int64())
}
