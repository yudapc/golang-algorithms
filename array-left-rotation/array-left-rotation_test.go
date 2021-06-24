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

func BenchmarkArrayLeftRotation(b *testing.B) {
	input := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i := 0; i < b.N; i++ {
		if i > len(input) {
			break
		}
		ArrayLeftRotation(input, int32(i))
	}
}

func BenchmarkArrayLeftRotationSimple(b *testing.B) {
	input := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i := 0; i < b.N; i++ {
		if i > len(input) {
			break
		}
		ArrayLeftRotationSimple(input, int32(i))
	}
}
