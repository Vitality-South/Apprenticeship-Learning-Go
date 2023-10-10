// An app to assist in the NIST elliptic curves seeds bounty.
// https://words.filippo.io/dispatches/seeds-bounty/
package main

import (
	"crypto/sha1"
	"fmt"
	"os"
	"slices"
	"strings"
	"unicode"

	"github.com/square/exit"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/Vitality-South/Apprenticeship-Learning-Go/pkg/fileutil"
	"github.com/Vitality-South/Apprenticeship-Learning-Go/pkg/sliceutil"
)

const MULTI_HASH_ITERATIONS = 1000

type Transform func(string) string

func transform(s string, t []Transform) []string {
	slice := make([]string, len(t))

	for i, v := range t {
		slice[i] = v(s)
	}

	return slice
}

func hex(b [20]byte) string {
	return fmt.Sprintf("%x", b)
}

func found(hashesToSearch []string, hash [20]byte) bool {
	return slices.Contains(hashesToSearch, strings.ToUpper(hex(hash)))
}

func original(s string) string {
	return s
}

func noSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func wrapWithDoubleQuotes(s string) string {
	return fmt.Sprintf(`"%s"`, s)
}

func spacesAsDashes(s string) string {
	return strings.ReplaceAll(s, " ", "-")
}

func spacesAsCRLF(s string) string {
	return strings.ReplaceAll(s, " ", "\r\n")
}

func spacesAsNewLines(s string) string {
	return strings.ReplaceAll(s, " ", "\n")
}

func spacesAsTabs(s string) string {
	return strings.ReplaceAll(s, " ", "\t")
}

func tryVaryingPunctuation(s string) string {
	return strings.ReplaceAll(s, ".", "!")
}

func removeTrailingPunctuation(s string) string {
	runes := []rune(s)

	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsPunct(runes[i]) {
			runes = runes[:i]
		} else {
			break
		}
	}

	return string(runes)
}

func titleCase(s string) string {
	return cases.Title(language.English).String(s)
}

func lineNotEmpty(s string) bool {
	return len(s) > 0
}

func main() {
	os.Exit(_main())
}

func _main() int {
	hashesToCrack, err := fileutil.Lines("hashes.txt")

	if err != nil {
		return exit.NotOK
	}

	hashesToCrack = sliceutil.Filter(hashesToCrack, lineNotEmpty)

	guesses, err := fileutil.Lines("guesses.txt")

	if err != nil {
		return exit.NotOK
	}

	guesses = sliceutil.Filter(guesses, lineNotEmpty)

	variations := []Transform{
		original,
		removeTrailingPunctuation,
		tryVaryingPunctuation,
	}

	versions := []Transform{
		original,
		noSpaces,
		spacesAsDashes,
		spacesAsCRLF,
		spacesAsNewLines,
		spacesAsTabs,
	}

	iterations := []Transform{
		original,
		strings.ToLower,
		strings.ToUpper,
		titleCase,
		wrapWithDoubleQuotes,
	}

	multiHashText := []Transform{
		strings.ToLower,
		strings.ToUpper,
	}

	var guessCount uint64

	for _, g := range guesses {
		varies := transform(g, variations)

		for _, v := range varies {
			vers := transform(v, versions)

			for _, vv := range vers {
				iters := transform(vv, iterations)

				for _, i := range iters {
					guessCount++

					// hash our guess
					hashBytes := sha1.Sum([]byte(i))

					if found(hashesToCrack, hashBytes) {
						fmt.Println("Found match", hex(hashBytes), i)
					}

					// try multi-hashing the bytes representation of our guess
					multi := hashBytes
					for j := 0; j < MULTI_HASH_ITERATIONS; j++ {
						guessCount++

						multi = sha1.Sum(multi[:])

						if found(hashesToCrack, multi) {
							fmt.Println("Found multi-hashing bytes match", hex(multi), j, i)
						}
					}

					// try multi-hashing the text/hex representation of our guess
					for _, f := range multiHashText {
						textHash := f(hex(hashBytes))
						for j := 0; j < MULTI_HASH_ITERATIONS; j++ {
							guessCount++

							b := sha1.Sum([]byte(textHash))
							textHash = f(hex(b))

							if found(hashesToCrack, b) {
								fmt.Println("Found multi-hashing text match", textHash, j, i)
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("Guesses:", guessCount)

	return exit.OK
}
