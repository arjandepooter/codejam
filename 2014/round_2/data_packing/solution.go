package main

import (
	"fmt"
	"github.com/arjandepooter/codejam/utils"
	"sort"
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
	capacity int
	sizes    []int
}

func (problem *Problem) Solve() string {
	sizes := problem.sizes
	sort.Ints(sizes)
	r := 0
	for len(sizes) > 0 {
		cur := sizes[len(sizes)-1]
		m := problem.capacity - cur
		bi := -1
		for i, n := range sizes[:len(sizes)-1] {
			if n > m {
				break
			}
			bi = i
		}
		if bi == -1 {
			sizes = sizes[:len(sizes)-1]
		} else {
			sizes = append(sizes[:bi], sizes[bi+1:len(sizes)-1]...)
		}
		r++
	}
	return fmt.Sprint(r)
}

func getProblems(input *utils.Input) []*Problem {
	n := input.ReadInt()

	problems := make([]*Problem, n)
	for i := 0; i < n; i++ {
		problem := Problem{}
		problem.sizes = make([]int, input.ReadInt())
		problem.capacity = input.ReadInt()

		for i := 0; i < len(problem.sizes); i++ {
			problem.sizes[i] = input.ReadInt()
		}

		problems[i] = &problem
	}

	return problems
}
