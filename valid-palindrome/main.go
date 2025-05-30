package main

import (
	"fmt"
	"strings"
)

func isAlphanumeric(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	s = strings.ToLower(s)

	for i < j {
		if !isAlphanumeric(s[i]) {
			i++
			continue
		}
		if !isAlphanumeric(s[j]) {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true

}

func main() {
	input := "A man, a plan, a canal: Panama"

	fmt.Println(isPalindrome(input)) // Expected: true
}
