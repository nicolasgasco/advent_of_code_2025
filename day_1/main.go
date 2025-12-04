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
	calcPassword(false)

	calcPassword(true)
}

func calcPassword(isSecurePassword bool) {
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

		if isSecurePassword {
			count += int(value / maxNum)
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

		if isSecurePassword {
			wentThroughZero := rawNewNum < minNum || rawNewNum > maxNum
			isAtZero := currentNum == 0
			if !isAtZero && wentThroughZero {
				count++
			}
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

	if isSecurePassword {
		fmt.Printf("The solution with complex password is: %d\n", count)
	} else {
		fmt.Printf("The solution with simple password is: %d\n", count)
	}
}
