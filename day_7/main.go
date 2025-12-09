package main

import (
	"fmt"
	"os"
	"strings"
)

const startChar rune = 'S'
const splitChar rune = '^'
const beamChar rune = '|'

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := string(data)
	manifoldRows := strings.Split(input, "\n")

	splitCounter := 0
	for y, row := range manifoldRows {
		for x, char := range row {
			if char == startChar && y < len(manifoldRows)-1 {
				rowRunes := []rune(manifoldRows[y+1])
				rowRunes[x] = beamChar
				manifoldRows[y+1] = string(rowRunes)
			} else if char == splitChar {
				if y > 0 && manifoldRows[y-1][x] == uint8(beamChar) {
					rowRunes := []rune(manifoldRows[y])
					rowRunes[x-1] = beamChar
					rowRunes[x+1] = beamChar
					manifoldRows[y] = string(rowRunes)
					splitCounter++
				}
			} else {
				if y > 0 && manifoldRows[y-1][x] == uint8(beamChar) {
					rowRunes := []rune(manifoldRows[y])
					rowRunes[x] = beamChar
					manifoldRows[y] = string(rowRunes)
				}
			}
		}
	}

	fmt.Printf("Solution to Day 7 - Part 1 is %d\n", splitCounter)

	scenariosCounter := 0

	for _, row := range manifoldRows {
		if strings.Contains(row, string(splitChar)) {
			for _, char := range row {
				if char == beamChar {
					scenariosCounter++
				}
			}
		}
	}
	fmt.Printf("Solution to Day 7 - Part 2 is %d\n", scenariosCounter)

}
