package sort_numbers

import "sort"

func SortNumbers(numbers []int) []int {
	sort.Ints(numbers)

	return numbers
}
