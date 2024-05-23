package main

import (
	"fmt"
	"sort"

	"israeldahan.co.il/CrackingTheCodingInterviewGO/utils"
)

// given tow strings, write methods to decide if one is a permutation of the other?

func main() {
	strA := "asklflfkgnlakfgnladfkn"
	strB := "asklflfkgnlakfgnladfkn"
	fmt.Println("Check Permutation:", strA, strB)
	IsPermutation := CheckPermutation(strA, strB)
	fmt.Println(IsPermutation)
}

func CheckPermutation(str1 string, str2 string) bool {
	w1 := SortString(str1)
	w2 := SortString(str2)

	if w1 != w2 {
		return false
	}
	return true
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(utils.SortRunes(r))
	return string(r)
}
