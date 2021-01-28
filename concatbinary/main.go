package concatbinary

import (
	"fmt"
	"math"
	"math/big"
)

func concatenatedBinary(n int) int {
	s := concatenatedBinaryHelper(n)
	//fmt.Println(s)
	i := new(big.Int)
	i.SetString(s, 2)
	//fmt.Println(i)
	modValue := big.NewInt(int64(int(math.Pow10(9)) + 7))
	result := i.Mod(i, modValue)
	/*
		for _, _ := range s {
			//fmt.Println(d)
		}
	*/
	// lookup formula
	finalResult := int(result.Int64())
	fmt.Println(finalResult)
	return finalResult
}

func concatenatedBinaryHelper(n int) string {
	s := ""
	for i := n; i > 0; i-- {
		sn := fmt.Sprintf("%b", i)
		s = sn + s

	}
	return s
}
