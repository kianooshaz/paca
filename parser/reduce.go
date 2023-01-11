package parser

import (
	"github.com/kianooshaz/paca/grammar"
)

func (p *parser) reduce(ruleNumber string) {
	production := grammar.Productions[ruleNumber] // get grammar rule by rule-number
	p.Emit(production)                            // add to derivation list

	// pop states as the body size of production
	for i := 0; i < production.BodySize(); i++ {
		_, err := p.stack.Pop()
		if err != nil {
			panic(err.Error())
		}
	}

	// getting top of the stack for goto-table
	top, _ := p.stack.Top()
	state, err := p.getGoto(top, production)

	if err != nil { // checking has any state for goto
		panic(err.Error())
	}

	p.stack.Push(state)
}
