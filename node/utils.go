package node

import (
	"regexp"
	"strings"
)

func cleanText(s string) string {
	temp := strings.Trim(
		strings.ReplaceAll(
			removeDoubleSpaces(s),
			"\n",
			"",
		),
		" ",
	)
	if temp == " " {
		return ""
	}
	return temp
}

func removeDoubleSpaces(s string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(s, " ")
}
