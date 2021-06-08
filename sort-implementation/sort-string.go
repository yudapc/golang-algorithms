package sortimplementation

import "sort"

func SortString(values []string) []string {
	sort.Strings(values)
	return values
}
