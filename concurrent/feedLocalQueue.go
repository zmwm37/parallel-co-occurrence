package concurrent

import (
	"math"
	"math/rand"
	"proj3/context"
	"proj3/cooc"
	"proj3/document"
)

func FeedLocalQueue(id int, deqArray *DEQArray, ctx *context.GoContext) {
	localQ := deqArray.A[id]
	if ctx.Mode == "ps" {
		for {
			if localQ.IsEmpty() {
				if !deqArray.CriminalFlags[id] { // if not previously stealing, set CriminalFlag to true
					deqArray.CriminalFlags[id] = true
				}
				var potentialVictims []int
				for i := 0; i < (ctx.Capacity - 1); i++ {
					if !deqArray.CriminalFlags[i] {
						potentialVictims = append(potentialVictims, i)
					}
					if len(potentialVictims) == 0 { // if all workers are stealing, indicate  that work is done and workers can stop
						ctx.WorkLeft = false
					}
				}
				if ctx.WorkLeft {
					victimId := SelectVictim(potentialVictims)
					StealWork(id, victimId, deqArray.A, ctx)
				} else { // trying to steal, but no work left for anyone that isn't stolen
					return
				}
			}
			doc := localQ.PopBottom()
			if doc == document.NewDocument() { // if failed to steal or popped an empty document
				continue
			}
			tkns := cooc.Tokenize(doc.Text)
			m := cooc.CreateMap(tkns, ctx.WindowSize, ctx.VocabMap)
			ctx.ReduceChannel <- m
			//fmt.Println(id, "Added map to channel")
		}
	} else if ctx.Mode == "pb" {
		c := 0
		//
		for {
			if (c % 5) == 0 {
				var balance bool
				//flip coin
				n := localQ.Size()
				if n == 0 {
					balance = true
				} else {
					x := rand.Intn(n)
					y := rand.Intn(n)
					if x == y {
						balance = true
					} else {
						balance = false
					}
				}
				if balance {
					// select a random victim
					var potentialVictims []int
					for i := 0; i < (ctx.Capacity - 1); i++ { // TODO - add a lock while thread reads CriminalFlags
						if !deqArray.CriminalFlags[i] {
							potentialVictims = append(potentialVictims, i)
						}
						if len(potentialVictims) == 0 { // if all workers are stealing, indicate  that work is done and workers can stop
							ctx.WorkLeft = false
						}
					}
					if ctx.WorkLeft {
						victimId := SelectVictim(potentialVictims)
						absoluteDelta := math.Abs(float64(localQ.Size() - deqArray.A[victimId].Size()))
						if absoluteDelta >= ctx.BalanceThreshold {
							balanceAmount := int(absoluteDelta) / 2
							if localQ.Size() < deqArray.A[victimId].Size() {

								for i := 0; i < int(balanceAmount); i++ {
									StealWork(id, victimId, deqArray.A, ctx)
								}
							} else {
								StealWork(victimId, id, deqArray.A, ctx)
							}
						}
					} else { // trying to steal, but no work left for anyone that isn't stolen
						return
					}
				}
			}
			doc := localQ.PopBottom()
			if doc == document.NewDocument() { // if failed to steal or popped an empty document
				continue
			}
			tkns := cooc.Tokenize(doc.Text)
			m := cooc.CreateMap(tkns, ctx.WindowSize, ctx.VocabMap)
			ctx.ReduceChannel <- m
			c++
		}
	}

}

// WORK BALANCING NOTES
// every x tasks, flip a coin. prob of coin is 1 / (size of queue)
// if coin is true
// pick a random victim
// lock queues and compare sizes.
// if size delta greater than threshold, rebalance
//
