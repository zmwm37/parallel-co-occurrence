package main

import (
	"fmt"
	"proj3/tokens"
	"regexp"
	"strings"
)

func main() {
	doc := "All dogs are good dogs. Dogs like cats. Cats hate dogs."
	str := strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(doc, ""))
	tkns := strings.Split(str, " ")
	vocab := tokens.GetVocab(tkns)
	fmt.Println("Vocab: \n", vocab)
	fmt.Println("Vocab size:", len(vocab))
}
