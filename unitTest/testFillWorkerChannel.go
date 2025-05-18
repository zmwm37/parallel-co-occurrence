package main

import (
	"encoding/json"
	"fmt"
	"os"
	"proj3/concurrent"
)

func main() {

	// create a channel
	// popualate channel with documents
	corpusPath := "../data/small/small.txt"
	//DocChannel := make(chan document.Document)
	f, err := os.Open(corpusPath)
	if err != nil {
		panic("Could not read filepath")
	}
	reader := json.NewDecoder(f)
	const nWorkers int = 2
	// RECEIVE FROM CHANNEL
	DEQarray := concurrent.FeedWorkers(nWorkers, reader)
	for i, Q := range DEQarray {
		n := Q.Head.Next
		for {
			fmt.Println(i, "-", n.Val.Title)
			if n.Next == nil {
				break
			}
			n = n.Next
		}
		//fmt.Println("DEQ", i, "head", Q.Head.Next.Val.Title)
		//fmt.Println("DEQ Tail", Q.Tail.Val.Title)
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
