package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	ans := "2020"
	correct := "2025"

	userVal, err1 := strconv.ParseFloat(ans, 64)
	correctVal, err2 := strconv.ParseFloat(correct, 64)
	if err1 != nil || err2 != nil {
		panic("Invalid input")
	}

	diff := userVal - correctVal // positive means guess is higher, negative means lower
	//absDiff := math.Abs(diff)

	// Define the max difference for scaling
	const maxDiff = 100.0

	// Scale difference to -10..10, clamp within that range
	var scaledDiff float64

	if diff > 0 {
		scaledDiff = math.Min(diff/maxDiff*10, 10)
	} else {
		scaledDiff = math.Max(diff/maxDiff*10, -10)
	}

	fmt.Printf("Scaled difference (guess vs correct): %.2f\n", scaledDiff)
}
