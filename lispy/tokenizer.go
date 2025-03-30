package lispy

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
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

// Tokenizer is an iterator that generates tokens from an input stream
type Tokenizer struct {
	reader    *bufio.Reader
	current   rune
	peekRune  rune
	hasPeek   bool
	eof       bool
	lastError error
}

// NewTokenizer creates a new tokenizer from an io.Reader
func NewTokenizer(r io.Reader) *Tokenizer {
	t := &Tokenizer{
		reader: bufio.NewReader(r),
	}
	// Read the first character
	t.advance()
	return t
}

// advance reads the next character from the input
func (t *Tokenizer) advance() {
	if t.hasPeek {
		t.current = t.peekRune
		t.hasPeek = false
		return
	}

	r, _, err := t.reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			t.eof = true
		} else {
			t.lastError = err
		}
		t.current = 0 // NUL character
		return
	}
	t.current = r
}

// peek looks ahead at the next character without consuming it
func (t *Tokenizer) peek() rune {
	if t.hasPeek {
		return t.peekRune
	}

	r, _, err := t.reader.ReadRune()
	if err != nil {
		return 0 // Return NUL character on error or EOF
	}

	t.peekRune = r
	t.hasPeek = true
	return r
}

// skipWhitespace skips any whitespace characters
func (t *Tokenizer) skipWhitespace() {
	for !t.eof && unicode.IsSpace(t.current) {
		t.advance()
	}
}

// NextToken returns the next token from the input
func (t *Tokenizer) NextToken() (Token, error) {
	// Return any previous error
	if t.lastError != nil {
		return Token{}, t.lastError
	}

	// Skip whitespace
	t.skipWhitespace()

	// Check for EOF
	if t.eof {
		return Token{TOKEN_EOF, ""}, nil
	}

	// Process tokens
	switch t.current {
	case '(':
		t.advance()
		return Token{TOKEN_LPAREN, "("}, nil
	case ')':
		t.advance()
		return Token{TOKEN_RPAREN, ")"}, nil
	case '+', '-', '*', '/':
		op := string(t.current)
		t.advance()
		return Token{TOKEN_OPERATOR, op}, nil
	}

	// Process numbers
	if unicode.IsDigit(t.current) {
		return t.readNumber()
	}

	// Process symbols
	if isSymbolStart(t.current) {
		return t.readSymbol()
	}

	// Unknown token
	unknown := string(t.current)
	t.advance()
	return Token{}, fmt.Errorf("unknown token: %s", unknown)
}

// readNumber reads a number token
func (t *Tokenizer) readNumber() (Token, error) {
	var digits []rune

	// Collect digits
	for !t.eof && unicode.IsDigit(t.current) {
		digits = append(digits, t.current)
		t.advance()
	}

	// Create number string
	numStr := string(digits)

	// Validate as a number
	_, err := strconv.Atoi(numStr)
	if err != nil {
		return Token{}, fmt.Errorf("invalid number: %s", numStr)
	}

	return Token{TOKEN_NUMBER, numStr}, nil
}

// readSymbol reads a symbol token
func (t *Tokenizer) readSymbol() (Token, error) {
	var chars []rune

	// Collect symbol characters
	for !t.eof && isSymbolChar(t.current) {
		chars = append(chars, t.current)
		t.advance()
	}

	symStr := string(chars)
	return Token{TOKEN_SYMBOL, symStr}, nil
}

// TokenizeReader tokenizes an entire input stream
func TokenizeReader(r io.Reader) ([]Token, error) {
	tokenizer := NewTokenizer(r)
	var tokens []Token

	for {
		token, err := tokenizer.NextToken()
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)

		// Stop at EOF token
		if token.Type == TOKEN_EOF {
			break
		}
	}

	return tokens, nil
}

// Tokenize function (for backward compatibility)
func Tokenize(input string) ([]Token, error) {
	return TokenizeReader(strings.NewReader(input))
}

// Helper functions
func isSymbolStart(r rune) bool {
	return unicode.IsLetter(r)
}

func isSymbolChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

// Legacy helper functions (for backward compatibility)
func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isSymbol(s string) bool {
	if len(s) == 0 {
		return false
	}
	r := []rune(s)[0]
	return unicode.IsLetter(r)
}
