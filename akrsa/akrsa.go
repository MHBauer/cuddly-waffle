package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

func main() {
	base := big.NewInt(2)
	expt := big.NewInt(4096)
	bitsize := big.NewInt(0)
	bitsize = bitsize.Exp(base, expt, nil)
	e := big.NewInt(0)
	d := big.NewInt(0)
	m := big.NewInt(0)
	x := big.NewInt(0)
	r := big.NewInt(0)
	c := big.NewInt(0)
	y1 := big.NewInt(0)
	y2 := big.NewInt(0)
	rng := rand.New(rand.NewSource(0))
	e.Rand(rng, bitsize)
	d.Rand(rng, bitsize)
	m.Rand(rng, bitsize)
	x.Rand(rng, bitsize)
	r.Mul(e, d)
	c.Mul(e, x)
	y1.Mul(c, d)
	y2.Mul(r, x)
	r.Mod(r, m)
	c.Mod(c, m)
	y1.Mod(y1, m)
	y2.Mod(y2, m)
	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y1.Cmp(y2))
}
