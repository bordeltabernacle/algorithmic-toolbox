package main

import "testing"

type testCase struct {
	a        int
	b        int
	expected int
}

var euclidGCDTestCases = []testCase{
	{18, 35, 1},
	{28851538, 1183019, 17657},
}

func TestEuclidGCD(t *testing.T) {
	for _, test := range euclidGCDTestCases {
		actual := EuclidGCD(test.a, test.b)
		if actual != test.expected {
			t.Errorf("Euclid GCD test with: %d & %d\n\texpected: %d\n\t  actual: %d", test.a, test.b, test.expected, actual)
		}
	}

}
