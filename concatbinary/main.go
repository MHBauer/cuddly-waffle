package concatbinary

const mod int = 1_000_000_007
const upperLimit int = 10_000

// the trick is, we have these strings of bits available already, they're called 'numbers'
func concatenatedBinary(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		shift := numToShift(i)
		result = result << shift
		result = result | i // we know the bits are cleared, so logical-or them into the result
		//for result > mod {  // reduce if bigger. this is okay, becuase the biggest shift possible is for 10^5, which is less than 14 bits, with 10^9 being less than 20, 20+14, reduced by 20 bits, brings it back under. (19+ 14) = 33, -20 => 13.
		//	result = result - mod
		//}
		result = result % mod
	}
	return result
}

// the trick is, we have these strings of bits available already, they're called 'numbers'
func concatenatedBinarySubtraction(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		shift := numToShift(i)
		result = result << shift
		result = result | i // we know the bits are cleared, so logical-or them into the result
		for result > mod {  // reduce if bigger. this is okay, becuase the biggest shift possible is for 10^5, which is less than 14 bits, with 10^9 being less than 20, 20+14, reduced by 20 bits, brings it back under. (19+ 14) = 33, -20 => 13.
			result = result - mod
		}
	}
	return result
}

// the trick is, we have these strings of bits available already, they're called 'numbers'
func concatenatedBinaryConditionalMod(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		shift := numToShift(i)
		result = result << shift
		result = result | i // we know the bits are cleared, so logical-or them into the result
		if result > mod {   // reduce if bigger. this is okay, becuase the biggest shift possible is for 10^5, which is less than 14 bits, with 10^9 being less than 20, 20+14, reduced by 20 bits, brings it back under. (19+ 14) = 33, -20 => 13.
			result = result % mod
		}
	}
	return result
}

// the trick is, we have these strings of bits available already, they're called 'numbers'
func concatenatedBinaryAlwaysMod(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		shift := numToShift(i)
		result = result << shift
		result = result | i // we know the bits are cleared, so logical-or them into the result
		//for result > mod {  // reduce if bigger. this is okay, becuase the biggest shift possible is for 10^5, which is less than 14 bits, with 10^9 being less than 20, 20+14, reduced by 20 bits, brings it back under. (19+ 14) = 33, -20 => 13.
		//	result = result - mod
		//}
		result = result % mod
	}
	return result
}

// numToShift returns the position of the highest set bit. It is the
// amount of right shifts that would be needed to zero out the number.
func numToShift(n int) int {
	r := 0
	for ; n != 0; n = n >> 1 {
		r++
	}
	return r
}
