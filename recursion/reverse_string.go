package main

import "fmt"

// iterative approach
func iterReverse(s string) string {
	newString := ""
	for _, val := range(s) {
		newString = string(val) + newString
	}

	return newString
}

// recursive in approach
func recursiveInReverse(s string) string {
	return inPrivateReverse(0, s)
}
func inPrivateReverse(index int, s string) string {
	if index == len(s) {
		return ""
	}
	return inPrivateReverse(index+1, s) + string(s[index])
}

// outsite recursive approach
func recursiveOutReverse(s string) string {
	return outPrivateReverse(s)
}
func outPrivateReverse(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}
	left, right := s[0], s[length-1]
	return string(right) + outPrivateReverse(s[1:length-1]) + string(left)
}

func main() {
	s := "abcde"
	fmt.Println(iterReverse(s))
	fmt.Println(recursiveInReverse(s))
	fmt.Println(recursiveOutReverse(s))
}