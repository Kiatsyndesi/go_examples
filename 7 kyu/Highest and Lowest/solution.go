package Highest_and_Lowest

import (
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	result := ""
	sliceFromIn := strings.Split(in, " ")

	maxNumber, _ := strconv.Atoi(sliceFromIn[0])
	minNumber, _ := strconv.Atoi(sliceFromIn[0])

	for _, value := range sliceFromIn {
		currentValue, _ := strconv.Atoi(value)

		if currentValue > maxNumber {
			maxNumber = currentValue
		}

		if currentValue < minNumber {
			minNumber = currentValue
		}
	}

	result = strconv.Itoa(maxNumber) + " " + strconv.Itoa(minNumber)

	return result
}
