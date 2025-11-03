// Pattern: Two Pointers
// Problem: Two Sum II - Input Array Is Sorted
// URL: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
//
// Time:  O(n)
// Space: O(1)

// Step-by-Step Algorithm:

// Initialization
// Set left pointer to start (index 0)
// Set right pointer to end (index n-1)

// Two Pointer Traversal
// While left pointer <= right pointer:
//   Calculate current sum = numbers[left] + numbers[right]
//   If sum equals target:
//     Return [left+1, right+1] (1-indexed indices)
//   Else if sum > target:
//     Move right pointer backward (decrease sum)
//   Else (sum < target):
//     Move left pointer forward (increase sum)

// Termination
// Return [-1, -1] if no solution found (though solution is guaranteed)

// Key Insights:
// Sorted array advantage: Enables efficient two-pointer approach
// Single pass: O(n) time complexity
// Constant space: Only two pointers used, no extra data structures
// 1-indexed result: Add 1 to indices before returning
// Guaranteed solution: Problem states solution always exists
// Directional movement: Move pointers based on sum comparison to target

package main

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	l, r := 0, n-1

	for l <= r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}
	}
	return []int{-1, -1}
}

// cpu: AMD Ryzen 7 PRO 3700 8-Core Processor
// BenchmarkTwoSumSmall-16         16565154     76.77 ns/op         0 B/op     0 allocs/op
// BenchmarkTwoSumMedium-16          140874      7111 ns/op         0 B/op     0 allocs/op
// BenchmarkTwoSumLarge-16             1590    713635 ns/op         0 B/op     0 allocs/op
