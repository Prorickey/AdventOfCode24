package advent

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func Day7() {
	dat, _ := os.ReadFile("./input/day7")
	input := string(dat)

	var data [][][]int
	split := strings.Split(input, "\n")
	for _, line := range split {
		split2 := strings.Split(line, ": ")
		n2, _ := strconv.Atoi(split2[0])
		answer := []int{n2}
		var nums []int
		split3 := strings.Split(split2[1], " ")
		for _, n := range split3 {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}
		data = append(data, [][]int{answer, nums})
	}

	start := time.Now()
	num := d7part1(data)
	end := time.Now()

	fmt.Println("Part 1 / Sum of bs: ", num)
	fmt.Println("Took ", end.Sub(start))

	start2 := time.Now()
	num2 := d7part2(data)
	end2 := time.Now()

	fmt.Println("Part 2 / Sum of bs again: ", num2)
	fmt.Println("Took ", end2.Sub(start2))
}

func genCombos(length int) [][]int {
	total := 1 << length
	result := make([][]int, total)

	for i := 0; i < total; i++ {
		combination := make([]int, length)
		for j := 0; j < length; j++ {
			if i&(1<<(length-j-1)) != 0 {
				combination[j] = 1
			} else {
				combination[j] = 0
			}
		}
		result[i] = combination
	}

	return result
}

func d7part1(data [][][]int) int {
	sum := 0
	for _, eq := range data {
		possible := false
		answer := eq[0][0]
		nums := eq[1]
		combos := genCombos(len(nums) - 1)
		for _, combo := range combos {
			result := nums[0]
			for i, o := range combo {
				if o == 0 {
					result = result + nums[i+1]
				} else {
					result = result * nums[i+1]
				}
			}
			if result == answer {
				possible = true
				break
			}
		}
		if possible {
			sum += answer
		}
	}

	return sum
}

func genCombos2(length int) [][]int {
	total := int(math.Pow(3, float64(length)))
	result := make([][]int, total)

	for i := 0; i < total; i++ {
		combination := make([]int, length)
		val := i
		for j := length - 1; j >= 0; j-- {
			combination[j] = val % 3
			val /= 3
		}
		result[i] = combination
	}

	return result
}

func d7part2(data [][][]int) int {
	sum := 0
	for _, eq := range data {
		possible := false
		answer := eq[0][0]
		nums := eq[1]
		combos := genCombos2(len(nums) - 1)
		for _, combo := range combos {
			result := nums[0]
			for i, o := range combo {
				if o == 0 {
					result = result + nums[i+1]
				} else if o == 1 {
					result = result * nums[i+1]
				} else {
					t, _ := strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(nums[i+1]))
					result = t
				}
			}
			if result == answer {
				possible = true
				break
			}
		}
		if possible {
			sum += answer
		}
	}

	return sum
}
