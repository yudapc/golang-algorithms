package sortimplementation

import "sort"

func SortInteger(values []int) []int {
	sort.Ints(values)
	return values
}
