package advent

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Day4() {
	dat, _ := os.ReadFile("./input/day4")
	input := string(dat)

	var grid [][]string

	rows := strings.Split(input, "\n")
	for i := range len(rows) {
		grid = append(grid, strings.Split(rows[i], ""))
	}

	start := time.Now()
	count := d4part1(grid)
	end := time.Now()

	fmt.Println("Part 1 / XMAS's: ", count)
	fmt.Println("Took ", end.Sub(start))

	start2 := time.Now()
	count2 := d4part2(grid)
	end2 := time.Now()

	fmt.Println("Part 2 / X-MAS's: ", count2)
	fmt.Println("Took ", end2.Sub(start2))

}

func d4part1(grid [][]string) int {
	count := 0
	for r := range len(grid) {
		for c := range len(grid[r]) {
			if grid[r][c] == "X" {
				if r > 2 && grid[r-1][c] == "M" && grid[r-2][c] == "A" && grid[r-3][c] == "S" { // vertical going up
					count++
				}
				if r < len(grid)-3 && grid[r+1][c] == "M" && grid[r+2][c] == "A" && grid[r+3][c] == "S" { // vertical going down
					count++
				}
				if c > 2 && grid[r][c-1] == "M" && grid[r][c-2] == "A" && grid[r][c-3] == "S" { // horizontal going left
					count++
				}
				if c < len(grid[r])-3 && grid[r][c+1] == "M" && grid[r][c+2] == "A" && grid[r][c+3] == "S" { // horizontal going right
					count++
				}
				if r > 2 && c > 2 && grid[r-1][c-1] == "M" && grid[r-2][c-2] == "A" && grid[r-3][c-3] == "S" { // diagonal going north-west
					count++
				}
				if r > 2 && c < len(grid[r])-3 && grid[r-1][c+1] == "M" && grid[r-2][c+2] == "A" && grid[r-3][c+3] == "S" { // diagonal going north-east
					count++
				}
				if r < len(grid)-3 && c < len(grid[r])-3 && grid[r+1][c+1] == "M" && grid[r+2][c+2] == "A" && grid[r+3][c+3] == "S" { // diagonal going south-east
					count++
				}
				if r < len(grid)-3 && c > 2 && grid[r+1][c-1] == "M" && grid[r+2][c-2] == "A" && grid[r+3][c-3] == "S" { // diagonal going south-west
					count++
				}

			}
		}
	}

	return count
}

func d4part2(grid [][]string) int {
	count := 0
	for r := range len(grid) {
		for c := range len(grid[r]) {
			if grid[r][c] == "A" && r > 0 && c > 0 && r < len(grid)-1 && c < len(grid[r])-1 {
				if grid[r-1][c-1] == "M" && grid[r+1][c-1] == "M" && grid[r+1][c+1] == "S" && grid[r-1][c+1] == "S" {
					count++
					// M . S
					// . A .
					// M . S
				} else if grid[r-1][c-1] == "S" && grid[r+1][c-1] == "S" && grid[r+1][c+1] == "M" && grid[r-1][c+1] == "M" {
					count++
					// S . M
					// . A .
					// S . M
				} else if grid[r-1][c-1] == "M" && grid[r+1][c-1] == "S" && grid[r+1][c+1] == "S" && grid[r-1][c+1] == "M" {
					count++
					// M . M
					// . A .
					// S . S
				} else if grid[r-1][c-1] == "S" && grid[r+1][c-1] == "M" && grid[r+1][c+1] == "M" && grid[r-1][c+1] == "S" {
					count++
					// S . S
					// . A .
					// M . M
				}
			}
		}
	}

	return count
}
