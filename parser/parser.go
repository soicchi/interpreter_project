package parser

import (
	"fmt"

	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	lexer *lexer.Lexer
	currentToken token.Token
	peekToken token.Token
	errors []string
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer: lexer,
		errors: []string{},
	}
	parser.nextToken()
	parser.nextToken()

	return parser
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOF {
		statement := p.parserLetStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parserLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.isCurrentToken(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) isCurrentToken(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) isPeekToken(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.isPeekToken(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}