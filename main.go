package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func parseFlags() error {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, `usage: chomp [-h]

Receives input from STDIN
Removes extra whitespace characters from the beginning/end of lines
Removes lines which have just whitespace (no content)`)
		flag.PrintDefaults()
	}
	flag.Parse()
	return nil
}

// wrapper for 'main' code, to return single err to main
func chomp() error {
	if err := parseFlags(); err != nil {
		return err
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		if len(txt) != 0 {
			fmt.Println(txt)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := chomp(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
