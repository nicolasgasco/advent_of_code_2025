package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := string(data)
	tilesRows := strings.Split(input, "\n")

	maxX := 0
	maxY := 0
	coordinates := make([][]int, 0)
	for _, row := range tilesRows {
		rowChunks := strings.Split(row, ",")
		x, _ := strconv.Atoi(rowChunks[0])
		if x > maxX {
			maxX = x
		}

		y, _ := strconv.Atoi(rowChunks[1])
		if y > maxY {
			maxY = y
		}

		coordinates = append(coordinates, []int{x, y})
	}

	largestArea := 0
	for i, coordinate := range coordinates {
		x := coordinate[0]
		y := coordinate[1]
		for otherI, otherCoordinate := range coordinates {
			if otherI == i {
				continue
			}

			otherX := otherCoordinate[0]
			otherY := otherCoordinate[1]

			horizontalSide := 1
			if otherX > x {
				horizontalSide += otherX - x
			} else {
				horizontalSide += x - otherX
			}

			verticalSide := 1
			if otherY > y {
				verticalSide += otherY - y
			} else {
				verticalSide += y - otherY
			}

			area := horizontalSide * verticalSide

			if area > largestArea {
				largestArea = area
			}
		}
	}

	fmt.Printf("Solution to Day 9 - Part 1 is %d\n", largestArea)
}
