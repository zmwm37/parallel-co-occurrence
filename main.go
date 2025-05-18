package main

import (
	"concurrent"
	"context"
	"cooc"
	"fmt"
	"os"
	"sequential"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	args := os.Args
	usageString := "Usage: go run main.go <mode (s/ps/pb)> <corpus (small/medium/big)> <vocabSize (500/1000/5000)> <windowSize> <capacity> <balanceThreshold>"
	if len(args) < 5 {
		panic(usageString)
	}
	mode := args[1]
	corpus := args[2]
	corpusPath := "./data/" + corpus + "/" + corpus + ".txt"

	vocabSize := args[3]
	vocabMapPath := "./util/vocab_map_" + vocabSize + ".txt"
	vocabMap := cooc.LoadVocab(vocabMapPath)

	windowSize, wsErr := strconv.Atoi(args[4])

	if wsErr != nil {
		panic(usageString)
	}

	if mode == "s" {
		sequential.RunSequential(corpusPath, vocabMap, windowSize)
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Println(elapsed)
	} else {
		capacity, cErr := strconv.Atoi(args[5])
		if cErr != nil {
			panic(usageString)
		}
		var balanceThreshold int
		var btErr error
		if len(args) == 7 {
			balanceThreshold, btErr = strconv.Atoi(args[6])
			if btErr != nil {
				panic(usageString)
			}
		} else {
			balanceThreshold = 25 // default of 25 if none provided
		}

		ctx := context.GoContext{
			Mode:             mode,
			CorpusPath:       corpusPath,
			WindowSize:       windowSize,
			VocabMap:         vocabMap,
			Capacity:         capacity,
			BalanceThreshold: float64(balanceThreshold),
			WorkLeft:         true,
			ReduceChannel:    make(chan map[string]map[string]int),
		}
		if mode == "ps" || mode == "pb" {
			concurrent.RunConcurrent(&ctx)
			//CM := concurrent.RunWorkStealing(corpusPath, vocabMapPath)
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Println(elapsed)
		} else {
			panic(usageString)
		}
	}
}
