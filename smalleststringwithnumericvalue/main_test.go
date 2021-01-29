package smalleststringwithnumericvalue

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	n      int
	k      int
	answer string
}{
	{3, 27, "aay"},
	{5, 73, "aaszz"},
	{5, 31, "aaabz"},
}

func TestGetSmallestString(t *testing.T) {
	for _, tc := range testCases {
		result := getSmallestString(tc.n, tc.k)
		if result != tc.answer {
			t.Logf("%v should be %v", result, tc.answer)
			t.Fail()
		} else {
			t.Log(tc.answer, "ok")
		}
	}
}

func BenchmarkConcatBinarySubtraction(t *testing.B) {
	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(fmt.Sprint(tc.n), func(t *testing.B) {
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
			}
		})

	}
}
