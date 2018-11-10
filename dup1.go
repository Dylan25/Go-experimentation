// dup1 prints each line that appears more than once
// on stdin preceded by it's count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, lcount := range counts {
		if lcount > 1 {
			fmt.Printf("%d\t%s\n", lcount, line)
		}
	}
}
