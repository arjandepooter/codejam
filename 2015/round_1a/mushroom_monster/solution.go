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
	plates []int
}

func (problem *Problem) Solve() string {
	r1 := 0
	max := 0
	for i, p2 := range problem.plates[1:] {
		p1 := problem.plates[i]
		if p2 < p1 {
			r1 += p1 - p2
			if p1-p2 > max {
				max = p1 - p2
			}
		}
	}
	r2 := -0
	for i, _ := range problem.plates[1:] {
		p1 := problem.plates[i]
		if p1 < max {
			r2 += p1
		} else {
			r2 += max
		}
	}

	return fmt.Sprintf("%d %d", r1, r2)
}

func getProblems(input *utils.Input) []*Problem {
	n := input.ReadInt()

	problems := make([]*Problem, n)
	for i := 0; i < n; i++ {
		problem := Problem{}
		l := input.ReadInt()
		problem.plates = input.ReadInts(l)
		problems[i] = &problem
	}

	return problems
}
