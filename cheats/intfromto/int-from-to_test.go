package intfromto

import (
	"reflect"
	"testing"
)

func TestIntFromTo(t *testing.T) {
	result := IntFromTo(1, 5)
	expected := []int{1, 2, 3, 4, 5}
	if reflect.DeepEqual(result, expected) == false {
		t.Error("Error not same")
	}
}

func TestInt32FromTo(t *testing.T) {
	result := Int32FromTo(1, 5)
	expected := []int32{1, 2, 3, 4, 5}
	if reflect.DeepEqual(result, expected) == false {
		t.Error("Error not same")
	}
}

func TestInt64FromTo(t *testing.T) {
	result := Int64FromTo(1, 5)
	expected := []int64{1, 2, 3, 4, 5}
	if reflect.DeepEqual(result, expected) == false {
		t.Error("Error not same")
	}
}
