/*
	The Abstract Syntax Tree is a heirarchial collection of 'nodes' that contain structured information about the source code
	Tokens are to be parsed one-by-one and structured respectively
	You can even view an AST more simply as a JSON structure, where the 'root' of the JSON object is "Program"
	Structuring source code information like this allows for use to interpret the source recursively and much more simply
*/

package ast

import "github.com/ajtroup1/interpreters/parsing/token"

/*
	The AST consists solely of 'nodes', which are higher-level building blocks that form the coded program
	Each node is created differently since they serve much different purposes
		Ex. a 'let' statement holds completely different values and keys than an expression statement
	However, every node must be able to provide a literal value for its token(s) and its respective type differentiator
*/
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode() // Enforces type security for statements
}

type Expression interface {
	Node
	expressionNode() // Enforces type security for expressions
}

/* 
	---------------------------------------------------------------------------------------------------------------------
	**ALL STATEMENTS**     **ALL STATEMENTS**     **ALL STATEMENTS**     **ALL STATEMENTS**
	---------------------------------------------------------------------------------------------------------------------
*/

/*
	The Program node is the root of the AST
	Since Clear is essentially just a collection of statements, this holds a list of every statement declared in the program
	Additionally, since all expressions can be encapsulated into statements, this holds the entire program
*/
type Program struct {
	Statements []Statement
}

// Returns the first literal value in the entire Program node, giving a rough description of the start of the program
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

/* 
	---------------------------------------------------------------------------------------------------------------------
	**ALL EXPRESSIONS**     **ALL EXPRESSIONS**     **ALL EXPRESSIONS**     **ALL EXPRESSIONS**
	---------------------------------------------------------------------------------------------------------------------
*/