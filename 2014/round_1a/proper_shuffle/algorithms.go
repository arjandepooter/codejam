package main

import (
	"math/rand"
)

func GoodShuffle(n int) []int {
	perm := make([]int, n)

	for i := 0; i < n; i++ {
		perm[i] = i
	}
	for i := 0; i < n; i++ {
		p := rand.Intn(n)
		perm[p], perm[i] = perm[i], perm[p]
	}

	return perm
}

func main() {
	n, l := 1000000, 1000

	results := make(chan []int, n)
	quit := make(chan bool)
	histogram := make([][]int, l)
	for i := 0; i < 1000; i++ {
		histogram[i] = make([]int, l)
	}

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				results <- GoodShuffle(l)
			}
		}
	}()

	for i := 0; i < n; i++ {
		perm := <-results
		for n, x := range perm {
			histogram[n][x]++
		}
	}
	quit <- true
}
