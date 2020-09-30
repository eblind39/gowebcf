package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	a := scanMeasure("acceleration")
	v0 := scanMeasure("initial velocity")
	s0 := scanMeasure("initial displacement")

	S := GenDisplaceFn(a, v0, s0)

	fmt.Println("Displacement: ", S(3))
	fmt.Println("Displacement: ", S(5))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
	return fn
}

/*
 * read a float64
 */
func scanMeasure(strMeasure string) float64 {
	var strNumber string

	fmt.Println("Please enter", strMeasure)
	// "infinite" loop to read the float64
	for {
		fmt.Scan(&strNumber)

		// attemp to convert the input as an float64
		fltTmp, err := strconv.ParseFloat(strNumber, 64)
		if err != nil {
			// error on convertion, it's not a float64
			fmt.Println("Enter a valid", strMeasure)
		} else {
			return fltTmp
		}
	}
}
