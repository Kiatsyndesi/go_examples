package Vowel_Count

import "strings"

func GetCount(str string) (count int) {
	vowels := "aeiou"

	for _, letter := range str {
		if strings.ContainsRune(vowels, letter) {
			count++
		}
	}

	return count
}
