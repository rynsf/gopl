// Dup2 prints the count, filename and text of lines that appear
// more than once in the input. It reads from stdin or from a
// list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	inFile := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, inFile, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, inFile, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t", n)
			for _, file := range inFile[line] {
				fmt.Printf("%s\t", file)
			}
			fmt.Printf("%s\n", line)
		}
	}
}
func appendUnique(inFile map[string][]string, line string, fileName string) {
	for _, s := range inFile[line] {
		if s == fileName {
			return
		}
	}
	inFile[line] = append(inFile[line], fileName)
}

func countLines(f *os.File, counts map[string]int, inFile map[string][]string, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		appendUnique(inFile, line, fileName)
	}
}
