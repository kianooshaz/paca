package reader

import "errors"

type Source string

var seek int
var source Source

func New(src Source) {
	source = src
	seek = 0
}

func ReadRune() (rune, error) {
	if seek == len(source) {
		seek++
		return rune(' '), nil // for EOF
	}

	if seek > len(source) {
		return rune(0), errors.New("out of bound")
	}
	currentByte := source[seek]
	seek++
	return rune(currentByte), nil
}

func UnreadRune() {
	if seek <= 0 {
		return
	}
	seek--
}
