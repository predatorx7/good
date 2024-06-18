package main

import (
	"fmt"
	"good/pkg/leetcode5"
)

func print(input string, expectation string) {
	result := leetcode5.Solve(input)
	fmt.Printf("input: %s, result: %s, expectation: %s\n", input, result, expectation)
}

func main() {
	print("", "")
	print("babad", "bab") // or aba
	print("cbbd", "bb")
	print("bb", "bb")
}
