package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const rollRune = '@'
const removedRune = 'x'
const maxAdjacentRolls = 4

func main() {
	now := time.Now()

	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)

	rows := strings.Split(input, "\n")
	tempRows := make([]string, len(rows))
	copy(tempRows, rows)

	accessibleRollsCount := 0

	for y, row := range rows {
		for x, roll := range row {
			if roll == rollRune {
				calculateAccessibleRolls(&accessibleRollsCount, x, y, &rows, &tempRows, len(row))
			}
		}
	}

	fmt.Printf("Solution to Day 4 - Part 1 is %d\n", accessibleRollsCount)

	rows = strings.Split(input, "\n")
	tempRows = make([]string, len(rows))
	copy(tempRows, rows)

	accessibleRollsCount = 0

	for {
		prevAccessibleRollsCount := accessibleRollsCount
		for y, row := range rows {
			for x, roll := range row {
				if roll == rollRune {
					calculateAccessibleRolls(&accessibleRollsCount, x, y, &rows, &tempRows, len(row))
				}
			}
		}
		rows = make([]string, len(tempRows))
		copy(rows, tempRows)

		if prevAccessibleRollsCount == accessibleRollsCount {
			break
		}
	}

	fmt.Printf("Solution to Day 4 - Part 2 is %d\n", accessibleRollsCount)

	fmt.Printf("Execution Time: %s\n", time.Since(now))
}

func calculateAccessibleRolls(accessibleRollsCount *int, x int, y int, rows *[]string, tempRows *[]string, rowLen int) {
	adjacentRollsCount := 0

	// Above
	if y > 0 {
		// Above
		aboveTile := (*rows)[y-1][x]
		if aboveTile == rollRune {
			adjacentRollsCount++
		}

		// Above Left
		if x > 0 {
			aboveLeftTile := (*rows)[y-1][x-1]
			if aboveLeftTile == rollRune {
				adjacentRollsCount++
			}
		}

		// Above right
		if x < rowLen-1 {
			aboveRightTile := (*rows)[y-1][x+1]
			if aboveRightTile == rollRune {
				adjacentRollsCount++
			}
		}
	}

	// Left
	if x > 0 {
		leftTile := (*rows)[y][x-1]
		if leftTile == rollRune {
			adjacentRollsCount++
		}
	}

	// Right
	if x < rowLen-1 {
		rightTile := (*rows)[y][x+1]
		if rightTile == rollRune {
			adjacentRollsCount++
		}
	}

	// Below
	if y < len(*rows)-1 {
		// Below
		belowTile := (*rows)[y+1][x]
		if belowTile == rollRune {
			adjacentRollsCount++
		}

		// Below Left
		if x > 0 {
			belowLeftTile := (*rows)[y+1][x-1]
			if belowLeftTile == rollRune {
				adjacentRollsCount++
			}
		}

		// Below right
		if x < rowLen-1 {
			belowRightTile := (*rows)[y+1][x+1]
			if belowRightTile == rollRune {
				adjacentRollsCount++
			}
		}
	}

	if adjacentRollsCount < maxAdjacentRolls {
		rowRunes := []rune((*tempRows)[y])
		rowRunes[x] = removedRune
		(*tempRows)[y] = string(rowRunes)

		*accessibleRollsCount++
	}
}
