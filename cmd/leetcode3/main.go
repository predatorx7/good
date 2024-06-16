package main

import (
	"fmt"
	"good/pkg/leetcode3"
)

func check(input string, expected int) {
	result := leetcode3.Solve(input)
	fmt.Printf("input = %s, expected = %d, actual = %d\n", input, expected, result)
}

func main() {
	check("abcabcbb", 3)
	check("dvdf", 3)
}
