package lexer

import (
	"io"

	"github.com/kianooshaz/paca/character"
	"github.com/kianooshaz/paca/tokens"
)

func (l *Lexer) lexString(r rune) {
	var value string
	var err error

	for {
		r, _, err = l.buffer.ReadRune()
		if err == io.EOF {
			panic("String not finished error")
		}

		if r == character.SingleQuotation {
			break
		}

		value += string(r)
	}

	l.emit(tokens.STRING, value)
}
