// dup3 prints the count of any duplicated line from stdin or any given files followed
// by the text of that line.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n,", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, lcount := range counts {
		if lcount > 1 {
			fmt.Printf("%d\t%s\n", lcount, line)
		}
	}
}
