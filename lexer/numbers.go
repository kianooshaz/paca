package lexer

import (
	"unicode"

	"github.com/kianooshaz/paca/reader"
)

func LexNumber(r rune) (string, error) {
	var str string
	var hasDot bool
	var hasE bool
	var err error

	for {
		if r == '.' {
			if hasDot || hasE {
				reader.UnreadRune() //reject
				break
			}
			hasDot = true
		}

		if r == 'E' {
			var r1, r2 rune
			r1, err = reader.ReadRune()
			if err != nil {
				reader.UnreadRune()
				break
			}
			if r1 == '+' || r1 == '-' {
				r2, err = reader.ReadRune()
				if err != nil || !unicode.IsDigit(r2) {
					reader.UnreadRune()
					reader.UnreadRune()
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

		r, err = reader.ReadRune()
		if err != nil {
			reader.UnreadRune()
			break
		}
	}
	return str, nil
}
