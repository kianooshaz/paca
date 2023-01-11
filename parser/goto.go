package parser

import "errors"

func (p *parser) getGoto(state string, productionHead string) (string, error) {
	goTo := p.gotoTable[state][productionHead]
	if goTo == "" {
		return goTo, errors.New("goto not exist")
	}
	return goTo, nil
}