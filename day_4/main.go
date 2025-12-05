package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const rollRune = '@'
const maxAdjacentRolls = 4

func main() {
	now := time.Now()

	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	rows := strings.Split(input, "\n")

	accessibleRollsCount := 0
	for y, row := range rows {
		var wg sync.WaitGroup
		var mutex sync.Mutex

		for x, roll := range row {
			if roll != rollRune {
				continue
			}

			wg.Add(1)
			go calculateAccessibleRolls(&accessibleRollsCount, x, y, &rows, len(row), &wg, &mutex)
		}

		wg.Wait()
	}

	fmt.Printf("Solution to Day 4 - Part 1 is %d\n", accessibleRollsCount)
	fmt.Printf("Execution Time: %s\n", time.Since(now))
}

func calculateAccessibleRolls(accessibleRollsCount *int, x int, y int, rows *[]string, rowLen int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	adjacentRollsCount := 0

	// Above
	if y > 0 {
		// Above
		if (*rows)[y-1][x] == rollRune {
			adjacentRollsCount++
		}

		// Above Left
		if x > 0 && (*rows)[y-1][x-1] == rollRune {
			adjacentRollsCount++
		}

		// Above right
		if x < rowLen-1 && (*rows)[y-1][x+1] == rollRune {
			adjacentRollsCount++
		}
	}

	// Left
	if x > 0 && (*rows)[y][x-1] == rollRune {
		adjacentRollsCount++
	}

	// Right
	if x < rowLen-1 && (*rows)[y][x+1] == rollRune {
		adjacentRollsCount++
	}

	// Below
	if y < len(*rows)-1 {
		// Below
		if (*rows)[y+1][x] == rollRune {
			adjacentRollsCount++
		}

		// Below Left
		if x > 0 && (*rows)[y+1][x-1] == rollRune {
			adjacentRollsCount++
		}

		// Below right
		if x < rowLen-1 && (*rows)[y+1][x+1] == rollRune {
			adjacentRollsCount++
		}
	}

	if adjacentRollsCount < maxAdjacentRolls {
		mutex.Lock()
		*accessibleRollsCount++
		mutex.Unlock()
	}
}
