package main

import (
	"fmt"
	"github.com/arjandepooter/codejam/utils"
	"strconv"
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
	N         int
	L         int
	outlets   []int64
	devices   []int64
	deviceMap map[int64]bool
}

type patterns []int64

func (p patterns) Len() int {
	return len(p)
}

func (p patterns) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p patterns) Less(i, j int) bool {
	return p[i] < p[j]
}

func getProblems(input *utils.Input) []*Problem {
	n := input.ReadInt()

	problems := make([]*Problem, n)
	for i := 0; i < n; i++ {
		problems[i] = &Problem{}

		N := input.ReadInt()
		L := input.ReadInt()
		outletStrings := input.ReadStrings(N)
		deviceStrings := input.ReadStrings(N)

		problems[i].N = N
		problems[i].L = L
		problems[i].outlets = make(patterns, N)
		for j := 0; j < N; j++ {
			problems[i].outlets[j], _ = strconv.ParseInt(outletStrings[j], 2, 64)
		}
		problems[i].devices = make(patterns, N)
		problems[i].deviceMap = make(map[int64]bool)
		for j := 0; j < N; j++ {
			val, _ := strconv.ParseInt(deviceStrings[j], 2, 64)
			problems[i].devices[j] = val
			problems[i].deviceMap[val] = true
		}
	}

	return problems
}

func (problem *Problem) Solve() string {
	outlet := problem.outlets[0]
	switches := make(patterns, len(problem.devices))

	for i, device := range problem.devices {
		switches[i] = outlet ^ device
	}

	shortestSwitch := -1
	for _, swits := range switches {
		n := numberOfBits(swits)
		if (shortestSwitch == -1 || n < shortestSwitch) && problem.CheckSwitch(swits) {
			shortestSwitch = n
		}
	}

	if shortestSwitch == -1 {
		return "NOT POSSIBLE"
	}
	return fmt.Sprint(shortestSwitch)
}

func (problem *Problem) CheckSwitch(swits int64) bool {
	newOutlets := make(patterns, len(problem.outlets))
	for i, outlet := range problem.outlets {
		newOutlets[i] = outlet ^ swits
	}

	for _, outlet := range newOutlets {
		if _, ok := problem.deviceMap[outlet]; !ok {
			return false
		}
	}

	return true
}

func numberOfBits(n int64) int {
	b := 0
	for n > 0 {
		m := n >> 1
		if m<<1 != n {
			b++
		}

		n = m
	}
	return b
}
