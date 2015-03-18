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
	A, B, K int
}

func (problem *Problem) Solve() string {
	return fmt.Sprint(getN(problem.A-1, problem.B-1, problem.K-1))
}

func getN(A, B, K int) int {
	if K >= A && K >= B {
		return (A + 1) * (B + 1)
	}
	m := max(A, B)
	msb := msb(m)

	bA := A >> uint(msb)
	bB := B >> uint(msb)
	bK := K >> uint(msb)
	res := 0
	for a := 0; a <= bA; a++ {
		for b := 0; b <= bB; b++ {
			if a&b <= bK {
				var newA, newB, newK int
				if a == 1 {
					newA = A - pow(2, msb)
				} else {
					newA = min(A, pow(2, msb)-1)
				}
				if b == 1 {
					newB = B - pow(2, msb)
				} else {
					newB = min(B, pow(2, msb)-1)
				}
				if a&b == 1 {
					newK = K - pow(2, msb)
				} else {
					newK = min(K, pow(2, msb)-1)
				}
				res += getN(newA, newB, newK)
			}
		}
	}

	return res
}

func max(n ...int) int {
	m := 0
	for _, a := range n {
		if a > m {
			m = a
		}
	}
	return m
}

func min(n ...int) int {
	m := pow(2, 32)
	for _, a := range n {
		if a < m {
			m = a
		}
	}
	return m
}

func msb(n int) int {
	for i := 0; i < 32; i++ {
		if pow(2, i) > n {
			return i - 1
		}
	}
	return 0
}

func pow(a, b int) int {
	n := 1
	for i := 0; i < b; i++ {
		n *= a
	}
	return n
}

func getProblems(input *utils.Input) []*Problem {
	n := input.ReadInt()

	problems := make([]*Problem, n)
	for i := 0; i < n; i++ {
		problem := Problem{}
		problem.A = input.ReadInt()
		problem.B = input.ReadInt()
		problem.K = input.ReadInt()

		problems[i] = &problem
	}

	return problems
}
