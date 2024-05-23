package main

// URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string
// has sufficient space at the end to hold the additional characters, and that you are given the "true"
// length of the string. (Note: If implementing in Java, please use a character array so that you can
// perform this operation in place.)

import (
	"fmt"
	"strings"
)

func main() {
	str := "Mr John Smith    "

	length := 13
	fmt.Println(URLify(str, length))
	fmt.Println(URLifyInPlace(str, length))
}

func URLify(str string, length int) string {
	s := str[:length]
	s = strings.Replace(str, " ", "%20", -1)
	return s
}

func URLifyInPlace(str string, length int) string {
	spaceCount := 0
	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			spaceCount++
		}
	}
	newIndex := length + spaceCount*2
	newStr := make([]string, newIndex)

	if length < len(str) {
		str = str[:length]
	}
	for i := length - 1; i >= 0; i-- {
		if str[i] == ' ' {
			newStr[newIndex-1] = "0"
			newStr[newIndex-2] = "2"
			newStr[newIndex-3] = "%"
			newIndex -= 3
		} else {
			newStr[newIndex-1] = string(str[i])
			newIndex--
		}
	}
	return strings.Join(newStr, "")
}
