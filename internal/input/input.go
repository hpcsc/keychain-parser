package input

import (
	"bufio"
	"io"
)

func From(r io.Reader) []string {
	var lines []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			lines = append(lines, line)
		}
	}

	return lines
}
