package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	now := time.Now()

	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := string(data)
	rows := strings.Split(input, "\n")

	// calculating all possible connections with distances
	connections := make([]string, 0)
	for i, row := range rows {
		rowCoordinates := strings.Split(row, ",")
		rowX, _ := strconv.Atoi(rowCoordinates[0])
		rowY, _ := strconv.Atoi(rowCoordinates[1])
		rowZ, _ := strconv.Atoi(rowCoordinates[2])

		rowConnectionsString := make([]string, 0)
		for innerI, innerRow := range rows {
			if innerI == i {
				continue
			}

			innerRowCoordinates := strings.Split(innerRow, ",")
			innerRowX, _ := strconv.Atoi(innerRowCoordinates[0])
			innerRowY, _ := strconv.Atoi(innerRowCoordinates[1])
			innerRowZ, _ := strconv.Atoi(innerRowCoordinates[2])

			linearDistance := math.Sqrt(
				math.Pow(float64(rowX-innerRowX), 2) +
					math.Pow(float64(rowY-innerRowY), 2) +
					math.Pow(float64(rowZ-innerRowZ), 2),
			)

			innerRowString := fmt.Sprintf("%s;%s:%.2f", row, innerRow, linearDistance)
			rowConnectionsString = append(rowConnectionsString, innerRowString)
		}
		connections = append(connections, rowConnectionsString...)
	}

	// sorting by shortest distance first
	sort.Slice(connections, func(i, j int) bool {
		iParts := strings.Split(connections[i], ":")
		iDistance, _ := strconv.ParseFloat(iParts[1], 64)

		jParts := strings.Split(connections[j], ":")
		jDistance, _ := strconv.ParseFloat(jParts[1], 64)

		return iDistance < jDistance
	})

	const numberOfConnections = 1000
	circuits := make([]string, 0)
	// iterating 2 by 2 to skip duplicate connections (e,g. a;b and b;a)
	for i := 0; i < numberOfConnections*2; i += 2 {
		circuitParts := strings.Split(connections[i], ":")
		junctionBoxes := strings.Split(circuitParts[0], ";")

		firstBoxIndex := -1
		secondBoxIndex := -1
		for circuitI, circuit := range circuits {
			if strings.Contains(circuit, junctionBoxes[0]) {
				firstBoxIndex = circuitI
			}
			if strings.Contains(circuit, junctionBoxes[1]) {
				secondBoxIndex = circuitI
			}
		}

		// both not found
		if firstBoxIndex == -1 && secondBoxIndex == -1 {
			circuitString := fmt.Sprintf("%s;%s", junctionBoxes[0], junctionBoxes[1])
			circuits = append(circuits, circuitString)
			continue
			// both found
		} else if firstBoxIndex != -1 && secondBoxIndex != -1 {
			// already in the same circuit
			if firstBoxIndex == secondBoxIndex {
				continue
				// merging circuits
			} else {
				circuits[firstBoxIndex] += ";" + circuits[secondBoxIndex]
				circuits = append(circuits[:secondBoxIndex], circuits[secondBoxIndex+1:]...)
				continue
			}
		}

		// first box found
		if firstBoxIndex != -1 {
			if !strings.Contains(circuits[firstBoxIndex], junctionBoxes[1]) {
				circuits[firstBoxIndex] += fmt.Sprintf(";%s", junctionBoxes[1])
			}
			continue
		}

		if !strings.Contains(circuits[secondBoxIndex], junctionBoxes[0]) {
			circuits[secondBoxIndex] += fmt.Sprintf(";%s", junctionBoxes[0])
		}
	}

	// sorting circuits by length descending
	sort.Slice(circuits, func(i, j int) bool {
		partsI := strings.Split(circuits[i], ";")
		partsJ := strings.Split(circuits[j], ";")

		return len(partsI) > len(partsJ)
	})

	total := 1
	for i := 0; i < 3; i++ {
		circuitParts := strings.Split(circuits[i], ";")
		total *= len(circuitParts)
	}

	fmt.Printf("Solution to Day 8 - Part 1 is %d\n", total)

	fmt.Printf("Time elapsed: %s\n", time.Since(now))
}
