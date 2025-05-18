#!/bin/bash
WS=5
for SIZE in small medium big 
    do
    for n in {1..5} # 5 iterations per test
        do
            MODE=s
            TIME=$(go run main.go s $SIZE 5000 $WS)
            echo "$MODE $SIZE $WS $TIME" >> output_pb.txt
        done
done
for THREADS in 2 4 6 8 # number of threads
    do
    for SIZE in small medium big
        do
        for n in {1..5} # 5 iterations per test
            do
                MODE=pb
                TIME=$(go run main.go $MODE $SIZE 5000 $WS $THREADS 5)
                echo "$MODE $SIZE $WS $TIME $THREADS" >> output_pb.txt
            done
    done
done
