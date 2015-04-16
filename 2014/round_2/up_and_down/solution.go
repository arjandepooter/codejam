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
	list []int
}

func (problem Problem) Solve() string {
	l := problem.list
	s := 0
	for len(l) > 2 {
		minI := min(l)
		if minI < len(l)-minI-1 {
			s += minI
		} else {
			s += len(l) - minI - 1
		}
		l = append(l[:minI], l[minI+1:]...)
	}
	return fmt.Sprint(s)
}

func min(list []int) int {
	minI := 0
	for i, n := range list {
		if n < list[minI] {
			minI = i
		}
	}
	return minI
}

func getProblems(input *utils.Input) []*Problem {
	n := input.ReadInt()

	problems := make([]*Problem, n)
	for i := 0; i < n; i++ {
		problem := Problem{}
		problem.list = make([]int, input.ReadInt())
		for i := 0; i < len(problem.list); i++ {
			problem.list[i] = input.ReadInt()
		}

		problems[i] = &problem
	}

	return problems
}
