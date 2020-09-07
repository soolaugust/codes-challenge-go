package main

import (
	"fmt"
	"sort"
)

func rearrangeBarcodes(barcodes []int) []int {
	count := map[int]int{}
	keys := []int{}
	for _, b := range barcodes {
		count[b]++
		if count[b] == 1 {
			keys = append(keys, b)
		}
	}
	sort.Slice(keys, func(i, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})
	start, index := 0, 0
	for index < len(keys) {
		k := keys[index]
		for i := start; i < len(barcodes); i += 2 {
			barcodes[i] = k
			count[k]--
			if count[k] == 0 {
				index++
				if index >= len(keys) {
					break
				}
				k = keys[index]
			}
		}
		start++

	}
	return barcodes
}

func main() {
	intput := []int{1, 1, 1, 2, 2, 2, 3}
	output := rearrangeBarcodes(intput)
	for _, value := range output {
		fmt.Printf("%d ", value)
	}
}
