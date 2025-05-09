package lexer

import (
	"rococcoon/go_json_test/token"
	"unicode"
)

type Lexer struct {
	Input    []rune
	Position int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		Input:    []rune(input),
		Position: 0,
	}
}

func (l *Lexer) TokenizeInput() []token.Token {
	var tokenSlice []token.Token

	for l.Position < len(l.Input) {
		char := l.Input[l.Position]

		// check if end of file
		if l.Position >= len(l.Input)-1 {
			tokenSlice = append(tokenSlice, token.NewToken(token.EOF, "EOF"))
			return tokenSlice
		}

		switch char {
		// handle structures
		case '{':
			tokenSlice = append(tokenSlice, token.NewToken(token.LeftBrace, "{"))
		case '}':
			tokenSlice = append(tokenSlice, token.NewToken(token.RightBrace, "}"))
		case '[':
			tokenSlice = append(tokenSlice, token.NewToken(token.LeftBracket, "["))
		case ']':
			tokenSlice = append(tokenSlice, token.NewToken(token.RightBracket, "]"))
		case ':':
			tokenSlice = append(tokenSlice, token.NewToken(token.Colon, ":"))
		case ',':
			tokenSlice = append(tokenSlice, token.NewToken(token.Comma, ","))
		case ' ', '\n', '\t':
			l.Position++
			continue
			// handle literals
		case 't':
			tokenSlice = append(tokenSlice, l.handleTrue())
		case 'f':
			tokenSlice = append(tokenSlice, l.handleFalse())
		case 'n':
			tokenSlice = append(tokenSlice, l.handleNull())
			// handle values
		case '"':
			tokenSlice = append(tokenSlice, l.handleString())
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
			tokenSlice = append(tokenSlice, l.handleNumber())
		default:
			tokenSlice = append(tokenSlice, l.handleDefault(char))
		}

		l.Position++
	}
	return tokenSlice
}

// Handle Literals

func (l *Lexer) handleTrue() token.Token {
	// Make sure there are enough characters left
	if l.Position+3 >= len(l.Input) {
		return token.NewToken(token.Illegal, "ILLEGAL")
	}

	if l.Input[l.Position+1] == 'r' &&
		l.Input[l.Position+2] == 'u' &&
		l.Input[l.Position+3] == 'e' {
		// Advance position past "true"
		l.Position += 3
		return token.NewToken(token.True, "true")
	}

	// If it doesn't match, return ILLEGAL and optionally adjust position
	return token.NewToken(token.Illegal, "ILLEGAL")
}

func (l *Lexer) handleFalse() token.Token {
	// Make sure there are enough characters left
	if l.Position+4 >= len(l.Input) {
		return token.NewToken(token.Illegal, "ILLEGAL")
	}

	if l.Input[l.Position+1] == 'a' &&
		l.Input[l.Position+2] == 'l' &&
		l.Input[l.Position+3] == 's' &&
		l.Input[l.Position+4] == 'e' {
		// Advance position past "false"
		l.Position += 4
		return token.NewToken(token.False, "false")
	}

	// If it doesn't match, return ILLEGAL and optionally adjust position
	return token.NewToken(token.Illegal, "ILLEGAL")
}

func (l *Lexer) handleNull() token.Token {
	// Make sure there are enough characters left
	if l.Position+3 >= len(l.Input) {
		return token.NewToken(token.Illegal, "ILLEGAL")
	}

	if l.Input[l.Position+1] == 'u' &&
		l.Input[l.Position+2] == 'l' &&
		l.Input[l.Position+3] == 'l' {
		// Advance position past "null"
		l.Position += 3
		return token.NewToken(token.Null, "null")
	}

	// If it doesn't match, return ILLEGAL and optionally adjust position
	return token.NewToken(token.Illegal, "ILLEGAL")
}

// Handle Values

func (l *Lexer) handleString() token.Token {
	var str string

	// increment past initial quotation mark
	l.Position++

	for l.Position < len(l.Input) && l.Input[l.Position] != '"' {
		str += string(l.Input[l.Position])
		l.Position++
	}

	return token.NewToken(token.String, str)
}

func (l *Lexer) handleNumber() token.Token {
	var num string

	// handle minus character
	if l.Input[l.Position] == '-' {
		num += string(l.Input[l.Position])
		l.Position++
	}

	// append numbers before a decimal
	for l.Position < len(l.Input) && unicode.IsDigit(l.Input[l.Position]) {
		num += string(l.Input[l.Position])
		l.Position++
	}

	// append numbers after decimal
	if l.Position < len(l.Input) && l.Input[l.Position] == '.' {
		num += string(l.Input[l.Position])
		l.Position++

		// Accept at least one digit after the decimal
		for l.Position < len(l.Input) && unicode.IsDigit(l.Input[l.Position]) {
			num += string(l.Input[l.Position])
			l.Position++
		}
	}

	l.Position--
	return token.NewToken(token.Number, num)
}

func (l *Lexer) handleDefault(char rune) token.Token {
	return token.NewToken(token.Illegal, string(char))
}
