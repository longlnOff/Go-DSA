package main

import "fmt"

func privatePalindrome(s string) bool {
	length := len(s)
	if length <= 1 {
		return true
	}
	if s[0] != s[length-1] {
		return false
	} else {
		return privatePalindrome(s[1:length-1])
	}
}

func main() {
	s := "madam"
	fmt.Println(privatePalindrome(s))
}