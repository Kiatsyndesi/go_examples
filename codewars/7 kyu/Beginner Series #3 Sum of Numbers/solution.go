package Beginner_Series__3_Sum_of_Numbers

func GetSum(a, b int) int {
	sum := 0

	if a > b {
		temp := a
		a = b
		b = temp
	}

	for a <= b {
		sum += a
		a++
	}

	return sum
}
