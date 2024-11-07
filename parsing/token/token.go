/*
  token
  Defines all tokens that build expressions, statements, etc. in the parser
  Tokens act as building blocks for bigger building blocks
  Contains all unique characters and combinations of them used in Clear
*/
package token

type TokenType string

/*
  A token is a structured piece of information regarding a "word" or segment of text
  For simple chars like "+", "=", "/", just store the type of character it is to evaluate later
    Ex. Token = { type: PLUS, value: nil }
  For more complex data like identifiers, store the type as the corresponding data type and a Literal value of the token
    Ex. Token = { type: IDENTIFIER, value: "myValue" }
    Ex. Token = { type: INT, value: "123" }
*/
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // indentifier for variables
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	STAR   = "*"
	SLASH  = "/"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// This map defines all keywords in the Clear language and maps them to their respective token
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

/*
  Function used to check if a string is a keyword
  Checks through all keywords from the map above and if it matches one:
    It returns the corresponding token
  If it does not match any keywords
    It is simply an identifier
*/
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
