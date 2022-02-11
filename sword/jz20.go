package sword

import (
	"strconv"
	"strings"
)

func isNumber(s string) bool {
	if _, err := strconv.Atoi(strings.ReplaceAll(s, " ", "")); err != nil {
		return false
	}
	return true
}
