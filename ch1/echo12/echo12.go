// Echo12
// Exercise 1.2: Modify the echo program to print the index
// and the value of each of its arguments, one per line.
// See page 8.

package main

import (
	"fmt"
	"os"
)

// TODO: Is it possible to replace lines 16 and 17 with strings.Join()?
//		 fmt.println(strings.Join(os.Args[1:], " "))
func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, ":", arg)
	}
}
