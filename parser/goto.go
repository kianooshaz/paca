package parser

import (
	"errors"
	"fmt"
)

func (p *parser) getGoto(state string, productionHead string) (string, error) {
	goTo := p.gotoTable[state][productionHead]
	if goTo == "" {
		message := fmt.Sprintf("goto not exist. state: %s, head: %s", state, productionHead)
		return goTo, errors.New(message)
	}
	return goTo, nil
}
