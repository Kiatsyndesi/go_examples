package Moving_Zeros_To_The_End

func MoveZeros(arr []int) []int {
	nonZeros := []int{}
	zeros := []int{}

	for _, value := range arr {
		if value == 0 {
			zeros = append(zeros, value)
		} else {
			nonZeros = append(nonZeros, value)
		}
	}

	return append(nonZeros, zeros...)
}
