package utils

import (
	"fmt"
	"testing"
)

func TestGcd(t *testing.T) {
	gcd := Gcd(500, 8)
	fmt.Println(gcd)

}
