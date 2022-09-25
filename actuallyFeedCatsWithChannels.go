package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func FeedCatsWithChannels(cats [NUM_OF_CATS]cat) {
	start := time.Now()
	batchNo := 1

	dirtyDishes := make(chan bool)
	quitDishwashing := make(chan bool)

	catsFedWg := sync.WaitGroup{}
	catsFedWg.Add(NUM_OF_CATS)

	dishWashingWg := sync.WaitGroup{}
	go washDishes(dirtyDishes, quitDishwashing, &dishWashingWg)

	catsFed := 0
	//this loop sends the cats to be fed in batches
	for ; catsFed+NUM_OF_DISHES < NUM_OF_CATS; catsFed += NUM_OF_DISHES {
		//sending cats in batches
		fmt.Println("Starting batch no. ", batchNo)
		feedCatsWithChan(cats[catsFed:catsFed+NUM_OF_DISHES], dirtyDishes, &catsFedWg, &dishWashingWg)
		batchNo++
		dishWashingWg.Wait()
	}
	//if there are cats remaining, then send them in the final batch
	if catsFed < NUM_OF_CATS {
		fmt.Println("Starting final batch no. ", batchNo)
		feedCats(cats[catsFed:NUM_OF_CATS], batchNo, &catsFedWg, catsFed)
	}
	//wait until all cats are fed
	catsFedWg.Wait()
	//Don't want to leave the dishes dirty so do 1 last wash =]
	dishWashingWg.Add(1)
	dirtyDishes <- true
	quitDishwashing <- true
	fmt.Println("Actual time to feed all cats (seconds): ", time.Since(start))
}

func feedCatsWithChan(cats []cat, dirtyDishes chan bool, catsFedWg *sync.WaitGroup, dishWashingWg *sync.WaitGroup) {
	currBatchWg := sync.WaitGroup{}
	currBatchWg.Add(len(cats))
	//Get cats eating on all dishes
	for i := 0; i < len(cats); i++ {
		go func(i int) {

			fmt.Println("Cat ", strconv.Itoa(cats[i].id), " has started eating...")
			time.Sleep(time.Duration(cats[i].eatingTime) * time.Second)
			fmt.Println("Cat ", cats[i].id, " has finished eating")
			currBatchWg.Done()
			catsFedWg.Done()
		}(i)
	}
	//wait until all cats in the current batch have finished eating
	currBatchWg.Wait()
	//trigger the washDishes function via this channel
	dirtyDishes <- true
	// fmt.Println(len(dirtyDishes))
	//Add 1 to the dishwashing wait group so the next batch of cats is blocked from eating
	dishWashingWg.Add(1)
	return
}

func washDishes(dirtyDishes chan bool, quit chan bool, dishWashingWg *sync.WaitGroup) {
	for {
		fmt.Println(len(dirtyDishes))
		select {
		case <-dirtyDishes:
			fmt.Println("Washing Dishes...")
			time.Sleep(time.Duration(DISHWASHING_TIME) * time.Second)
			fmt.Println("Finished washing dishes!")
			//Set dishwashingWg to Done so that the next batch of cats can start eating
			dishWashingWg.Done()
			fmt.Println(len(dirtyDishes))
		case <-quit:
			return
		}
	}
}
