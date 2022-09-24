package main

import (
	"math/rand"
	"sort"
	"strconv"
	"time"
)

const NUM_OF_CATS int = 50
const NUM_OF_DISHES int = 3
const DISHWASHING_TIME int = 2
const MIN_CAT_EATING_TIME int = 1
const MAX_CAT_EATING_TIME int = 3

func main() {
	var feedingTimes [NUM_OF_CATS]int
	//generating array of random ints between MAX_CAT_EATING_TIME & MIN_CAT_EATING_TIME
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < NUM_OF_CATS; i++ {
		feedingTimes[i] = rand.Intn(MAX_CAT_EATING_TIME-MIN_CAT_EATING_TIME) + MIN_CAT_EATING_TIME
	}
	// fmt.Println("Total Feeding Time In Seconds: " + strconv.Itoa(calcFeedingTime(feedingTimes)))
	sortedCats := sortCatsDesc(feedingTimes)
	FeedCats(sortedCats)
}

type cat struct {
	id         int
	eatingTime int
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

func calcFeedingTime(feedingTimes [NUM_OF_CATS]int) int {
	//Create a new array of cat to keep track of the cats' IDs
	cats := sortCatsDesc(feedingTimes)

	batchNumber := 1
	catsFed := 0
	feedingTime := 0
	for i := NUM_OF_DISHES - 1; i < NUM_OF_CATS; i += NUM_OF_DISHES {
		log := "Feeding cat "
		//TODO: the following loop is duplicated and can be abstracted to a function
		//loop through the cats in the current batch and add their IDs for logging
		for j := 0; j < NUM_OF_DISHES; j++ {
			log += strconv.Itoa(cats[i-j].id) + " and "
		}
		//crop the trailing 'and' and add the batchNumber
		log = log[:len(log)-4] + "in batch no. " + strconv.Itoa(batchNumber)
		// fmt.Println(log)
		//TODO: consider using math.Max here
		//since it's sorted descending, the left-most cat will always take the longest or equal longest in the batch
		feedingTime += cats[i-NUM_OF_DISHES+1].eatingTime + DISHWASHING_TIME
		batchNumber++
		catsFed += NUM_OF_DISHES
	}
	//if there are cats remaining because NUM_OF_CATS % NUM_OF_DISHES > 0
	if catsFed < NUM_OF_CATS {
		lastLog := "Feeding cat "
		feedingTime += cats[catsFed].eatingTime
		//loop through the remaining cats and add their IDs for logging
		for i := catsFed; i < NUM_OF_CATS; i++ {
			lastLog += strconv.Itoa(cats[i].id) + " and "
		}
		//crop the last 'and'
		lastLog = lastLog[:len(lastLog)-5]
		// fmt.Print(lastLog + " in the last batch (" + strconv.Itoa(batchNumber) + ")\n")
	}

	return feedingTime
}
