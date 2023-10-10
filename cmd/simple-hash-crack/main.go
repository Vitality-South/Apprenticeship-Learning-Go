package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/Vitality-South/Apprenticeship-Learning-Go/pkg/fileutil"
	"github.com/square/exit"
)

func hex(b [20]byte) string {
	return fmt.Sprintf("%x", b)
}

func found(hashesToSearch []string, hash [20]byte) bool {
	return slices.Contains(hashesToSearch, strings.ToUpper(hex(hash)))
}

func main() {
	os.Exit(_main())
}

func _main() int {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: %s [input-file-hashes-to-match] [input-file-password-guesses]\n", os.Args[0])

		return exit.UsageError
	}

	hashes, err := fileutil.Lines(os.Args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening hashes to match file: %v\n", err)

		return exit.NotOK
	}

	file, err := os.Open(os.Args[2])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening password guesses file: %v\n", err)

		return exit.NotOK
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		guess := sc.Text()

		hash := sha1.Sum([]byte(guess))

		if found(hashes, hash) {
			fmt.Printf("Found match for hash %s: %s\n", hex(hash), guess)
		}
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading password guess file line by line: %v\n", err)

		return exit.NotOK
	}

	return exit.OK
}
