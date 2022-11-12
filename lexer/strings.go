package lexer

import (
	"github.com/kianooshaz/paca/tokens"
)

func (l *lexer) lexString(r rune) {
	str := string(r)
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

		str += string(r)

		if r == '"' && prevRune != '\\' {
			break
		}
	}

	l.emit(tokens.STRING, str)
}
