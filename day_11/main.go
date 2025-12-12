package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := string(data)
	rows := strings.Split(input, "\n")

	instructions := make(map[string][]string, len(rows))
	for _, row := range rows {
		chunks := strings.Split(row, ": ")

		machine := chunks[0]
		outputs := strings.Split(chunks[1], " ")

		instructions[machine] = outputs
	}

	startOutputs := instructions["you"]
	sum := followOutputs(&startOutputs, &instructions)

	fmt.Printf("Result of Dat 11 - Part 1 is %d\n", sum)
}

func followOutputs(outputs *[]string, instructions *map[string][]string) int {
	sum := 0

	if len(*outputs) == 1 && (*outputs)[0] == "out" {
		return sum + 1
	}

	for _, output := range *outputs {
		innerOutputs := (*instructions)[output]
		sum += followOutputs(&innerOutputs, instructions)
	}

	return sum
}
