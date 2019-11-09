// Other:
//    - CR/LF becomes LF.
//    - Go prefers tabs, not spaces.
//
// Capitalization convention is also enforced, due to its meaning to the compiler/linker.
package main

// Example code from an introductory Go talk.
// For full video/slides/blog, please see:
// https://github.com/219-design/golang-giveth

// try out of order: fmt math log errors io
import (
	"fmt"
	"math"
	"log"
	"errors" //     the imported package names will be sorted alphabetically by go fmt
	"io"
)

type Address struct {
	heading string
	street string //    members of the struct will be column-aligned by go fmt
	apt string
	code int
	isUSA bool
}

func main() {
	flag := true
	if(flag){ // parentheses will be removed and spacing repaired by go fmt
		fmt.Println("true"); // semi-colon will be removed by go fmt
	    } // indentation will be repaired

fmt.Println("another thing") //   indentation will be repaired by go fmt

	
	x := []int{1, 2, 3} //    the 2 blank lines above here will be reduced to 1 by go fmt

	fmt.Println(x)
}

// to suppress errors about unused imports
var _ = errors.New
var _ = io.EOF
var _ = log.Fatal
var _ = math.Floor
