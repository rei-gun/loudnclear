package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

const NUM_OF_CATS int = 10
const NUM_OF_DISHES int = 3
const DISHWASHING_TIME int = 2
const MIN_CAT_EATING_TIME int = 1
const MAX_CAT_EATING_TIME int = 3

type cat struct {
	id         int
	eatingTime int
}

func main() {
	var feedingTimes [NUM_OF_CATS]int
	//generating array of random ints between MAX_CAT_EATING_TIME & MIN_CAT_EATING_TIME
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < NUM_OF_CATS; i++ {
		feedingTimes[i] = rand.Intn(MAX_CAT_EATING_TIME-MIN_CAT_EATING_TIME) + MIN_CAT_EATING_TIME
	}
	fmt.Println("Calculated Feeding Time In Seconds: " + strconv.Itoa(calcFeedingTime(feedingTimes)))
	sortedCats := sortCatsDesc(feedingTimes)
	// FeedCats(sortedCats)
	FeedCatsWithChannels(sortedCats)
}

func sortCatsDesc(feedingTimes [NUM_OF_CATS]int) [NUM_OF_CATS]cat {
	//Create a new array of cat to keep track of the cats' IDs
	var cats [NUM_OF_CATS]cat
	for i := 0; i < NUM_OF_CATS; i++ {
		cats[i] = cat{i, feedingTimes[i]}
	}

	//sort the cats in descending order of their eatingTime
	sort.Slice(cats[:], func(i, j int) bool {
		return cats[i].eatingTime > cats[j].eatingTime
	})
	return cats
}
