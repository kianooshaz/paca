package lexer

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/kianooshaz/paca/tokens"
)

type lexer struct {
	buffer *bufio.Reader
	tokens []tokens.Token
}

func New(source string) *lexer {
	buffer := bufio.NewReader(strings.NewReader(source))
	return &lexer{
		buffer: buffer,
		tokens: make([]tokens.Token, 0),
	}
}

func (l lexer) Lex() {
	r, _, err := l.buffer.ReadRune()
	if err != nil {
		log.Fatal("lex error")
	}

	fmt.Println(l.LexString(r))
}

func (l *lexer) emit(t tokens.Type, v string) {
	l.tokens = append(l.tokens, tokens.Token{t, v})
}

func (l *lexer) PrintTokens() {
	for _, token := range l.tokens {
		fmt.Println(token)
	}
}
