package main

import "testing"

type testCase struct {
	a        int
	b        int
	expected int
}

var LCMTestCases = []testCase{
	{6, 8, 24},
	{28851538, 1183019, 1933053046},
}

func TestLCM(t *testing.T) {
	for _, test := range LCMTestCases {
		actual := LCM(test.a, test.b)
		if actual != test.expected {
			t.Errorf("Lowest Common Multiple test with: %d & %d\n\texpected: %d\n\t  actual: %d", test.a, test.b, test.expected, actual)
		}
	}

}
