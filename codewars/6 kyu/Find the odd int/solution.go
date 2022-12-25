package Find_the_odd_int

func FindOdd(seq []int) int {
	counters := make(map[int]int)
	result := 0

	for _, value := range seq {
		counters[value] += 1
	}

	for key, value := range counters {
		if value%2 != 0 {
			result = key
		}
	}
	return result
}
