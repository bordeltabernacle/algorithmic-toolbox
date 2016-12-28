package main

import "testing"

type testCase struct {
	input, expected int
}

var testCases = []testCase{
	{3, 2},
	{331, 9},
	{327305, 5},
}

func TestMaxPairProd(t *testing.T) {
	for _, test := range testCases {
		actual := FibLastDigit(test.input)
		if actual != test.expected {
			t.Errorf("Max Pairwise Product test with: %d\n\texpected: %d\n\t  actual: %d", test.input, test.expected, actual)
		}
	}

}
