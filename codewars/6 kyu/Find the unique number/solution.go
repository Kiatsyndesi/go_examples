package Find_the_unique_number

func FindUniq(arr []float32) float32 {
	// Do the magic
	countOfValues := make(map[float32]float32)
	result := float32(0)

	for _, value := range arr {
		countOfValues[value] += 1.0
	}

	for key, value := range countOfValues {
		if value == 1.0 {
			result = key
		}
	}

	return result
}