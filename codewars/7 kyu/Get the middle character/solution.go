package Get_the_middle_character

import "strings"

func GetMiddle(s string) string {

	letters := strings.Split(s, "")

	if isEven(s) {
		return letters[len(letters)/2-1] + letters[len(letters)/2]
	} else {
		return letters[len(letters)/2]
	}
}

func isEven(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	return true
}
