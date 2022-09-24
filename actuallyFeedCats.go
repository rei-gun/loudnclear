package main

import (
	"fmt"
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
*/
func FeedCats(sortedCats [NUM_OF_CATS]cat) int {
	catsEating := make(chan int)
	catsEating <- 0
	washingDishes := make(chan bool)
	washingDishes <- false
	inactiveDishes := make(chan int)
	inactiveDishes <- 0
	catsFed := make(chan int)
	catsFed <- 0

	for i := 0; i < NUM_OF_DISHES; i++ {
		go feedAndWash(sortedCats, inactiveDishes, catsEating, catsFed, washingDishes)
	}
	return 0
}

func feedAndWash(cats [NUM_OF_CATS]cat, inactiveDishes chan int, catsEating chan int, catsFed chan int, washingDishes chan bool) time.Duration {
	start := time.Now()
	for {
		// quitVal := <-quit
		// select {
		// case _catsFed := <-catsFed:
		_catsFed := <-catsFed
		if _catsFed+1 >= len(cats) {
			_inactiveDishes := <-inactiveDishes
			_inactiveDishes++
			inactiveDishes <- _inactiveDishes
		}
		// case inactiveDishes := <-inactiveDishes:
		inactiveDishes := <-inactiveDishes
		if inactiveDishes == NUM_OF_DISHES {
			return time.Since(start)
		}
		// case _washingDishes := <-_washingDishes:
		_washingDishes := <-washingDishes
		if _washingDishes {
			time.Sleep(time.Duration(DISHWASHING_TIME) / 2) //another GR is washing dishes so wait half the time
		}

		// case _catsEating := <-catsEating:
		_catsEating := <-catsEating
		if _catsEating < NUM_OF_DISHES {
			_catsEating++
			catsEating <- _catsEating
			_catsFed := <-catsFed
			_catsFed++
			catsFed <- _catsFed
			fmt.Printf("Feeding cat %d, cats eating: %d \n", cats[_catsFed-1].id, _catsEating)
			time.Sleep(time.Duration(cats[_catsFed-1].eatingTime))

		} else if _catsEating == NUM_OF_DISHES {
			washingDishes <- true
			catsEating <- 0
			fmt.Println("WASHING DISHES!")
			time.Sleep(time.Duration(DISHWASHING_TIME))
		}
		// }
	}
}
