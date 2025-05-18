package main

import (
	"fmt"
	"proj3/concurrent"
	"proj3/document"
)

func main() {
	// Test pushing bottom
	fmt.Println("Add one item")
	Q := concurrent.NewUnBoundedDEQueue()
	d1 := document.NewDocument("1")
	d1.Text = "Hello queue!"
	Q.PushBottom(d1)
	fmt.Println("Head Next correct?", Q.Head.Next.Val.Id)
	fmt.Println("Tail correct?", Q.Tail.Val)

	fmt.Println("Add a second item")
	d2 := document.NewDocument("2")
	Q.PushBottom(d2)
	fmt.Println("Head Next correct?", Q.Head.Next.Val)
	fmt.Println("Tail correct?", Q.Tail.Val)

}
