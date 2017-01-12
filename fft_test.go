package main

import (
	"fmt"
	"testing"
)

func TestFFT(t *testing.T) {
	x := []float64{94, 94, 112, 112, 112, 94, 94, 94, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 130, 112, 112, 112, 130, 130, 130, 130, 130, 130, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 166, 148, 148, 148, 166, 166, 166, 148, 148, 148, 166, 166, 166, 166, 166, 166, 166, 166, 184, 166, 166, 166, 166, 166, 166, 166, 184, 166, 166, 166, 184, 184, 184, 184, 184, 166, 166, 166, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 166, 166, 184, 220, 184, 166, 184, 202, 184, 166, 166, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 166, 166, 166, 184, 166, 166, 166, 166, 166, 166, 166, 166, 148, 148, 148, 166, 166, 166, 166, 166, 148, 148, 148, 166, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 130, 130, 130, 148, 148, 148, 148, 148, 130, 130, 130, 148, 130, 130, 130, 130, 130, 130, 112, 112, 112, 112, 112, 112, 112, 112, 112, 130, 112, 112, 112, 112, 94, 94, 94, 112, 112, 112, 112, 112, 112, 112, 112, 112, 94, 94, 94, 112, 94, 94, 94, 94, 94, 112, 94, 94, 94, 94, 94, 112, 94, 94, 94, 112, 94, 94, 94, 94, 94, 94, 94, 94, 76, 76, 76, 76, 76, 76, 58, 58, 58, 76, 58, 58, 58, 58, 58, 58, 40, 40, 40, 40, 40, 40, 40, 58, 40, 40, 40, 58, 58, 76, 40, 22, 40, 76, 58, 40, 40, 40, 40, 58, 40, 40, 40, 58, 40, 40, 40, 58, 40, 40, 40, 58, 58, 58, 40, 40, 40, 58, 40, 40, 40, 58, 58, 58, 58, 58, 58, 58, 58, 76, 58, 58, 58, 76, 58, 58, 58, 76, 76, 76, 76, 76, 58, 58, 58, 76, 58, 58, 58, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 94, 76, 76, 76, 76, 76, 76, 76, 94, 94, 94, 94, 94, 76, 76, 76, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 130, 112, 112, 112, 130, 112, 112, 112, 130, 130, 130, 130, 148, 130, 130, 130, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 166, 202, 166, 148, 166, 184, 166, 166, 166, 184, 166, 166, 166, 166, 166, 184, 166, 166, 166, 184, 184, 184, 166, 166, 166, 166, 166, 184, 184, 184, 166, 166, 166, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 166, 166, 166, 184, 166, 166, 166, 166, 166, 184, 166, 166, 166, 184, 166, 166, 166, 184, 166, 166, 166, 184, 184, 184, 184, 184, 184, 184, 184, 184, 166, 166, 166, 184, 184, 184, 184, 184, 166, 166, 166, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 166, 166, 166, 184, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 148, 148, 148, 166, 166, 166, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 130, 130, 130, 148, 130, 112, 112, 130, 130, 130, 130, 130, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 76, 76, 76, 94, 94, 94, 94, 94, 94, 94, 76, 76, 76, 94, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 76, 58, 58, 58, 76, 76, 76, 76, 76, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 76, 58, 58, 58, 76, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 40, 40, 40, 58, 58, 58, 58, 58, 58, 58, 40, 40, 40, 58, 40, 40, 58, 76, 58, 58, 58, 58, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 76, 76, 76, 58, 58, 58, 76, 58, 58, 58, 76, 58, 58, 58, 58, 58, 76, 76, 94, 76, 58, 76, 94, 76, 76, 94, 112, 112, 112, 94, 94, 94, 112, 94, 94, 94, 112, 94, 94, 112, 130, 112, 112, 94, 94, 94, 94, 94, 94, 94, 112, 112, 112, 112, 112, 112, 130, 112, 112, 112, 130, 112, 112, 112, 130, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 130, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 112, 130, 130, 130, 130, 130, 112, 112, 112, 130, 112, 112, 112, 130, 130, 130, 130, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 166, 166, 166, 148, 148, 148, 166, 166, 166, 166, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 202, 184, 184, 184, 184, 184, 202, 202, 202, 184, 184, 184, 202, 184, 184, 184, 202, 202, 202, 202, 220, 202, 184, 184, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 202, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 184, 166, 166, 166, 166, 166, 166, 166, 166, 166, 166, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 148, 130, 130, 130, 148, 130, 130, 130, 148, 130, 130, 130, 148, 130, 130, 130, 148, 130, 130, 130}
	x = ZeroPadArray(x, 1024)
	fmt.Println(x)
	y := make([]complex128, len(x))
	DFFT(x, y, len(x), 1)
	for _, val := range y {
		fmt.Println(val)
	}
}
