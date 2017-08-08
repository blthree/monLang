package lexer

import (
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType	token.TokenType
		expectedLiteral	string
	}{
		// continue here on page 18
	}
}
