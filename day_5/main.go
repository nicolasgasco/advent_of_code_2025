package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

	freshIngredientsCount := 0
	for _, id := range ids {
		isFresh := false

		for _, r := range ranges {
			rangeValues := strings.Split(r, "-")

			start, _ := strconv.Atoi(rangeValues[0])
			end, _ := strconv.Atoi(rangeValues[1])

			idValue, _ := strconv.Atoi(id)

			if idValue >= start && idValue <= end {
				isFresh = true
				break
			}
		}

		if isFresh {
			freshIngredientsCount++
		}
	}

	fmt.Printf("Solution to Day 5 - Part 1 is %d\n", freshIngredientsCount)
}
