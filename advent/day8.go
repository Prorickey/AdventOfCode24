package advent

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Day8() {
	dat, _ := os.ReadFile("./input/day8")
	input := string(dat)

	grid := [][]string{}
	split := strings.Split(input, "\n")
	for _, line := range split {
		var row []string
		rawRow := strings.Split(line, "")
		for _, space := range rawRow {
			row = append(row, space)
		}
		grid = append(grid, row)
	}

	start := time.Now()
	num := d8part1(grid)
	end := time.Now()

	fmt.Println("Part 1 / : ", num)
	fmt.Println("Took ", end.Sub(start))

	start2 := time.Now()
	num2 := d8part2(grid)
	end2 := time.Now()

	fmt.Println("Part 2 / : ", num2)
	fmt.Println("Took ", end2.Sub(start2))
}

func d8part1(grid [][]string) int {

	var locs []location

	for r := range grid {
		for c, loc := range grid[r] {
			if loc != "." {
				locs = append(locs, location{
					letter:   loc,
					location: []int{r, c},
				})
			}
		}
	}

	var individualLocs [][]int

	for i, loc := range locs {
		for j, loc2 := range locs {
			if i != j && loc2.letter == loc.letter {

				rise := loc.location[0] - loc2.location[0]
				run := loc.location[1] - loc2.location[1]

				if rise < 0 {
					break
				}

				// test point on loc
				if loc.location[0]+rise >= 0 && loc.location[0]+rise < len(grid) && loc.location[1]+run < len(grid[0]) && loc.location[1]+run >= 0 {
					if !Contains(individualLocs, []int{loc.location[0] + rise, loc.location[1] + run}) {
						individualLocs = append(individualLocs, []int{loc.location[0] + rise, loc.location[1] + run})
					}
				}

				// test point on loc2
				if loc2.location[0]-rise >= 0 && loc2.location[0]-rise < len(grid) && loc2.location[1]-run < len(grid[0]) && loc2.location[1]-run >= 0 {
					if !Contains(individualLocs, []int{loc2.location[0] - rise, loc2.location[1] - run}) {
						individualLocs = append(individualLocs, []int{loc2.location[0] - rise, loc2.location[1] - run})
					}
				}
			}
		}
	}

	return len(individualLocs)
}

type location struct {
	letter   string
	location []int
}

func d8part2(grid [][]string) int {
	var locs []location

	for r := range grid {
		for c, loc := range grid[r] {
			if loc != "." {
				locs = append(locs, location{
					letter:   loc,
					location: []int{r, c},
				})
			}
		}
	}

	var individualLocs [][]int

	for i, loc := range locs {
		for j, loc2 := range locs {
			if i != j && loc2.letter == loc.letter {

				rise := loc.location[0] - loc2.location[0]
				run := loc.location[1] - loc2.location[1]

				if rise < 0 {
					break
				}

				// test point on loc
				cont := true
				n := 0
				for cont {
					if loc.location[0]+(n*rise) >= 0 && loc.location[0]+(n*rise) < len(grid) && loc.location[1]+(n*run) < len(grid[0]) && loc.location[1]+(n*run) >= 0 {
						if !Contains(individualLocs, []int{loc.location[0] + (n * rise), loc.location[1] + (n * run)}) {
							individualLocs = append(individualLocs, []int{loc.location[0] + (n * rise), loc.location[1] + (n * run)})
						}
						n++
					} else {
						cont = false
					}
				}

				// test point on loc2
				cont2 := true
				n = 0
				for cont2 {
					if loc2.location[0]-(n*rise) >= 0 && loc2.location[0]-(n*rise) < len(grid) && loc2.location[1]-(n*run) < len(grid[0]) && loc2.location[1]-(n*run) >= 0 {
						if !Contains(individualLocs, []int{loc2.location[0] - (n * rise), loc2.location[1] - (n * run)}) {
							individualLocs = append(individualLocs, []int{loc2.location[0] - (n * rise), loc2.location[1] - (n * run)})
						}
						n++
					} else {
						cont2 = false
					}
				}
			}
		}
	}

	return len(individualLocs)
}
