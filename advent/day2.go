package advent

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func Day2() {
	dat, _ := os.ReadFile("./input/day2")
	input := string(dat)

	reports := [][]int{}

	split := strings.Split(input, "\n")
	for _, line := range split {
		report := []int{}
		splitLine := strings.Split(line, " ")
		for _, value := range splitLine {
			num, _ := strconv.Atoi(value)
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	start := time.Now()
	count := d2part1(reports)
	end := time.Now()

	fmt.Println("Part 1 / Safe reports: ", count)
	fmt.Println("Took ", end.Sub(start))

	start2 := time.Now()
	count2 := d2part2(reports)
	end2 := time.Now()

	fmt.Println("Part 2 / Safe reports with removal: ", count2)
	fmt.Println("Took ", end2.Sub(start2))
}

func d2part1(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		}
	}

	return safeReports
}

func isSafe(report []int) bool {
	safe := true
	var inc *bool

	for j := 1; j < len(report); j++ {
		if inc != nil {
			if (report[j] > report[j-1] && !*inc) || (report[j] < report[j-1] && *inc) {
				safe = false
				fmt.Println(report)
				break
			}
		} else {
			if report[j] > report[j-1] {
				inc = new(bool)
				*inc = true
			} else if report[j] < report[j-1] {
				inc = new(bool)
				*inc = false
			}
		}

		if math.Abs(float64(report[j]-report[j-1])) > 3 || math.Abs(float64(report[j]-report[j-1])) < 1 {
			safe = false
			break
		}
	}

	return safe
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func d2part2(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		} else {
			fmt.Println("TEst")
			fmt.Println(report)
			for i := range len(report) {
				newReport := remove(report, i)
				fmt.Println(newReport)
				if isSafe(newReport) {
					safeReports++
					break
				}
			}
		}
	}

	return safeReports
}
