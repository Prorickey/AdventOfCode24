package advent

import (
	"fmt"
	"os"
	"strings"
	"sync"
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

type count struct {
	mu    sync.Mutex
	count int
}

func d6part2(defaultDirection int, defaultLoc []int, defaultBarrier [][]int, rowLength int, columnLength int) int {

	var wg sync.WaitGroup
	num := count{
		mu:    sync.Mutex{},
		count: 0,
	}

	for row := range rowLength {
		for col := range columnLength {
			if !Contains(defaultBarrier, []int{row, col}) && !(defaultLoc[0] == row && defaultLoc[1] == col) {
				wg.Add(1)
				go func(row int, col int) {
					defer wg.Done()
					barrier := make(map[string]struct{})
					barrier[fmt.Sprintf("%d,%d", row, col)] = struct{}{}
					for _, barr := range defaultBarrier {
						barrier[fmt.Sprintf("%d,%d", barr[0], barr[1])] = struct{}{}
					}

					loc := defaultLoc
					direction := defaultDirection
					visited := make(map[string]struct{})
					running := true

					for running {

						r, c := loc[0], loc[1]
						key := fmt.Sprintf("%d,%d,%d", r, c, direction)
						if _, found := visited[key]; found {
							num.mu.Lock()
							num.count++
							num.mu.Unlock()
							running = false
							break
						}
						visited[key] = struct{}{}

						switch direction {
						case 4:
							{
								if r == 0 {
									running = false
								} else if _, blocked := barrier[fmt.Sprintf("%d,%d", r-1, c)]; !blocked {
									loc = []int{r - 1, c}
								} else {
									direction = 5
								}
							}

						case 5:
							{
								if c == rowLength-1 {
									running = false
								} else if _, blocked := barrier[fmt.Sprintf("%d,%d", r, c+1)]; !blocked {
									loc = []int{r, c + 1}
								} else {
									direction = 6
								}
							}

						case 6:
							{
								if r == columnLength-1 {
									running = false
								} else if _, blocked := barrier[fmt.Sprintf("%d,%d", r+1, c)]; !blocked {
									loc = []int{r + 1, c}
								} else {
									direction = 7
								}
							}

						case 7:
							{
								if c == 0 {
									running = false
								} else if _, blocked := barrier[fmt.Sprintf("%d,%d", r, c-1)]; !blocked {
									loc = []int{r, c - 1}
								} else {
									direction = 4
								}
							}

						}
					}
				}(row, col)

			}
		}
	}

	wg.Wait()

	return num.count
}
