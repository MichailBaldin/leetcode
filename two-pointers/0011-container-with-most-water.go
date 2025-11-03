// Pattern: Two Pointers
// Problem: Container With Most Water
// URL: https://leetcode.com/problems/container-with-most-water/
//
// Time: O(n)
// Space: O(1)

// Step-by-Step Algorithm:

// Initialization
// Set left pointer to start (index 0)
// Set right pointer to end (index n-1)
// Initialize maxArea to 0

// Two Pointer Traversal
// While left < right:
//   Calculate current container height = min(height[left], height[right])
//   Calculate current area = height * (right - left)
//   Update maxArea if current area is larger
//   If height[left] < height[right]:
//     Move left pointer forward
//   Else:
//     Move right pointer backward

// Termination
// Return maxArea

// Key Insights:
// Two-pointer efficiency: Single pass from both ends
// Greedy approach: Always move the shorter pointer to potentially find taller line
// Area calculation: Limited by shorter line and distance between lines
// No need to check all pairs: O(n) vs O(n²) brute force
// Width decreases: As pointers move inward, width decreases so we seek taller lines

package main

import "fmt"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maximumArea := 0

	for l < r {
		currMaxHeight := minHeight(height[l], height[r])
		currArea := currMaxHeight * (r - l)

		if currArea > maximumArea {
			maximumArea = currArea
		}

		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return maximumArea
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}

func minHeight(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// cpu: AMD Ryzen 7 PRO 3700 8-Core Processor
// BenchmarkMaxAreaSmall-16        17224658     73.55 ns/op         0 B/op     0 allocs/op
// BenchmarkMaxAreaMedium-16         186889      6276 ns/op         0 B/op     0 allocs/op
// BenchmarkMaxAreaLarge-16            1850    606888 ns/op         0 B/op     0 allocs/op
