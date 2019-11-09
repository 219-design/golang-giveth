package main

// Example code from an introductory Go talk.
// For full video/slides/blog, please see:
// https://github.com/219-design/golang-giveth

import (
	"fmt"
	"math" //                      error: imported and not used: "math"
)

const b byte = 256 //              error: constant 256 overflows byte

func decider(i int, j int) bool {
	if i < j {
	}
} //                               error: missing return at end of function

func main() {

	numbers := []int{1, 2, 3}

	var idx bool = true
	x := numbers[idx] //           error: non-integer slice index idx

	fmt.Println("Hello")
}
