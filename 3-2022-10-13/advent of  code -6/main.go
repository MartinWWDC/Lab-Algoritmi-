package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
func main() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

	arr := strings.Split(string(content), ",")
	arr2 := []int{}
	fmt.Println((arr))

	for _, i := range arr {
		j, _ := strconv.Atoi(i)
		arr2 = append(arr2, j)
	}
	fmt.Println(arr2)
	//arr2 := []int{3, 4, 3, 1, 2}
	n := 256
	for g := 0; g < n; g++ {
		for i, el := range arr2 {
			el--
			if el == -1 {
				el = 6
				arr2 = append(arr2, 8)
			}
			arr2[i] = el

		}
		//fmt.Println("day ", g, " ", arr2)

	}
	//fmt.Println(arr2)
	fmt.Println(len(arr2))

}*/

const MAX_DAY = 256
const MAX_TIMER = 10
const TIMER_AFTER_RESET = 6
const TIMER_AFTER_SPAWN = 8

func main() {
	var buffer string
	fmt.Scanln(&buffer)
	var splitBuffer = strings.Split(buffer, ",")
	var timerFishCountBuckets [MAX_TIMER]int
	for _, numString := range splitBuffer {
		initialTimer, _ := strconv.Atoi(numString)
		timerFishCountBuckets[initialTimer]++
		// fmt.Println(initialTimer, timerFishCountBuckets)
	}
	// fmt.Println(timerFishCountBuckets)
	for i := 0; i < MAX_DAY; i++ {
		fishToBearChildCount := timerFishCountBuckets[0]
		for n := 1; n < MAX_TIMER; n++ {
			timerFishCountBuckets[n-1] = timerFishCountBuckets[n]
		}
		timerFishCountBuckets[TIMER_AFTER_RESET] += fishToBearChildCount
		timerFishCountBuckets[TIMER_AFTER_SPAWN] += fishToBearChildCount

		// fmt.Println(timerFishCountBuckets)
	}

	var totalFishCount = 0
	for _, fishWithTimerTimeCount := range timerFishCountBuckets {
		totalFishCount += fishWithTimerTimeCount
	}
	fmt.Println(totalFishCount)
}
