package utils

import "strings"

func EscapeNonUrlChars(s string) string {
	s = strings.ReplaceAll(s, "=", "")
	s = strings.ReplaceAll(s, "+", "-")
	s = strings.ReplaceAll(s, "/", "_")
	return s
}
