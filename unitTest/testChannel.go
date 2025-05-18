package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"proj3/document"
	"sync"
)

func main() {

	// create a channel
	// popualate channel with documents
	corpusPath := "../data/medium/medium.txt"
	DocChannel := make(chan document.Document)
	f, err := os.Open(corpusPath)
	if err != nil {
		panic("Could not read filepath")
	}
	reader := json.NewDecoder(f)
	// RECEIVE FROM CHANNEL
	var w sync.WaitGroup
	const nWorkers int = 4
	w.Add(nWorkers)

	for i := 1; i <= nWorkers; i++ { // spawn 2 workers to iterate over channel
		// GO FUNCTION RECEIVES FROM CHANNEL AND DOES WORK WITH ITEM
		go func(i int, c <-chan document.Document) {
			j := 1
			for d := range c { // GET DOCUMENT FROM GQ, PLACE IN OWN DEQ
				fmt.Printf("%d.%d got %s\n", i, j, d.Title)
				j++
			}
		}(i, DocChannel)
		w.Done()
	}

	// SEND TO CHANNEL
	for {
		var doc document.Document
		err = reader.Decode(&doc)
		if err == io.EOF {
			close(DocChannel)
			break
		} else if err != nil {
			// handle error
			panic("Could not decode")
		}
		DocChannel <- doc
	}

	//spawn go routines to RECEIVE

}

// EXAMPLE 1
// function that will populate channel
//func fibonacci(n int, c chan int) {
//	x, y := 0, 1
//	for i := 0; i < n; i++ {
//		c <- x
//		x, y = y, x+y
//	}
//	close(c)
//}
//
//func main() {
//	c := make(chan int, 10) // create channel
//	go fibonacci(cap(c), c) // spawn sender function
//	for i := range c {      // iterate over channel
//		fmt.Println(i) // print
//	}
//}
//
//// EXAMPLE 2
//// function to SEND to value channel or RECEIVE from quit channel
//func fibonacci(c, quit chan int) {
//	x, y := 0, 1
//	for {
//		select {
//		case c <- x:
//			x, y = y, x+y
//		case <-quit:
//			fmt.Println("quit")
//			return
//		}
//	}
//}
//
//func main() {
//	// make channels
//	c := make(chan int)
//	quit := make(chan int)
//	go func() { // call anonymous function to RECEIVE from value channel
//		for i := 0; i < 10; i++ {
//			fmt.Println(<-c)
//		}
//		quit <- 0 // populate quit channel when done
//	}()
//	// run fibonacci
//	fibonacci(c, quit)
//}
//
