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
	strings []string
}

type block struct {
	char  rune
	count int
}

type blocks []*block

func (b *blocks) String() string {
	output := ""
	for _, blk := range *b {
		output += string(blk.char)
	}
	return output
}

func (problem *Problem) Solve() string {
	numberOfStrings := len(problem.strings)
	allBlocks := make([]blocks, numberOfStrings)

	for i, s := range problem.strings {
		blocks := blocks{}
		currentBlock := &block{' ', 0}
		for _, char := range s {
			if char != currentBlock.char {
				currentBlock = &block{char, 0}
				blocks = append(blocks, currentBlock)
			}
			currentBlock.count++
		}
		allBlocks[i] = blocks
	}

	s := allBlocks[0].String()
	for _, blocks := range allBlocks[1:] {
		if blocks.String() != s {
			return "Fegla Won"
		}
	}

	totalOperations := 0
	for i := 0; i < len(allBlocks[0]); i++ {
		counts := make([]int, numberOfStrings)
		for j, blocks := range allBlocks {
			counts[j] = blocks[i].count
		}

		sort.Ints(counts)
		median := counts[numberOfStrings/2]

		operations := 0
		for _, count := range counts {
			if count < median {
				operations += median - count
			} else {
				operations += count - median
			}
		}
		totalOperations += operations
	}

	return fmt.Sprint(totalOperations)
}

func getProblems(input *utils.Input) []*Problem {
	n := input.ReadInt()

	problems := make([]*Problem, n)
	for i := 0; i < n; i++ {
		problem := Problem{}

		numberOfStrings := input.ReadInt()
		problem.strings = input.ReadStrings(numberOfStrings)

		problems[i] = &problem
	}

	return problems
}
