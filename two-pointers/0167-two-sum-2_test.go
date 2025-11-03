package main

import (
	"math/rand"
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
			expected: []int{1, 2},
		},
		{
			name:     "empty_case",
			nums:     nil,
			target:   100,
			expected: []int{-1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("twoSum(%v, %d) = %v, want %v\n", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}

func BenchmarkTwoSumSmall(b *testing.B) {
	nums := generateSortedArray(100)
	l, r := generateTwoRandPointers(100)
	target := nums[l] + nums[r]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}

func BenchmarkTwoSumMedium(b *testing.B) {
	nums := generateSortedArray(10000)
	l, r := generateTwoRandPointers(10000)
	target := nums[l] + nums[r]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}

func BenchmarkTwoSumLarge(b *testing.B) {
	nums := generateSortedArray(1000000)
	l, r := generateTwoRandPointers(1000000)
	target := nums[l] + nums[r]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}

func generateTwoRandPointers(size int) (int, int) {
	return rand.Intn(size / 2), rand.Intn(size/2) + size/2
}
