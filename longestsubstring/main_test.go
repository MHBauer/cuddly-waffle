package longestsubstring

import (
	"fmt"
	"testing"
)

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
		{"xxyxxesufntb", 1, "xx"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprint(tc.input, tc.k), func(t *testing.T) {
			if longestSubstring(tc.input, tc.k) == tc.output {
				t.Error("")
			}
		})
	}

}
