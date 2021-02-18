package longestsubstring

// brute force implementation, O(input^3)
func longestSubstring(input string, k int) string {

	var maxLength, start, end int

	for i := 0; i < len(input); i++ { // O(input)
		for j := i + 1; j < len(input); j++ { // O(input)
			length := j - i
			segment := input[i:j]
			// count unique chars in segment
			count := countChars(segment) // O(input)
			// if new result, save it
			if count <= k && length > maxLength {
				start, end = i, j
			}
		}
	}
	return input[start:end]
}

func countChars(segment string) int {
	counts := map[rune]int{}
	for _, c := range segment {
		if _, ok := counts[c]; !ok {
			counts[c] = 0
		}
		counts[c]++
	}
	return len(counts)
}
