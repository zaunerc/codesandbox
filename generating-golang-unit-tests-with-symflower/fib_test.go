// file fib_test.go
package fib

import "testing"

var tests = []struct {
	input    int64
	expected int64
}{
	{
		0, 0,
	},
	{
		1, 1,
	},
	{
		2, 1,
	},
	{
		3, 2,
	},
	{
		4, 3,
	},
	{
		5, 5,
	},
	{
		13, 233,
	},
}

func TestFastFibonacci(t *testing.T) {
	for _, test := range tests {
		actual := FastFibonacci(test.input)
		if test.expected != actual {
			t.Fatalf("expected %d got %d", test.expected, actual)
		}
	}
}

func TestNaiveFibonacci(t *testing.T) {
	for _, test := range tests {
		actual := NaiveFibonacci(test.input)
		if test.expected != actual {
			t.Fatalf("expected %d got %d", test.expected, actual)
		}
	}

}

// TestBothFunctions implements a simple stress test
func TestBothFunctions(t *testing.T) {
	for i := int64(0); i < 20; i++ {
		naiveFibResult := NaiveFibonacci(i)
		fastFibResult := FastFibonacci(i)
		if naiveFibResult != fastFibResult {
			t.Fatalf("different results for fibonacci number with index %d", i)
		}
	}
}
