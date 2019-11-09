package main

// Example code from an introductory Go talk.
// For full video/slides/blog, please see:
// https://github.com/219-design/golang-giveth

import (
	"fmt"
)

func MakeCounter() func() int {
	counterValue := 0
	return func() int {
		counterValue++
		return counterValue
	}
}

func main() {
	counter := MakeCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
