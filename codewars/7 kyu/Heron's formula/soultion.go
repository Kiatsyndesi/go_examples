package Heron_s_formula

import "math"

func Heron(a, b, c float64) (area float64) {
	s := (a + b + c) / 2

	area = math.Sqrt(s * (s - a) * (s - b) * (s - c))

	return area
}
