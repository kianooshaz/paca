package parser

import (
	"fmt"

	"github.com/kianooshaz/paca/lexer"
)

type Table map[string]map[string]string

type parser struct {
	actionTable Table
	gotoTable   Table
	stack       Stack
}

func New(actionTable Table, gotoTable Table) *parser {
	return &parser{
		actionTable: actionTable,
		stack:       NewStack(),
	}
}

func (p *parser) Parse(l *lexer.Lexer) {
	token, err := l.GetToken()
	if err != nil {
		panic("parser error")
	}
	p.stack.Push("0")

	for {
		top, err := p.stack.Top()
		if err != nil { // check stack is not empty
			panic("parser error: " + err.Error())
		}

		action, err := p.getAction(top, string(token.Type))
		if err != nil { // check has action
			message := fmt.Sprintf("parser error: no action. state: %s, token: %s", top, token)
			panic(message)
		}

		if action == "r0" {
			break // accept
		}

		actionType := string(action[0])
		actionValue := string(action[1:])

		// check action type is 'r' to reduce
		if actionType == "r" {
			p.reduce(actionValue)
		}

		// shift
		p.shift(actionValue)
		token, err = l.GetToken()
		if err != nil { // check has token
			panic("parser error: no token")
		}
	}

}
