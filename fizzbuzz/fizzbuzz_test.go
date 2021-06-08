package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	input := 16
	expected := []string{
		"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16",
	}
	result := FizzBuzz(input)
	if !reflect.DeepEqual(result, expected) {
		t.Error("The result should ", expected)
	}
}
