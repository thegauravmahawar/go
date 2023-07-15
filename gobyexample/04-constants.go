package main

import (
	"fmt"
	"math"
)

// const declares a constant value.
const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 500000000

	// Constant expressions perform arithmetic with arbitrary precision.
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it's given one.
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
