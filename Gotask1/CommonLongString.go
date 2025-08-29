/*
*最长公共前缀
 */
package main

import "fmt"

func lcp(str1, str2 string) string {
	length, index := min(len(str1), len(str2)), 0
	for {
		if index < length && str1[index] == str2[index] {

			index++
		} else {
			break
		}
	}
	return str1[:index]
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix, count := strs[0], len(strs)
	for i := range count - 1 {
		prefix = lcp(prefix, strs[i+1])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func main() {
	strss := []string{"flower", "flow", "flight"}
	prefixStr := longestCommonPrefix(strss)
	fmt.Println(prefixStr)
}
