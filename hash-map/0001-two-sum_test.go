package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "general_case",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "not_solution",
			nums:     []int{1, 2, 3},
			target:   7,
			expected: []int{},
		},
		{
			name:     "empty_slice",
			nums:     []int{},
			target:   9,
			expected: []int{},
		},
		{
			name:     "nil_slice",
			nums:     nil,
			target:   100,
			expected: []int{},
		},
		{
			name:     "solution_at_beggining",
			nums:     []int{3, 2, 4},
			target:   5,
			expected: []int{0, 1},
		},
		{
			name:     "solution_at_end",
			nums:     []int{3, 2, 4},
			target:   5,
			expected: []int{0, 1},
		},
		{
			name:     "two_elements_only",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "mixed",
			nums:     []int{-3, 4, 3, 90},
			target:   0,
			expected: []int{0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}

func BenchmarkTwoSumBig(b *testing.B) {
	largeNums := make([]int, 10000)
	for i := range largeNums {
		largeNums[i] = i
	}
	target := 9998 + 9999

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum(largeNums, target)
	}
}

func BenchmarkTwoSumMiddle(b *testing.B) {
	largeNums := make([]int, 100)
	for i := range largeNums {
		largeNums[i] = i
	}
	target := 98 + 99

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum(largeNums, target)
	}
}

func BenchmarkTwoSumSmall(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}
