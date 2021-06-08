package sortimplementation

import (
	"reflect"
	"testing"
)

func TestSortInteger(t *testing.T) {
	input := []int{3, 4, 2, 1, 6, 5, 8, 7, 9, 10}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := SortInteger(input)
	if !reflect.DeepEqual(result, expected) {
		t.Error("The result should sorted ", expected)
	}
}
