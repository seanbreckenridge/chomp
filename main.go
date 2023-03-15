package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Flags struct {
	maxCapacity int
}

func parseFlags() *Flags {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, `usage: chomp [-h] [FLAGS]

Receives input from STDIN
Removes extra whitespace characters from the beginning/end of lines
Removes lines which have just whitespace (no content)`)
		fmt.Fprintln(os.Stderr, "\nFlags:")
		flag.PrintDefaults()
	}
	maxCapacity := flag.Int("max-capacity", 64, "Maximum capacity for each line in kilobytes (default 64K)")
	flag.Parse()
	return &Flags{
		maxCapacity: *maxCapacity,
	}
}

// wrapper for 'main' code, to return single err to main
func chomp() error {
	flags := parseFlags()

	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, flags.maxCapacity*1024)
	scanner.Buffer(buffer, int(flags.maxCapacity*1024))
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		if len(txt) != 0 {
			fmt.Println(txt)
		}
	}
	if err := scanner.Err(); err != nil {
		if err == bufio.ErrTooLong {
			return fmt.Errorf("line too long (max %dK). You can increase this limit with the -max-capacity flag", flags.maxCapacity)
		}
		return err
	}
	return nil
}

func main() {
	if err := chomp(); err != nil {
		fmt.Fprintf(os.Stderr, "chomp: %s\n", err.Error())
		os.Exit(1)
	}
}
