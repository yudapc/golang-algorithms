package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	input := []int{3, 1, 4, 2, 7, 5, 6, 8, 9, 10}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := BubbleSort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Error("The result should ", result)
	}
}
