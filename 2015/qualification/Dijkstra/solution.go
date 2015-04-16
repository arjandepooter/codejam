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

type quat struct {
	r   rune
	neg bool
}

type Problem struct {
	block  []quat
	repeat int
}

var mapping map[rune]map[rune]quat

func init() {
	mapping = map[rune]map[rune]quat{
		'1': map[rune]quat{
			'1': quat{'1', false},
			'i': quat{'i', false},
			'j': quat{'j', false},
			'k': quat{'k', false},
		},
		'i': map[rune]quat{
			'1': quat{'i', false},
			'i': quat{'1', true},
			'j': quat{'k', true},
			'k': quat{'j', false},
		},
		'j': map[rune]quat{
			'1': quat{'j', false},
			'i': quat{'k', false},
			'j': quat{'1', true},
			'k': quat{'i', true},
		},
		'k': map[rune]quat{
			'1': quat{'k', false},
			'i': quat{'j', true},
			'j': quat{'i', false},
			'k': quat{'1', true},
		},
	}
}

func (problem *Problem) Solve() string {
	cur := quat{'1', false}
	foundI, foundJ := false, false
	for i := 0; i < problem.repeat && i < 8 && !(foundI && foundJ); i++ {
		for _, r := range problem.block {
			cur = r.multiply(cur)
			if !foundI {
				if cur.equals(quat{'i', false}) {
					cur = quat{'1', false}
					foundI = true
				}
			} else if !foundJ {
				if cur.equals(quat{'j', false}) {
					cur = quat{'1', false}
					foundJ = true
					break
				}
			}
		}
	}
	if foundI && foundJ && (quat{'1', true}).equals(total(problem.block, problem.repeat)) {
		return "YES"
	}
	return "NO"
}

func (self quat) equals(other quat) bool {
	return self.r == other.r && self.neg == other.neg
}

func (self quat) multiply(with quat) quat {
	neg := self.neg != with.neg
	result := mapping[self.r][with.r]
	result.neg = result.neg != neg
	return result
}

func (self quat) power(n int) quat {
	if n == 0 {
		return quat{'1', false}
	}
	if n == 1 {
		return self
	}
	if n%2 == 1 {
		return self.power(n - 1).multiply(self)
	}
	return self.multiply(self).power(n / 2)
}

func total(block []quat, repeat int) quat {
	cur := quat{'1', false}
	for _, q := range block {
		cur = q.multiply(cur)
	}

	return cur.power(repeat)
}

func (q quat) String() string {
	if q.neg {
		return fmt.Sprintf("-%s", string(q.r))
	}
	return fmt.Sprintf("%v", string(q.r))
}

func getProblems(input *utils.Input) []*Problem {
	n := input.ReadInt()

	problems := make([]*Problem, n)
	for i := 0; i < n; i++ {
		problem := Problem{}
		l := input.ReadInt()
		problem.block = make([]quat, l)
		problem.repeat = input.ReadInt()
		block := input.ReadString()
		for j, r := range block {
			problem.block[j] = quat{r, false}
		}
		problems[i] = &problem
	}

	return problems
}
