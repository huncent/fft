package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func ZeroPadArray(x []float64, padToSize int) []float64 {
	if len(x) >= padToSize {
		return x
	}
	output := make([]float64, padToSize)
	for i := 0; i < len(x); i++ {
		output[i] = x[i]
	}
	for j := len(x); j < padToSize; j++ {
		output[j] = 0.0
	}
	return output
}

// func NextPowerOfTwo(x int64) {
// 	next := math.Pow(2, math.Ceil(math.Log2(x)/math.Log2(2)))
// 	return
// }

func DFFT(x []float64, y []complex128, n, s int) {
	if n == 1 {
		y[0] = complex(x[0], 0)
		return
	}
	DFFT(x, y, n/2, 2*s)
	DFFT(x[s:], y[n/2:], n/2, 2*s)
	for i := 0; i < n/2; i++ {
		theta := -2.0 * math.Pi * float64(i) / float64(n)
		term := cmplx.Rect(1.0, theta) * y[i+n/2]
		temp := y[i]
		y[i] = temp + term
		y[i+n/2] = temp - term
	}
	return
}

func main() {
	x := []float64{math.Pi / 4, math.Pi / 2, math.Pi * 3 / 4, math.Pi}
	x = ZeroPadArray(x, 32)
	fmt.Println(x)
	y := make([]complex128, len(x))
	DFFT(x, y, len(x), 1)
	for _, val := range y {
		fmt.Println(val)
	}
}
