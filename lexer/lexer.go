package lexer

import "github.com/agp745/Interpreter-Go/token"

type Lexer struct {
	input string
	position int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	char byte // current char under examination

}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

//Gives the next character and advances position in the input string.
func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.char = 0
	} else {
		lex.char = lex.input[lex.readPosition]
	}

	lex.position = lex.readPosition
	lex.readPosition += 1
}

func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	switch lex.char {
	case '=':
		tok = newToken(token.ASSIGN, lex.char)
	case '+':
		tok = newToken(token.PLUS, lex.char)
	case ';':
		tok = newToken(token.SEMICOLON, lex.char)
	case ',':
		tok = newToken(token.COMMA, lex.char)
	case '(':
		tok = newToken(token.LPAREN, lex.char)
	case ')':
		tok = newToken(token.RPAREN, lex.char)
	case '{':
		tok = newToken(token.LBRACE, lex.char)
	case '}':
		tok = newToken(token.RBRACE, lex.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lex.char) {
			tok.Literal = lex.readIdentifier()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lex.char)
		}
	}

	lex.readChar()
	return tok
}


func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type: tokenType,
		Literal: string(ch),
	}
}

func (lex *Lexer) readIdentifier() string {
	postion := lex.position
	for isLetter(lex.char) {
		lex.readChar()
	}
	return lex.input[postion:lex.position]
}

func isLetter(char byte) bool {
	return 'a' <= char && char >= 'z' || 'A' <= char && char >= 'Z' || char == '_'
}