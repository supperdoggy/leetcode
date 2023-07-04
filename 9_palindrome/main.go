package main

import (
	"fmt"
)

// with string convertation
// func isPalindrome(x int) bool {
// 	if x < 0 {
// 		return false
// 	}

// 	s := strconv.Itoa(x)

// 	for i := 0; i < len(s)/2; i++ {
// 		if s[i] != s[len(s)-i-1] {
// 			return false
// 		}

// 	}

// 	return true
// }

// func reverseNumber(num int) int {
// 	res := 0
// 	for num > 0 {
// 		remainder := num % 10
// 		res = (res * 10) + remainder
// 		num /= 10
// 	}
// 	return res
// }

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	num := x
	y := 0
	for num > 0 {
		remainder := num % 10
		y = (y * 10) + remainder
		num /= 10
	}

	if x != y {
		return false
	}

	return true
}

func main() {

	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(4004))
}
