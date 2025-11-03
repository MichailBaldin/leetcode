// Pattern: Two Pointers
// Problem: Valid Palindrome
// URL: https://leetcode.com/problems/valid-palindrome/
//
// Time: O(n)
// Space: O(1)

// Step-by-Step Algorithm:

// Initialization
// Set left pointer to start (index 0)
// Set right pointer to end (index n-1)

// Two Pointer Traversal
// While left < right:
//   Skip non-alphanumeric characters from left with bounds check
//   Skip non-alphanumeric characters from right with bounds check
//   If valid characters remain for comparison:
//     Convert both characters to lowercase
//     If characters are not equal, return false
//     Move left pointer forward and right pointer backward

// Termination
// Return true if all compared characters matched

// Key Insights:
// Two-pointer efficiency: Single pass from both ends
// Alphanumeric filtering: Skip spaces, punctuation and special characters
// Case insensitivity: Convert to lowercase before comparison
// Bounds checking: Essential when skipping characters to prevent index out of range
// Early termination: Return immediately when mismatch found
// Memory efficiency: No additional data structures required

package main

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlphanumericByte(s[l]) {
			l++
		}

		for l < r && !isAlphanumericByte(s[r]) {
			r--
		}

		if l < r {
			if toLower(s[l]) != toLower(s[r]) {
				return false
			}
			l++
			r--
		}
	}
	return true
}

func isAlphanumericByte(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

// cpu: AMD Ryzen 7 PRO 3700 8-Core Processor
// BenchmarkIsPalindromeSmall-16           36700718             35.25 ns/op               0 B/op          0 allocs/op
// BenchmarkIsPalindromeMedium-16             99675             12338 ns/op               0 B/op          0 allocs/op
// BenchmarkIsPalindromeLarge-16                876           1216402 ns/op               0 B/op          0 allocs/op
// BenchmarkIsPalindromeWorstCase-16           1658            718049 ns/op               0 B/op          0 allocs/op
