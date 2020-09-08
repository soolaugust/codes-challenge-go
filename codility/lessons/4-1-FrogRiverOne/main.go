package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func solution(X int, A []int) int {
	// write your code in Go 1.4
	set := make(map[int]bool, X)
	correctSum := (1 + X) * X / 2
	currentSum := 0
	for index, val := range A {
		if !set[val] {
			set[val] = true
			currentSum += val
		}
		if currentSum == correctSum {
			return index
		}
	}
	return -1
}
