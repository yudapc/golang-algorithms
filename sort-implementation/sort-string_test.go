package sortimplementation

import (
	"reflect"
	"testing"
)

func TestSortString(t *testing.T) {
	input := []string{"c", "g", "a", "v", "b"}
	expected := []string{"a", "b", "c", "g", "v"}
	result := SortString(input)
	if !reflect.DeepEqual(result, expected) {
		t.Error("The result should sorted ", expected)
	}
}
