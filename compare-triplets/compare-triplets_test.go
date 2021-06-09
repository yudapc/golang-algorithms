package comparetriplets

import (
	"reflect"
	"testing"
)

func TestCompareTriplets(t *testing.T) {
	t.Run("positive case 1", func(t *testing.T) {
		inputA := []int32{5, 6, 7}
		inputB := []int32{3, 6, 10}
		expected := []int32{1, 1}
		result := compareTriplets(inputA, inputB)
		if !reflect.DeepEqual(result, expected) {
			t.Error("The result should be ", expected)
		}
	})
	t.Run("positive case 2", func(t *testing.T) {
		inputA := []int32{5, 7, 7}
		inputB := []int32{3, 6, 10}
		expected := []int32{2, 1}
		result := compareTriplets(inputA, inputB)
		if !reflect.DeepEqual(result, expected) {
			t.Error("The result should be ", expected)
		}
	})
	t.Run("nagative case", func(t *testing.T) {
		inputA := []int32{5, 7}
		inputB := []int32{3, 6, 10}
		expected := []int32{2, 1}
		result := compareTriplets(inputA, inputB)
		if result != nil {
			t.Error("The result should be ", expected)
		}
	})
}
