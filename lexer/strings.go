package lexer

import (
	"github.com/kianooshaz/paca/tokens"
)

func (l *lexer) lexString(r rune) {
	var str string

	for {
		prevRune := r
		r, _, _ = l.buffer.ReadRune()

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
