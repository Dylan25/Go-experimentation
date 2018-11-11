// dup2 prints the count of any duplicated line from stdin or any given files followed
// by the text of that line.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dp2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, filemap := range counts {
		printNumberOfOccurences(filemap, line)
	}
}

func countLines(f *os.File, counts map[string]map[string]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][filename]++
	}
}

func printNumberOfOccurences(filemap map[string]int, line string) {
	var total int
	for _, occurences := range filemap {
		total += occurences
	}
	if total > 1 {
		fmt.Printf("%d total occurences of '%s'***************\n", total, line)
		for filename, occurences := range filemap {
			fmt.Printf("   %d occurences of '%s' in file: '%s'\n", occurences, line, filename)
		}
		fmt.Println("")
	}
}
