package parser

import (
	"rococcoon/go_json_test/ast"
	"rococcoon/go_json_test/token"
	"strconv"
)

type Parser struct {
	currentToken token.Token
	peekToken    token.Token
	Tokens       []token.Token
	position     int
}

func NewParser(tokens []token.Token) *Parser {
	p := &Parser{
		Tokens:   tokens,
		position: 0,
	}
	return p
}

func (p *Parser) nextToken() {
	p.position++
	if p.position < len(p.Tokens) {
		p.currentToken = p.Tokens[p.position]
		if p.position+1 < len(p.Tokens) {
			p.peekToken = p.Tokens[p.position+1]
		} else {
			p.peekToken = token.Token{Type: token.EOF, Literal: "EOF"}
		}
	} else {
		p.currentToken = token.Token{Type: token.EOF, Literal: "EOF"}
		p.peekToken = token.Token{Type: token.EOF, Literal: "EOF"}
	}
}

func (p *Parser) ParseRoot() *ast.RootNode {
	p.currentToken = p.Tokens[p.position]
	p.peekToken = p.Tokens[p.position+1]

	val := p.ParseValue()

	return &ast.RootNode{
		Value: val,
	}
}

func (p *Parser) ParseValue() ast.Value {
	switch p.currentToken.Type {
	case token.String:
		return p.handleString()
	case token.Number:
		return p.handleNumber()
	case token.True, token.False:
		return p.handleBool()
	case token.Null:
		return p.handleNull()
	case token.LeftBrace:
		return p.handleObject()
	case token.LeftBracket:
		return p.handleArray()
	case token.EOF:
		return p.handleEOF()
	default:
		return nil
	}
}

func (p *Parser) handleString() ast.Value {
	str := &ast.StringLiteral{
		Value: p.currentToken.Literal,
	}
	p.nextToken()
	return str
}

func (p *Parser) handleNumber() ast.Value {
	val, err := strconv.ParseFloat(p.currentToken.Literal, 64)
	if err != nil {
		return &ast.IllegalLiteral{
			Message: "Error parsing number",
		}
	}

	num := &ast.NumberLiteral{
		Value: val,
	}

	p.nextToken()
	return num
}

func (p *Parser) handleBool() ast.Value {
	var val bool

	switch p.currentToken.Literal {
	case "true":
		val = true
	case "false":
		val = false
	default:
		return &ast.IllegalLiteral{
			Message: "Error parsing boolean",
		}
	}

	p.nextToken()
	return &ast.BooleanLiteral{
		Value: val,
	}
}

func (p *Parser) handleNull() ast.Value {
	if p.currentToken.Literal != "null" {
		return &ast.IllegalLiteral{
			Message: "Error handling null",
		}
	}

	p.nextToken()
	return &ast.NullLiteral{}
}

func (p *Parser) handleObject() ast.Value {
	obj := &ast.Object{}
	pairs := []ast.Property{}

	// Advance past the '{'
	p.nextToken()

	for p.currentToken.Type != token.RightBrace && p.currentToken.Type != token.EOF {
		// Expect string key
		if p.currentToken.Type != token.String {
			break
		}
		key := p.currentToken.Literal

		p.nextToken() // move to colon

		if p.currentToken.Type != token.Colon {
			break // malformed JSON
		}

		p.nextToken() // move to value
		value := p.ParseValue()

		pairs = append(pairs, ast.Property{
			Key:   key,
			Value: value,
		})

		// If there's a comma, move past it and continue
		if p.currentToken.Type == token.Comma {
			p.nextToken()
		}
	}

	// Expect closing '}'
	if p.currentToken.Type == token.RightBrace {
		p.nextToken()
	}

	obj.Pairs = pairs
	return obj
}

func (p *Parser) handleArray() ast.Value {
	arr := &ast.Array{}
	elements := []ast.Value{}

	// skip ove the bracket "["
	p.nextToken()

	for p.currentToken.Type != token.RightBracket && p.currentToken.Type != token.EOF {
		elem := p.ParseValue()
		if elem != nil {
			elements = append(elements, elem)
		}

		if p.currentToken.Type == token.Comma {
			p.nextToken()
		} else {
			break
		}

		if p.currentToken.Type == token.RightBracket {
			p.nextToken()
		}
	}

	p.nextToken()
	arr.Elements = elements
	return arr
}

func (p *Parser) handleEOF() ast.Value {
	return nil
}
