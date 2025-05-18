package document

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

const nSmallDocs = 100
const nMediumDocs = 1000
const nBigDocs = 10000

func TestParseJson(t *testing.T) {
	pathFile := "../data/test/test.txt"
	// pathFile := "../data/small/small.txt"
	f, err := os.Open(pathFile)
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
		fmt.Println("DOC:", doc)
	}
}

func TestFeedDocs(t *testing.T) {
	t.Run("1=Small", func(t *testing.T) {
		pathFile := "../data/small/small.txt"
		GQ := FeedDocs(pathFile)
		if len(GQ) != nSmallDocs {
			t.Errorf("Queue size incorrect.\ngot:%v\nexpected:%v", len(GQ), nSmallDocs)
		}
	})
	t.Run("2=Medium", func(t *testing.T) {
		pathFile := "../data/medium/medium.txt"
		GQ := FeedDocs(pathFile)
		if len(GQ) != nMediumDocs {
			t.Errorf("Queue size incorrect.\ngot:%v\nexpected:%v", len(GQ), nMediumDocs)
		}
	})
	t.Run("3=Big", func(t *testing.T) {
		pathFile := "../data/big/big.txt"
		GQ := FeedDocs(pathFile)
		if len(GQ) != nBigDocs {
			t.Errorf("Queue size incorrect.\ngot:%v\nexpected:%v", len(GQ), nBigDocs)
		}
	})

}
