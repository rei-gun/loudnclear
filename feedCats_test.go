package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSortCatsDesc(t *testing.T) {
	var feedingTimes [NUM_OF_CATS]int
	//generating array of random ints between MAX_CAT_EATING_TIME & MIN_CAT_EATING_TIME
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		feedingTimes[i] = rand.Intn(MAX_CAT_EATING_TIME-MIN_CAT_EATING_TIME) + MIN_CAT_EATING_TIME
	}
	sortedCats := sortCatsDesc(feedingTimes)
	for i := 0; i < len(sortedCats)-1; i++ {
		if sortedCats[i].eatingTime < sortedCats[i+1].eatingTime {
			t.Errorf("index %d is less than index %d", i, i+1)
		}
	}
}

func TestWashDishes(t *testing.T) {
	dirtyDishes := make(chan bool)
	quitDishwashing := make(chan bool)
	dishWashingWg := sync.WaitGroup{}

	// dirtyDishes <- true
	dishWashingWg.Add(1)
	go washDishes(dirtyDishes, quitDishwashing, &dishWashingWg)
	time.Sleep(time.Duration(DISHWASHING_TIME) * time.Second)
	quitDishwashing <- true
}

func TestFeedCats(t *testing.T) {
	return
}
