package main

import (
	"fmt"
	"os"
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
	mathProblemsStrings := strings.Split(input, "\n")

	operatorsRunes := parseOperators(&mathProblemsStrings)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	totalMaps := make(map[int]int)
	for _, problem := range mathProblemsStrings {
		wg.Add(1)
		go func() {
			defer wg.Done()

			problemChunks := strings.Split(problem, " ")

			operatorInt := 0
			for _, chunk := range problemChunks {
				if chunk == "" {
					continue
				}

				chunkInt, err := strconv.Atoi(chunk)
				if err != nil {
					break
				}

				mu.Lock()
				if operatorsRunes[operatorInt] == '+' {
					totalMaps[operatorInt] += chunkInt
				} else if operatorsRunes[operatorInt] == '*' {
					existingValue, exists := totalMaps[operatorInt]
					if exists {
						totalMaps[operatorInt] = existingValue * chunkInt
					} else {
						totalMaps[operatorInt] = chunkInt
					}
				}
				mu.Unlock()

				operatorInt++
			}
		}()
	}
	wg.Wait()

	total := 0
	for _, v := range totalMaps {
		total += v
	}

	fmt.Printf("Solution to Day 6 - Part 1 is %d\n", total)
}

func parseOperators(mathProblemsStrings *[]string) []rune {
	operatorsRunes := make([]rune, 0)

	operatorsString := (*mathProblemsStrings)[len(*mathProblemsStrings)-1]
	problemChunks := strings.Split(operatorsString, " ")

	operatorsIndex := 0
	for _, chunk := range problemChunks {
		if chunk == "" {
			continue
		}

		chunkRune := rune(chunk[0])
		if chunkRune == '+' || chunkRune == '*' {
			operatorsRunes = append(operatorsRunes, chunkRune)
			operatorsIndex++
		}
	}

	return operatorsRunes
}
