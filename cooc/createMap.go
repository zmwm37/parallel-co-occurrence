package cooc

// For a given document, represented as a slice of tokens, create a nested map of words
// to the coounts of their context words.
func CreateMap(tkns []string, contextSize int, vocabMap map[string]int) map[string]map[string]int {
	var documentMap = map[string]map[string]int{}
	for i := 0; i < len(tkns); i++ {
		iWord := tkns[i]

		_, iInVocab := vocabMap[iWord]
		if !iInVocab { // if word not in vocab, continue
			continue
		}
		_, iPrs := documentMap[iWord]
		if !iPrs {
			documentMap[iWord] = map[string]int{}
		}

		for delta := 0; delta < contextSize; delta++ {
			j := i - contextSize + delta
			if j >= 0 {
				jWord := tkns[j]
				_, jInVocab := vocabMap[jWord]
				if !jInVocab { // if context word not in vocab, continue
					continue
				}
				_, jPrs := documentMap[iWord][jWord]

				if !jPrs {
					documentMap[iWord][jWord] = 1
				} else {
					documentMap[iWord][jWord]++
				}
			}
		}
	}

	return documentMap
}
