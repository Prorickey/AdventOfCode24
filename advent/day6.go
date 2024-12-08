package advent

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Day6() {
	dat, _ := os.ReadFile("./input/day6")
	input := string(dat)

	tempGrid := strings.Split(input, "\n")
	var grid [][]string
	for i := range len(tempGrid) {
		grid = append(grid, strings.Split(tempGrid[i], ""))
	}

	direction := 4
	var loc []int
	var barrier [][]int
	for r := range len(grid) {
		for c := range len(grid[r]) {
			switch grid[r][c] {
			case "#":
				barrier = append(barrier, []int{r, c})
			case "^":
				loc = []int{r, c}
			}
		}
	}

	start := time.Now()
	num := d6part1(direction, loc, barrier, len(grid[0]), len(grid))
	end := time.Now()

	fmt.Println("Part 1 / Distinct Locations: ", num)
	fmt.Println("Took ", end.Sub(start))

	start2 := time.Now()
	num2 := d6part2(direction, loc, barrier, len(grid[0]), len(grid))
	end2 := time.Now()

	fmt.Println("Part 2 / Possibilities: ", num2)
	fmt.Println("Took ", end2.Sub(start2))

}

func Contains(visited [][]int, location []int) bool {
	for i := range len(visited) {
		if visited[i][0] == location[0] && visited[i][1] == location[1] {
			return true
		}
	}
	return false
}

func containsWithDirection(visited [][]int, location []int) bool {
	for i := range len(visited) {
		if visited[i][0] == location[0] && visited[i][1] == location[1] && visited[i][2] == location[2] {
			return true
		}
	}
	return false
}

func d6part1(direction int, loc []int, barrier [][]int, rowLength int, columnLength int) int {
	var visited [][]int
	running := true

	for running {

		r := loc[0]
		c := loc[1]

		if !Contains(visited, []int{r, c}) {
			visited = append(visited, []int{r, c})
		}

		switch direction {
		case 4:
			{
				if r == 0 {
					running = false
				} else if !Contains(barrier, []int{r - 1, c}) {
					loc = []int{r - 1, c}
				} else {
					direction = 5
				}
			}

		case 5:
			{
				if c == rowLength-1 {
					running = false
				} else if !Contains(barrier, []int{r, c + 1}) {
					loc = []int{r, c + 1}
				} else {
					direction = 6
				}
			}

		case 6:
			{
				if r == columnLength-1 {
					running = false
				} else if !Contains(barrier, []int{r + 1, c}) {
					loc = []int{r + 1, c}
				} else {
					direction = 7
				}
			}

		case 7:
			{
				if c == 0 {
					running = false
				} else if !Contains(barrier, []int{r, c - 1}) {
					loc = []int{r, c - 1}
				} else {
					direction = 4
				}
			}

		}
	}

	return len(visited)
}

func d6part2(defaultDirection int, defaultLoc []int, defaultBarrier [][]int, rowLength int, columnLength int) int {

	num := 0

	for row := range rowLength {
		for col := range columnLength {
			if !Contains(defaultBarrier, []int{row, col}) && !(defaultLoc[0] == row && defaultLoc[1] == col) {
				barrier := append(defaultBarrier, []int{row, col})

				loc := defaultLoc
				direction := defaultDirection
				var visited [][]int
				running := true

				for running {

					r := loc[0]
					c := loc[1]

					if containsWithDirection(visited, []int{r, c, direction}) {
						num++
						running = false
						break
					} else {
						visited = append(visited, []int{r, c, direction})
					}

					switch direction {
					case 4:
						{
							if r == 0 {
								running = false
							} else if !Contains(barrier, []int{r - 1, c}) {
								loc = []int{r - 1, c}
							} else {
								direction = 5
							}
						}

					case 5:
						{
							if c == rowLength-1 {
								running = false
							} else if !Contains(barrier, []int{r, c + 1}) {
								loc = []int{r, c + 1}
							} else {
								direction = 6
							}
						}

					case 6:
						{
							if r == columnLength-1 {
								running = false
							} else if !Contains(barrier, []int{r + 1, c}) {
								loc = []int{r + 1, c}
							} else {
								direction = 7
							}
						}

					case 7:
						{
							if c == 0 {
								running = false
							} else if !Contains(barrier, []int{r, c - 1}) {
								loc = []int{r, c - 1}
							} else {
								direction = 4
							}
						}

					}
				}
			}
		}
	}

	return num
}
