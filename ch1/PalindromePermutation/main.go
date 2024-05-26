package main

// Palindrome Permutation: Given a string, write a function to check if it is a permutation of
// a palindrome. A palindrome is a word or phrase that is the same forwards and backwards. A
// permutation is a rearrangement of letters. The palindrome does not need to be limited to just
// dictionary words.
// EXAMPLE
// Input: Tact Coa
// Output: True (permutations: "taco cat", "atco cta", etc.)

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcdefghijklmnopqrstuvwxyz zyxwvutsrqponmlkjihgfedcba"
	fmt.Println(isPalindromePermutation(str))
	fmt.Println(isPalindromePermutationByTable(str))
	fmt.Println(isPalindromePermutationByBit(str))
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

func isPalindromePermutationByTable(str string) bool {
	str = strings.ToLower(str)
	str = strings.Replace(str, " ", "", -1)

	table := make([]int, 26)
	oddCount := 0

	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			table[char-'a']++
			if table[char-'a']%2 != 0 {
				oddCount++
			} else {
				oddCount--
			}
		}
	}
	return oddCount <= 1
}

func isPalindromePermutationByBit(str string) bool {
	str = strings.ToLower(str)
	str = strings.Replace(str, " ", "", -1)

	bitVector := 0

	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			bitVector = toggle(bitVector, char-'a')
		}
	}
	//fmt.Printf("Bit vector: %b\n", bitVector)
	//fmt.Printf("Bit vector-1: %b\n", bitVector-1)
	return bitVector == 0 || (bitVector&(bitVector-1)) == 0
}

func toggle(bitVector int, index rune) int {
	if index < 0 {
		return bitVector
	}

	mask := 1 << uint(index)

	if (bitVector & mask) == 0 {
		bitVector |= mask
	} else {
		bitVector &= ^mask
	}
	//fmt.Printf("Mask: %b\n", mask) // Print bitVector as a binary string
	//fmt.Printf("Bit vector: %b\n", bitVector) // Print bitVector as a binary string
	return bitVector
}
