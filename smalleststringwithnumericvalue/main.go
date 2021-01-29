package smalleststringwithnumericvalue

import "fmt"

// sounds like greedy bin packing.
// we'll want to fill up on z's then move on, is it optimal?
// maybe recursive, memoizable, dp type problem. reminds me of coins from SICP.
// assumes such a string must exist. for example,  n = 5 k = 131, zzzzz is 130, not possible.

//

// strings up to 10000 long

func getSmallestString(n int, k int) string {
	chars := make([]byte, n)
	for fillLeft, spaceLeft := k, n; spaceLeft > 0; {
		charToAdd := 26
		fmt.Println("is ", fillLeft, "geq than", spaceLeft+26)
		if fillLeft >= spaceLeft+26 { // spaceleft + 26 represents the ability to add a z and fill with a
			fmt.Println("yes, add a z")

			// use z
			fillLeft -= 26
			// a z
		} else {
			// use fillLeft - spaceLeft
			charToAdd = (fillLeft - spaceLeft) + 1
			fillLeft -= charToAdd
			fmt.Println("no, add a ", string(rune(charToAdd+96)))
		}
		spaceLeft -= 1
		fmt.Println(charToAdd)
		chars[spaceLeft] = byte(charToAdd) + 96 // lowercases are +140?
	}
	return string(chars)
}
