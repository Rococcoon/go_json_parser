package token

type TokenType string

const (
	Illegal TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Litterals
	String TokenType = "STRING"
	Number TokenType = "NUMBER"

	// Structure
	LeftBrace    TokenType = "LEFTBRACE"    // '{'
	RightBrace   TokenType = "RIGHTBRACE"   // '}'
	LeftBracket  TokenType = "LEFTBRACKET"  // '['
	RightBracket TokenType = "RIGHTBRACKET" // ']'
	Comma        TokenType = "COMMA"
	Colon        TokenType = "COLON"

	//Values
	True  TokenType = "TRUE"
	False TokenType = "FALSE"
	Null  TokenType = "NULL"
)

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}
