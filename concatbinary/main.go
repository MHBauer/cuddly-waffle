package concatbinary

import (
	"fmt"
	"math"
	"math/big"
)

var upperLimit int = 100000 // 10^5 upperLimit
var intString map[int]string = make(map[int]string, upperLimit)

func concatenatedBinary(n int) int {
	if _, ok := intString[n]; !ok {
		initializeTable()
	}
	s := intString[n]
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

func initializeTable() {
	intString[1] = "1"
	for i := 2; i < upperLimit; i++ {
		s := ""
		sn := fmt.Sprintf("%b", i)
		s = sn + intString[i-1]
		intString[i] = s
	}
}
