package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/*
- 1. Create 1 function for feeding a cat, receives the feeding time and sleeps for that duration
- 2. Start goroutines (GR) of function == NUM_OF_DISHES
- 3. Initiate a channel with int = 0
- 4. channel++ when cat finishes
- 5. If channel == NUM_OF_DISHES then sleep for DISHWASHING_TIME, set channel = -1 (2 remaining should wait)
- 6. When dishwashing (GR) finishes, set channel = 0, which sets all 3 GRs to start the next batch eating
- 7. If GR tries to pick up cat above len(cats) then set quitChannel++
- 8. When quitChannel -- len(cats) then exit

- 1. a function that feeds cats, run a GR for each cat. use waitgroup to ensure
*/
func FeedCats2(cats [NUM_OF_CATS]cat) time.Duration {
	start := time.Now()
	catsFed := 0
	batchNo := 1
	catsFedWg := sync.WaitGroup{}
	catsFedWg.Add(len(cats))
	i := 0
	for ; i+NUM_OF_DISHES < NUM_OF_CATS; i += NUM_OF_DISHES {
		feedCats(cats[i:i+NUM_OF_DISHES], batchNo, &catsFedWg, catsFed)
		fmt.Println("Washing Dishes")
		time.Sleep(time.Duration(DISHWASHING_TIME))
		batchNo++
	}
	if i < NUM_OF_CATS {
		feedCats(cats[i:NUM_OF_CATS], batchNo, &catsFedWg, catsFed)
	}
	catsFedWg.Wait()
	fmt.Println("Time to feed all cats (seconds): ", time.Since(start))
	return time.Since(start)
}

func feedCats(cats []cat, batchNo int, catsFedWg *sync.WaitGroup, catsFed int) {
	wg := sync.WaitGroup{}
	wg.Add(len(cats))
	for i := 0; i < len(cats); i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("Feeding cat " + strconv.Itoa(cats[i].id) + " in batch " + strconv.Itoa(batchNo))
			time.Sleep(time.Duration(cats[i].eatingTime))
			catsFed++
			catsFedWg.Done()
		}(i)
	}
	// time.Sleep(time.Duration(eatingTime))
	wg.Wait()
	return
}
