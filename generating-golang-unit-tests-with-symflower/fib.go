// file fib.fio
package fib

func NaiveFibonacci(n int64) int64 {

	// fib(-1), fib(-2), fib(-3), ... is undefined
	if n < 0 {
		return -1
	}

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a := NaiveFibonacci(n - 1)
	b := NaiveFibonacci(n - 2)
	return a + b
}

func FastFibonacci(n int64) int64 {

	// If n >= 9223372036854775807 (math.MaxInt)
	// this would lead to an index out of bounds error later on
	// when accessing the numbers helper array.
	if n >= 9223372036854775807 {
		return -1
	}

	// fib(-1), fib(-2), fib(-3), ... is undefined
	if n < 0 {
		return -1
	}

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	numbers := make([]int64, n+1)

	numbers[0] = 0
	numbers[1] = 1

	for i := 2; i < len(numbers); i++ {
		numbers[i] = numbers[i-1] + numbers[i-2]
	}

	return numbers[n]
}
