package Interlocking_Binary_Pairs

import (
	"strconv"
	"strings"
)

func findMaxLen(a []string, b []string) int {
	if len(a) >= len(b) {
		return len(a)
	} else {
		return len(b)
	}
}

func Interlockable(a, b uint64) bool {
	binaryA := strings.Split(strconv.FormatUint(a, 2), "")
	binaryB := strings.Split(strconv.FormatUint(b, 2), "")

	for i := findMaxLen(binaryA, binaryB); i >= 0; i-- {
		if binaryA[i]
	}
}
