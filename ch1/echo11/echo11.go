// Echo11 
// Exercise 1.1: Modify the echo program to also print os.Args[0],
// the name of the command that invoked it.
// See page 8.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
