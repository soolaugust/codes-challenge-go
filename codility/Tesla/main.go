package main

// you can also use imports, for example:
// import "fmt"
// import "os"
import (
	"fmt"
	"strings"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// Solution for Tesla OA
// odd - 1
// even divide by 2
// find how many steps from given binary number to 0
// e.g. 111 ->  7
// step 1: 7 - 1 = 6
// step 2: 6 / 2 = 3
// step 3: 3 - 1 = 2
// step 4: 2 / 2 = 1
// step 5: 1 - 1 = 0
// so output should be 5
func Solution(S string) int {
	// write your code in Go 1.4
	sLen := len(S)
	firstOneIndex := strings.Index(S, "1")
	if sLen == 0 && firstOneIndex == -1 {
		return 0
	}
	oneCounts := 0
	for _, val := range S {
		if val == '1' {
			oneCounts++
		}
	}
	firstOneDistance := sLen - 1 - firstOneIndex
	return firstOneDistance + oneCounts
}

func main() {
	fmt.Println(Solution("111"))
}
