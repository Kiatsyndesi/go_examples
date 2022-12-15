package Sum_of_Digits__Digital_Root

func DigitalRoot(n int) int {
	var sum int

	for n >= 10 {
		sum = 0

		for n > 0 {
			sum += n % 10
			n /= 10
		}

		n = sum
	}

	return sum
}

/* Recursive Way
func DigitalRoot(n int) int {
	digits, sum := findAndSumDigits(n)

	if len(digits) == 1 {
		return sum
	} else {
		n = sum

		return DigitalRoot(n)
	}
}

func findAndSumDigits(n int) ([]int, int) {
	digits := []int{}
	digit := 0
	sum := 0

	for n > 0 {
		digit = n % 10
		sum += digit
		digits = append(digits, digit)

		n = n / 10
	}

	return digits, sum
}
*/
