package main

import "testing"

type testCaseA struct {
	input, expected int
}

type testCaseB struct {
	input, mod, expected int
}

var naiveFibTestCases = []testCaseA{
	{3, 2},
	{10, 55},
}

var fibLastDigitTestCases = []testCaseA{
	{3, 2},
	{331, 9},
	{327305, 5},
}

var hugeFibModTestCases = []testCaseB{
	{1, 239, 1},
	{239, 1000, 161},
	// this is too slow
	// {2816213588, 30524, 10249},
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

func TestHugeFibMod(t *testing.T) {
	for _, test := range hugeFibModTestCases {
		actual := HugeFibMod(test.input, test.mod)
		if actual != test.expected {
			t.Errorf("Huge Fibonacci test with: %d and modulo: %d\n\texpected: %d\n\t  actual: %d", test.input, test.mod, test.expected, actual)
		}
	}
}
