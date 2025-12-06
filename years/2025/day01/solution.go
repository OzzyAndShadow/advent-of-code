package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := GetInput()

	part1Password := Part1(input)
	part2Password := Part2(input)
	fmt.Printf("Part 1 Password: %v\n", part1Password)
	fmt.Printf("Part 2 Password: %v\n", part2Password)
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	dial := 50
	password := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		direction := line[0]
		i64, err := strconv.ParseInt(line[1:], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing step size: %v\n", err)
			return -1
		}

		currentStepSize := int(i64)
		tick := 1
		if direction == 'L' {
			tick = -1
		}
		for i := 0; i < currentStepSize; i++ {
			dial += tick
			if dial < 0 {
				dial = 99
			}
			if dial > 99 {
				dial = 0
			}
		}

		if dial == 0 {
			password++
		}
	}
	return password
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	dial := 50
	password := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		direction := line[0]
		i64, err := strconv.ParseInt(line[1:], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing step size: %v\n", err)
			return -1
		}

		currentStepSize := int(i64)
		tick := 1
		if direction == 'L' {
			tick = -1
		}
		for i := 0; i < currentStepSize; i++ {
			dial += tick
			if dial < 0 {
				dial = 99
			}
			if dial > 99 {
				dial = 0
			}
			if dial == 0 {
				password++
			}
		}
	}
	return password
}

func GetInput() string {
	content, err := os.ReadFile("years/2025/day01/input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return ""
	}
	return strings.ReplaceAll(string(content), "\r", "")
}
