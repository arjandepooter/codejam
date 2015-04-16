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
	return fmt.Sprint(getSwaps(problem.list))
}

func getSwaps(list []int) int {
	if len(list) <= 2 {
		return 0
	}
	minI := min(list)
	s := getSwaps(append(list[:minI], list[minI+1:]...))
	if minI < len(list)-minI-1 {
		return s + minI
	}
	return s + (len(list) - minI - 1)
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
