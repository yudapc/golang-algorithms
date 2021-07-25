package logicthird

func calculateSumOfArray(numbers []int) int {
	length := len(numbers)
	maxIndex := length - 1
	result := 0

	for index, item := range numbers {
		nextIndex := index + 1
		if nextIndex > maxIndex {
			nextIndex = maxIndex
		}

		nextItem := numbers[nextIndex]
		if index == maxIndex {
			nextItem = numbers[0]
		}

		if item == nextItem {
			result += item
		}
	}

	return result
}
