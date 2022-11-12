package lexer

import (
	"github.com/kianooshaz/paca/tokens"
)

func (l lexer) LexSymbol(r rune) {
	symbol, ok := tokens.Symbols[string(r)]
	if !ok {
		panic("invalid symbol") // TODO refactor message
	}

	totalRunes := string(r)

	for ok && symbol.HasNext {
		r, _, _ := l.buffer.ReadRune()

		s, ok := tokens.Symbols[totalRunes+string(r)]
		if !ok {
			l.buffer.UnreadRune()
			break
		}

		symbol = s
		totalRunes += string(r)
	}

	l.emit(symbol.Name, totalRunes)
}
