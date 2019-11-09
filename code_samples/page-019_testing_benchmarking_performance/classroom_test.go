package main

// Example code from an introductory Go talk.
// For full video/slides/blog, please see:
// https://github.com/219-design/golang-giveth

import (
	"fmt"
	"testing"
)

func TestClassroom(t *testing.T) {
	const startVal = 2
	room := &Classroom{deskCount: startVal}
	room.AddOneDesk()
	if startVal == room.deskCount {
		t.Error("AddOneDesk did not change desk count")
	}
}
func ExampleAddOneDesk() {
	room := &Classroom{deskCount: 20}
	room.AddOneDesk()
	fmt.Println(room.deskCount)
	// Output:
	// 21
}
func BenchmarkAddOneDesk(t *testing.B) {
	const startVal = 2
	room := &Classroom{deskCount: startVal}
	room.AddOneDesk()
}
