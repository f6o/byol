package lispy

import (
	"strings"
	"testing"
)

func TestTokenizerIterator(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []Token
		wantErr  bool
	}{
		{
			name:     "Simple Addition",
			input:    "(+ 10 20)",
			expected: []Token{{TOKEN_LPAREN, "("}, {TOKEN_OPERATOR, "+"}, {TOKEN_NUMBER, "10"}, {TOKEN_NUMBER, "20"}, {TOKEN_RPAREN, ")"}, {TOKEN_EOF, ""}},
			wantErr:  false,
		},
		{
			name:     "Invalid Token",
			input:    "(+ 10 $)",
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a tokenizer from the input string
			tokenizer := NewTokenizer(strings.NewReader(tc.input))

			// Collect tokens using the iterator
			var tokens []Token
			var err error

			for {
				var token Token
				token, err = tokenizer.NextToken()

				// If we expect an error, break on the first error
				if err != nil {
					break
				}

				tokens = append(tokens, token)

				// Break at EOF
				if token.Type == TOKEN_EOF {
					break
				}
			}

			// Check error expectations
			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Check token expectations
			if len(tokens) != len(tc.expected) {
				t.Errorf("Expected %d tokens, got %d", len(tc.expected), len(tokens))
				return
			}

			for i, expectedToken := range tc.expected {
				if tokens[i].Type != expectedToken.Type || tokens[i].Literal != expectedToken.Literal {
					t.Errorf("Token %d: expected %v, got %v", i, expectedToken, tokens[i])
				}
			}
		})
	}
}

func TestTokenizeReader(t *testing.T) {
	input := "(+ 10 (* 5 2))"
	expected := []Token{
		{TOKEN_LPAREN, "("},
		{TOKEN_OPERATOR, "+"},
		{TOKEN_NUMBER, "10"},
		{TOKEN_LPAREN, "("},
		{TOKEN_OPERATOR, "*"},
		{TOKEN_NUMBER, "5"},
		{TOKEN_NUMBER, "2"},
		{TOKEN_RPAREN, ")"},
		{TOKEN_RPAREN, ")"},
		{TOKEN_EOF, ""},
	}

	// Test TokenizeReader function
	tokens, err := TokenizeReader(strings.NewReader(input))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	if len(tokens) != len(expected) {
		t.Errorf("Expected %d tokens, got %d", len(expected), len(tokens))
		return
	}

	for i, expectedToken := range expected {
		if tokens[i].Type != expectedToken.Type || tokens[i].Literal != expectedToken.Literal {
			t.Errorf("Token %d: expected %v, got %v", i, expectedToken, tokens[i])
		}
	}
}

// Example of processing a large input in chunks
func TestLargeInputProcessing(t *testing.T) {
	// Create a large input (in a real scenario, this could be a file or network stream)
	var largeInput strings.Builder
	for i := 0; i < 1000; i++ {
		largeInput.WriteString("(+ 1 2) ")
	}

	// Create a tokenizer
	tokenizer := NewTokenizer(strings.NewReader(largeInput.String()))

	// Process tokens in batches
	tokenCount := 0
	batchSize := 10
	batch := make([]Token, 0, batchSize)

	for {
		token, err := tokenizer.NextToken()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			return
		}

		batch = append(batch, token)
		tokenCount++

		// Process batch when it's full or at EOF
		if len(batch) >= batchSize || token.Type == TOKEN_EOF {
			// In a real application, you would process the batch here
			// For this test, we just clear the batch
			batch = batch[:0]

			// Break at EOF
			if token.Type == TOKEN_EOF {
				break
			}
		}
	}

	// We expect 6 tokens per "(+ 1 2)" expression: (, +, 1, 2, ), plus spaces between expressions
	// Plus a final EOF token
	expectedMinTokens := 1000*5 + 1
	if tokenCount < expectedMinTokens {
		t.Errorf("Expected at least %d tokens, got %d", expectedMinTokens, tokenCount)
	}
}
