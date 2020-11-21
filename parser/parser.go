package parser

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

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) ParseProgram() *ast.Program {

	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
	// program = newProgramASTNode()

	// advanceTokens()

	// for currentToken() != EOF_TOKEN {
	// 	statement = null

	// 	if currentToken() == LET_TOKEN {
	// 		statement = parseLetStatement()
	// 	} else if currentToken() == RETURN_TOKEN {
	// 		statement = parseReturnStatement()
	// 	} else if currentToken() == IF_TOKEN {
	// 		statement = parseIfStatement()
	// 	}

	// 	if statement != null {
	// 		program.Statements.push(statement)
	// 	}

	// 	advanceTokens()
	// }

	// return program
}

// Pseudo-code we're going for
// function parseLetStatement() {
// 	advanceTokens()

// 	identifier = parseIdentifier()

// 	advanceTokens()

// 	if currentToken() != EQUAL_TOKEN {
// 		parseError("no equal sign")
// 		return null
// 	}

// 	advanceTokens()

// 	value = parseExpression()

// 	variableStatement = newVariableStatementASTNode()
// 	variableStatement.identifier = identifier
// 	variableStatement.value = value

// 	return variableStatement
// }

// function parseIdentifier() {
// 	identifier = newIdentifierASTNode()
// 	identifier.token = currentToken()

// 	return identifier
// }

// function parseExpression() {
// 	if (currentToken() == INTERGER_TOKEN) {
// 		if (nextToken() == PLUS_TOKEN) {
// 			return parseOperationExpression()
// 		} else if (nextToken() == SEMICOLON_TOKEN) {
// 			return parseIntegerLiteral()
// 		}
// 	} else if (currentToken() == LEFT_PAREN) {
// 		return parsedGroupedExpression()
// 	}
// }

// function parseOperatorExpression() {
// 	operatorExpression = newOperatorExpression()

// 	operatorExpression.left = parseIntegerLiteral()
// 	advanceTokens()
// 	operatorExpression.operator = currentToken()
// 	advanceTokens()
// 	operatorExpression.right = parseExpression()

// 	return operatorExpression
// }
