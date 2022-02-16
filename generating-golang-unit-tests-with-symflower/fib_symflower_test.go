package fib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSymflowerFastFibonacci1(t *testing.T) {
	n := int64(-1)
	actual := FastFibonacci(n)
	expected := int64(-1)
	assert.Equal(t, expected, actual)
}
func TestSymflowerFastFibonacci2(t *testing.T) {
	n := int64(0)
	actual := FastFibonacci(n)
	expected := int64(0)
	assert.Equal(t, expected, actual)
}
func TestSymflowerFastFibonacci3(t *testing.T) {
	n := int64(1)
	actual := FastFibonacci(n)
	expected := int64(1)
	assert.Equal(t, expected, actual)
}
func TestSymflowerFastFibonacci4(t *testing.T) {
	n := int64(2)
	actual := FastFibonacci(n)
	expected := int64(1)
	assert.Equal(t, expected, actual)
}
func TestSymflowerFastFibonacci5(t *testing.T) {
	n := int64(9223372036854775807)
	actual := FastFibonacci(n)
	expected := int64(-1)
	assert.Equal(t, expected, actual)
}
func TestSymflowerNaiveFibonacci6(t *testing.T) {
	n := int64(-1)
	actual := NaiveFibonacci(n)
	expected := int64(-1)
	assert.Equal(t, expected, actual)
}
func TestSymflowerNaiveFibonacci7(t *testing.T) {
	n := int64(0)
	actual := NaiveFibonacci(n)
	expected := int64(0)
	assert.Equal(t, expected, actual)
}
func TestSymflowerNaiveFibonacci8(t *testing.T) {
	n := int64(1)
	actual := NaiveFibonacci(n)
	expected := int64(1)
	assert.Equal(t, expected, actual)
}
func TestSymflowerNaiveFibonacci9(t *testing.T) {
	n := int64(2)
	actual := NaiveFibonacci(n)
	expected := int64(1)
	assert.Equal(t, expected, actual)
}
