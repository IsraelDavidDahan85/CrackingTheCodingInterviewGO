package main

// Palindrome Permutation: Given a string, write a function to check if it is a permutation of
// a palindrome. A palindrome is a word or phrase that is the same forwards and backwards. A
// permutation is a rearrangement of letters. The palindrome does not need to be limited to just
// dictionary words.
// EXAMPLE
// Input: Tact Coa
// Output: True (permutations: "taco cat'; "atc o eta·; etc.)

import (
	"fmt"
	"strings"
)

func main() {
	str := "Tact Coa"
	fmt.Println(isPalindromePermutation(str))
}

func isPalindromePermutation(str string) bool {
	str = strings.ToLower(str)
	str = strings.Replace(str, " ", "", -1)

	charCount := make(map[rune]int)

	for _, char := range str {
		charCount[char]++
	}

	oddCount := 0
	for _, count := range charCount {
		if count%2 != 0 {
			oddCount++
		}
	}

	return oddCount <= 1
}

// Output:
// true

// Time complexity: O(n)
// Space complexity: O(n)
