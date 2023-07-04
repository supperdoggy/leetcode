package main

import "fmt"

// initial approach, but fails on timeout when s is really big
// func lengthOfLongestSubstring(s string) int {
// 	longest := 0

// 	for i := range s {
// 		substring := map[rune]struct{}{}
// 		for ; i < len(s); i++ {
// 			_, exists := substring[rune(s[i])]
// 			if exists {
// 				substring = make(map[rune]struct{})
// 			}

// 			substring[rune(s[i])] = struct{}{}

// 			if len(substring) > longest {
// 				longest = len(substring)
// 			}
// 		}
// 	}

//		return longest
//	}

func lengthOfLongestSubstring(s string) int {
	longest := 0

	substring := []rune{}
	for _, v := range s {
		index := indexOf(substring, v)
		if index != -1 {
			substring = substring[index+1:]
		}

		substring = append(substring, v)

		if len(substring) > longest {
			longest = len(substring)
		}

	}

	return longest
}

func indexOf(s []rune, a rune) int {
	for k, v := range s {
		if v == a {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabca"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("asjrgapa"))
}
