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
	barbers  []int64
	position int64
}

func (problem *Problem) Solve() string {
	return fmt.Sprint(problem.binSearch(0, (2<<60)) + 1)
}

func (problem *Problem) binSearch(first, last int64) int {
	m := (first + last) / 2
	clients, possibleBarbers := problem.getClientsOnM(m)
	if len(possibleBarbers) > 0 && problem.position <= clients+int64(len(possibleBarbers)) && problem.position > clients {
		return possibleBarbers[int(problem.position-clients)-1]
	}
	if clients >= problem.position {
		return problem.binSearch(first, m)
	}
	return problem.binSearch(m, last)
}

func (problem *Problem) getClientsOnM(m int64) (int64, []int) {
	var val int64
	possibleBarbers := []int{}
	for i, barber := range problem.barbers {
		val += (m - 1) / barber
		if m%barber == 0 {
			possibleBarbers = append(possibleBarbers, i)
		}
	}

	return val + int64(len(problem.barbers)), possibleBarbers
}

func max(list []int64) int64 {
	var m int64
	for _, n := range list {
		if n > m {
			m = n
		}
	}
	return m
}

func getProblems(input *utils.Input) []*Problem {
	n := input.ReadInt()

	problems := make([]*Problem, n)
	for i := 0; i < n; i++ {
		problem := Problem{}
		problem.barbers = make([]int64, input.ReadInt())
		problem.position = int64(input.ReadInt())
		for i := 0; i < len(problem.barbers); i++ {
			problem.barbers[i] = int64(input.ReadInt())
		}

		problems[i] = &problem
	}

	return problems
}
