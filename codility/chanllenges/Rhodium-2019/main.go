package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// Solution define solution for Rhodium 2019
func Solution(T []int) int {
	// write your code in Go 1.4
	tLen := len(T)
	result := 0
	for i, val := range T {
		road := make([]int, tLen)
		roadCount := 1
		if val != i {
			road[val] = 1
		}

		for j := i + 1; j < tLen; j++ {
			if road[j] > 0 {
				roadCount -= road[j]
				road[j] = 0
			}

			if T[j] < i || T[j] > j {
				road[T[j]]++
				roadCount++
			} else if j == T[j] {
				roadCount++
			}

			if roadCount <= 1 {
				result++
			}
		}
	}
	return result + tLen
}

func main() {
	T := []int{2, 0, 2, 2, 1, 0}
	fmt.Println(Solution(T))
}
