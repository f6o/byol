package lispy

type Parser struct {
	src string
	ast *AST
}

func (p *Parser) Parse() error {
	if p.ast != nil {
		return new Error("already parsed")
	}
	// TODO
}
