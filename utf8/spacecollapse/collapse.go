//go:build !solution

package spacecollapse

import (
	"strings"
	"unicode"
)

func CollapseSpaces(input string) string {
	var res strings.Builder
	res.Grow(len(input))

	ongoingSpace := false

	for _, char := range input {
		if ongoingSpace {
			if unicode.IsSpace(char) {
				continue
			}

			res.WriteRune(' ')
			ongoingSpace = false
		}

		if unicode.IsSpace(char) {
			ongoingSpace = true
		} else {
			res.WriteRune(char)
		}
	}

	if ongoingSpace {
		res.WriteRune(' ')
	}

	return res.String()
}
