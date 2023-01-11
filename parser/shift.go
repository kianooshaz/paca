package parser

func (p *parser) shift(state string) {
	p.stack.Push(state)
}
