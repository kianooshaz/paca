package lexer

import (
	"strings"
	"unicode"

	"github.com/kianooshaz/paca/character"
	"github.com/kianooshaz/paca/tokens"
)

func (l *Lexer) lexIdent(r rune) {
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

	if t, ok := tokens.Keywords[strings.ToLower(value)]; ok {
		if t == tokens.END {
			r, _, _ = l.buffer.ReadRune()
			if r == character.Dot {
				l.emit(tokens.ENDSTOP, string(tokens.ENDSTOP))
				return
			}
			l.buffer.UnreadRune()
		}
		l.emit(t, value)
		return
	}

	id := l.identTable.GetID(value)
	l.emit(tokens.IDENT, id)
}
