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
	now := time.Now()

	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := string(data)
	tilesRows := strings.Split(input, "\n")

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	coordinates := make([][]int, 0)
	for _, row := range tilesRows {
		rowChunks := strings.Split(row, ",")
		x, _ := strconv.Atoi(rowChunks[0])
		y, _ := strconv.Atoi(rowChunks[1])

		coordinates = append(coordinates, []int{x, y})
	}

	largestArea := 0
	for i, coordinate := range coordinates {
		x := coordinate[0]
		y := coordinate[1]
		for otherI, otherCoordinate := range coordinates {
			wg.Add(1)
			go func() {
				if otherI == i {
					wg.Done()
					return
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
					mu.Lock()
					largestArea = area
					mu.Unlock()
				}
				wg.Done()
			}()
			wg.Wait()
		}
	}

	fmt.Printf("Solution to Day 9 - Part 1 is %d\n", largestArea)

	fmt.Printf("Execution took %s\n", time.Since(now))
}
