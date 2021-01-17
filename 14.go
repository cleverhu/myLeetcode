package main

import (
	"fmt"
	"strings"
)

//输入：strs = ["flower","flow","flight"]
//输出："fl"

//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) < 1 {
		return ""
	}

	s := ""
	b := true

	for k, _ := range []rune(strs[0]) {
		tmp := strs[0][0:k+1]
		for _, v := range strs {
			b = strings.HasPrefix(v, tmp)
			if !b {
				break
			}
		}
		if b {

			s = tmp
		} else {
			break
		}
	}
	return s
}

func main() {
	strs := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(strs))
}
