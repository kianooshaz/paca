package lexer

import (
	"io"

	"github.com/kianooshaz/paca/tokens"
)

func (l *lexer) lexString(r rune) {
	var str string
	var err error

	for {
		prevRune := r
		r, _, err = l.buffer.ReadRune()
		if err == io.EOF {
			panic("String not finished error")
		}

		if prevRune == '\\' {
			if r == '\\' {
				str += string(r)
				r = rune(0)
				continue
			}
		}

		if r == '"' && prevRune != '\\' {
			break
		}

		str += string(r)
	}

	l.emit(tokens.STRING, str)
}
