package main

// Example code from an introductory Go talk.
// For full video/slides/blog, please see:
// https://github.com/219-design/golang-giveth

import (
	"fmt"
	"net/mail"
	"os/user"
	"time"
)

func ConvertToStringWeWant(group *user.Group) string {
	return "TODO"
}

func GetGroupInformation(groupName string) (string, error) {
	var grp *user.Group
	var err error
	if grp, err = user.LookupGroup("PREFIX" + groupName); err != nil {
		// os/user.LookupGroup knows its input, but it doesn't know the
		// original value of the 'groupName' parameter passed to GetGroupInformation,
		// so we add that information in our error message:
		return "", fmt.Errorf("my module: lookup for %s: %v", groupName, err)
	}

	s := ConvertToStringWeWant(grp)
	// Do other arbitrary logic here...
	return s, nil
}

func main() {
	var loc *time.Location
	var addr *mail.Address

	var err error

	loc, err = time.LoadLocation("America/Nova_Yorkia")
	fmt.Println(err)

	addr, err = mail.ParseAddress("xyz@lmn@ijk@qrs")
	fmt.Println(err)

	s, err := GetGroupInformation("abcdefghijklmnop")
	fmt.Println(err)

	if err == nil {
		fmt.Println(loc, addr, s)
	}
}
