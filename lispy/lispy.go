package lispy

import (
	"fmt"
)

type AST struct {
	Tag      string
	Contents string
	Children []AST
}

func (ast AST) Print(depth int) {
	if ast.Contents != "" {
		fmt.Printf("%*s '%s'\n", depth*2+len(ast.Tag), ast.Tag, ast.Contents)
	} else {
		fmt.Printf("%*s\n", depth*2+len(ast.Tag), ast.Tag)
	}

	if len(ast.Children) > 0 {
		for _, child := range ast.Children {
			child.Print(depth + 1)
		}
	}
}
