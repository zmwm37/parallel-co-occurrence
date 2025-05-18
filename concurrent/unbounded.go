package concurrent

import (
	"proj3/document"
	"sync"
)

/**** YOU CANNOT MODIFY ANY OF THE FOLLOWING INTERFACES/TYPES ********/
type Task interface{}

type DEQueue interface {
	PushBottom(task Task)
	IsEmpty() bool //returns whether the queue is empty
	PopTop() Task
	PopBottom() Task
	Size() int
}

/******** DO NOT MODIFY ANY OF THE ABOVE INTERFACES/TYPES *********************/

// NewUnBoundedDEQueue returns an empty UnBoundedDEQueue
func NewUnBoundedDEQueue() *UnBoundedDEQueue {
	sentinel := NewNode()
	var m sync.Mutex
	q := UnBoundedDEQueue{Head: sentinel, Tail: sentinel, L: &m}
	return &q
}

type Node struct {
	Val  document.Document // To avoid ABA, must NOT be a pointer
	Pred *Node
	Next *Node
}

func NewNode() *Node { // maybe this shouldn't be a pointer
	return &Node{Pred: nil, Next: nil}
}

type UnBoundedDEQueue struct {
	Head *Node
	Tail *Node
	L    *sync.Mutex
}

func (q *UnBoundedDEQueue) PushBottom(task document.Document) {
	q.L.Lock()
	new := NewNode()
	new.Val = task
	if q.IsEmpty() {
		new.Pred = q.Head
		new.Next = nil
		q.Tail = new
		q.Tail.Pred = q.Head
		q.Head.Next = new
	} else {
		new.Next = q.Head.Next
		q.Head.Next.Pred = new
		if q.Head.Next == q.Tail {
			q.Tail.Pred = new
		}
		q.Head.Next = new
		new.Pred = q.Head
	}
	q.L.Unlock()
}

func (q *UnBoundedDEQueue) PopBottom() document.Document {
	q.L.Lock()
	var t *Node
	if q.IsEmpty() {
		t = NewNode()
	} else {
		t = q.Head.Next
		new_head := q.Head.Next.Next
		if new_head == nil {
			q.Tail = q.Head
			q.Head.Next = nil
		} else {
			new_head.Pred = q.Head
			if new_head == q.Tail {
				q.Tail.Pred = q.Head
			}
			q.Head.Next = new_head
		}

	}
	q.L.Unlock()
	return t.Val
}

func (q *UnBoundedDEQueue) PopTop() document.Document {

	var d document.Document
	q.L.Lock()
	if !q.IsEmpty() {
		d = q.Tail.Val
		q.Tail = q.Tail.Pred
	}
	q.L.Unlock()
	return d
}

func (q *UnBoundedDEQueue) IsEmpty() bool {
	if q.Head == q.Tail {
		return true
	} else {
		return false
	}
}

func (q *UnBoundedDEQueue) Size() int {
	q.L.Lock()
	counter := 0
	current := q.Head
	for current.Next != nil {
		counter++
		current = current.Next
	}
	q.L.Unlock()
	return counter
}
