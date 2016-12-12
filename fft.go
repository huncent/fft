package main

import "math"

func ZeroPadArray(x []float64, padToSize int) []float64 {
	if length(x) <= padToSize {
		return x
	}
	output = make(float64, padToSize)
	for i := 0; i < length(x); i++ {
		output[i] = x[i]
	}
	for j := length(x); j < padToSize; j++ {
		output[j] = 0.0
	}
	return output
}

func NextPowerOfTwo(x int64) {
	next = math.Pow(2, math.Ceil(math.Log2(x)/math.Log2(2)))
}

func DFFT(x []float64) {
	// if len(x) != power of 2 then pad
	if length(x)%2 == 0 {
		x = ZeroPadArray(x)
	}
	// TODO: the rest
}
