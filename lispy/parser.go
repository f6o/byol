package lispy

import "errors"

type Parser struct {
	src string
	ast *AST
}

func (p *Parser) Parse() error {
	if p.ast != nil {
		return errors.New("already parsed")
	}
	// TODO
	return nil
}
