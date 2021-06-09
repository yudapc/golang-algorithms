package sockmerchant

func sockMerchant(length int32, numbers []int32) int32 {
	var count int32
	m := make(map[int32]bool)
	for i := int32(0); i < length; i++ {
		value := numbers[i]
		if m[value] {
			count += 1
			m[value] = false
		} else {
			m[value] = true
		}
	}
	return count
}
