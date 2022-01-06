package pointer

import (
	"strings"
)

// 反转指针
func reverseWords(s string) string {
	var res strings.Builder
	for _, s2 := range strings.Split(s, " ") {
		res.WriteString(reverseWord(s2))
		res.WriteString(" ")
	}
	return res.String()[:len(res.String())-1]
}

func reverseWord(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
