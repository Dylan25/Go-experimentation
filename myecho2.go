package main

import (
	"fmt"
	"os"
)

// myecho2 echos command line arguments
// range produces a pair, the index, then value at index
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
