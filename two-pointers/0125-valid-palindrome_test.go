package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "valid_palindrome",
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "invalid_palindrome",
			s:        "race a car",
			expected: false,
		},
		{
			name:     "single_space",
			s:        " ",
			expected: true,
		},
		{
			name:     "all_non_alphanumeric",
			s:        ".,!",
			expected: true,
		},
		{
			name:     "single_letter",
			s:        "a",
			expected: true,
		},
		{
			name:     "single_number",
			s:        "1",
			expected: true,
		},
		{
			name:     "with_numbers",
			s:        "0P0",
			expected: true, // "0p0" vs "0P0" - разные в lowercase
		},
		{
			name:     "only_numbers_palindrome",
			s:        "12321",
			expected: true,
		},
		{
			name:     "only_numbers_not_palindrome",
			s:        "12345",
			expected: false,
		},
		{
			name:     "mixed_case_palindrome",
			s:        "Aa",
			expected: true,
		},
		{
			name:     "long_string_single_alpha",
			s:        "......a......",
			expected: true,
		},
		{
			name:     "unicode_chars",
			s:        "Привет, 世界!",
			expected: true, // не палиндром после фильтрации
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.s)
			if result != tt.expected {
				t.Errorf("isPalindrome(%q) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}

func BenchmarkIsPalindromeSmall(b *testing.B) {
	s := "A man, a plan, a canal: Panama"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isPalindrome(s)
	}
}

func BenchmarkIsPalindromeMedium(b *testing.B) {
	s := generatePalindromeString(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isPalindrome(s)
	}
}

func BenchmarkIsPalindromeLarge(b *testing.B) {
	s := generatePalindromeString(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isPalindrome(s)
	}
}

func BenchmarkIsPalindromeWorstCase(b *testing.B) {
	s := generateNonAlphanumericString(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isPalindrome(s)
	}
}

func generateNonAlphanumericString(size int) string {
	bytes := make([]byte, size)
	symbols := []byte{'.', ',', ' ', '!', '?', ';', ':', '-', '_', '(', ')', '[', ']', '{', '}'}

	for i := range bytes {
		bytes[i] = symbols[i%len(symbols)]
	}
	return string(bytes)
}

func generatePalindromeString(size int) string {
	if size < 1 {
		return ""
	}

	bytes := make([]byte, size)

	for i := 0; i < size/2; i++ {
		// Use alphanumeric characters only
		switch i % 3 {
		case 0:
			bytes[i] = byte('a' + i%26)
			bytes[size-1-i] = bytes[i]
		case 1:
			bytes[i] = byte('A' + i%26)
			bytes[size-1-i] = bytes[i]
		case 2:
			bytes[i] = byte('0' + i%10)
			bytes[size-1-i] = bytes[i]
		}
	}

	// Handle middle character for odd length
	if size%2 == 1 {
		middle := size / 2
		bytes[middle] = 'm'
	}

	return string(bytes)
}
