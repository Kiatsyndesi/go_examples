package Bit_Counting

import (
	"strconv"
	"strings"
)

func CountBits(number uint) int {
	binaryNumber := strings.Split(strconv.FormatUint(uint64(number), 2), "")
	result := 0

	for _, value := range binaryNumber {
		if value == "1" {
			result++
		}
	}

	return result
}
