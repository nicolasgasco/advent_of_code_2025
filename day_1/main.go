package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PasswordMode int

const (
	SimplePassword PasswordMode = iota
	SecurePassword
)

const startingNum = 50
const minNum = 0
const maxNum = 99 + 1 // +1 to compensate for 100 being equal to 0

func main() {
	count, err := calculatePassword(SimplePassword)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Solution for Part 1: %d\n", count)

	count, err = calculatePassword(SecurePassword)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Solution for Part 2: %d\n", count)
}

func calculatePassword(mode PasswordMode) (uint, error) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}

	input := string(data)
	rotations := strings.Split(input, "\n")

	var count uint = 0
	currentNum := startingNum
	for _, rotation := range rotations {
		valueStr := strings.TrimSpace(rotation[1:])
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			panic(err)
		}

		if mode == SecurePassword {
			count += uint(value / maxNum) // in case there are several full rotations
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
			return 0, errors.New("invalid rotation direction")
		}

		if mode == SecurePassword {
			wentThroughZero := rawNewNum < minNum || rawNewNum > maxNum
			isAtZero := currentNum == 0
			if !isAtZero && wentThroughZero {
				count++
			}
		}
		rawNewNum = rawNewNum % maxNum

		if rawNewNum < minNum {
			currentNum = maxNum - (-rawNewNum)
		} else if rawNewNum > minNum {
			currentNum = rawNewNum % maxNum
		} else {
			currentNum = rawNewNum
			count++
		}
	}

	return count, nil
}
