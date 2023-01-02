package lexer

import (
	"log"

	"github.com/kianooshaz/paca/tokens"
)

func (l *lexer) lexSymbol(r rune) {
	symbol, ok := tokens.Symbols[string(r)]
	if !ok {
		log.Fatalf("syntax error: line (%d): character (%s)", l.line, string(r))
	}

	if r == rune(47) {
		rr, _, _ := l.buffer.ReadRune()
		if rr == rune(47) {
			// comment
			l.buffer.ReadLine()
			return
		}
		l.buffer.UnreadRune()
	}

	value := string(r)

	for ok && symbol.HasNext {
		r, _, _ = l.buffer.ReadRune()

		s, ok := tokens.Symbols[value+string(r)]
		if !ok {
			l.buffer.UnreadRune()
			break
		}

		symbol = s
		value += string(r)
	}

	l.emit(symbol.Name, value)
}
