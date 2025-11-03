// Pattern: Two Pointers
// Problem: Squares of a Sorted Array
// URL: https://leetcode.com/problems/squares-of-a-sorted-array/
//
// Time:  O(n)
// Space: O(n)

// Step-by-Step Algorithm:

// Initialization
// Check for edge cases: if input array is empty, return empty slice
// Initialize two pointers: left at start (0) and right at end (n-1)
// Create result array with same length as input
// Initialize insertion index at end of result array (n-1)

// Two Pointer Traversal
// While left pointer <= right pointer:
//   Calculate squares of elements at both pointers
//   Compare absolute values (or compare squares directly):
//   If left square >= right square:
//     Place left square at current insertion position
//     Move left pointer forward
//   Else:
//     Place right square at current insertion position
//     Move right pointer backward
//   Move insertion index backward

// Termination
// Return the fully populated result array

// Key Insights:
// Two-pointer efficiency: Processes array in single pass from both ends
// Sorted input utilization: Leverages the sorted property of input array
// Reverse filling: Builds result from largest to smallest squares
// No sorting required: Avoids O(n log n) sort operation
// Handles negatives gracefully: Negative numbers become positive when squared

package main

func sortedSquares(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	for i, num := range nums {
		nums[i] = num * num
	}

	l, r := 0, n-1
	result := make([]int, n)
	index := n - 1

	for l <= r {
		if nums[l] > nums[r] {
			result[index] = nums[l]
			l++
		} else {
			result[index] = nums[r]
			r--
		}
		index--
	}
	return result
}

// cpu: AMD Ryzen 7 PRO 3700 8-Core Processor
// BenchmarkSortedSquaresSmall-16   3378448     377.2 ns/op       896 B/op     1 allocs/op
// BenchmarkSortedSquaresMedium-16    39148     28240 ns/op     81920 B/op     1 allocs/op
// BenchmarkSortedSquaresLarge-16       382   2968951 ns/op   8003588 B/op     1 allocs/op
