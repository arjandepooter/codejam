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
	levels []int
}

func (problem *Problem) Solve() string {
	added := 0
	total := 0
	for level, count := range problem.levels {
		if count > 0 && level > (added+total) {
			added += level - (added + total)
		}
		total += count
	}

	return fmt.Sprint(added)
}

func getProblems(input *utils.Input) []*Problem {
	n := input.ReadInt()

	problems := make([]*Problem, n)
	for i := 0; i < n; i++ {
		problem := Problem{}

		maxLevel := input.ReadInt()
		levelString := input.ReadString()
		problem.levels = make([]int, maxLevel+1)

		for j, r := range levelString {
			problem.levels[j] = int(r - '0')
		}

		problems[i] = &problem
	}

	return problems
}
