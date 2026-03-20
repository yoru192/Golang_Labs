package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var y float64
	x := float64(rand.Int31n(1000)) // [0, 1000)

	if x > 100 {
		y = (math.Pow(x, 3) - 3) * (math.Pow(x, 2) + 3)
	} else {
		y = x/math.Pow(x, 2) - 4
	}

	fmt.Printf("x: %.0f\n", x)
	fmt.Printf("Result: %.2f\n", y)

	switch {
	case x > 100:
		y = (math.Pow(x, 3) - 3) * (math.Pow(x, 2) + 3)
	default:
		y = x/math.Pow(x, 2) - 4
	}

	fmt.Printf("x: %.0f\n", x)
	fmt.Printf("Result: %.2f\n", y)
}
