package parser

import (
	"fmt"

	"github.com/kianooshaz/paca/grammar"
	"github.com/kianooshaz/paca/lexer"
	"github.com/kianooshaz/paca/stack"
)

type Table map[string]map[string]string

type parser struct {
	actionTable Table
	gotoTable   Table
	stack       stack.Stack
	productions []grammar.Production
}

func New(actionTable Table, gotoTable Table) *parser {
	return &parser{
		actionTable: actionTable,
		gotoTable:   gotoTable,
		stack:       stack.New(),
		productions: make([]grammar.Production, 0),
	}
}

func (p *parser) Parse(l *lexer.Lexer) {
	token, err := l.GetToken()
	if err != nil {
		message := fmt.Sprintf("parser error, %s", err.Error())
		panic(message)
	}
	p.stack.Push("0") // push state 0 into stack

	for {
		top, err := p.stack.Top()
		if err != nil { // check stack is empty
			panic("parser error: " + err.Error())
		}

		action, err := p.getAction(top, token) // shift | reduce | Error
		if err != nil {                        // check has action
			message := fmt.Sprintf("parser error: no action. state: %s, token: %s", top, token)
			panic(message)
		}

		if action == "r0" {
			break // accept
		}

		actionType := string(action[0])   // `s` for shift 							| `r` for reduce
		actionValue := string(action[1:]) // `state number` for shift 	| `rule number` for reduce

		// check action type is 'r' to reduce
		if actionType == "r" {
			p.reduce(actionValue)
			continue
		}

		// shift
		p.shift(actionValue)
		token, err = l.GetToken()
		if err != nil { // check has token
			message := fmt.Sprintf("parser error, %s", err.Error())
			panic(message)
		}
	}
}

func (p *parser) Emit(production grammar.Production) {
	p.productions = append(p.productions, production)
}

func (p *parser) PrintDerivation() {
	for i := len(p.productions) - 1; i >= 0; i-- {
		fmt.Printf("%d)\t", len(p.productions)-i)
		fmt.Print(p.productions[i].ToString())
		fmt.Print("\n")
	}
}
