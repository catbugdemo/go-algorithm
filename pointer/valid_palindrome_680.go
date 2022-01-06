package pointer

//给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
//
//示例 1:
//
//输入: "aba"
//输出: True

//validPalindrome
func validPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s[left:right]) || isPalindrome(s[left+1:right+1])
		}
		left++
		right--
	}
	return true
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
