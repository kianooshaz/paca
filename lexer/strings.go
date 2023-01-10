package lexer

import (
	"io"

	"github.com/kianooshaz/paca/tokens"
)

func (l *lexer) lexString(r rune) {
	var value string
	var err error

	for {
		r, _, err = l.buffer.ReadRune()
		if err == io.EOF {
			panic("String not finished error")
		}

		if r == rune(39) {
			break
		}

		value += string(r)
	}

	l.emit(tokens.STRING, value)
}
