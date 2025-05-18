package concurrent

import (
	"encoding/json"
	"os"
	"proj3/context"
	"proj3/cooc"
	"proj3/document"
	"testing"
)

const nSmallDocs = 100
const nMediumDocs = 1000
const nBigDocs = 10000

// Push single task to bottom
func TestPushBottom(t *testing.T) {
	t.Run("A=1", func(t *testing.T) {
		Q := NewUnBoundedDEQueue()
		d1 := document.NewDocument()
		d1.Id = "1"
		Q.PushBottom(d1)
		if Q.Head.Next.Val.Id != "1" {
			t.Errorf("Head value incorrect.\ngot:%v\nexpected:%v", Q.Head.Next.Val.Id, "1")
		}
		if Q.Tail.Val.Id != "1" {
			t.Errorf("Tail value incorrect.\ngot:%v\nexpected:%v", Q.Tail.Val.Id, "1")
		}
	})
	t.Run("A=2", func(t *testing.T) {
		Q := NewUnBoundedDEQueue()
		d1 := document.NewDocument()
		d1.Id = "1"
		d2 := document.NewDocument()
		d2.Id = "2"
		Q.PushBottom(d1)
		Q.PushBottom(d2)
		if Q.Head.Next.Val.Id != "2" {
			t.Errorf("Head value incorrect.\ngot:%v\nexpected:%v", Q.Head.Next.Val.Id, "2")
		}
		if Q.Tail.Val.Id != "1" {
			t.Errorf("Tail value incorrect.\ngot:%v\nexpected:%v", Q.Tail.Val.Id, "1")
		}
	})

}

func TestPopTop(t *testing.T) {
	// Test that empty queue returns empty document
	t.Run("Empty", func(t *testing.T) {
		Q := NewUnBoundedDEQueue()
		res := Q.PopTop()
		var emptyDocument document.Document
		if res != emptyDocument {
			t.Errorf("Result incorrect.\ngot:%v\nexpected:%v", res, emptyDocument)
		}
	})

	// Test that pop returns tail of queue
	t.Run("A=1", func(t *testing.T) {
		Q := NewUnBoundedDEQueue()
		d1 := document.NewDocument()
		d1.Id = "1"
		d2 := document.NewDocument()
		d2.Id = "2"
		Q.PushBottom(d1)
		Q.PushBottom(d2)
		res := Q.PopTop()
		if res.Id != "1" {
			t.Errorf("Result incorrect.\ngot:%v\nexpected:%v", res.Id, "1")
		}
	})

	// Push 2, pop 3
	t.Run("A=2", func(t *testing.T) {
		Q := NewUnBoundedDEQueue()
		d1 := document.NewDocument()
		d1.Id = "1"
		d2 := document.NewDocument()
		d2.Id = "2"
		Q.PushBottom(d1)
		Q.PushBottom(d2)
		Q.PopTop()
		Q.PopTop()
		res := Q.PopTop()
		var emptyDocument document.Document
		if res != emptyDocument {
			t.Errorf("Result incorrect.\ngot:%v\nexpected:%v", res, emptyDocument)
		}
		if Q.Tail != Q.Head {
			t.Errorf("Tail and Head do not match.\ntail:%v\nhead:%v", Q.Tail, Q.Head)
		}
	})

}

func TestPopBottom(t *testing.T) {
	t.Run("InitialEmpty", func(t *testing.T) {
		Q := NewUnBoundedDEQueue()
		res := Q.PopBottom()
		var emptyDocument document.Document
		if res != emptyDocument {
			t.Errorf("Result incorrect.\ngot:%v\nexpected:%v", res, emptyDocument)
		}
	})

	// Push and pop one element
	t.Run("A=1", func(t *testing.T) {
		Q := NewUnBoundedDEQueue()
		d1 := document.NewDocument()
		d1.Id = "1"
		Q.PushBottom(d1)
		res := Q.PopBottom()
		if res.Id != "1" {
			t.Errorf("Result incorrect.\ngot:%v\nexpected:%v", res.Id, "1")
		}
		if !Q.IsEmpty() {
			t.Errorf("Result incorrect.\ngot:%v\nexpected:%v", !Q.IsEmpty(), true)
		}
	})

	// Push 2, pop 1
	t.Run("A=2", func(t *testing.T) {
		Q := NewUnBoundedDEQueue()
		d1 := document.NewDocument()
		d1.Id = "1"
		d2 := document.NewDocument()
		d2.Id = "2"
		Q.PushBottom(d1)
		Q.PushBottom(d2)
		res := Q.PopBottom()
		if res.Id != "2" {
			t.Errorf("Result incorrect.\ngot:%v\nexpected:%v", res.Id, "2")
		}
	})
	// Push 2, pop 2
	t.Run("A=3", func(t *testing.T) {
		Q := NewUnBoundedDEQueue()
		d1 := document.NewDocument()
		d1.Id = "1"
		d2 := document.NewDocument()
		d2.Id = "2"
		Q.PushBottom(d1)
		Q.PushBottom(d2)
		Q.PopBottom()
		res := Q.PopBottom()
		if res.Id != "1" {
			t.Errorf("Result incorrect.\ngot:%v\nexpected:%v", res.Id, "1")
		}
		if !Q.IsEmpty() {
			t.Errorf("Result incorrect.\ngot:%v\nexpected:%v", !Q.IsEmpty(), true)
		}
	})
}

func TestSize(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		Q := NewUnBoundedDEQueue()
		if Q.Size() != 0 {
			t.Errorf("Result incorrect.\ngot:%v\nexpected:%v", Q.Size(), 0)
		}
	})
	t.Run("A=1", func(t *testing.T) {
		Q := NewUnBoundedDEQueue()
		d1 := document.NewDocument()
		d1.Id = "1"
		d2 := document.NewDocument()
		d2.Id = "2"
		Q.PushBottom(d1)
		Q.PushBottom(d2)
		Q.PopBottom()
		if Q.Size() != 1 {
			t.Errorf("Result incorrect.\ngot:%v\nexpected:%v", Q.Size(), 1)
		}
	})

	t.Run("A=2", func(t *testing.T) {
		Q := NewUnBoundedDEQueue()
		d1 := document.NewDocument()
		d1.Id = "1"
		d2 := document.NewDocument()
		d2.Id = "2"
		d3 := document.NewDocument()
		d3.Id = "3"
		Q.PushBottom(d1)
		Q.PushBottom(d2)
		Q.PushBottom(d3)
		if Q.Size() != 3 {
			t.Errorf("Result incorrect.\ngot:%v\nexpected:%v", Q.Size(), 3)
		}
	})

}

func TestFeedWorkers(t *testing.T) {
	t.Run("Small Corpus", func(t *testing.T) {
		corpusPath := "../data/small/small.txt"
		f, err := os.Open(corpusPath)
		if err != nil {
			panic("Could not read filepath")
		}
		reader := json.NewDecoder(f)
		ctx := &context.GoContext{Capacity: 2}
		workerQs := FeedWorkers(ctx, reader)
		c := 0
		for _, q := range workerQs.A {
			c = c + q.Size()
		}
		if c != nSmallDocs {
			t.Errorf("Result incorrect.\ngot:%v\nexpected:%v", c, nSmallDocs)
		}
		if ctx.NDocuments != nSmallDocs {
			t.Errorf("NDocuments incorrect.\ngot:%v\nexpected:%v", ctx.NDocuments, nSmallDocs)
		}

	})

	t.Run("Medium Corpus", func(t *testing.T) {
		corpusPath := "../data/medium/medium.txt"
		f, err := os.Open(corpusPath)
		if err != nil {
			panic("Could not read filepath")
		}
		reader := json.NewDecoder(f)
		ctx := &context.GoContext{Capacity: 2}
		workerQs := FeedWorkers(ctx, reader)
		c := 0
		for _, q := range workerQs.A {
			c = c + q.Size()
		}
		if c != nMediumDocs {
			t.Errorf("Result incorrect.\ngot:%v\nexpected:%v", c, nMediumDocs)
		}
		if ctx.NDocuments != nMediumDocs {
			t.Errorf("NDocuments incorrect.\ngot:%v\nexpected:%v", ctx.NDocuments, nMediumDocs)
		}

	})

	t.Run("Big Corpus", func(t *testing.T) {
		corpusPath := "../data/big/big.txt"
		f, err := os.Open(corpusPath)
		if err != nil {
			panic("Could not read filepath")
		}
		reader := json.NewDecoder(f)
		ctx := &context.GoContext{Capacity: 2}
		workerQs := FeedWorkers(ctx, reader)
		c := 0
		for _, q := range workerQs.A {
			c = c + q.Size()
		}
		if c != nBigDocs {
			t.Errorf("Result incorrect.\ngot:%v\nexpected:%v", c, nBigDocs)
		}
		if ctx.NDocuments != nBigDocs {
			t.Errorf("NDocuments incorrect.\ngot:%v\nexpected:%v", ctx.NDocuments, nBigDocs)
		}

	})
}

func TestFeedLocalQueue(t *testing.T) {
	filePath := "../util/vocab_map_5.txt"
	VocabMap := cooc.LoadVocab(filePath)
	corpusPath := "../data/small/small.txt"
	f, err := os.Open(corpusPath)
	if err != nil {
		panic("Could not read filepath")
	}
	reader := json.NewDecoder(f)

	ctx := context.GoContext{Mode: "ps", Capacity: 2, WorkLeft: false, WindowSize: 3, VocabMap: VocabMap,
		ReduceChannel: make(chan map[string]map[string]int)}

	DEQArray := FeedWorkers(&ctx, reader)
	for i := 0; i < (ctx.Capacity - 1); i++ {
		go FeedLocalQueue(i, DEQArray, &ctx)
	}

	c := 0
	for i := 1; i <= ctx.NDocuments; i++ {
		<-ctx.ReduceChannel
		c++
	}
	if c != nSmallDocs {
		t.Errorf("NDocuments incorrect.\ngot:%v\nexpected:%v", c, nSmallDocs)
	}
}

func TestStealWork(t *testing.T) {
	filePath := "../util/vocab_map_5.txt"
	VocabMap := cooc.LoadVocab(filePath)
	ctx := context.GoContext{Mode: "ps", WorkLeft: true, Capacity: 2, NDocuments: 4, VocabMap: VocabMap,
		ReduceChannel: make(chan map[string]map[string]int)}
	deqArray := NewDEQArray(ctx.Capacity)
	deqArray.CriminalFlags = []bool{false, true}

	nExpected := 4
	// create somes docs
	d1 := document.NewDocument()
	d1.Id = "1"
	d1.Text = "This is the first document."
	d2 := document.NewDocument()
	d2.Id = "2"
	d2.Text = "This is the second document."
	d3 := document.NewDocument()
	d3.Id = "3"
	d3.Text = "This is the third document."
	d4 := document.NewDocument()
	d4.Id = "4"
	d4.Text = "This is the fourth document."
	deqArray.A[0] = NewUnBoundedDEQueue()
	deqArray.A[0].PushBottom(d1)
	deqArray.A[0].PushBottom(d2)
	deqArray.A[0].PushBottom(d3)
	deqArray.A[0].PushBottom(d4)
	deqArray.A[1] = NewUnBoundedDEQueue()

	for i := 0; i < ctx.Capacity; i++ {
		go FeedLocalQueue(i, deqArray, &ctx)
	}
	c := 0
	for i := 1; i <= ctx.NDocuments; i++ {

		<-ctx.ReduceChannel
		c++
	}

	if c != nExpected {
		t.Errorf("NDocuments incorrect.\ngot:%v\nexpected:%v", c, nExpected)
	}

}
