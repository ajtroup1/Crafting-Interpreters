/*
	The Parser interprets tokens one-by-one, using the current token and one of the upcoming tokens to make decisions about structuring source code information
	Heavily relies on recursive descent parsing, or "Pratt Parsing", which uses precedences to evaluate expressions in their correct order
	Starting with a new Program node, the parser creates respective "child nodes" that form the tree and hold information about the program's instructions
*/

package parser

import (
	"github.com/ajtroup1/interpreters/parsing/ast"
	"github.com/ajtroup1/interpreters/parsing/lexer"
	"github.com/ajtroup1/interpreters/parsing/token"
)

/*
Overall state for the Parser
Simply keeps track of the current token and the token X positions away
*/
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}

// Small helper function to advance both the current and peek token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

/*
Encounters each new statement one by one and parses it then adds it to the Program's statement list
*/
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

/*
	**STATEMENT PARSING**
 */

/*
	Parses every statement in Clear
	Similarly to the lexer, it switches the type of statement and calls the respective function to assign its node
	Returns the structured node and appends to Program in ParseProgram
*/
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// Handles assigning let statement information to a corresponding Let node
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}


// -----------------------------------------------------------------------------------------


// Conditional functions that act as type checks for current and peek token
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
// Conditional returning whether the peeked token is a token type
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}
