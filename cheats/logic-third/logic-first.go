package logicthird

func smallestInteger(numbers []int) int {
	min := 0
	for index, item := range numbers {
		if index == 0 || item < min {
			min = item
		}
	}

	return min
}
