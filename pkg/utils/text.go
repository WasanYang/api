package utils

import "strings"

type stringUtils struct {
}
type IStringUtils interface {
	ReplaceSpaceWithDash(text string) string
}

func NewStringUtils() IStringUtils {
	return &stringUtils{}
}

func (s *stringUtils) ReplaceSpaceWithDash(text string) string {
	return strings.Replace(text, " ", "-", -1)
}
