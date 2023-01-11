package parser

import (
	"github.com/kianooshaz/paca/grammar"
)

func (p *parser) reduce(rule string) {
	production := grammar.Productions[rule]
	p.Emit(production)
	for i := 0; i < production.BodySize(); i++ {
		_, err := p.stack.Pop()
		if err != nil {
			panic(err.Error())
		}
	}
	top, _ := p.stack.Top()
	state, err := p.getGoto(top, production)
	if err != nil {
		panic(err.Error())
	}
	p.stack.Push(state)
}
