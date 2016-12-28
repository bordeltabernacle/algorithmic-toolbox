package main

import "testing"

type testCase struct {
	input    []int
	expected int
}

var testCases = []testCase{
	{[]int{1, 2, 3}, 6},
	{[]int{7, 5, 14, 2, 8, 8, 10, 1, 2, 3}, 140},
	{[]int{4, 6, 2, 6, 1}, 36},
	{[]int{100000, 90000}, 9000000000},
}

func TestMaxPairProd(t *testing.T) {
	for _, test := range testCases {
		actual := MaxPairwiseProduct(test.input)
		if actual != test.expected {
			t.Errorf("Max Pairwise Product test with: %v\n\texpected: %d\n\t  actual: %d", test.input, test.expected, actual)
		}
	}

}
