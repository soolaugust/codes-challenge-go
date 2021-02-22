package main

import "fmt"

// Test case
// [1, 2, 5, 3, 1, 3] => 2
// [3, 3, 3, 5, 4] => 3
// [6, 5, 5, 6, 2, 2] => 4
func main() {
	A := []int{1, 2, 5, 3, 1, 3}
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	// write your code in Go 1.4
	count, left, right := 1, 0, len(A)
	for {
		if left > right {
			break
		}
		mid := (left + right) >> 1
		if check(mid, A) {
			count = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return count
}

func check(val int, a []int) bool {
	count := 0
	for _, v := range a {
		if v >= val {
			count++
			if count == val {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}
