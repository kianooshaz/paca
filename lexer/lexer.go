package lexer

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/kianooshaz/paca/tokens"
)

type lexer struct {
	buffer *bufio.Reader
	tokens []tokens.Token
}

func New(source string) *lexer {
	buffer := bufio.NewReader(strings.NewReader(source))
	l := &lexer{
		buffer: buffer,
		tokens: make([]tokens.Token, 0),
	}
	l.lex()
	return l
}

func (l *lexer) lex() {
	r, _, err := l.buffer.ReadRune()
	if err == io.EOF {
		l.emit(tokens.EOF, "")
		return
	}

	switch {
	case r == rune(32):
		l.lex()
	case r == '"':
		l.lexString(r)
	case unicode.IsDigit(r):
		l.lexNumber(r)
	case unicode.IsLetter(r):
		// l.lexIdent(r)
	default:
		l.lexSymbol(r)
	}

	l.lex()
}

func (l *lexer) emit(t tokens.Type, v string) {
	l.tokens = append(l.tokens, tokens.Token{Type: t, Value: v})
}

func (l *lexer) PrintTokens() {
	for _, token := range l.tokens {
		fmt.Println(token)
	}
}
