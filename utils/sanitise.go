package utils

import (
	"regexp"
	"strings"
)

func SanitisePromptInput(s string) []string {
	alpha := regexp.MustCompile("[^A-Za-z ]")
	whitespace := regexp.MustCompile("\\s+")

	s = alpha.ReplaceAllString(s, " ")
	s = whitespace.ReplaceAllString(s, " ")
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return []string{}
	}

	return strings.Split(s, " ")
}
