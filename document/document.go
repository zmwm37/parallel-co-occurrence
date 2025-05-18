package document

import (
	"encoding/json"
	"io"
	"os"
	"proj3/context"
)

type Document struct {
	Id    string
	Title string
	Text  string
	ctx   *context.GoContext
}

func NewDocument() Document {
	return Document{}
}

// Parse json and feed into Global Queue
func FeedDocs(corpusPath string) []Document {
	var GlobalQueue []Document
	f, err := os.Open(corpusPath)
	if err != nil {
		panic("Could not read filepath")
	}
	reader := json.NewDecoder(f)
	for {
		var doc Document
		err = reader.Decode(&doc)
		if err == io.EOF {
			break
		} else if err != nil {
			// handle error
			panic("Could not decode")
		}
		GlobalQueue = append(GlobalQueue, doc)
	}
	return GlobalQueue
}

// adapated from https://www.geeksforgeeks.org/queue-in-go-language/
func Dequeue(q []Document) (Document, []Document) {
	element := q[0] // The first element is the one to be dequeued.
	if len(q) == 1 {
		var tmp = []Document{}
		return element, tmp
	}
	return element, q[1:] // Slice off the element once it is dequeued.
}
