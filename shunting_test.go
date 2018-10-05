package sidecar

import (
	"fmt"
	"math"
	"testing"
)

func TestIt(t *testing.T) {
	data := []struct {
		input  string
		result float64
	}{
		{
			"2 + 2",
			4,
		},
		{
			"2 - 1",
			1,
		},
		{
			"2 * 3",
			6,
		},
		{
			"5 / 1",
			5,
		},
		{
			"1 + 2 * 3 + 4",
			11,
		},
		{
			"1 + 2 * ( 3 + 4 )",
			15,
		},

		{
			"1 + -2 * ( 3 + 4 )",
			-13,
		},
		{
			"3 * 8 / 3 / 8",
			1,
		},
		{
			"8 / ( 3 - ( 8 / 3 ) )",
			24,
		},
	}
	for _, v := range data {
		result := Calc(v.input)
		fmt.Println(result, v.result, result-v.result)
		if math.Abs(result-v.result) > 0.000001 {
			t.Errorf("Expected %f for %s, got %f", v.result, v.input, result)
		}
	}
}
