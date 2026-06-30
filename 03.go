package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	// charLastIndex tracks the last seen index of each character.
	charLastIndex := make(map[rune]int)
	longest := 0
	left := 0

	for right, char := range s {
		// If the character is inside our current window, we move the left
		// pointer past the last occurrence of the duplicate character.
		if lastIdx, ok := charLastIndex[char]; ok && lastIdx >= left {
			left = lastIdx + 1
		}

		// Update the last seen index of the current character.
		charLastIndex[char] = right

		// Calculate window size and update longest if it's larger.
		windowSize := right - left + 1 // the +1 is to count the actual char, since we start at 0
		if windowSize > longest {
			longest = windowSize
		}
	}

	return longest
}

func main3() {
	fmt.Printf("Longest substring of 'dvdf': %d\n", lengthOfLongestSubstring("dvdf"))
	fmt.Printf("Longest substring of 'abcabcbb': %d\n", lengthOfLongestSubstring("abcabcbb"))
}
