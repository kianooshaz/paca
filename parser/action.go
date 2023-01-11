package parser

import (
	"errors"

	"github.com/kianooshaz/paca/tokens"
)

func (p *parser) getAction(state string, token tokens.Token) (string, error) {
	action := p.actionTable[state][string(token.Type)]
	if action == "" {
		return action, errors.New("action not exist")
	}
	return action, nil
}