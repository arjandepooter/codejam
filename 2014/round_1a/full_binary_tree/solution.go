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
		fmt.Printf("Case #%d: %d\n", i+1, problem.Solve())
	}
}

type Node struct {
	nodes []*Node
}

func (node *Node) getTreeSize(from *Node) int {
	numberOfChildren := len(node.nodes)
	if from != nil {
		numberOfChildren--
	}
	if numberOfChildren <= 1 {
		return 1
	}

	childSizes := make([]int, numberOfChildren)
	i := 0
	for _, neighbour := range node.nodes {
		if neighbour == from {
			continue
		}
		childSizes[i] = neighbour.getTreeSize(node)
		i++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(childSizes)))

	return 1 + childSizes[0] + childSizes[1]
}

type Problem struct {
	nodes map[int]*Node
}

func (problem *Problem) Solve() int {
	maxN := 0
	for _, root := range problem.nodes {
		n := root.getTreeSize(nil)
		if n > maxN {
			maxN = n
		}
	}

	return len(problem.nodes) - maxN
}

func getProblems(input *utils.Input) []*Problem {
	n := input.ReadInt()

	problems := make([]*Problem, n)
	for i := 0; i < n; i++ {
		problem := Problem{}
		problem.nodes = make(map[int]*Node)

		N := input.ReadInt()
		for j := 1; j < N; j++ {
			id1, id2 := input.ReadInt(), input.ReadInt()
			node1, exists := problem.nodes[id1]
			if !exists {
				node1 = &Node{[]*Node{}}
			}
			node2, exists := problem.nodes[id2]
			if !exists {
				node2 = &Node{[]*Node{}}
			}
			node1.nodes = append(node1.nodes, node2)
			node2.nodes = append(node2.nodes, node1)
			problem.nodes[id1] = node1
			problem.nodes[id2] = node2
		}

		problems[i] = &problem
	}

	return problems
}
