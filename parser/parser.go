package parser

import (
	"fmt"

	"github.com/agp745/Interpreter-Go/ast"
	"github.com/agp745/Interpreter-Go/lexer"
	"github.com/agp745/Interpreter-Go/token"
)

type Parser struct {
	lex *lexer.Lexer

	currToken token.Token
	peekToken token.Token
  errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
    lex: l,
    errors: []string{},
  }

	// read 2 tokens so that currToken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
  return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
  msg := fmt.Sprintf("Expected next token to be %s, got %s instead", t, p.peekToken.Type)
  p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		currStatement := p.parseStatement()
		if currStatement != nil {
			program.Statements = append(program.Statements, currStatement)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
  case token.RETURN:
    return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{ Token: p.currToken }

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{ Token: p.currToken, Value: p.currToken.Literal }

	if !p.expectPeek((token.ASSIGN)) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
  stmt := &ast.ReturnStatement{
    Token: p.currToken,
  }

  p.nextToken()

  // TODO: We're skipping the expressions until we encounter a semicolon
  for !p.currTokenIs(token.SEMICOLON) {
    p.nextToken()
  }
  
  return stmt
}

func (p *Parser) currTokenIs(t token.TokenType) bool {
	return p.currToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
    p.peekError(t)
		return false
	}
}

