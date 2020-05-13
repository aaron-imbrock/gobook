// Echo12
// Exercise 1.2: Modify the echo program to print the index
// and the value of each of its arguments, one per line.
// See page 8.

package main

import (
	"fmt"
	"os"
)

//!+
func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, ":", arg)
	}
}

//!-
