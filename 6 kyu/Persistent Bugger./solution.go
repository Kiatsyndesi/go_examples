package Persistent_Bugger_

func Persistence(n int) int {
	numberOfDigits := countDigits(n)
	digits := []int{}
	product := 1
	counter := 0

	for numberOfDigits > 1 {
		digits = findDigits(n)

		for _, value := range digits {
			product *= value
		}

		n = product
		product = 1
		counter += 1

		numberOfDigits = countDigits(n)
	}

	return counter
}

func countDigits(n int) int {
	count := 0

	for n > 0 {
		n = n / 10
		count += 1
	}

	return count
}

func findDigits(n int) []int {
	digits := []int{}
	digit := 0

	for n > 0 {
		digit = n % 10
		digits = append(digits, digit)

		n = n / 10
	}

	return digits
}
