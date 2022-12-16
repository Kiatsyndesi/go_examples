package Grasshopper___Summation

func Summation(n int) int {
	start := 1
	result := 0

	for start <= n {
		result += start
		start++
	}

	return result
}