package lexer

import (
	"unicode"
)

func (l lexer) LexNumber(r rune) (string, error) {
	var str string
	var hasDot bool
	var hasE bool
	var err error

	for {
		if r == '.' {
			if hasDot || hasE {
				l.buffer.UnreadRune()
				break
			}
			hasDot = true
		}

		if r == 'E' {
			var r1, r2 rune
			r1, _, err = l.buffer.ReadRune()
			if err != nil {
				l.buffer.UnreadRune()
				break
			}
			if r1 == '+' || r1 == '-' {
				r2, _, err = l.buffer.ReadRune()
				if err != nil || !unicode.IsDigit(r2) {
					l.buffer.UnreadRune()
					l.buffer.UnreadRune()
					break
				}
				str += string(r) + string(r1) + string(r2)
			}
			if unicode.IsDigit(r1) {
				str += string(r) + string(r1)
			}
			hasE = true
		}

		if unicode.IsDigit(r) || r == '.' {
			str += string(r)
		}

		r, _, err = l.buffer.ReadRune()
		if err != nil {
			l.buffer.UnreadRune()
			break
		}
	}
	return str, nil
}
