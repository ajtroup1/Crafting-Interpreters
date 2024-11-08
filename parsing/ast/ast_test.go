package ast

import (
	"fmt"
	"testing"

	"github.com/ajtroup1/interpreters/parsing/token"
	"github.com/ajtroup1/interpreters/util"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf(util.RedText(fmt.Sprintf("program.String() wrong. got=%q", program.String())))
	}
}
