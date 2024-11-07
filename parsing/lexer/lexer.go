package lexer

import "github.com/ajtroup1/interpreters/parsing/token"

/*
	Overall state and structure for the Lexer for Clear
	The Lexer processes source code char by char and forms a string of tokens to be parsed
	The Lexer recieves raw input text, processes the chars individually, and makes decisions based on the read char
	Clear's Lexer does not consider whitespace
		Why does Python's consider whitespace?
*/
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

/*
	Instantiates and returns a new instance of Lexer
	Recieves the entire source code to assign to the Lexer state
*/
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Simple (but crucial) helper function to either return the current char, update the Lexer state, and check for EOF
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

/*
	The core decision-making function of the Lexer to be called repeatedly
	Reads the current char and returns its corresponding token
	If no 'direct' match is found, the result is read as an identifier or integer
		If it does not match either of these conditions, it is an illegal character
*/
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Clear does not consider whitespace
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
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}
	l.readChar()
	return tok
}

// Helper function to abstract creating a token object
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Helper function to read identifiers (starting with char) until a non-alphanumeric character is encountered
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isAlphanumeric(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Returns whether the char is the in alphabet (upper or lower) or '_' (bool)
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

/*
	Returns true if the character is in the alphabet (upper or lower)
	Also returns true if the character is a digit (0 -> 9)
	Since identifiers cannot begin with a number, but can continue with a nunmber, a separate check is needed than isLetter()
*/
func isAlphanumeric(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ('0' <= ch && ch <= '9') || ch == '_'
}

// Simple function to remove any whitespace character in front of the Lexer until a non-whitespace character is encountered
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Tracks character positions until a non-digit is encountered
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Returns true if the character is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
