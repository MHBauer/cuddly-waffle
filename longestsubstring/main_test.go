package longestsubstring

import "testing"

func TestLongestSubstring(t *testing.T) {

	// longest string with k unique characters
	// xxyxxgehfls , k=2
	// xxyxx
	//
	//
	testCases := []struct {
		input  string
		k      int
		output string
	}{
		{"xxyxxesufntb", 2, "xxyxx"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			if longestSubstring(tc.input, tc.k) == tc.output {
				t.Error("")
			}
		})
	}

}
