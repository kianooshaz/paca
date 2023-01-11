package lexer

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/kianooshaz/paca/character"
	"github.com/kianooshaz/paca/tokens"
)

type Lexer struct {
	buffer     *bufio.Reader
	identTable *tokens.IdentTable
	tokens     []tokens.Token
	nextToken  int
}

func New(source string, identTable *tokens.IdentTable) *Lexer {
	buffer := bufio.NewReader(strings.NewReader(source))
	l := &Lexer{
		buffer:     buffer,
		identTable: identTable,
		tokens:     make([]tokens.Token, 0),
		nextToken:  0,
	}
	l.lex()
	return l
}

func (l *Lexer) lex() {
	r, _, err := l.buffer.ReadRune()
	if err == io.EOF {
		l.emit(tokens.EOF, "$")
		return
	}

	switch {
	// space or tab or new line
	case r == character.Space || r == character.Tab || r == character.NewLine || r == character.CarriageReturn:
		break
	case r == rune(39):
		// 39 = '
		l.lexString(r)
	case unicode.IsDigit(r):
		l.lexNumber(r)
	case unicode.IsLetter(r):
		l.lexIdent(r)
	default:
		l.lexSymbol(r)
	}

	l.lex()
}

func (l *Lexer) emit(t tokens.Type, v string) {
	l.tokens = append(l.tokens, tokens.Token{Type: t, Value: v})
}

func (l *Lexer) GetToken() (tokens.Token, error) {
	if l.nextToken >= len(l.tokens) {
		return tokens.Token{}, errors.New("no token")
	}
	token := l.tokens[l.nextToken]
	l.nextToken++
	return token, nil
}

func (l *Lexer) PrintTokens() {
	for _, token := range l.tokens {
		fmt.Println(token)
	}
}

func (l *Lexer) PrintIdentTable() {
	l.identTable.Print()
}
