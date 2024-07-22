package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && !isPrefix(prefix, strs[i]) {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return ""
		}
	}
	return prefix
}

func isPrefix(prefix, str string) bool {
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

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // Output: "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "cecar", "car"}))      // Output: ""
}
