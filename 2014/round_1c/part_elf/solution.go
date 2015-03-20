package main

import (
	"fmt"
	"github.com/arjandepooter/codejam/utils"
	"strconv"
	"strings"
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

func gcd(a, b int64) int64 {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if b > a {
		a, b = b, a
	}

	return gcd(b, a%b)
}

func lcm(a, b int64) int64 {
	return (a * b) / gcd(a, b)
}

func isPowerOf2(a int64) bool {
	return a == 1 || (a%2 == 0 && isPowerOf2(a/2))
}

type Problem struct {
	P, Q int64
}

func (problem Problem) Solve() string {
	if !isPowerOf2(problem.Q) {
		return "impossible"
	}

	n := 0
	for problem.P < problem.Q {
		problem.P *= 2
		n++
	}

	return fmt.Sprintf("%d", n)
}

func getProblems(input *utils.Input) []*Problem {
	n := input.ReadInt()

	problems := make([]*Problem, n)
	for i := 0; i < n; i++ {
		problem := Problem{}

		data := strings.Split(input.ReadString(), "/")
		P, _ := strconv.ParseInt(data[0], 10, 64)
		Q, _ := strconv.ParseInt(data[1], 10, 64)
		g := gcd(P, Q)
		P = P / g
		Q = Q / g
		problem.P = P
		problem.Q = Q

		problems[i] = &problem
	}

	return problems
}
