package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const startingNum = 50
const minNum = 0
const maxNum = 99 + 1 // +1 to compensate for 100 being equal to 0

func main() {
	partOne()

	partTwo()
}

func partOne() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	rotations := strings.Split(input, "\n")

	count := 0
	currentNum := startingNum
	for _, rotation := range rotations {
		valueStr := strings.TrimSpace(rotation[1:])
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			panic(err)
		}
		value = value % maxNum // in case of large rotations

		rawNewNum := currentNum

		direction := rune(rotation[0])
		switch direction {
		case 'L':
			rawNewNum -= value
		case 'R':
			rawNewNum += value
		default:
			panic("unknown direction")
		}

		rawNewNum = rawNewNum % maxNum

		if rawNewNum < minNum {
			currentNum = maxNum + rawNewNum
		} else if rawNewNum > minNum {
			currentNum = rawNewNum % maxNum
		} else {
			currentNum = rawNewNum
			count++
		}
	}

	fmt.Printf("The solution to part one of day 1 is: %d\n", count)
}

func partTwo() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(data)
	rotations := strings.Split(input, "\n")

	count := 0
	currentNum := startingNum
	for _, rotation := range rotations {
		valueStr := strings.TrimSpace(rotation[1:])
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			panic(err)
		}

		count += int(value / maxNum)
		value = value % maxNum // in case of large rotations

		rawNewNum := currentNum

		direction := rune(rotation[0])
		switch direction {
		case 'L':
			rawNewNum -= value
		case 'R':
			rawNewNum += value
		default:
			panic("unknown direction")
		}

		wentThroughZero := rawNewNum < minNum || rawNewNum > maxNum
		isAtZero := currentNum == 0
		if !isAtZero && wentThroughZero {
			count++
		}

		rawNewNum = rawNewNum % maxNum

		if rawNewNum < minNum {
			currentNum = maxNum + rawNewNum
		} else if rawNewNum > minNum {
			currentNum = rawNewNum % maxNum
		} else {
			currentNum = rawNewNum
			count++
		}
	}

	fmt.Printf("The solution to part two of day 1 is: %d\n", count)
}
