package averybigsum

func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	sum := int64(0)
	for _, i := range ar {
		sum += i
	}
	return sum
}
