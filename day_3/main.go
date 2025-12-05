package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	banks := strings.Split(input, "\n")

	calculateSimpleJoltage(&banks)

	calculateComplexJoltage(&banks)
}

func calculateSimpleJoltage(banks *[]string) {
	joltage := 0

	for _, bank := range *banks {
		if bank == "" {
			continue
		}

		var batteries = []rune(bank)

		highestValue := 0
		secondHighestValue := 0
		for i, battery := range batteries {
			batteryValue := int(battery - '0')
			hasOnlyOneBatteryLeft := i == len(batteries)-1
			if batteryValue > highestValue && !hasOnlyOneBatteryLeft {
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

func calculateComplexJoltage(banks *[]string) {
	const bankJoltageBatteries = 12
	joltage := 0

	for _, bank := range *banks {
		if bank == "" {
			continue
		}

		numBatteriesToRemove := len(bank) - bankJoltageBatteries

		var batteries = []rune(bank)
		fmt.Printf("Bank before removal: %c\n", batteries)

		sortedBatteries := make([]rune, len(batteries))
		copy(sortedBatteries, batteries)
		sort.Slice(sortedBatteries, func(i, j int) bool {
			return sortedBatteries[j] > sortedBatteries[i]
		})

		batteriesToRemove := sortedBatteries[:numBatteriesToRemove]

		fmt.Printf("Removing batteries: %c\n", batteriesToRemove)

		var bankJoltage = []rune{}
		batteriesToRemoveIndex := 0
		hasRemovedAllBatteries := false
		for _, battery := range batteries {
			if battery == batteriesToRemove[batteriesToRemoveIndex] && !hasRemovedAllBatteries {
				if batteriesToRemoveIndex < len(batteriesToRemove)-1 {
					batteriesToRemoveIndex++
				} else {
					hasRemovedAllBatteries = true
				}
			} else {
				bankJoltage = append(bankJoltage, battery)
			}
		}

		fmt.Printf("Final bank: %c\n", bankJoltage)
	}

	fmt.Println("Result of Day 3 - Part 1 is", joltage)

}
