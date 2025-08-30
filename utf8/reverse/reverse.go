//go:build !solution

package reverse

import (
	"strings"
	"unicode/utf8"
)

func Reverse(input string) string {
	builder := strings.Builder{}
	builder.Grow(len(input))

	for len(input) > 0 {
		r, size := utf8.DecodeLastRuneInString(input)
		if r == utf8.RuneError && size == 1 {
			// невалидный байт → заменяем
			r = '\uFFFD'
		}
		builder.WriteRune(r)
		input = input[:len(input)-size]
	}

	return builder.String()
}
