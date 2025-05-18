package cooc

import (
	"regexp"
	"strings"
)

func Tokenize(doc string) []string {
	str := strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(doc, ""))
	tokens := strings.Split(str, " ")
	return tokens
}
