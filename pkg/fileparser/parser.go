package fileparser

import (
	"regexp"
	"strings"
)

func ParseString(input string) []string {
	wordRegex := regexp.MustCompile(`\b\w+\b`)

	matches := wordRegex.FindAllString(input, -1)

	for i, match := range matches {
		matches[i] = strings.ToLower(match)
	}

	return matches
}