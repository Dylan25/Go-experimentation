package main

import (
	"fmt"
	"os"
	"strconv"
)

// myecho4 prints it's name, the command line arguments, and their indexes in os.Args
func main() {
	var s, sep string
	fmt.Println(os.Args[0], " ")
	for index, argument := range os.Args[1:] {
		s += sep + strconv.Itoa(index) + ": " + argument
		sep = " "
	}
	fmt.Println(s)
}
