package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	banks := strings.Split(input, "\n")

	joltage := 0

	for _, bank := range banks {
		if bank == "" {
			continue
		}

		var batteries = []rune(bank)

		highestValue := 0
		secondHighestValue := 0
		for i, battery := range batteries {
			batteryValue := int(battery - '0')
			hasOneCharLeft := i == len(batteries)-1
			if batteryValue > highestValue && !hasOneCharLeft {
				highestValue = batteryValue
				secondHighestValue = 0
			} else if batteryValue > secondHighestValue {
				secondHighestValue = batteryValue
			}
		}
		joltage += highestValue*10 + secondHighestValue
	}

	fmt.Println("Result of Day 3 - Part 1 is", joltage)
}
