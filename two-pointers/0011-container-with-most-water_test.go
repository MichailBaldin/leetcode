package main

import (
	"math/rand"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "example_1",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "example_2",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "increasing_height",
			height:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 25, // min(1,10)*9 = 9 OR min(5,10)*5 = 25
		},
		{
			name:     "decreasing_height",
			height:   []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: 25, // min(10,1)*9 = 9 OR min(10,6)*5 = 25
		},
		{
			name:     "all_same_height",
			height:   []int{5, 5, 5, 5, 5},
			expected: 20, // 5 * 4 = 20
		},
		{
			name:     "valley_shape",
			height:   []int{1, 2, 1},
			expected: 2, // min(1,1)*2 = 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxArea(tt.height)
			if result != tt.expected {
				t.Errorf("maxArea(%v) = %d, want %d", tt.height, result, tt.expected)
			}
		})
	}
}

func BenchmarkMaxAreaSmall(b *testing.B) {
	height := generateHeightArray(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxArea(height)
	}
}

func BenchmarkMaxAreaMedium(b *testing.B) {
	height := generateHeightArray(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxArea(height)
	}
}

func BenchmarkMaxAreaLarge(b *testing.B) {
	height := generateHeightArray(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxArea(height)
	}
}

// Generate random height array with values between 0 and 1000
func generateHeightArray(size int) []int {
	height := make([]int, size)
	for i := range height {
		height[i] = rand.Intn(1001) // random heights from 0 to 1000
	}
	return height
}
