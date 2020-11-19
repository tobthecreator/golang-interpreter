import (
	"github.com/tobthecreator/golang-interpreter/ast"
	"github.com/tobthecreator/golang-interpreter/lexer"
	"github.com/tobthecreator/golang-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken() // only sets peekToken to the first token
	p.nextToken() // currentToken and peekToken set

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}