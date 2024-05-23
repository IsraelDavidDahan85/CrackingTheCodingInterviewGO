package main

import (
	"fmt"
	"os"
)

// implement an algorithm to determine if a string has all unique character,
// what if you cannot use additional data structures?

func main() {
	str := "abra cadabra hokus pocus"
	isUniqueA := isUnique(str)
	if isUniqueA {
		fmt.Println("check is done")
	}

	s1 := "tes"
	isUniqueB := isUnique(s1)
	if isUniqueB {
		fmt.Println("check is done")
	}

}

func isUnique(s string) bool {
	heap := make(map[string]int)
	isUnique := true
	for i := 0; i < len(s); i++ {
		heap[string(s[i])]++
		if heap[string(s[i])] > 1 {
			isUnique = false
			break
		}
	}
	if isUnique {
		fmt.Fprintln(os.Stdout, "the '", s, "' is unique")
	} else {
		fmt.Fprintln(os.Stdout, "the '", s, "' is not unique")
	}
	return isUnique
}
