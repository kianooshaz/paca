package lexer

import (
	"unicode"

	"github.com/kianooshaz/paca/tokens"
	"github.com/kianooshaz/paca/unicodes"
)

func (l *lexer) lexIdent(r rune) {
	var value string
	for {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			value += string(r)
		} else {
			l.buffer.UnreadRune()
			break
		}
		r, _, _ = l.buffer.ReadRune()
	}

	if t, ok := tokens.Keywords[value]; ok {
		if t == tokens.END {
			r, _, _ = l.buffer.ReadRune()
			if r == unicodes.Dot {
				l.emit(tokens.ENDSTOP, "end.")
				return
			}
			l.buffer.UnreadRune()
		}
		l.emit(t, value)
	} else {
		id := l.identTable.GetID(value)
		l.emit(tokens.IDENT, id)
	}
}
