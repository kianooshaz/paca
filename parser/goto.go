package parser

import (
	"errors"
	"fmt"

	"github.com/kianooshaz/paca/grammar"
)

func (p *parser) getGoto(state string, production grammar.Production) (string, error) {
	goTo := p.gotoTable[state][production.Head]
	if goTo == "" {
		message := fmt.Sprintf("goto not exist. state: %s, head: %s", state, production.Head)
		message = fmt.Sprintf("%s\n%s", message, production)
		return goTo, errors.New(message)
	}
	return goTo, nil
}
