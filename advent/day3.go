package advent

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func Day3() {
	dat, _ := os.ReadFile("./input/day3")
	input := string(dat)

	start := time.Now()
	count := d3part1(input)
	end := time.Now()

	fmt.Println("Part 1 / Multiplied: ", count)
	fmt.Println("Took ", end.Sub(start))

	start2 := time.Now()
	count2 := d3part2(input)
	end2 := time.Now()

	fmt.Println("Part 2 / Multiplied with Instructions: ", count2)
	fmt.Println("Took ", end2.Sub(start2))
}

func d3part1(input string) int {
	regex := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)

	matches := regex.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}

	return sum
}

func d3part2(input string) int {

	regex := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)|do\(\)|don\'t\(\)`)

	matches := regex.FindAllStringSubmatch(input, -1)

	sum := 0
	do := true
	for _, match := range matches {
		if match[0] == "do()" {
			do = true
		} else if match[0] == "don't()" {
			do = false
		} else if do {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			sum += num1 * num2
		}
	}

	return sum
}
