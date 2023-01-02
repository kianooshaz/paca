package lexer

import (
	"unicode"

	"github.com/kianooshaz/paca/tokens"
)

func (l *lexer) lexNumber(r rune) {
	var token = tokens.INTEGER
	var str string
	for {
		if r == 'E' {
			if token == tokens.SCIENTIFIC {
				l.buffer.UnreadRune()
				break
			}

			r1, _, _ := l.buffer.ReadRune()
			if r1 == '+' || r1 == '-' {
				r2, _, _ := l.buffer.ReadRune()
				if !unicode.IsDigit(r2) {
					l.buffer.UnreadRune()
					l.buffer.UnreadRune()
					break
				}
				str += string(r) + string(r1) + string(r2) // str = str + "E+3"
			} else if unicode.IsDigit(r1) {
				str += string(r) + string(r1) // str = str + "E3"
			} else {
				l.buffer.UnreadRune()
				break
			}

			token = tokens.SCIENTIFIC
		} else if r == '.' {
			r2, _, _ := l.buffer.ReadRune()
			if r2 == '.' {
				l.emit(token, str)
				l.emit(tokens.DOTDOT, "..")
				return
			}
			l.buffer.UnreadRune()

			if token == tokens.FLOAT || token == tokens.SCIENTIFIC {
				l.buffer.UnreadRune()
				break
			}

			token = tokens.FLOAT
			str += string(r)
		} else if unicode.IsDigit(r) {
			str += string(r)
		} else {
			l.buffer.UnreadRune()
			break
		}

		r, _, _ = l.buffer.ReadRune()
	}

	l.emit(token, str)
}
