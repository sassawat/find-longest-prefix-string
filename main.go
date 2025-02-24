package main

import (
	"fmt"
	"unicode"
)

func isPrefix(prefix, str string) bool {
	if len(str) > 200 {
		return false
	}

	if len(prefix) > len(str) {
		return false
	}

	for i := 0; i < len(prefix); i++ {
		if prefix[i] != str[i] {
			return false
		}
	}
	return true
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) > 200 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		str := strs[i]

		for _, char := range str {
			r := rune(char)
			if !unicode.IsLower(r) {
				return ""
			}
		}
		for len(prefix) > 0 && !isPrefix(prefix, str) {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return ""
		}
	}
	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // Output: "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "cecar", "car"}))      // Output: ""
}
