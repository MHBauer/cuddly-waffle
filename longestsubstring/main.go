package longestsubstring

import "fmt"

// brute force implementation, O(input^3)
func longestSubstring(input string, k int) string {
	var maxLength, start, end int

	i := 0
	counts := map[byte]int{}
	for j := i; j < len(input); j++ {
		length := j - i
		// count unique chars in segment
		c := input[j]
		if _, ok := counts[c]; !ok {
			counts[c] = 0
		}
		counts[c]++
		count := len(counts)

		fmt.Println("adding", counts, input[i:j+1], count)
		// if new result, save it
		if count <= k && length > maxLength {
			fmt.Println("new result found", input[i:j+1])
			start, end = i, j
		}
		for count > k {
			// remove characters we add first
			c := input[i]
			counts[c]--
			if counts[c] == 0 {
				delete(counts, c)
			}
			i++
			count = len(counts)
			fmt.Println("removing", counts, input[i:j+1], count)
		}
	}
	return input[start:end]
}

func countChars(segment string) int {
	counts := map[rune]int{}
	for _, c := range segment { // O(input)
		if _, ok := counts[c]; !ok {
			counts[c] = 0
		}
		counts[c]++
	}
	return len(counts)
}
