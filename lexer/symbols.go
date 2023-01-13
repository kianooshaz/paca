package lexer

import (
	"io"
	"log"

	"github.com/kianooshaz/paca/character"
	"github.com/kianooshaz/paca/tokens"
)

func (l *Lexer) lexSymbol(r rune) {

	// one line comment with {}
	if r == character.LCurlyBrace {
		for {
			r, _, err := l.buffer.ReadRune()
			if err == io.EOF {
				log.Fatalln("comment should be end with }")
				return
			}

			if r == character.RCurlyBrace {
				return
			}
		}
	}

	// one line comment with //
	if r == character.Slash {
		rr, _, _ := l.buffer.ReadRune()
		if rr == character.Slash {
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
		if value == "(" && r == character.Star {
			for {
				r, _, err := l.buffer.ReadRune()
				if err == io.EOF {
					log.Fatalln("comment should be end with }")
					return
				}

				if r == character.Star {
					r, _, _ = l.buffer.ReadRune()
					if r == character.RParent {
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
