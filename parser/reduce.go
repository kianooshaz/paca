package parser

import (
	"strings"

	"github.com/kianooshaz/paca/grammar"
)

func (p *parser) reduce(rule string) {
	production := grammar.Productions[rule]
	production.Print()
	bodySize := len(strings.Split(production.Body, " "))
	for i := 0; i < bodySize; i++ {
		_, err := p.stack.Pop()
		if err != nil {
			panic(err.Error())
		}
	}
	top, _ := p.stack.Pop()
	state, _ := p.getGoto(top, production.Head)
	p.stack.Push(state)
}
