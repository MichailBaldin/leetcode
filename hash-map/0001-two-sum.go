// Pattern: Hash Map
// Problem: Two Sum
// URL: https://leetcode.com/problems/two-sum/
//
// Time:  O(n)
// Space: O(n)

// Step-by-Step Algorithm:

// Initialization
// Check for edge cases: if input array is nil or empty, return empty slice
// Create a hash map numMap with pre-allocated capacity equal to array length
// The map will store: number → index mappings

// Single Iteration
// Iterate through each number in array with its index
// For current number num at position i:
// Calculate complement = target - num
// Check if complement exists in hash map
// If found: return [map[complement], i] immediately
// If not found: store num → i in hash map

// Termination
// If no solution found after full iteration, return empty slice

// Key Insights:
// One-pass efficiency: Finds solution in single iteration without nested loops
// Hash map lookup: O(1) time complexity for complement checks
// Early termination: Returns immediately when solution is found
// Memory optimization: Pre-allocated map reduces dynamic allocations

package main

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	numMap := make(map[int]int, len(nums))

	for i, num := range nums {
		compliment := target - num
		if idx, exists := numMap[compliment]; exists {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return []int{}
}

// without memory optimization

// cpu: AMD Ryzen 7 PRO 3700 8-Core Processor
// BenchmarkTwoSumBig-16               1366   1001470 ns/op    591502 B/op    80 allocs/op
// BenchmarkTwoSumMiddle-16          135165      8516 ns/op      4472 B/op    10 allocs/op
// BenchmarkTwoSumSmall-16         21539698     65.59 ns/op        16 B/op     1 allocs/op

// with memory optimization
// make(map[int]int, len(nums))

// cpu: AMD Ryzen 7 PRO 3700 8-Core Processor
// BenchmarkTwoSumBig-16               3842    418962 ns/op    295570 B/op    34 allocs/op
// BenchmarkTwoSumMiddle-16          278400      4015 ns/op      2360 B/op     4 allocs/op
// BenchmarkTwoSumSmall-16         18143608     72.09 ns/op        16 B/op     1 allocs/op

// allocations decreased by 2 times
