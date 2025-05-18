package main

import (
	"fmt"
	"os"
	"proj3/cooc"
	"strconv"
)

func main() {
	size := os.Args[1]
	contextSize, _ := strconv.Atoi(os.Args[2])
	smallDoc := "All dogs are good dogs."
	mediumDoc := "All dogs are good dogs. Dogs like cats. Cats hate dogs."
	largeDoc := "The Australian Shepherd is a breed of herding dog from the United States. The name of the breed is technically a misnomer, as it was developed in California in the 19th century, although it has its origins in Asturias, in the northwest of Spain; the breed was unknown in Australia at the time."
	var doc string
	if size == "s" {
		doc = smallDoc
	} else if size == "m" {
		doc = mediumDoc
	} else if size == "l" {
		doc = largeDoc
	} else {
		panic("Argument of s/m/l must be passed in first position")
	}
	tkns := cooc.GetTokens(doc)
	docMap := cooc.CreateMap(tkns, contextSize)
	fmt.Println(docMap)
}
