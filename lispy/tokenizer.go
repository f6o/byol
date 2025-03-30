package lispy

import (
	"fmt"
	"strconv"
	"strings"
)

// Token Types
type TokenType int

const (
	TOKEN_NUMBER TokenType = iota
	TOKEN_OPERATOR
	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_SYMBOL
	TOKEN_EOF
)

// Token Structure
type Token struct {
	Type    TokenType
	Literal string
}

// Tokenize function
func Tokenize(input string) ([]Token, error) {
	var tokens []Token
	input = strings.TrimSpace(input)
	for len(input) > 0 {
		switch {
		case strings.HasPrefix(input, "("):
			tokens = append(tokens, Token{TOKEN_LPAREN, "("})
			input = input[1:]
		case strings.HasPrefix(input, ")"):
			tokens = append(tokens, Token{TOKEN_RPAREN, ")"})
			input = input[1:]
		case strings.ContainsAny(input[:1], "+-*/"):
			tokens = append(tokens, Token{TOKEN_OPERATOR, input[:1]})
			input = input[1:]
		case isDigit(input[:1]):
			var num string
			for len(input) > 0 && isDigit(input[:1]) {
				num += input[:1]
				input = input[1:]
			}
			tokens = append(tokens, Token{TOKEN_NUMBER, num})
		case isSymbol(input[:1]):
			var sym string
			for len(input) > 0 && isSymbol(input[:1]) {
				sym += input[:1]
				input = input[1:]
			}
			tokens = append(tokens, Token{TOKEN_SYMBOL, sym})
		case strings.HasPrefix(input, " "):
			input = input[1:]
		default:
			return nil, fmt.Errorf("unknown token: %s", input[:1])
		}
		input = strings.TrimLeft(input, " ")
	}
	tokens = append(tokens, Token{TOKEN_EOF, ""})
	return tokens, nil
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isSymbol(s string) bool {
	return strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
