package arrayleftrotation

import (
	"reflect"
	"testing"
)

func TestArrayLeftRotation(t *testing.T) {
	input := []int32{1, 2, 3, 4, 5}
	expected := []int32{3, 4, 5, 1, 2}
	result := ArrayLeftRotation(input, 2)
	if !reflect.DeepEqual(result, expected) {
		t.Error("The result should be ", expected)
	}
}

func TestArrayLeftRotationSimple(t *testing.T) {
	input := []int32{1, 2, 3, 4, 5}
	expected := []int32{3, 4, 5, 1, 2}
	result := ArrayLeftRotationSimple(input, 2)
	if !reflect.DeepEqual(result, expected) {
		t.Error("The result should be ", expected)
	}
}
