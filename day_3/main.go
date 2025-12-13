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

		var batteries = []rune(bank)
		var bankJoltage = make([]rune, 0)

		numBatteriesToRemove := len(batteries) - bankJoltageBatteries
		loopStart := 0
		for {
			highestBattery := '0'
			removedBatteries := 0
			for i := 0; i < numBatteriesToRemove+1; i++ {
				if (i + loopStart) > len(batteries)-1 {
					break
				}

				if batteries[i+loopStart] > highestBattery {
					highestBattery = batteries[i+loopStart]
					removedBatteries = i
				}
			}

			loopStart += removedBatteries + 1
			bankJoltage = append(bankJoltage, highestBattery)
			numBatteriesToRemove -= removedBatteries

			if numBatteriesToRemove <= 0 || len(bankJoltage) == bankJoltageBatteries {
				break
			}

		}

		missingBatteries := bankJoltageBatteries - len(bankJoltage)
		bankJoltage = append(bankJoltage, batteries[len(batteries)-missingBatteries:]...)

		bankJoltageString := string(bankJoltage)
		bankJoltageValue, _ := strconv.Atoi(bankJoltageString)
		joltage += bankJoltageValue
	}

	fmt.Println("Result of Day 3 - Part 2 is", joltage)
}
