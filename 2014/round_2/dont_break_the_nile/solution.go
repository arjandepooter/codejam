package main

import (
	"fmt"
	"github.com/arjandepooter/codejam/utils"
)

func main() {
	input, err := utils.GetInput()
	if err != nil {
		fmt.Errorf("Error occured: %x", err)
	}

	problems := getProblems(input)

	for i, problem := range problems {
		fmt.Printf("Case #%d: %s\n", i+1, problem.Solve())
	}
}

type Problem struct {
	width     int
	height    int
	buildings []*Building
}

type Building struct {
	x0, x1, y0, y1 int
}

type Direction struct {
	x, y int
}

func (problem *Problem) Solve() string {
	result := 0
	grid := problem.GetGrid()
	for i := 0; i < problem.width; i++ {
		if walkGrid(&grid, i, 0, Direction{0, 1}) {
			result++
		}
	}
	return fmt.Sprint(result)
}

func (problem *Problem) GetGrid() [][]bool {
	grid := make([][]bool, problem.height)
	for i := 0; i < problem.height; i++ {
		grid[i] = make([]bool, problem.width)
	}
	for _, building := range problem.buildings {
		for x := building.x0; x <= building.x1; x++ {
			for y := building.y0; y <= building.y1; y++ {
				grid[y][x] = true
			}
		}
	}

	return grid
}

func (direction *Direction) next() Direction {
	result := Direction{}
	result.x, result.y = direction.y, -1*direction.x
	return result
}

func walkGrid(grid *[][]bool, x, y int, dir Direction) bool {
	if y >= len(*grid) {
		return true
	}
	if x < 0 || y < 0 || x >= len((*grid)[0]) {
		return false
	}
	if (*grid)[y][x] {
		return false
	}

	(*grid)[y][x] = true

	dir.x, dir.y = -1*dir.x, -1*dir.y
	for i := 0; i < 3; i++ {
		dir = dir.next()
		if walkGrid(grid, x+dir.x, y+dir.y, dir) {
			return true
		}
	}

	return false
}

func getProblems(input *utils.Input) []*Problem {
	n := input.ReadInt()

	problems := make([]*Problem, n)
	for i := 0; i < n; i++ {
		problem := Problem{}
		problem.width = input.ReadInt()
		problem.height = input.ReadInt()
		problem.buildings = make([]*Building, input.ReadInt())

		for i := 0; i < len(problem.buildings); i++ {
			b := Building{}
			b.x0 = input.ReadInt()
			b.y0 = input.ReadInt()
			b.x1 = input.ReadInt()
			b.y1 = input.ReadInt()
			problem.buildings[i] = &b
		}

		problems[i] = &problem
	}

	return problems
}
