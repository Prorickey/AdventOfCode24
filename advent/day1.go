package advent

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Day1() {
	dat, _ := os.ReadFile("./input/day1")
	input := string(dat)

	var list1 []int
	var list2 []int

	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(split[0])
		num2, _ := strconv.Atoi(split[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	start := time.Now()
	count := d1part1(list1, list2)
	end := time.Now()

	fmt.Println("Part 1 / Cumulative Difference: ", count)
	fmt.Println("Took ", end.Sub(start))

	start2 := time.Now()
	count2 := d1part2(list1, list2)
	end2 := time.Now()

	fmt.Println("Part 2 / Similarity Score: ", count2)
	fmt.Println("Took ", end2.Sub(start2))
}

func d1part1(list1 []int, list2 []int) int {
	totalDiff := 0
	for i, num := range list1 {
		totalDiff += int(math.Abs(float64(num - list2[i])))
	}

	return totalDiff
}

func d1part2(list1 []int, list2 []int) int {
	similarity := 0
	for _, num := range list1 {
		occurrences := 0
		for _, num2 := range list2 {
			if num2 == num {
				occurrences++
			}
		}
		similarity += num * occurrences
	}

	return similarity
}
