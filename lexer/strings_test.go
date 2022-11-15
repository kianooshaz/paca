package lexer

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/kianooshaz/paca/tokens"
)

func TestLexString(t *testing.T) {
	testCases := []struct {
		desc   string
		buffer string
		want   tokens.Token
	}{
		{
			desc:   "one",
			buffer: `hello world"`,
			want:   tokens.Token{Type: tokens.STRING, Value: "hello world"},
		},
		{
			desc:   "two",
			buffer: `hello \\ // world"`,
			want:   tokens.Token{Type: tokens.STRING, Value: `"hello \\ // world"`},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			l := &lexer{
				buffer: bufio.NewReader(bytes.NewReader([]byte(tC.buffer))),
				tokens: make([]tokens.Token, 0),
			}

			l.lexString('"')

			if tC.want.Type != l.tokens[0].Type {
				t.Error()
			}

			if tC.want.Value != l.tokens[0].Value {
				t.Error()
			}
		})
	}
}
