package main

import "testing"

type testCase struct {
	input, expected int
}

var naiveFibTestCases = []testCase{
	{3, 2},
	{10, 55},
}

var fibLastDigitTestCases = []testCase{
	{3, 2},
	{331, 9},
	{327305, 5},
}

func TestNaiveFib(t *testing.T) {
	for _, test := range naiveFibTestCases {
		actual := NaiveFib(test.input)
		if actual != test.expected {
			t.Errorf("Naive Fibonacci test with: %d\n\texpected: %d\n\t  actual: %d", test.input, test.expected, actual)
		}
	}

}

func TestFibLastDigit(t *testing.T) {
	for _, test := range fibLastDigitTestCases {
		actual := FibLastDigit(test.input)
		if actual != test.expected {
			t.Errorf("Fibonacci Last Digit test with: %d\n\texpected: %d\n\t  actual: %d", test.input, test.expected, actual)
		}
	}

}
