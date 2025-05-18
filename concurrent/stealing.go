package concurrent

import (
	"math/rand"
	"proj3/context"
	"proj3/document"
)

// NewWorkStealingExecutor returns an ExecutorService that is implemented using the work-stealing algorithm.
// @param capacity - The number of goroutines in the pool
// @param threshold - The number of items that a goroutine in the pool can
// grab from the executor in one time period. For example, if threshold = 10
// this means that a goroutine can grab 10 items from the executor all at
// once to place into their local queue before grabbing more items. It's
// not required that you use this parameter in your implementation.
func NewWorkStealingExecutor(capacity, threshold int) ExecutorService {

	/** TODO: Remove the return nil and implement this function **/
	return nil
}

func StealWork(thiefId int, victimId int, A []*UnBoundedDEQueue, ctx *context.GoContext) {
	stolenDoc := A[victimId].PopTop() // TODO - what happens in simultaneous popping?
	emptyDoc := document.NewDocument()
	if stolenDoc == emptyDoc {
		return
	}
	A[thiefId].PushBottom(stolenDoc)
}

func SelectVictim(potentialVictims []int) int {
	victimIdx := rand.Intn(len(potentialVictims))
	return potentialVictims[victimIdx]
}
