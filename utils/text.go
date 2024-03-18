package utils

import "unicode"

// Capitalise the first character in a given string
// Does not care about sentences, just the very first letter
func UcFirst(s string) string {
	if s == "" {
		return ""
	}

	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
