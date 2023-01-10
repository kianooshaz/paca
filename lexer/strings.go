package lexer

import (
	"io"

	"github.com/kianooshaz/paca/tokens"
	"github.com/kianooshaz/paca/unicodes"
)

func (l *lexer) lexString(r rune) {
	var value string
	var err error

	for {
		r, _, err = l.buffer.ReadRune()
		if err == io.EOF {
			panic("String not finished error")
		}

		if r == unicodes.Backslash {
			rr, _, _ := l.buffer.ReadRune()
			if rr == unicodes.SingleQuotation {
				value += string(r) + string(rr)
				continue
			}
			l.buffer.UnreadRune()
		}

		if r == unicodes.SingleQuotation {
			break
		}

		value += string(r)
	}

	l.emit(tokens.STRING, value)
}
