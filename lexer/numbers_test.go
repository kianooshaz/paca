package lexer

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/kianooshaz/paca/tokens"
)

func TestLexNumber(t *testing.T) {
	testCases := []struct {
		desc       string
		firstDigit rune
		buffer     string
		want       tokens.Token
	}{
		{
			desc:       "one",
			firstDigit: '1',
			buffer:     `2345`,
			want:       tokens.Token{Type: tokens.INT, Value: "12345"},
		},
		{
			desc:       "two",
			firstDigit: '1',
			buffer:     `2.456`,
			want:       tokens.Token{Type: tokens.FLOAT, Value: `12.456`},
		},
		{
			desc:       "three",
			firstDigit: '1',
			buffer:     `.234E12`,
			want:       tokens.Token{Type: tokens.SCIENTIFIC, Value: `1.234E12`},
		},
		{
			desc:       "four",
			firstDigit: '1',
			buffer:     `.234E+12`,
			want:       tokens.Token{Type: tokens.SCIENTIFIC, Value: `1.234E+12`},
		},
		{
			desc:       "five",
			firstDigit: '1',
			buffer:     `.234E-12`,
			want:       tokens.Token{Type: tokens.SCIENTIFIC, Value: `1.234E-12`},
		},
		{
			desc:       "six",
			firstDigit: '1',
			buffer:     `2345..466`,
			want:       tokens.Token{Type: tokens.INT, Value: "12345"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			l := &lexer{
				buffer: bufio.NewReader(bytes.NewReader([]byte(tC.buffer))),
				tokens: make([]tokens.Token, 0),
			}

			l.lexNumber(tC.firstDigit)

			if tC.want.Type != l.tokens[0].Type {
				t.Error()
			}

			if tC.want.Value != l.tokens[0].Value {
				t.Error()
			}
		})
	}
}
