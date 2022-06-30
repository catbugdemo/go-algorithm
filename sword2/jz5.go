package sword2

import "strings"

// 替换
func replaceSpace(s string) string {
	var res string
	for _, buf := range s {
		if buf == ' ' {
			res += "%20"
		} else {
			res += string(buf)
		}
	}
	return res
}

func replaceSpace2(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
