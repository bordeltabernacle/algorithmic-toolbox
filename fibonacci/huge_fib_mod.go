package main

import "math/big"

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
	mod := rem(curr, big.NewInt(int64(m)))
	return int(mod.Int64())
}

func rem(n, m *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Div(n, m)
	r.Mul(r, m)
	r.Sub(n, r)
	return r
}
