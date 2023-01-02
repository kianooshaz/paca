package lexer

import (
	"io"
	"log"

	"github.com/kianooshaz/paca/tokens"
)

func (l *lexer) lexSymbol(r rune) {

	// one line comment with {}
	if r == rune(123) {
		for {
			r, _, err := l.buffer.ReadRune()
			if err == io.EOF {
				log.Fatalln("comment should be end with }")
				return
			}

			if r == rune(125) {
				return
			}
		}
	}

	// one line comment with //
	if r == rune(47) {
		rr, _, _ := l.buffer.ReadRune()
		if rr == rune(47) {
			// comment
			l.buffer.ReadLine()
			return
		}
		l.buffer.UnreadRune()
	}

	symbol, ok := tokens.Symbols[string(r)]
	if !ok {
		log.Fatalf("syntax error: character (%s)\n", string(r))
	}

	value := string(r)

	for ok && symbol.HasNext {
		r, _, _ = l.buffer.ReadRune()

		// multi line comment
		if value == "(" && r == rune(42) {
			for {
				r, _, err := l.buffer.ReadRune()
				if err == io.EOF {
					log.Fatalln("comment should be end with }")
					return
				}

				if r == rune(42) {
					r, _, _ = l.buffer.ReadRune()
					if r == rune(41) {
						return
					}
				}
			}
		}

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
