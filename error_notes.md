## Errors 
1. Updates to tail weren't sticking
    - I wasn't passing the receiver as a pointer to the method, so changes to values do not persist  
    - https://stackoverflow.com/questions/73494229/ineffective-assignment-to-field-when-trying-to-update-a-struct-in-go
2. Trying to populate the worker nodes DEQs, getting the following error:  
    ```panic: runtime error: invalid memory address or nil pointer dereference```
The error is going back to my `Lock` call in `PushBottom`. This makes me think that lock is either not unlocked in time, or that lock not initialized?  
    - **Solution:** I had not initialized the queues in the array, so there was no Lock method to call.  

3. Populating worker node DEQs and the Tail is showing up as `nil` after populating.   
 - Fixed, see git comment on 5/20

4. FeedLocalWorkerQueue isn't terminating
    - The channel is open and since multiple workers are writing to it, one can't close it without risking blocking other workers
    - Trying to read from channel with finite number of documents, but number of documents is not correct  
    - **Solution:** Add nDocuments to context and only iterate that many times while pulling from channel

5. How to indicate that all work has been done for the DEQArray?
    - If victim is empty, lock the DEQarray and check if all are empty. If yes, flip flag
    - Different solution: in DEQArray, create boolean flag for each worker. first time
    worker queue is empty, flip flag.
        - once all flags are flipped, flip context flag for remaining work  
6. More tail pointer problems when PopTop in StealWork testing.
 - `panic: runtime error: invalid memory address or nil pointer dereference` when grabbing q.Tail.Val 
 - IsEmpty wasn't running right, messing up pointers. 



## Questions
- if I want to return nil from a function that normally
