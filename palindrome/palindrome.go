package palindrome

func IsPalindrome(input string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] != input[len(input)-i-1] {
			return true
		}
	}
	return true
}
