package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "general_case",
			nums:     []int{-3, -2, 0, 1, 3, 5},
			expected: []int{0, 1, 4, 9, 9, 25},
		},
		{
			name:     "empty_slice",
			nums:     nil,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sortedSquares(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sortedSquares(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func BenchmarkSortedSquaresSmall(b *testing.B) {
	nums := generateSortedArray(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sortedSquares(nums)
	}
}

func BenchmarkSortedSquaresMedium(b *testing.B) {
	nums := generateSortedArray(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sortedSquares(nums)
	}
}

func BenchmarkSortedSquaresLarge(b *testing.B) {
	nums := generateSortedArray(1000000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sortedSquares(nums)
	}
}

// generate random numbers between -1000 and 1000
func generateSortedArray(size int) []int {
	nums := make([]int, size)

	for i := range nums {
		nums[i] = rand.Intn(2000) - 1000
	}
	return nums
}
