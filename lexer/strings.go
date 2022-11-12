package lexer

import (
	"errors"

	"github.com/kianooshaz/paca/tokens"
)

// func (l lexer) LexString1(r rune) {
// 	str, _ := l.buffer.ReadString('"')
// 	str = strings.TrimRight(str, `"`)
// 	fmt.Println(str, len(str))
// }

func (l lexer) LexString(r rune) (tokens.Token, error) {
	str := string(r)
	for {
		prevRune := r
		var err error
		r, _, err = l.buffer.ReadRune()
		if err != nil {
			l.buffer.UnreadRune()
			return tokens.Token{}, errors.New("invalid string")
		}

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
	return tokens.Token{Type: "string", Value: str}, nil
}
