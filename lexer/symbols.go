package lexer

import (
	"github.com/kianooshaz/paca/tokens"
)

func (l lexer) LexSymbol(r rune) (tokens.Token, error) {

	symbol, ok := tokens.Symbols[string(r)]

	if !ok {
		return tokens.Token{}, nil
	}

	totalRunes := string(r)
	for ok && symbol.IsNextRequire {
		r, _, err := l.buffer.ReadRune()
		if err != nil {
			break
		}
		symbol, ok = tokens.Symbols[totalRunes+string(r)]
		if ok {
			totalRunes += string(r)
		}
	}

	return tokens.Token{
		Type:  "symbol", 
		Value: string(symbol.Name),
	}, nil
}
