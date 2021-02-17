package longestsubstring

// import "fmt"

// brute force implementation
func longestSubstring(input string, k int) string {

	var maxLength, start, end int
	
	for i:= 0; i < len (input); i++ {
		for j:= i+1; j < len (input); j++ {
			length := j - i 
			segment := input[i:j]
			//fmt.Println(segment)
			// count unique chars in segment
			count := countChars(segment)
			//fmt.Println(count)
			// if new result, save it
			if count <= k && length > maxLength {
				start = i
				end = j
			}
		}
	}
	return input[start:end]
}

func countChars(segment string) int {
	counts := map[rune]int {}
	for _, c := range segment {
		if _,ok := counts[c]; !ok {
			counts[c]=0
		}
		counts[c]++
	}
	return len(counts)
}
