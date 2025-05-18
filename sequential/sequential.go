package sequential

import (
	"proj3/cooc"
	"proj3/document"
)

func RunSequential(corpusPath string, vocabMap map[string]int, windowSize int) *cooc.CooccurrenceMatrix {
	CM := cooc.NewCoocMatrix(len(vocabMap))
	CM.VocabMap = vocabMap
	CM.InitC()
	gq := document.FeedDocs(corpusPath)

	for {
		var doc document.Document
		doc, gq = document.Dequeue(gq)
		//fmt.Println("Title:", doc.Title)
		tkns := cooc.Tokenize(doc.Text)
		m := cooc.CreateMap(tkns, windowSize, vocabMap)
		CM.ReduceDocMap(m)
		if len(gq) == 0 {
			break
		}
	}
	return CM

}
