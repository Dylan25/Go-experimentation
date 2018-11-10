package main

import (
	"fmt"
	"os"
)

// my echo prints it's command line args
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
