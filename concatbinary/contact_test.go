package concatbinary

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	n               int
	binaryExpansion string
	answer          int
}{
	{1, "1", 1},
	// {2, "110", 6},
	// {3, "11011", 27},
	// {4, "11011100", 220},
	// {5, "", 1765},
	// {6, "", 14126},
	// {7, "", 113_015},
	// {8, "", 1_808_248},
	// {9, "", 28_931_977},
	// {10, "", 462_911_642},
	// {11, "", 406_586_234},
	{12, "1101110010111011110001001101010111100", 505379714},
	{1000, "", 499361981},
	{8552, "", 65181209},
	{100000, "", 757631812},
}

func TestShfitCount(t *testing.T) {
	testCases := []struct {
		n      int
		answer int
	}{
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 3},
		{5, 3},
		{1024, 11},
	}
	for _, tc := range testCases {
		result := numToShift(tc.n)
		if result != tc.answer {
			t.Logf("%v should be %v", result, tc.answer)
			t.Fail()
		}
	}
}

func TestConcatBinary(t *testing.T) {
	for _, tc := range testCases {
		result := concatenatedBinary(tc.n)
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
				concatenatedBinarySubtraction(tc.n)
			}
		})

	}
}

func BenchmarkConcatBinaryAlwaysMod(t *testing.B) {
	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(fmt.Sprint(tc.n), func(t *testing.B) {
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
				concatenatedBinaryAlwaysMod(tc.n)
			}
		})

	}
}

func BenchmarkConcatBinaryConditionalMod(t *testing.B) {
	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(fmt.Sprint(tc.n), func(t *testing.B) {
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
				concatenatedBinaryConditionalMod(tc.n)
			}
		})
	}
}

// n=1 is 1
// n=2 is 1,10 -> 110 -> 2+4 -> 6
// n=3 is 1,10,11 -> 11011 -> 1+2+8+16 => 27
// n=4 is 1,10,11,100 -> 11011100 -> 4+8+16+64+128 => 220
