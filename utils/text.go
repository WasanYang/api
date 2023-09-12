package utils

import "strings"

func ReplaceSpaceWithDash(text string) string {
	return strings.Replace(text, " ", "-", -1)
}
