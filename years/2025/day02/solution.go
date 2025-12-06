package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := GetInput()

	part1Sum := Part1(input)
	part2Sum := Part2(input)
	fmt.Printf("Part 1 Sum: %v\n", part1Sum)
	fmt.Printf("Part 2 Sum: %v\n", part2Sum)
}

func Part1(input string) int {
	sum := 0
	sections := strings.Split(input, ",")

	for _, section := range sections {
		bounds := strings.Split(section, "-")
		lowerBound, errL := strconv.ParseInt(bounds[0], 10, 64)
		upperBound, errU := strconv.ParseInt(bounds[1], 10, 64)
		if errL != nil {
			fmt.Printf("Error parsing lower bound: %v\n", errL)
			return -1
		}
		if errU != nil {
			fmt.Printf("Error parsing lower bound: %v\n", errU)
			return -1
		}

		lowerBoundI := int(lowerBound)
		upperBoundI := int(upperBound)
		for i := lowerBoundI; i <= upperBoundI; i++ {
			chars := strconv.FormatInt(int64(i), 10)
			if len(chars)%2 != 0 {
				continue
			}
			firstHalf := chars[:len(chars)/2]
			secondHalf := chars[len(chars)/2:]
			if firstHalf == secondHalf {
				sum += i
			}
		}
	}

	return sum
}

func Part2(input string) int {
	sum := 0
	sections := strings.Split(input, ",")

	for _, section := range sections {
		bounds := strings.Split(section, "-")
		lowerBound, errL := strconv.ParseInt(bounds[0], 10, 64)
		upperBound, errU := strconv.ParseInt(bounds[1], 10, 64)
		if errL != nil {
			fmt.Printf("Error parsing lower bound: %v\n", errL)
			return -1
		}
		if errU != nil {
			fmt.Printf("Error parsing lower bound: %v\n", errU)
			return -1
		}

		lowerBoundI := int(lowerBound)
		upperBoundI := int(upperBound)
		for i := lowerBoundI; i <= upperBoundI; i++ {
			chars := strconv.FormatInt(int64(i), 10)
			for a := 2; a <= len(chars); a++ {
				if len(chars)%a != 0 {
					continue
				}

				partLength := len(chars) / a
				parts := make([]string, a)
				for j := 0; j < a; j++ {
					parts[j] = chars[j*partLength : (j+1)*partLength]
				}
				valid := true
				for j := 1; j < a; j++ {
					if parts[j] != parts[j-1] {
						valid = false
						break
					}
				}
				if valid {
					sum += i
					break
				}
			}
		}
	}

	return sum
}

func GetInput() string {
	content, err := os.ReadFile("years/2025/day02/input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return ""
	}
	return strings.ReplaceAll(string(content), "\r", "")
}
