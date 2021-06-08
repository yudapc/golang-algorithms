package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	input := "ana"
	expected := true
	result := IsPalindrome(input)
	if result != expected {
		t.Error("Should return true is palindrome")
	}
}

func TestIsNotPalindrome(t *testing.T) {
	input := "anbba"
	expected := true
	result := IsPalindrome(input)
	if result != expected {
		t.Error("Should return true is not palindrome")
	}
}
