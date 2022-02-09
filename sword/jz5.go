package sword

import "strings"

func replaceSpace(s string) string {
	// write code here
	return strings.ReplaceAll(s, " ", "%20")
}
