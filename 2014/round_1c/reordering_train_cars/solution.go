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
		fmt.Printf("Case #%d: %d\n", i+1, problem.Solve())
	}
}

const CHARS string = "abcdefghijklmnopqrstuvwxyz"

type Problem struct {
	sets []*set
}

type set []rune

func (problem Problem) Solve() int {
	uniformSets := map[rune]int{}
	for i := 0; i < len(problem.sets); {
		if problem.sets[i].isUniform() {
			uniformSets[(*problem.sets[i])[0]]++
			problem.sets = append(problem.sets[:i], problem.sets[i+1:]...)
		} else {
			i++
		}
	}

	subSets := [][]*set{}
	for len(problem.sets) > 0 {
		start := problem.sets[0]
		subSet := []*set{start}
		current := start
		problem.sets = problem.sets[1:]

		for {
			tails := getPossibleTails(current, problem.sets)
			if len(tails) > 1 {
				return 0
			}
			if len(tails) == 0 {
				break
			}
			i := tails[0]
			subSet = append(subSet, problem.sets[i])
			current = problem.sets[i]
			problem.sets = append(problem.sets[:i], problem.sets[i+1:]...)
		}
		current = start
		for {
			heads := getPossibleHeads(current, problem.sets)
			if len(heads) > 1 {
				return 0
			}
			if len(heads) == 0 {
				break
			}
			i := heads[0]
			subSet = append([]*set{problem.sets[i]}, subSet...)
			current = problem.sets[i]
			problem.sets = append(problem.sets[:i], problem.sets[i+1:]...)
		}

		subSets = append(subSets, subSet)
	}

	usedChars := map[rune]bool{}
	for _, c := range CHARS {
		usedChars[c] = false
	}

	c := []rune{}
	for _, subSet := range subSets {
		for _, s := range subSet {
			c = append(c, *s...)
		}
	}

	if len(c) > 0 {
		cur := c[0]
		usedChars[cur] = true
		for _, r := range c[1:] {
			if r == cur {
				continue
			}
			if usedChars[r] {
				return 0
			}
			usedChars[r] = true
			cur = r
		}
	}

	l := len(subSets)
	result := 1
	for c, n := range uniformSets {
		if !usedChars[c] {
			l++
		}
		result = (result * fac(n, 1000000007)) % 1000000007
	}
	result = (result * fac(l, 1000000007)) % 1000000007

	return result
}

func fac(n, m int) int {
	if n == 0 {
		return 1
	}
	return (n * (fac(n-1, m) % m)) % m
}

func (s set) isUniform() bool {
	char := s[0]
	for _, c := range s[1:] {
		if c != char {
			return false
		}
	}
	return true
}

func getPossibleHeads(s *set, heads []*set) []int {
	result := []int{}
	for i, head := range heads {
		if (*head)[len(*head)-1] == (*s)[0] {
			result = append(result, i)
		}
	}
	return result
}

func getPossibleTails(s *set, tails []*set) []int {
	result := []int{}
	for i, tail := range tails {
		if (*tail)[0] == (*s)[len(*s)-1] {
			result = append(result, i)
		}
	}
	return result
}

func getProblems(input *utils.Input) []*Problem {
	n := input.ReadInt()

	problems := make([]*Problem, n)
	for i := 0; i < n; i++ {
		problem := Problem{}

		s := input.ReadInt()
		problem.sets = make([]*set, s)
		for j := 0; j < s; j++ {
			l := set(input.ReadString())
			problem.sets[j] = &l
		}

		problems[i] = &problem
	}

	return problems
}
