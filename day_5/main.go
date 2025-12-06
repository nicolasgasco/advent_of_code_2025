package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)

	databaseSections := strings.Split(input, "\n\n")
	ranges := strings.Split(databaseSections[0], "\n")
	ids := strings.Split(databaseSections[1], "\n")

	freshIngredientsCount := calculateFreshIngredientsCount(&ranges, &ids)
	fmt.Printf("Solution to Day 5 - Part 1 is %d\n", freshIngredientsCount)

	freshIds := calculateFreshIds(&ranges)
	fmt.Printf("Solution to Day 5 - Part 2 is %d\n", freshIds)
}

func calculateFreshIds(ranges *[]string) int {
	rangesInts := make([][2]int, len(*ranges))
	for i, r := range *ranges {
		rangeValues := strings.Split(r, "-")
		start, _ := strconv.Atoi(rangeValues[0])
		end, _ := strconv.Atoi(rangeValues[1])
		rangesInts[i] = [2]int{start, end}
	}
	sort.Slice(rangesInts, func(i, j int) bool {
		return rangesInts[i][0] < rangesInts[j][0]
	})

	uniqueRangesInts := make([][2]int, 0)

	for _, r := range rangesInts {
		overlapFound := false
		for ui, ur := range uniqueRangesInts {
			isStartInRange := r[0] >= ur[0] && r[0] <= ur[1]
			if isStartInRange {
				overlapFound = true

				isEndBeyondRange := r[1] > ur[1]
				if isEndBeyondRange {
					uniqueRangesInts[ui][1] = r[1]
				}
			}
		}
		if !overlapFound {
			uniqueRangesInts = append(uniqueRangesInts, r)
		}
	}

	totalIdsCount := 0
	for _, ur := range uniqueRangesInts {
		totalIdsCount += ur[1] - ur[0] + 1
	}

	return totalIdsCount
}

func calculateFreshIngredientsCount(ranges *[]string, ids *[]string) int {
	freshIngredientsCount := 0

	wc := sync.WaitGroup{}
	mu := sync.Mutex{}

	for _, id := range *ids {
		idValue, _ := strconv.Atoi(id)

		wc.Add(1)
		go func() {
			defer wc.Done()
			for _, r := range *ranges {
				rangeValues := strings.Split(r, "-")

				start, _ := strconv.Atoi(rangeValues[0])
				end, _ := strconv.Atoi(rangeValues[1])

				if idValue >= start && idValue <= end {
					mu.Lock()
					freshIngredientsCount++
					mu.Unlock()

					break
				}
			}
		}()
	}

	wc.Wait()

	return freshIngredientsCount
}
