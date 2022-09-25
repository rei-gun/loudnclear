package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func FeedCats(cats [NUM_OF_CATS]cat) {
	start := time.Now()
	catsFed := 0
	batchNo := 1
	allCatsFedWg := sync.WaitGroup{}
	allCatsFedWg.Add(len(cats))

	i := 0
	for ; i+NUM_OF_DISHES < NUM_OF_CATS; i += NUM_OF_DISHES {
		feedCats(cats[i:i+NUM_OF_DISHES], batchNo, &allCatsFedWg, catsFed)
		fmt.Println("Washing Dishes...")
		time.Sleep(time.Duration(DISHWASHING_TIME) * time.Second)
		fmt.Println("Finished washing dishes!")
		batchNo++
	}
	if i < NUM_OF_CATS {
		feedCats(cats[i:NUM_OF_CATS], batchNo, &allCatsFedWg, catsFed)
	}
	allCatsFedWg.Wait()
	fmt.Println("Time to feed all cats (seconds): ", time.Since(start))
}

func feedCats(cats []cat, batchNo int, catsFedWg *sync.WaitGroup, catsFed int) {
	currBatchFedWg := sync.WaitGroup{}
	currBatchFedWg.Add(len(cats))
	for i := 0; i < len(cats); i++ {
		go func(i int) {
			fmt.Println("Cat ", strconv.Itoa(cats[i].id), " in batch ", strconv.Itoa(batchNo), "has started eating...")
			time.Sleep(time.Duration(cats[i].eatingTime) * time.Second)
			catsFed++
			currBatchFedWg.Done()
			catsFedWg.Done()
			fmt.Println("Cat ", cats[i].id, " has finished eating")
		}(i)
	}
	currBatchFedWg.Wait()
	return
}
