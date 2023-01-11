package parser

import "errors"

func (p *parser) getAction(state string, tokenValue string) (string, error) {
	action := p.actionTable[state][tokenValue]
	if action == "" {
		return action, errors.New("action not exist")
	}
	return action, nil
}