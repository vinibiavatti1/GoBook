package main

import "fmt"

func main() {
	// Initialize a map for the integer values
	ints := map[string]Numero{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Generic Sums: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

type Numero int64

type Number interface {
	~int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
