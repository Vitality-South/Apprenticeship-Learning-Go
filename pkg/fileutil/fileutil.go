package fileutil

import (
	"bufio"
	"os"
)

// Lines returns a slice of strings, each string representing a line in the file
func Lines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
