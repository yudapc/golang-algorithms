package averybigsum

import "testing"

func TestAVeryBigSum(t *testing.T) {
	t.Run("input 5", func(t *testing.T) {
		input := []int64{5}
		expected := int64(5)
		result := aVeryBigSum(input)
		if expected != result {
			t.Error("The result should be ", expected, " received is ", result)
		}
	})
	t.Run("input 1000000001 1000000002 1000000003 1000000004 1000000005", func(t *testing.T) {
		input := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
		expected := int64(5000000015)
		result := aVeryBigSum(input)
		if expected != result {
			t.Error("The result should be ", expected, " received is ", result)
		}
	})
}
