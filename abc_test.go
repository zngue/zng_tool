package zng_tool

import (
	"fmt"
	"math/big"
	"testing"
)

func TestFloat(t *testing.T) {

	var slice = []float64{
		1,
		1,
		0.27,
		65.88,
		13.56,
		9.49,
		6.78,
		1.35,
		0.67,
	}
	var oldRate float64

	var rate = new(big.Float)
	for _, v := range slice {
		rate = rate.Add(rate, big.NewFloat(v))
		oldRate = oldRate + v
	}
	f, accuracy := rate.Float32()
	s := accuracy.String()
	fmt.Println(f, s, oldRate)

}
