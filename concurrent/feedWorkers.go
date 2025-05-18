package concurrent

import (
	"encoding/json"
	"io"
	"proj3/context"
	"proj3/document"
	"sync"
)

type DEQArray struct {
	A             []*UnBoundedDEQueue
	CriminalFlags []bool // indicates if a queue has at any point been ready to steal
}

func NewDEQArray(nDEQs int) *DEQArray {
	return &DEQArray{A: make([]*UnBoundedDEQueue, nDEQs), CriminalFlags: make([]bool, nDEQs)}
}

func FeedWorkers(ctx *context.GoContext, reader *json.Decoder) *DEQArray {
	var wg sync.WaitGroup
	ctx.NDocuments = 0
	deqArray := NewDEQArray(ctx.Capacity - 1)
	GlobalChannel := make(chan document.Document)

	for i := 0; i < (ctx.Capacity - 1); i++ {
		deqArray.A[i] = NewUnBoundedDEQueue()
		deqArray.CriminalFlags[i] = false
	}

	wg.Add(ctx.Capacity - 1)

	for i := 0; i < (ctx.Capacity - 1); i++ { // spawn 2 workers to iterate over channel
		// GO FUNCTION RECEIVES FROM CHANNEL AND DOES WORK WITH ITEM
		go func(i int, c <-chan document.Document, DEQArray []*UnBoundedDEQueue) {
			for d := range c { // GET DOCUMENT FROM GQ, PLACE IN OWN DEQ
				DEQArray[i].PushBottom(d)
			}
			wg.Done()
		}(i, GlobalChannel, deqArray.A)

	}

	for {
		var doc document.Document
		err := reader.Decode(&doc)

		if err == io.EOF {
			break
		} else if err != nil {
			// handle error
			panic("Could not decode")
		}
		ctx.NDocuments++
		GlobalChannel <- doc // This is submitting
	}
	close(GlobalChannel)
	wg.Wait()
	return deqArray
}
