package main

// Example code from an introductory Go talk.
// For full video/slides/blog, please see:
// https://github.com/219-design/golang-giveth

import (
	"fmt"
)

type Classroom struct { // Note: no declaration of implemented base interfaces.
	deskCount int
}

type Office struct {
	deskCount int
}

func (c *Classroom) AddOneDesk() {
	c.deskCount++
}

func (o *Office) AddOneDesk() {
	o.deskCount++
}

// DeskHolder interface is implemented by Classroom and Office.
type DeskHolder interface {
	AddOneDesk()
}

// AddDeskTo accepts any object that fulfills the DeskHolder interface.
func AddDeskTo(holder DeskHolder) {
	holder.AddOneDesk()
}

func main() {
	room := &Classroom{deskCount: 2}
	fmt.Println(room)

	room.AddOneDesk()
	fmt.Println(room)

	office := &Office{deskCount: 0}
	fmt.Println(office)

	office.AddOneDesk()
	fmt.Println(office)

	AddDeskTo(office)
	AddDeskTo(room)
}
