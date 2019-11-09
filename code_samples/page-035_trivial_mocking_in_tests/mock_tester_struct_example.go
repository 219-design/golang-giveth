package main

// Example code from an introductory Go talk.
// For full video/slides/blog, please see:
// https://github.com/219-design/golang-giveth

import (
	"errors"
	"fmt"
	"os"
)

type FakeTesterFileInfo struct {
	os.FileInfo
	forceFail bool
}

func (f FakeTesterFileInfo) Size() int64 {
	if f.forceFail {
		return 70000
	} else {
		return 1024
	}
}

func ProcessFile(fileInfo os.FileInfo) error {
	if fileInfo.Size() > 65535 {
		return errors.New("module_name: file too big")
	}

	// do some kind of processing here

	// calling this would crash:
	// fileInfo.IsDir()

	return nil
}

func main() {
	err := ProcessFile(FakeTesterFileInfo{forceFail: true})
	fmt.Println("When forceFail is true: ", err)
	err = ProcessFile(FakeTesterFileInfo{forceFail: false})
	fmt.Println("When forceFail is false: ", err)
	fmt.Println("done")
}

/*
https://golang.org/pkg/os/#FileInfo

type FileInfo interface {
        Name() string       // base name of the file
        Size() int64        // length in bytes for regular files; system-dependent for others
        Mode() FileMode     // file mode bits
        ModTime() time.Time // modification time
        IsDir() bool        // abbreviation for Mode().IsDir()
        Sys() interface{}   // underlying data source (can return nil)
}
*/
