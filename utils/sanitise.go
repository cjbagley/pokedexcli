package utils

import (
	"regexp"
	"strings"
)

func SanitisePromptInput(s string) []string {
	alpha := regexp.MustCompile("[^A-Za-z ]")
	s = alpha.ReplaceAllString(s, " ")
	s = strings.ToLower(s)

	return strings.Fields(s)
}
