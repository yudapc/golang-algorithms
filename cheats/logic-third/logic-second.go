package logicthird

import "strconv"

func countNumberTimes(numbers []int) map[string]int {
	result := make(map[string]int)
	for _, v := range numbers {
		result[strconv.Itoa(v)] += 1
	}
	return result
}
