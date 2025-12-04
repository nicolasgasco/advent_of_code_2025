package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	idRanges := strings.Split(input, ",")

	sumPartOne := calculateInvalidIds(
		calculateInvalidIdsParams{
			idRanges: idRanges, isComplexId: false,
		})
	fmt.Printf("Solution to Day 2 - Part 1 is %d\n", sumPartOne)

	sumPartTwo := calculateInvalidIds(
		calculateInvalidIdsParams{
			idRanges: idRanges, isComplexId: true,
		})
	fmt.Printf("Solution to Day 2 - Part 2 is %d\n", sumPartTwo)

	fmt.Printf("Execution time: %s\n", time.Since(startTime))
}

type calculateInvalidIdsParams struct {
	idRanges    []string
	isComplexId bool
}

func calculateInvalidIds(params calculateInvalidIdsParams) int {
	idRanges := params.idRanges
	isComplexId := params.isComplexId

	var wg sync.WaitGroup
	var mutex sync.Mutex
	sum := 0

	for _, idRange := range idRanges {
		wg.Add(1)
		go inspectIdRange(idRange, &sum, isComplexId, &wg, &mutex)
	}
	wg.Wait()

	return sum
}

func inspectIdRange(idRange string, sum *int, isComplexId bool, wg *sync.WaitGroup, mutex *sync.Mutex) int {
	defer wg.Done()

	idRangeChunks := strings.Split(idRange, "-")
	startId := idRangeChunks[0]
	endId := idRangeChunks[1]

	startIdInt, err := strconv.Atoi(startId)
	if err != nil {
		panic(err)
	}

	endIdInt, err := strconv.Atoi(endId)
	if err != nil {
		panic(err)
	}

	for i := startIdInt; i <= endIdInt; i++ {
		isInvalid := false

		idString := fmt.Sprintf("%d", i)
		if isComplexId {
			isInvalid = isInvalidIdPartTwo(idString)
		} else {
			isInvalid = isInvalidIdPartOne(idString)
		}

		if isInvalid {
			mutex.Lock()
			*sum += i
			mutex.Unlock()
		}
	}

	return 0
}

func isInvalidIdPartOne(id string) bool {
	idLength := len(id)

	patternLength := idLength / 2

	pattern := id[:patternLength]
	comparison := id[patternLength:]

	return pattern == comparison
}

func isInvalidIdPartTwo(id string) bool {
	const minPatternLength = 1

	idLength := len(id)

	highestPatternLength := idLength / 2

	for patternLength := highestPatternLength; patternLength >= minPatternLength; patternLength-- {
		if idLength%patternLength != 0 {
			continue
		}

		pattern := id[0:patternLength]

		var isInvalidId bool = false
		for i := patternLength; i <= len(id)-patternLength; i += patternLength {

			subPattern := id[i : i+patternLength]
			if subPattern == pattern {
				isInvalidId = true
			} else {
				isInvalidId = false
				break
			}
		}
		if isInvalidId {
			return true
		}
	}
	return false
}
