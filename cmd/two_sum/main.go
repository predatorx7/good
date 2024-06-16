package main

import (
	"fmt"
	"good/pkg/two_sum"
)

func main() {
	result := two_sum.TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("result: %v", result)
}
