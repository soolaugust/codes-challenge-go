package main

import "fmt"

func getPermutation(n int, k int) string {
	var outputValue string
	nArray := toArray(n)
	k=k-1
	for i := n - 1; i >= 0; i-- {
		index := k/getPermutationCount(i)
		outputValue = fmt.Sprintf("%s%d", outputValue, nArray[index])
		k = k%getPermutationCount(i)
		nArray = append(nArray[0:index], nArray[index+1:]...)
	}
	return outputValue
}

func getPermutationCount(n int) int {
	if n == 1 || n == 0{
		return 1
	}
	return n * getPermutationCount(n-1)
}

func toArray(number int) []int {
	var array []int
	for i:=1;i<=number;i++ {
		array = append(array,i)
	}
	return array
}

func main() {
	fmt.Println(getPermutation(3, 6))
}
