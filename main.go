//go:build !codeanalysis
// +build !codeanalysis

package main

import "fmt"

var (
	floats = []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	ints   = []int64{1, 2, 3, 4, 5}
)

func sumInts(numbers []int64) int64 {
	var s int64
	for _, n := range numbers {
		s += n
	}
	return s
}

func sumFloats(numbers []float64) float64 {
	var s float64
	for _, n := range numbers {
		s += n
	}
	return s
}

// nosyntax
func sumIntsOrFloats[V int64 | float64](v []V) V {
	var s V
	for _, v := range v {
		s += v
	}
	return s
}

func main() {
	fmt.Println("sumInts :", sumInts(ints))
	fmt.Println("generic sumInts :", sumIntsOrFloats(ints))

	fmt.Println("sumFloats :", sumFloats(floats))
	fmt.Println("generic sumFloats :", sumIntsOrFloats(floats))
}
