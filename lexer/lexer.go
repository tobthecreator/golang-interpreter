package lexer

import (
	"github.com/tobthecreator/golang-interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	// Init a new Lexer with the input string, other values will be defaults
	l := &Lexer{input: input}

	// read the first character and increment readPosition
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// If the current l.ch is a whitespace character, push past
	// to the next valuable character to parse
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// If not a single ch, and if current ch is a Letter
		if isLetter(l.ch) {
			// Iterate through until we hit a non-letter
			// the sequence of letters is the new literal
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // avoids the readChar() below
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			// If the above is not true, this character is unsupported
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	// We'll want to expand this for common string expressions
	// Ideas: ?!@$%
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9' // Int only, for now
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

/*
currentPosition
 v
[x][ ][=][ ][5][ ][+][ ][5];
	^
	readPosition

When readPosition exceeds the length of the file, we'll set the character to EOF
*/
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for NUL, will work as our EOF
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}
