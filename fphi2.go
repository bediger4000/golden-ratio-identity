package main

import (
	"fmt"
	"math"
)

func main() {
	phi := math.Phi
	fmt.Printf("Original phi := %e\n", phi)
	epsilon := machineEpsilon()

	phiCubed := phi * phi * phi
	numerator := math.Sqrt(phiCubed)
	denominator := math.Sqrt(phiCubed)
	newPhi := numerator / denominator
	i := 0
	for math.Abs(newPhi-phi) > epsilon {
		numerator = math.Sqrt(phiCubed + numerator)
		denominator = math.Sqrt(phiCubed - denominator)
		newPhi = numerator / denominator
		i++
		fmt.Printf("%d\t%e\t%e\n", i, newPhi, newPhi-phi)
	}
}

func machineEpsilon() float64 {
	e := 1.0
	for (1.0 + e) > 1.0 {
		e /= 2.0
	}
	// fmt.Printf("machine epsilon %g\n", e*2.0)
	return e * 2.0
}
