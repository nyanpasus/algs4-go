package alphabet

import (
	"fmt"
	"strconv"
	"strings"
)

var BINARY, _ = NewByAlpha("01")

var OCTAL, _ = NewByAlpha("01234567")

var DECIMAL, _ = NewByAlpha("0123456789")

var HEXADECIMAL, _ = NewByAlpha("0123456789ABCDEF")

var DNA, _ = NewByAlpha("ACGT")

var LOWERCASE, _ = NewByAlpha("abcdefghijklmnopqrstuvwxyz")

var UPPERCASE, _ = NewByAlpha("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

var PROTEIN, _ = NewByAlpha("ACDEFGHIKLMNPQRSTVWY")

var BASE64, _ = NewByAlpha("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

var ASCII, _ = NewByRadix(128)

var EXTENDED_ASCII, _ = NewByRadix(256)

var UNICODE16, _ = NewByRadix(65536)

type Alphabet struct {
	alphabet []rune
	inverse  []int
	r        int
}

func NewByAlpha(alpha string) (*Alphabet, error) {
	unicode := make([]bool, 65536) // rune more than 65536
	for _, v := range alpha {
		if unicode[v] {
			return nil, fmt.Errorf("Illegal alphabet: repeated character = '" + strconv.QuoteRune(v) + "'")
		}
		unicode[v] = true
	}

	inverse := make([]int, 65536) // rune more than 65536
	for i := 0; i < len(inverse); i++ {
		inverse[i] = -1
	}
	result := Alphabet{
		alphabet: []rune(alpha),
		inverse:  inverse,
		r:        len(alpha),
	}
	for c := 0; c < result.r; c++ {
		result.inverse[result.alphabet[c]] = c
	}

	return &result, nil
}

func New() (*Alphabet, error) {
	return NewByRadix(256)
}

func NewByRadix(radix int) (*Alphabet, error) {
	result := Alphabet{
		alphabet: make([]rune, radix),
		inverse:  make([]int, radix),
		r:        radix,
	}
	for i := 0; i < radix; i++ {
		result.alphabet[i] = rune(i)
		result.inverse[i] = i
	}
	return &result, nil
}

func (a *Alphabet) ToRune(index int) rune {
	if index < 0 || index >= a.r {
		return -1
	}
	return a.alphabet[index]
}

func (a *Alphabet) ToIndex(r rune) int {
	if int(r) >= len(a.inverse) || a.inverse[r] == -1 {
		return -1
	}
	return a.inverse[r]
}

func (a *Alphabet) Contains(r rune) bool {
	return a.inverse[r] != -1
}

func (a *Alphabet) R() int {
	return a.r
}

func (a *Alphabet) LgR() int {
	lgR := 0
	for t := a.r - 1; t >= 1; t /= 2 {
		lgR++
	}
	return lgR
}

func (a *Alphabet) ToIndices(s string) []int {
	source := []rune(s)
	target := make([]int, len(s))
	for i := range source {
		target[i] = a.ToIndex(source[i])
	}
	return target
}

func (a *Alphabet) ToRunes(indices []int) string {
	var s strings.Builder
	s.Grow(len(indices))
	for _, v := range indices {
		s.WriteRune(a.ToRune(v))
	}
	return s.String()
}
