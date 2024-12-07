package advent

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func Day5() {
	dat, _ := os.ReadFile("./input/day5")
	input := string(dat)

	sections := strings.Split(input, "\n\n")
	rawRules := strings.Split(sections[0], "\n")
	rawPages := strings.Split(sections[1], "\n")
	pages := [][]int{}
	for i := range len(rawPages) {
		page := []int{}
		rawPage := strings.Split(rawPages[i], ",")
		for j := range len(rawPage) {
			num, _ := strconv.Atoi(rawPage[j])
			page = append(page, num)
		}
		pages = append(pages, page)
	}

	rules := map[int][]int{} // Wow I hate this
	for i := range len(rawRules) {
		split := strings.Split(rawRules[i], "|")
		first, _ := strconv.Atoi(split[0])
		second, _ := strconv.Atoi(split[1])

		if rules[first] == nil {
			rules[first] = []int{
				second,
			}
		} else {
			rules[first] = append(rules[first], second)
		}
	}

	start := time.Now()
	sum := d5part1(rules, pages)
	end := time.Now()

	fmt.Println("Part 1 / Valid Pages Median Sum: ", sum)
	fmt.Println("Took ", end.Sub(start))

	start2 := time.Now()
	sum2 := d5part2(rules, pages)
	end2 := time.Now()

	fmt.Println("Part 2 / Pages Median Sum: ", sum2)
	fmt.Println("Took ", end2.Sub(start2))

}

func d5part1(rules map[int][]int, pages [][]int) int {
	validPages := [][]int{}
	for i := range len(pages) {
		valid := true
		before := []int{}
		page := pages[i]
		for j := range len(page) {
			num := page[j]
			for r := range len(rules[num]) {
				rule := rules[num][r] // This number can't be in before
				if slices.Contains(before, rule) {
					valid = false
					break
				}
			}
			before = append(before, num)
		}

		if valid {
			validPages = append(validPages, page)
		}
	}

	sum := 0
	for i := range len(validPages) {
		page := validPages[i]
		sum += page[(len(page)-1)/2]
	}

	return sum
}

func checkAndCorrectlyOrder(rules map[int][]int, page []int) []int {
	before := []int{}
	for j := range len(page) {
		num := page[j]
		for r := range len(rules[num]) {
			rule := rules[num][r] // This number can't be in before
			if slices.Contains(before, rule) {
				// put num before rule in page array
				newPage := []int{}
				for n := range len(page) {
					if page[n] == rule {
						newPage = append(newPage, num)
						newPage = append(newPage, rule)
					} else if page[n] != num {
						newPage = append(newPage, page[n])
					}
				}
				return checkAndCorrectlyOrder(rules, newPage)
			}
		}
		before = append(before, num)
	}

	return page
}

func d5part2(rules map[int][]int, pages [][]int) int {
	sum := 0
	for i := range len(pages) {
		before := []int{}
		page := pages[i]
		for j := range len(page) {
			num := page[j]
			valid := true
			for r := range len(rules[num]) {
				rule := rules[num][r] // This number can't be in before
				if slices.Contains(before, rule) {
					valid = false
					break
				}
			}
			before = append(before, num)
			if !valid {
				fixedPage := checkAndCorrectlyOrder(rules, pages[i])
				sum += fixedPage[(len(fixedPage)-1)/2]
				break
			}
		}
	}

	return sum
}
