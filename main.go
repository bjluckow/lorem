package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"
)

//go:embed corpus.txt
var corpus string

func main() {
	wordsPerLine := flag.Int("w", 10, "words per line")
	numLines := flag.Int("n", 10, "number of lines (>0)")
	delimiter := flag.String("d", "\n", "line delimiter")
	flag.Parse()

	if *wordsPerLine <= 0 {
		fmt.Fprintln(os.Stderr, "words per line must be > 0")
		os.Exit(1)
	}

	if *numLines <= 0 {
		fmt.Fprintln(os.Stderr, "number of lines must be >0")
		os.Exit(1)
	}

	fields := strings.Fields(corpus)
	if len(fields) == 0 {
		fmt.Fprintln(os.Stderr, "corpus is empty")
		os.Exit(1)
	}

	i := 0
	lineCount := 0

	for {
		if *numLines > 0 && lineCount >= *numLines {
			break
		}

		line := make([]string, *wordsPerLine)

		for j := 0; j < *wordsPerLine; j++ {
			line[j] = fields[i%len(fields)]
			i++
		}

		fmt.Print(strings.Join(line, " "))
		fmt.Print(*delimiter)

		lineCount++
	}
}
